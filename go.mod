module knative.dev/eventing-autoscaler-keda

go 1.16

require (
	github.com/kedacore/keda/v2 v2.3.0
	go.uber.org/zap v1.17.0
	k8s.io/api v0.20.7
	k8s.io/apiextensions-apiserver v0.20.7
	k8s.io/apimachinery v0.20.7
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.20.7
	knative.dev/eventing v0.24.1-0.20210702080639-90932eb671e0
	knative.dev/eventing-awssqs v0.24.0
	knative.dev/eventing-kafka v0.24.1-0.20210702151339-f6cdc01e6fbc
	knative.dev/eventing-redis v0.24.1-0.20210702190040-c72feb45ce15
	knative.dev/hack v0.0.0-20210622141627-e28525d8d260
	knative.dev/pkg v0.0.0-20210701025203-30f9568e894e
	knative.dev/reconciler-test v0.0.0-20210630182710-2a6d91dfee1e
)

replace (
	k8s.io/api => k8s.io/api v0.19.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.7
	k8s.io/apiserver => k8s.io/apiserver v0.19.7
	k8s.io/client-go => k8s.io/client-go v0.19.7
	k8s.io/code-generator => k8s.io/code-generator v0.19.7
)
