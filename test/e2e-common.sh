#!/usr/bin/env bash

# Copyright 2021 The Knative Authors
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

# This script runs the end-to-end tests against eventing-contrib built
# from source.

# If you already have the *_OVERRIDE environment variables set, call
# this script with the --run-tests arguments and it will use the cluster
# and run the tests.
# Note that local clusters often do not have the resources to run 12 parallel
# tests (the default) as the tests each tend to create their own namespaces and
# dispatchers.  For example, a local Docker cluster with 4 CPUs and 8 GB RAM will
# probably be able to handle 6 at maximum.  Be sure to adequately set the
# MAX_PARALLEL_TESTS variable before running this script, with the caveat that
# lowering it too much might make the tests run over the timeout that
# the go_test_e2e commands are using below.

# This script includes common functions for testing setup and teardown.

TEST_PARALLEL=${MAX_PARALLEL_TESTS:-12}

source $(dirname $0)/../vendor/knative.dev/hack/e2e-tests.sh

# If gcloud is not available make it a no-op, not an error.
which gcloud &> /dev/null || gcloud() { echo "[ignore-gcloud $*]" 1>&2; }

# Use GNU Tools on MacOS (Requires the 'grep' and 'gnu-sed' Homebrew formulae)
if [ "$(uname)" == "Darwin" ]; then
  sed=gsed
  grep=ggrep
fi

# main config path from HEAD.
readonly KAFKA_SOURCE_CONFIG="./config/source/single"
readonly KAFKA_MT_SOURCE_CONFIG="./config/source/multi"

# Vendored Eventing Test Images.
readonly VENDOR_EVENTING_TEST_IMAGES="vendor/knative.dev/eventing/test/test_images/"
# HEAD eventing test images.
readonly HEAD_EVENTING_TEST_IMAGES="${GOPATH}/src/knative.dev/eventing/test/test_images/"

# Config tracing config.
readonly CONFIG_TRACING_CONFIG="test/config/config-tracing.yaml"

# Strimzi Kafka Cluster Brokers URL (base64 encoded value for k8s secret)
readonly STRIMZI_KAFKA_NAMESPACE="kafka" # Installation Namespace
readonly STRIMZI_KAFKA_CLUSTER_BROKERS="my-cluster-kafka-bootstrap.kafka.svc:9092"

# Eventing Kafka Channel CRD Secret (Will be modified with Strimzi Cluster Brokers - No Authentication)
readonly EVENTING_KAFKA_SECRET_TEMPLATE="300-kafka-secret.yaml"

# Strimzi installation config template used for starting up Kafka clusters.
readonly STRIMZI_INSTALLATION_CONFIG_TEMPLATE="test/config/100-strimzi-cluster-operator-0.20.0.yaml"
# Strimzi installation config.
readonly STRIMZI_INSTALLATION_CONFIG="$(mktemp)"
# Kafka cluster CR config file.
readonly KAFKA_INSTALLATION_CONFIG="test/config/100-kafka-ephemeral-triple-2.6.0.yaml"
# Kafka TLS ConfigMap.
readonly KAFKA_TLS_CONFIG="test/config/config-kafka-tls.yaml"
# Kafka SASL ConfigMap.
readonly KAFKA_SASL_CONFIG="test/config/config-kafka-sasl.yaml"
# Kafka Users CR config file.
readonly KAFKA_USERS_CONFIG="test/config/100-strimzi-users-0.20.0.yaml"
# Kafka PLAIN cluster URL
readonly KAFKA_PLAIN_CLUSTER_URL="my-cluster-kafka-bootstrap.kafka.svc.cluster.local:9092"
# Kafka TLS cluster URL
readonly KAFKA_TLS_CLUSTER_URL="my-cluster-kafka-bootstrap.kafka.svc.cluster.local:9093"
# Kafka SASL cluster URL
readonly KAFKA_SASL_CLUSTER_URL="my-cluster-kafka-bootstrap.kafka.svc.cluster.local:9094"
# Kafka cluster URL for our installation, during tests
KAFKA_CLUSTER_URL=${KAFKA_PLAIN_CLUSTER_URL}

# KEDA
readonly KEDA_INSTALLATION_CONFIG="test/config/keda-2.2.0.yaml"
readonly KEDA_NAMESPACE="keda"

# KEDA-Eventing autoscaler
readonly EVENTING_AUTOSCALER_KEDA_CONFIG="./config"
readonly EVENTING_AUTOSCALER_KEDA_NAMESPACE="knative-eventing"

# Namespaces where we install Eventing components
# This is the namespace of knative-eventing itself
export EVENTING_NAMESPACE="knative-eventing"

