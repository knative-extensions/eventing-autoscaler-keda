module knative.dev/eventing-autoscaler-keda

go 1.16

require (
	github.com/kedacore/keda/v2 v2.4.0
	go.uber.org/zap v1.19.1
	k8s.io/api v0.21.4
	k8s.io/apiextensions-apiserver v0.21.4
	k8s.io/apimachinery v0.21.4
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.21.4
	knative.dev/eventing v0.26.1-0.20211028061727-2eb07cbddd07
	knative.dev/eventing-awssqs v0.26.1-0.20211028081826-881d593e2f36
	knative.dev/eventing-kafka v0.26.1-0.20211028065026-ee67c5013735
	knative.dev/eventing-redis v0.26.1-0.20211027133002-72b89ceebe7c
	knative.dev/hack v0.0.0-20211027200727-f1228dd5e3e6
	knative.dev/pkg v0.0.0-20211027171921-f7b70f5ce303
	knative.dev/reconciler-test v0.0.0-20211028081126-a4cd8030f736
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
