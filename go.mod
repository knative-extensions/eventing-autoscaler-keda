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
	knative.dev/eventing v0.29.1-0.20220228193110-f2045b0d65b7
	knative.dev/eventing-awssqs v0.29.1-0.20220228154610-ed0d1a15cf97
	knative.dev/eventing-kafka v0.28.1-0.20220301081542-37e8f99c63ac
	knative.dev/eventing-redis v0.29.1-0.20220228154712-9e4b67639ed4
	knative.dev/hack v0.0.0-20220224013837-e1785985d364
	knative.dev/pkg v0.0.0-20220228195509-fe264173447b
	knative.dev/reconciler-test v0.0.0-20220216192840-2c3291f210ce
	sigs.k8s.io/controller-runtime v0.10.3
)