# Namespace where we install eventing-kafka components (may be different than EVENTING_NAMESPACE)
readonly SYSTEM_NAMESPACE="knative-eventing"
export SYSTEM_NAMESPACE

# Zipkin setup
readonly KNATIVE_EVENTING_MONITORING_YAML="test/config/monitoring.yaml"

# Public latest nightly or release yaml files.
readonly KNATIVE_EVENTING_KAFKA_SOURCE_RELEASE="$(get_latest_knative_yaml_source "eventing-kafka" "source")"
readonly KNATIVE_EVENTING_KAFKA_MT_SOURCE_RELEASE="$(get_latest_knative_yaml_source "eventing-kafka" "mt-source")"

#
function knative_setup() {
  if [[ "$TEST_KAFKA_SOURCE" == 0 && "$TEST_KAFKA_MT_SOURCE" == 0 ]]; then
    echo "ERROR: missing or invalid test selector flag (either --kafka-source or --kafka-mt-source)"
    exit 1
  fi

  if is_release_branch; then
    if [[ "$TEST_KAFKA_SOURCE" == "1" ]]; then
      echo ">> Install Kafka source from ${KNATIVE_EVENTING_KAFKA_SOURCE_RELEASE}"
      kubectl apply -f ${KNATIVE_EVENTING_KAFKA_SOURCE_RELEASE}
    elif [[ "$TEST_KAFKA_MT_SOURCE" == "1" ]]; then
      echo ">> Install Kafka mt source from ${KNATIVE_EVENTING_KAFKA_MT_SOURCE_RELEASE}"
      kubectl apply -f ${KNATIVE_EVENTING_KAFKA_MT_SOURCE_RELEASE}
    fi
  else
    if [[ "$TEST_KAFKA_SOURCE" == "1" ]]; then
      echo ">> Install Kafka source from HEAD"
      pushd .
      cd ${GOPATH} && mkdir -p src/knative.dev && cd src/knative.dev
      if [[  ! -d "eventing-kafka" ]]; then
        git clone https://github.com/knative-sandbox/eventing-kafka
      fi
      cd eventing-kafka
      ko apply -f "${KAFKA_SOURCE_CONFIG}"
      popd
    elif [[ "$TEST_KAFKA_MT_SOURCE" == "1" ]]; then
      echo ">> Install Kafka source from HEAD"
      pushd .
      cd ${GOPATH} && mkdir -p src/knative.dev && cd src/knative.dev
      if [[  ! -d "eventing-kafka" ]]; then
        git clone https://github.com/knative-sandbox/eventing-kafka
      fi
      cd eventing-kafka
      ko apply -f "${KAFKA_MT_SOURCE_CONFIG}"
      popd
    fi
  fi
  wait_until_pods_running "${EVENTING_NAMESPACE}" || fail_test "Knative Kafka did not come up"

  kafka_setup
  keda_setup
  #zipkin_setup # not used yet.
}

# Setup zipkin
function zipkin_setup() {
  echo "Installing Zipkin..."
  sed "s/\${SYSTEM_NAMESPACE}/${SYSTEM_NAMESPACE}/g" < "${KNATIVE_EVENTING_MONITORING_YAML}" | kubectl apply -f -
  wait_until_pods_running "${SYSTEM_NAMESPACE}" || fail_test "Zipkin inside eventing did not come up"
  # Setup config tracing for tracing tests
  sed "s/\${SYSTEM_NAMESPACE}/${SYSTEM_NAMESPACE}/g" <  "${CONFIG_TRACING_CONFIG}" | kubectl apply -f -
}

# Remove zipkin
function zipkin_teardown() {
  echo "Uninstalling Zipkin..."
  sed "s/\${SYSTEM_NAMESPACE}/${SYSTEM_NAMESPACE}/g" <  "${KNATIVE_EVENTING_MONITORING_YAML}" | kubectl delete -f -
  wait_until_object_does_not_exist deployment zipkin "${SYSTEM_NAMESPACE}" || fail_test "Zipkin deployment was unable to be deleted"
  kubectl delete -n "${SYSTEM_NAMESPACE}" configmap config-tracing
}

