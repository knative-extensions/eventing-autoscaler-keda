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

# Calling this script without arguments will create a new cluster in
# project $PROJECT_ID, start Knative eventing system, install resources
# in eventing-contrib, run all of the test suites and delete the cluster.

# Variables supported by this script:
#   PROJECT_ID:  the GKR project in which to create the new cluster (unless using "--run-tests")
#   MAX_PARALLEL_TESTS:  The maximum number of go tests to run in parallel (via "-test.parallel", default 12)

TEST_PARALLEL=${MAX_PARALLEL_TESTS:-12}

source "$(dirname "$0")/e2e-common.sh"

# Create the system namespace if it doesn't already exist (may be the same as the EVENTING_NAMESPACE)
kubectl get namespace "${SYSTEM_NAMESPACE}" || kubectl create namespace "${SYSTEM_NAMESPACE}"

TEST_KAFKA_SOURCE=${TEST_KAFKA_SOURCE:-0}
TEST_KAFKA_MT_SOURCE=${TEST_KAFKA_MT_SOURCE:-0}

echo "e2e-tests.sh command line: $@"

# If you wish to use this script just as test setup, *without* teardown, add "--skip-teardowns" to the initialize command
initialize "$@" --skip-istio-addon --cluster-version=1.28

echo "e2e-tests.sh environment:"
echo "EVENTING_NAMESPACE: ${EVENTING_NAMESPACE}"
echo "SYSTEM_NAMESPACE: ${SYSTEM_NAMESPACE}"
echo "TEST_KAFKA_SOURCE: ${TEST_KAFKA_SOURCE}"
echo "TEST_KAFKA_MT_SOURCE: ${TEST_KAFKA_MT_SOURCE}"

if [[ "$TEST_KAFKA_SOURCE" == 1 || "$TEST_KAFKA_MT_SOURCE" == 1 ]]; then
  echo "Launching the KAFKA SOURCE TESTS:"
  test_kafka_source  || exit 1
fi

# Terminate any zipkin port-forward processes that are still present on the system
ps -e -o pid,command | grep 'kubectl port-forward zipkin[^*]*9411:9411 -n' | sed 's/^ *\([0-9][0-9]*\) .*/\1/' | xargs kill

# If you wish to use this script just as test setup, *without* teardown, just uncomment this line and comment all go_test_e2e commands
# trap - SIGINT SIGQUIT SIGTSTP EXIT

success
