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
	knative.dev/eventing v0.28.1-0.20211230085924-1ba7810fba5d
	knative.dev/eventing-awssqs v0.28.1-0.20211217050719-055902b9ed60
	knative.dev/eventing-kafka v0.28.1-0.20211222111320-3b85e59ad408
	knative.dev/eventing-redis v0.28.1-0.20211221060317-1f3f89538815
	knative.dev/hack v0.0.0-20211222071919-abd085fc43de
	knative.dev/pkg v0.0.0-20220104185830-52e42b760b54
	knative.dev/reconciler-test v0.0.0-20211222120418-816f2192fec9
	sigs.k8s.io/controller-runtime v0.10.3
)
