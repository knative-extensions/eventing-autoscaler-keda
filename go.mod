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
	knative.dev/eventing v0.20.1-0.20210215082944-d8468ca728b3
	knative.dev/eventing-awssqs v0.20.1-0.20210215083656-ab8b4f6eee15
	knative.dev/eventing-kafka v0.19.1-0.20210215131511-9022563cdde8
	knative.dev/eventing-redis v0.20.1-0.20210215083356-932221df6e85
	knative.dev/hack v0.0.0-20210203173706-8368e1f6eacf
	knative.dev/pkg v0.0.0-20210212203835-448ae657fb5f
)

replace (
	k8s.io/api => k8s.io/api v0.19.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.7
	k8s.io/apiserver => k8s.io/apiserver v0.19.7
	k8s.io/client-go => k8s.io/client-go v0.19.7
	k8s.io/code-generator => k8s.io/code-generator v0.19.7
)
