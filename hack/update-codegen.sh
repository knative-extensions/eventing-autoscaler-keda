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

# If we run with -mod=vendor here, then generate-groups.sh looks for vendor files in the wrong place.
export GOFLAGS=-mod=

echo "=== Update Codegen for $MODULE_NAME"

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.

# KEDA uses Kubebuilder
# Kubebuilder project layout has API under 'api/v1alpha1', ie. 'github.com/kedacore/keda/api/v1alpha1'
# client-go codegen expects group name (keda) in the path, ie. 'github.com/kedacore/keda/api/keda/v1alpha1'
# Because there's no way how to modify any of these settings, to enable client codegen,
# we need to reorganize things a little bit (copy to 'third_party/api/keda/v1alpha1')
rm -rf ${REPO_ROOT_DIR}/third_party/pkg/apis/keda
mkdir -p ${REPO_ROOT_DIR}/third_party/pkg/apis/keda
cp -R "${REPO_ROOT_DIR}/vendor/github.com/kedacore/keda/v2/api/v1alpha1" ${REPO_ROOT_DIR}/third_party/pkg/apis/keda/

group "Kubernetes Codegen"

# Generate our own client (otherwise injection won't work)
${CODEGEN_PKG}/generate-groups.sh "client,informer,lister" \
  knative.dev/eventing-autoscaler-keda/third_party/pkg/client knative.dev/eventing-autoscaler-keda/third_party/pkg/apis \
  "keda:v1alpha1" \
  --go-header-file ${REPO_ROOT_DIR}/hack/boilerplate/boilerplate.go.txt

group "Knative Codegen"

${KNATIVE_CODEGEN_PKG}/hack/generate-knative.sh "injection" \
  knative.dev/eventing-autoscaler-keda/third_party/pkg/client knative.dev/eventing-autoscaler-keda/third_party/pkg/apis \
  "keda:v1alpha1" \
  --go-header-file ${REPO_ROOT_DIR}/hack/boilerplate/boilerplate.go.txt

group "Update deps post-codegen"

# Make sure our dependencies are up-to-date
${REPO_ROOT_DIR}/hack/update-deps.sh
