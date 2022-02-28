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
	knative.dev/eventing v0.29.1-0.20220226194900-cbf1b0863ed6
	knative.dev/eventing-awssqs v0.29.0
	knative.dev/eventing-kafka v0.28.1-0.20220225123842-030679de0c2c
	knative.dev/eventing-redis v0.29.1-0.20220216210340-3371e1969d5a
	knative.dev/hack v0.0.0-20220224013837-e1785985d364
	knative.dev/pkg v0.0.0-20220225161142-708dc1cc48e9
	knative.dev/reconciler-test v0.0.0-20220216192840-2c3291f210ce
	sigs.k8s.io/controller-runtime v0.10.3
)
