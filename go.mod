module knative.dev/eventing-autoscaler-keda

go 1.16

require (
	github.com/kedacore/keda/v2 v2.3.0
	go.uber.org/zap v1.18.1
	k8s.io/api v0.20.7
	k8s.io/apiextensions-apiserver v0.20.7
	k8s.io/apimachinery v0.20.7
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.20.7
	knative.dev/eventing v0.24.1-0.20210713221030-0af15fd14a18
	knative.dev/eventing-awssqs v0.24.1-0.20210712140124-93961c1cf002
	knative.dev/eventing-kafka v0.24.1-0.20210714130431-7a5b04ed9093
	knative.dev/eventing-redis v0.24.1-0.20210714033730-9ec30de6f320
	knative.dev/hack v0.0.0-20210622141627-e28525d8d260
	knative.dev/pkg v0.0.0-20210713175550-7d1b0f19ef71
	knative.dev/reconciler-test v0.0.0-20210713145651-ebe91f96fea0
)

replace (
	k8s.io/api => k8s.io/api v0.19.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.7
	k8s.io/apiserver => k8s.io/apiserver v0.19.7
	k8s.io/client-go => k8s.io/client-go v0.19.7
	k8s.io/code-generator => k8s.io/code-generator v0.19.7
)
