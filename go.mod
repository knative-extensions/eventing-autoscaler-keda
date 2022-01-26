module knative.dev/eventing-autoscaler-keda

go 1.16

require (
	github.com/kedacore/keda/v2 v2.5.0
	go.uber.org/zap v1.19.1
	k8s.io/api v0.22.5
	k8s.io/apiextensions-apiserver v0.22.5
	k8s.io/apimachinery v0.22.5
	k8s.io/client-go v0.22.5
	k8s.io/code-generator v0.22.5
	knative.dev/eventing v0.29.0
	knative.dev/eventing-awssqs v0.29.0
	knative.dev/eventing-kafka v0.28.1-0.20220126011444-734561d578b9
	knative.dev/eventing-redis v0.29.0
	knative.dev/hack v0.0.0-20220118141833-9b2ed8471e30
	knative.dev/pkg v0.0.0-20220118160532-77555ea48cd4
	knative.dev/reconciler-test v0.0.0-20220118183433-c8bfbe66bada
	sigs.k8s.io/controller-runtime v0.10.3
)