function knative_teardown() {
  echo ">> Stopping Knative Eventing"
  if is_release_branch; then
  if [[ "$TEST_KAFKA_SOURCE" == "1" ]]; then
      echo ">> Uninstalling Kafka source from ${KNATIVE_EVENTING_KAFKA_SOURCE_RELEASE}"
      kubectl delete -f ${KNATIVE_EVENTING_KAFKA_SOURCE_RELEASE}
    elif [[ "$TEST_KAFKA_MT_SOURCE" == "1" ]]; then
      echo ">> Uninstalling Kafka mt source from ${KNATIVE_EVENTING_KAFKA_MT_SOURCE_RELEASE}"
      kubectl delete -f ${KNATIVE_EVENTING_KAFKA_MT_SOURCE_RELEASE}
    fi
  else
     if [[ "$TEST_KAFKA_SOURCE" == "1" ]]; then
      echo ">> Install Kafka source from HEAD"
      pushd .
      cd ${GOPATH}/src/knative.dev/eventing-kafka
      ko delete -f "${KAFKA_SOURCE_CONFIG}"
      popd
    elif [[ "$TEST_KAFKA_MT_SOURCE" == "1" ]]; then
      echo ">> Install Kafka source from HEAD"
      pushd .
      cd ${GOPATH}/src/knative.dev/eventing-kafka
      ko delete -f "${KAFKA_MT_SOURCE_CONFIG}"
      popd
    fi
  fi

  kafka_teardown
  keda_teardown
  #zipkin_teardown
}

# setup strimzi
function kafka_setup() {
  # Create The Namespace Where Strimzi Kafka Will Be Installed
  echo "Installing Kafka Cluster"
  kubectl get -o name namespace ${STRIMZI_KAFKA_NAMESPACE} || kubectl create namespace ${STRIMZI_KAFKA_NAMESPACE} || return 1

  # Install Strimzi Into The Desired Namespace (Dynamically Changing The Namespace)
  sed "s/namespace: .*/namespace: ${STRIMZI_KAFKA_NAMESPACE}/" ${STRIMZI_INSTALLATION_CONFIG_TEMPLATE} > "${STRIMZI_INSTALLATION_CONFIG}"

  # Create The Actual Kafka Cluster Instance For The Cluster Operator To Setup
  kubectl apply -f "${STRIMZI_INSTALLATION_CONFIG}" -n "${STRIMZI_KAFKA_NAMESPACE}"
  kubectl apply -f "${KAFKA_INSTALLATION_CONFIG}" -n "${STRIMZI_KAFKA_NAMESPACE}"

  # Delay Pod Running Check Until All Pods Are Created To Prevent Race Condition (Strimzi Kafka Instance Can Take A Bit To Spin Up)
  local iterations=0
  local progress="Waiting for Kafka Pods to be created..."
  while [[ $(kubectl get pods --no-headers=true -n ${STRIMZI_KAFKA_NAMESPACE} | wc -l) -lt 6 && $iterations -lt 60 ]] # 1 ClusterOperator, 3 Zookeeper, 1 Kafka, 1 EntityOperator
  do
    echo -ne "${progress}\r"
    progress="${progress}."
    iterations=$((iterations + 1))
    sleep 3
  done
  echo "${progress}"

  # Wait For The Strimzi Kafka Cluster Operator To Be Ready (Forcing Delay To Ensure CRDs Are Installed To Prevent Race Condition)
  wait_until_pods_running "${STRIMZI_KAFKA_NAMESPACE}" || fail_test "Failed to start up a Strimzi Kafka Instance"

  # Create some Strimzi Kafka Users
  kubectl apply -f "${KAFKA_USERS_CONFIG}" -n "${STRIMZI_KAFKA_NAMESPACE}"
}

function kafka_teardown() {
  echo "Uninstalling Kafka cluster"
  kubectl delete -f ${KAFKA_INSTALLATION_CONFIG} -n "${STRIMZI_KAFKA_NAMESPACE}"
  kubectl delete -f "${STRIMZI_INSTALLATION_CONFIG}" -n "${STRIMZI_KAFKA_NAMESPACE}"
  kubectl delete namespace "${STRIMZI_KAFKA_NAMESPACE}"
}

# setup keda
function keda_setup() {
  echo "Installing KEDA"
  kubectl apply -f "$KEDA_INSTALLATION_CONFIG"

  wait_until_pods_running "${KEDA_NAMESPACE}" || fail_test "Failed to start up KEDA"
}

# uninstall keda
function keda_teardown() {
  echo "Uninstalling KEDA"
  kubectl delete -f "$KEDA_INSTALLATION_CONFIG"
}

# Add function call to trap
# Parameters: $1 - Function to call
#             $2...$n - Signals for trap
function add_trap() {
  local cmd=$1
  shift
  for trap_signal in $@; do
    local current_trap="$(trap -p $trap_signal | cut -d\' -f2)"
    local new_cmd="($cmd)"
    [[ -n "${current_trap}" ]] && new_cmd="${current_trap};${new_cmd}"
    trap -- "${new_cmd}" $trap_signal
  done
}

