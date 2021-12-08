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
	knative.dev/eventing v0.27.1-0.20211207155857-e00a9f4df1cc
	knative.dev/eventing-awssqs v0.27.0
	knative.dev/eventing-kafka v0.27.1-0.20211207151905-587f1bedbb36
	knative.dev/eventing-redis v0.27.1-0.20211207072257-b4a73be30369
	knative.dev/hack v0.0.0-20211203062838-e11ac125e707
	knative.dev/pkg v0.0.0-20211206113427-18589ac7627e
	knative.dev/reconciler-test v0.0.0-20211207070557-0d138a88867b
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
