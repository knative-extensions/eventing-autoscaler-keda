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
	knative.dev/eventing v0.29.1-0.20220203222422-fb092415eb36
	knative.dev/eventing-awssqs v0.29.0
	knative.dev/eventing-kafka v0.28.1-0.20220203051220-507931bbaee0
	knative.dev/eventing-redis v0.29.1-0.20220203210023-29146ce3d96f
	knative.dev/hack v0.0.0-20220203160821-9b303d690fc9
	knative.dev/pkg v0.0.0-20220203020920-51be315ed160
	knative.dev/reconciler-test v0.0.0-20220202155955-6e47083645cf
	sigs.k8s.io/controller-runtime v0.10.3
)
