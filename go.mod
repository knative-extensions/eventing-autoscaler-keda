module knative.dev/eventing-autoscaler-keda

go 1.15

require (
	github.com/kedacore/keda/v2 v2.2.0
	go.uber.org/zap v1.16.0
	k8s.io/api v0.20.4
	k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.20.4
	knative.dev/eventing v0.22.1-0.20210511163447-55ec6f5801e2
	knative.dev/eventing-awssqs v0.22.1-0.20210420222937-9c231d2460b6
	knative.dev/eventing-kafka v0.22.1-0.20210510225737-4027f852ca6c
	knative.dev/eventing-redis v0.22.1-0.20210511155646-dbba73d10a75
	knative.dev/hack v0.0.0-20210428122153-93ad9129c268
	knative.dev/pkg v0.0.0-20210510175900-4564797bf3b7
	knative.dev/reconciler-test v0.0.0-20210506205310-ed3c37806817
)

replace (
	k8s.io/api => k8s.io/api v0.19.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.7
	k8s.io/apiserver => k8s.io/apiserver v0.19.7
	k8s.io/client-go => k8s.io/client-go v0.19.7
	k8s.io/code-generator => k8s.io/code-generator v0.19.7
)
