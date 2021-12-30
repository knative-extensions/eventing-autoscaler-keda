module knative.dev/eventing-autoscaler-keda

go 1.16

require (
	github.com/kedacore/keda/v2 v2.5.0
	go.uber.org/zap v1.19.1
	k8s.io/api v0.22.4
	k8s.io/apiextensions-apiserver v0.22.2
	k8s.io/apimachinery v0.22.4
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.22.4
	knative.dev/eventing v0.28.1-0.20211230085924-1ba7810fba5d
	knative.dev/eventing-awssqs v0.28.1-0.20211217050719-055902b9ed60
	knative.dev/eventing-kafka v0.28.1-0.20211222111320-3b85e59ad408
	knative.dev/eventing-redis v0.28.1-0.20211221060317-1f3f89538815
	knative.dev/hack v0.0.0-20211222071919-abd085fc43de
	knative.dev/pkg v0.0.0-20211216142117-79271798f696
	knative.dev/reconciler-test v0.0.0-20211222120418-816f2192fec9
	sigs.k8s.io/controller-runtime v0.10.3
)

replace (
	k8s.io/api => k8s.io/api v0.21.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.21.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.21.4
	k8s.io/apiserver => k8s.io/apiserver v0.21.4
	k8s.io/client-go => k8s.io/client-go v0.21.4
	k8s.io/code-generator => k8s.io/code-generator v0.21.4
)
