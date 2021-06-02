module knative.dev/eventing-autoscaler-keda

go 1.16

require (
	github.com/kedacore/keda/v2 v2.3.0
	go.uber.org/zap v1.16.0
	k8s.io/api v0.20.7
	k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery v0.20.7
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.20.7
	knative.dev/eventing v0.23.1-0.20210528215830-b813a54f049b
	knative.dev/eventing-awssqs v0.23.1-0.20210528155729-47d2aa10fba2
	knative.dev/eventing-kafka v0.23.1-0.20210528154330-f77841ca86d9
	knative.dev/eventing-redis v0.23.1-0.20210525202916-98f6a2a1597e
	knative.dev/hack v0.0.0-20210428122153-93ad9129c268
	knative.dev/pkg v0.0.0-20210528203030-47dfdcfaedfd
	knative.dev/reconciler-test v0.0.0-20210528174829-f667a8f5433e
)

replace (
	k8s.io/api => k8s.io/api v0.19.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.7
	k8s.io/apiserver => k8s.io/apiserver v0.19.7
	k8s.io/client-go => k8s.io/client-go v0.19.7
	k8s.io/code-generator => k8s.io/code-generator v0.19.7
)
