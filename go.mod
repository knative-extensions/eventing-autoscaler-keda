module knative.dev/eventing-autoscaler-keda

go 1.16

require (
	github.com/kedacore/keda/v2 v2.4.0
	go.uber.org/zap v1.19.0
	k8s.io/api v0.21.4
	k8s.io/apiextensions-apiserver v0.21.4
	k8s.io/apimachinery v0.21.4
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.21.4
	knative.dev/eventing v0.25.1-0.20210903113057-4a55e519013e
	knative.dev/eventing-awssqs v0.25.1-0.20210831131755-2918b790b116
	knative.dev/eventing-kafka v0.25.1-0.20210903091657-acf06a29bf44
	knative.dev/eventing-redis v0.25.1-0.20210901161033-19fc0b8bef7f
	knative.dev/hack v0.0.0-20210806075220-815cd312d65c
	knative.dev/pkg v0.0.0-20210903055557-dec98b426132
	knative.dev/reconciler-test v0.0.0-20210820180205-a25de6a08087
	sigs.k8s.io/controller-runtime v0.6.5
)

replace (
	k8s.io/api => k8s.io/api v0.21.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.21.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.21.4
	k8s.io/apiserver => k8s.io/apiserver v0.21.4
	k8s.io/client-go => k8s.io/client-go v0.21.4
	k8s.io/code-generator => k8s.io/code-generator v0.21.4
)
