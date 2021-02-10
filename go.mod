module knative.dev/eventing-autoscaler-keda

go 1.15

require (
	github.com/kedacore/keda/v2 v2.1.0
	go.uber.org/zap v1.16.0
	k8s.io/api v0.20.2
	k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.20.0
	knative.dev/eventing v0.20.1-0.20210209233432-7ce82834b39d
	knative.dev/eventing-awssqs v0.20.1-0.20210210082317-a6f4a1579bd9
	knative.dev/eventing-kafka v0.19.1-0.20210210072926-348670ebd780
	knative.dev/eventing-redis v0.20.1-0.20210209230002-a6b0f57a7af6
	knative.dev/hack v0.0.0-20210203173706-8368e1f6eacf
	knative.dev/pkg v0.0.0-20210208175252-a02dcff9ee26
)

replace (
	k8s.io/api => k8s.io/api v0.19.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.7
	k8s.io/apiserver => k8s.io/apiserver v0.19.7
	k8s.io/client-go => k8s.io/client-go v0.19.7
	k8s.io/code-generator => k8s.io/code-generator v0.19.7
)
