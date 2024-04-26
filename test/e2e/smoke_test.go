//go:build e2e
// +build e2e

/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	"testing"
	"time"

	ks "knative.dev/eventing-kafka-broker/test/rekt/resources/kafkasource"
	"knative.dev/pkg/system"
	_ "knative.dev/pkg/system/testing"
	"knative.dev/reconciler-test/pkg/environment"
	"knative.dev/reconciler-test/pkg/k8s"
	"knative.dev/reconciler-test/pkg/knative"
)

const (
	kafkaBootstrapUrlPlain = "my-cluster-kafka-bootstrap.kafka.svc:9092"
)

// TestSmoke_KafkaSource
func TestSmoke_KafkaSource(t *testing.T) {
	t.Parallel()

	ctx, env := global.Environment(
		knative.WithKnativeNamespace(system.Namespace()),
		knative.WithLoggingConfig,
		k8s.WithEventListener,
		environment.WithPollTimings(2*time.Second, 20*time.Second),
		environment.Managed(t),
	)

	env.Test(ctx, t,
		ks.Install("readysource", ks.WithBootstrapServers([]string{kafkaBootstrapUrlPlain})).AsFeature())
}
