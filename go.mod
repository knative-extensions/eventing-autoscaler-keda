module knative.dev/eventing-autoscaler-keda

go 1.16

require (
	github.com/kedacore/keda/v2 v2.6.1
	go.uber.org/zap v1.21.0
	k8s.io/api v0.23.4
	k8s.io/apiextensions-apiserver v0.23.4
	k8s.io/apimachinery v0.23.4
	k8s.io/client-go v0.23.4
	k8s.io/code-generator v0.23.4
	knative.dev/eventing v0.30.1
	knative.dev/eventing-kafka v0.30.0
	knative.dev/eventing-redis v0.30.1
	knative.dev/hack v0.0.0-20220318020218-14f832e506f8
	knative.dev/pkg v0.0.0-20220318185521-e6e3cf03d765
	knative.dev/reconciler-test v0.0.0-20220321082547-afd9ca7f6603
	sigs.k8s.io/controller-runtime v0.11.1
)
