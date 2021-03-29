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
	knative.dev/eventing v0.21.1-0.20210325205419-3ebce0d42aa2
	knative.dev/eventing-awssqs v0.21.1-0.20210324153618-f0d9f983ecc0
	knative.dev/eventing-kafka v0.21.1-0.20210326072819-cf8381199923
	knative.dev/eventing-redis v0.21.1-0.20210326191320-d0592ec7c87f
	knative.dev/hack v0.0.0-20210325223819-b6ab329907d3
	knative.dev/pkg v0.0.0-20210329065222-9d92ea16c0d3
)

replace (
	k8s.io/api => k8s.io/api v0.19.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.7
	k8s.io/apiserver => k8s.io/apiserver v0.19.7
	k8s.io/client-go => k8s.io/client-go v0.19.7
	k8s.io/code-generator => k8s.io/code-generator v0.19.7
)
