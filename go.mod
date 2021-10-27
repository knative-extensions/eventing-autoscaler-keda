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
	knative.dev/eventing v0.26.1-0.20211027135601-c804fd2b94fb
	knative.dev/eventing-awssqs v0.26.1-0.20211019131536-65d59f0f3c54
	knative.dev/eventing-kafka v0.26.1-0.20211027132002-a7255df49f63
	knative.dev/eventing-redis v0.26.1-0.20211027133002-72b89ceebe7c
	knative.dev/hack v0.0.0-20211026141922-a71c865b5f66
	knative.dev/pkg v0.0.0-20211027105800-3b33e02e5b9c
	knative.dev/reconciler-test v0.0.0-20211019133535-04ea7ddac7e3
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
