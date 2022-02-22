module knative.dev/eventing-autoscaler-keda

go 1.16

require (
	github.com/kedacore/keda/v2 v2.5.0
	go.uber.org/zap v1.19.1
	k8s.io/api v0.22.5
	k8s.io/apiextensions-apiserver v0.22.5
	k8s.io/apimachinery v0.22.5
	k8s.io/client-go v0.22.5
	k8s.io/code-generator v0.22.5
	knative.dev/eventing v0.29.1-0.20220221175003-2a69eec1e358
	knative.dev/eventing-awssqs v0.29.0
	knative.dev/eventing-kafka v0.28.1-0.20220222070004-b8ce9727a9c6
	knative.dev/eventing-redis v0.29.1-0.20220216210340-3371e1969d5a
	knative.dev/hack v0.0.0-20220218190734-a8ef7b67feec
	knative.dev/pkg v0.0.0-20220217155112-d48172451966
	knative.dev/reconciler-test v0.0.0-20220216192840-2c3291f210ce
	sigs.k8s.io/controller-runtime v0.10.3
)