function test_setup() {
  # Install kail if needed.
  if ! which kail > /dev/null; then
    bash <( curl -sfL https://raw.githubusercontent.com/boz/kail/master/godownloader.sh) -b "$GOPATH/bin"
  fi

  # Capture all logs.
  kail > "${ARTIFACTS}/k8s.log.txt" &
  local kail_pid=$!
  # Clean up kail so it doesn't interfere with job shutting down
  add_trap "kill $kail_pid || true" EXIT

  # Publish test images.
  echo ">> Publishing test images from eventing"
  # We vendor test image code from eventing, in order to use ko to resolve them into Docker images, the
  # path has to be a GOPATH.  The two slashes at the beginning are to anchor the match so that running the test
  # twice doesn't re-parse the yaml and cause errors.
  #sed -i 's@//knative.dev/eventing/test/test_images@//knative.dev/eventing-kafka/vendor/knative.dev/eventing/test/test_images@g' "${VENDOR_EVENTING_TEST_IMAGES}"*/*.yaml
  #$(dirname $0)/upload-test-images.sh "${VENDOR_EVENTING_TEST_IMAGES}" e2e || fail_test "Error uploading test images"
  #$(dirname $0)/upload-test-images.sh "test/test_images" e2e || fail_test "Error uploading test images"

  install_autoscaler
}

function test_teardown() {
  uninstall_autoscaler
}

function create_tls_secrets() {
  echo "Creating TLS Kafka secret"
  STRIMZI_CRT=$(kubectl -n ${STRIMZI_KAFKA_NAMESPACE} get secret my-cluster-cluster-ca-cert --template='{{index .data "ca.crt"}}' | base64 --decode )
  TLSUSER_CRT=$(kubectl -n ${STRIMZI_KAFKA_NAMESPACE} get secret my-tls-user --template='{{index .data "user.crt"}}' | base64 --decode )
  TLSUSER_KEY=$(kubectl -n ${STRIMZI_KAFKA_NAMESPACE} get secret my-tls-user --template='{{index .data "user.key"}}' | base64 --decode )

  kubectl create secret --namespace "${SYSTEM_NAMESPACE}" generic strimzi-tls-secret \
    --from-literal=ca.crt="$STRIMZI_CRT" \
    --from-literal=user.crt="$TLSUSER_CRT" \
    --from-literal=user.key="$TLSUSER_KEY"
}

function create_sasl_secrets() {
  echo "Creating SASL Kafka secret"
  STRIMZI_CRT=$(kubectl -n ${STRIMZI_KAFKA_NAMESPACE} get secret my-cluster-cluster-ca-cert --template='{{index .data "ca.crt"}}' | base64 --decode )
  SASL_PASSWD=$(kubectl -n ${STRIMZI_KAFKA_NAMESPACE} get secret my-sasl-user --template='{{index .data "password"}}' | base64 --decode )

  kubectl create secret --namespace "${SYSTEM_NAMESPACE}" generic strimzi-sasl-secret \
    --from-literal=ca.crt="$STRIMZI_CRT" \
    --from-literal=password="$SASL_PASSWD" \
    --from-literal=saslType="SCRAM-SHA-512" \
    --from-literal=user="my-sasl-user"
}

# Install eventing autoscaler keda
function install_autoscaler() {
  ko apply -f "${EVENTING_AUTOSCALER_KEDA_CONFIG}" || return 1
  wait_until_pods_running "${EVENTING_AUTOSCALER_KEDA_NAMESPACE}" || fail_test "Failed to install eventing autoscaler KEDA"
}

# Uninstall eventing autoscaler keda
function uninstall_autoscaler() {
  ko delete -f "${EVENTING_AUTOSCALER_KEDA_CONFIG}" || return 1
}

# Runs the KafkaSource tests
function test_kafka_source() {
  echo "Testing Kafka source"
  if [[ "$TEST_KAFKA_MT_SOURCE" == 1 ]]; then
     echo multi-tenancy enabled
  fi

  go_test_e2e -tags=e2e -timeout=20m -test.parallel=${TEST_PARALLEL} ./test/e2e/...  || fail_test

  # wait for all KafkaSources to be deleted
  local iterations=0
  local progress="Waiting for KafkaSources to be deleted..."
  while [[ "$(kubectl get kafkasources --all-namespaces)" != "" && $iterations -lt 60 ]]
  do
    echo -ne "${progress}\r"
    progress="${progress}."
    iterations=$((iterations + 1))
    kubectl get kafkasources --all-namespaces -oyaml
    sleep 3
  done
}

function parse_flags() {
  # This function will be called repeatedly by initialize() with one fewer
  # argument each time and expects a return value of "the number of arguments to skip"
  # so we can just check the first argument and return 1 (to have it redirected to the
  # test container) or 0 (to have initialize() parse it normally).
  case $1 in
    --kafka-source)
      TEST_KAFKA_SOURCE=1
      return 1
      ;;
    --kafka-mt-source)
      TEST_KAFKA_MT_SOURCE=1
      return 1
      ;;
  esac
  return 0
}
