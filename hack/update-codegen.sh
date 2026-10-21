#!/usr/bin/env bash

# Copyright 2020 The Knative Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

source $(dirname $0)/../vendor/knative.dev/hack/codegen-library.sh
source "${CODEGEN_PKG}/kube_codegen.sh"

# Run outside vendor mode so the code generators resolve packages from the module cache.
export GOFLAGS=-mod=

echo "=== Update Codegen for $MODULE_NAME"

cp -R "${REPO_ROOT_DIR}/vendor/github.com/kedacore/keda/v2/apis/keda/v1alpha1" ${REPO_ROOT_DIR}/third_party/pkg/apis/keda/

# Drop KEDA's admission webhook sources. They import sigs.k8s.io/controller-runtime,
# whose pkg/client/config init() registers a global "kubeconfig" flag. Our controller
# uses knative.dev/pkg, which also registers "kubeconfig", so importing these would
# panic at startup with "flag redefined: kubeconfig". We only need the API types here.
# A few string constants the *_types.go files still reference are kept in the
# checked-in webhook_constants.go (not overwritten by the cp above).
rm -f ${REPO_ROOT_DIR}/third_party/pkg/apis/keda/v1alpha1/*_webhook.go

# KEDA's generated deepcopy (zz_generated.deepcopy.go, re-copied above) still
# contains DeepCopy funcs for the *CustomValidator types that live only in the
# webhook sources we just deleted. Strip those blocks (and their doc comments)
# so the package compiles, then re-run gofmt to collapse the blank lines left
# behind.
ZZ_DEEPCOPY=${REPO_ROOT_DIR}/third_party/pkg/apis/keda/v1alpha1/zz_generated.deepcopy.go
awk '
  /^\/\// { comment = comment $0 "\n"; next }
  /^func \(in \*[A-Za-z]*CustomValidator\)/ { skip=1; comment=""; next }
  skip==1 { if ($0 ~ /^}/) { skip=0 }; next }
  { printf "%s", comment; comment=""; print }
  END { printf "%s", comment }
' "${ZZ_DEEPCOPY}" > "${ZZ_DEEPCOPY}.tmp" && mv "${ZZ_DEEPCOPY}.tmp" "${ZZ_DEEPCOPY}"
gofmt -w "${ZZ_DEEPCOPY}"

group "Kubernetes Codegen"

# Generate our own client (otherwise injection won't work).
# generate-groups.sh was removed in k8s.io/code-generator; use kube_codegen.sh.
# Only the clientset/listers/informers are generated here. The deepcopy funcs
# (zz_generated.deepcopy.go) come from KEDA via the cp -R above, so gen_helpers
# is intentionally not invoked.
kube::codegen::gen_client \
  --with-watch \
  --output-dir "${REPO_ROOT_DIR}/third_party/pkg/client" \
  --output-pkg "knative.dev/eventing-autoscaler-keda/third_party/pkg/client" \
  --boilerplate "${REPO_ROOT_DIR}/hack/boilerplate/boilerplate.go.txt" \
  "${REPO_ROOT_DIR}/third_party/pkg/apis"

group "Knative Codegen"

${KNATIVE_CODEGEN_PKG}/hack/generate-knative.sh "injection" \
  knative.dev/eventing-autoscaler-keda/third_party/pkg/client knative.dev/eventing-autoscaler-keda/third_party/pkg/apis \
  "keda:v1alpha1" \
  --go-header-file ${REPO_ROOT_DIR}/hack/boilerplate/boilerplate.go.txt

group "Update deps post-codegen"

# Make sure our dependencies are up-to-date
${REPO_ROOT_DIR}/hack/update-deps.sh
