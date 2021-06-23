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
	"strconv"
	"testing"

	"knative.dev/eventing-kafka/test/rekt/features/kafkasource"
	ks "knative.dev/eventing-kafka/test/rekt/resources/kafkasource"
	_ "knative.dev/pkg/system/testing"
)

// TestSmoke_KafkaSource
func TestSmoke_KafkaSource(t *testing.T) {
	t.Parallel()

	ctx, env := global.Environment()
	t.Cleanup(env.Finish)

	names := []string{
		"customname",
		"name-with-dash",
		"name1with2numbers3",
		"name63-0123456789012345678901234567890123456789012345678901234",
	}

	configs := [][]ks.CfgFn{
		{
			ks.WithBootstrapServers([]string{"my-cluster-kafka-bootstrap.kafka:9092"}),
		},
	}

	for _, name := range names {
		for i, cfg := range configs {
			n := name + strconv.Itoa(i) // Make the name unique for each config.
			if len(n) >= 64 {
				n = n[:63] // 63 is the max length.
			}
			env.Test(ctx, t, kafkasource.KafkaSourceGoesReady(n, cfg...))
		}
	}
}
