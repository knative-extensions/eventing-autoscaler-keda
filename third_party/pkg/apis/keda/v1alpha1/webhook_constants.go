/*
Copyright 2026 The KEDA Authors

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

package v1alpha1

// These constants originate in KEDA's admission webhook sources
// (scaledobject_webhook.go, triggerauthentication_webhook.go), which this project
// strips out during codegen because they import sigs.k8s.io/controller-runtime and
// would panic at controller startup (duplicate "kubeconfig" flag). The remaining
// *_types.go files still reference them, so they are kept here.
const (
	cpuString      = "cpu"
	memoryString   = "memory"
	workloadString = "workload"
)
