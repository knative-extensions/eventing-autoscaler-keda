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

export GO111MODULE=on
# If we run with -mod=vendor here, then generate-groups.sh looks for vendor files in the wrong place.
export GOFLAGS=-mod=

if [ -z "${GOPATH:-}" ]; then
  export GOPATH=$(go env GOPATH)
fi

REPO_ROOT=$(dirname ${BASH_SOURCE})/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${REPO_ROOT}; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

KNATIVE_CODEGEN_PKG=${KNATIVE_CODEGEN_PKG:-$(cd ${REPO_ROOT}; ls -d -1 ./vendor/knative.dev/pkg 2>/dev/null || echo ../pkg)}

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.

# KEDA uses Kubebuilder
# Kubebuilder project layout has API under 'api/v1alpha1', ie. 'github.com/kedacore/keda/api/v1alpha1'
# client-go codegen expects group name (keda) in the path, ie. 'github.com/kedacore/keda/api/keda/v1alpha1'
# Because there's no way how to modify any of these settings, to enable client codegen,
# we need to hack things a little bit (in vendor move temporarily api directory in 'api/keda/v1alpha1')
rm -rf ${REPO_ROOT}/vendor/github.com/kedacore/keda/api/keda
mkdir ${REPO_ROOT}/vendor/github.com/kedacore/keda/api/keda
mv ${REPO_ROOT}/vendor/github.com/kedacore/keda/api/v1alpha1 ${REPO_ROOT}/vendor/github.com/kedacore/keda/api/keda/v1alpha1
cp ${REPO_ROOT}/hack/keda-doc.go ${REPO_ROOT}/vendor/github.com/kedacore/keda/api/keda/v1alpha1/doc.go
cp ${REPO_ROOT}/hack/keda-register.go ${REPO_ROOT}/vendor/github.com/kedacore/keda/api/keda/v1alpha1/register.go

# Groups
chmod +x ${CODEGEN_PKG}/generate-groups.sh
${CODEGEN_PKG}/generate-groups.sh "client,informer,lister" \
  knative.dev/eventing-autoscaler-keda/pkg/client github.com/kedacore/keda/api \
  "keda:v1alpha1" \
  --go-header-file ${REPO_ROOT}/hack/boilerplate/boilerplate.go.txt

find ./pkg -type f -name '*.go' -exec sed -i '' -e 's@github.com/kedacore/keda/api/keda/v1alpha1@github.com/kedacore/keda/api/v1alpha1@g' {} +


# Knative Injection
chmod +x ${KNATIVE_CODEGEN_PKG}/hack/generate-knative.sh
OUTPUT_PKG="knative.dev/eventing-autoscaler-keda/pkg/client/injection/keda" \
VERSIONED_CLIENTSET_PKG="github.com/kedacore/keda/pkg/generated/clientset/versioned" \
EXTERNAL_INFORMER_PKG="github.com/kedacore/keda/pkg/generated/informers/externalversions" \
  ${KNATIVE_CODEGEN_PKG}/hack/generate-knative.sh "injection" \
    github.com/kedacore/keda \
    github.com/kedacore/keda/api \
    "keda:v1alpha1" \
    --go-header-file ${REPO_ROOT}/hack/boilerplate/boilerplate.go.txt \

# Move back api directory, which was moved temporarily for codegen
mv ${REPO_ROOT}/vendor/github.com/kedacore/keda/api/keda/v1alpha1 ${REPO_ROOT}/vendor/github.com/kedacore/keda/api/v1alpha1
rm -rf ${REPO_ROOT}/vendor/github.com/kedacore/keda/api/keda

# Make sure our dependencies are up-to-date
${REPO_ROOT}/hack/update-deps.sh
