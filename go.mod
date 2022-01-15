module knative.dev/eventing-autoscaler-keda

go 1.16

require (
	github.com/kedacore/keda/v2 v2.5.0
	go.uber.org/zap v1.19.1
	k8s.io/api v0.23.1
	k8s.io/apiextensions-apiserver v0.23.1
	k8s.io/apimachinery v0.23.1
	k8s.io/client-go v0.23.1
	k8s.io/code-generator v0.23.1
	knative.dev/eventing v0.28.1-0.20220107145225-eb4c06c8009d
	knative.dev/eventing-awssqs v0.28.1-0.20220105135633-cb44f263b413
	knative.dev/eventing-kafka v0.28.1-0.20220110121359-948216c3d32c
	knative.dev/eventing-redis v0.28.1-0.20220106062301-e8b78745d98e
	knative.dev/hack v0.0.0-20220111151514-59b0cf17578e
	knative.dev/pkg v0.0.0-20220114141842-0a429cba1c73
	knative.dev/reconciler-test v0.0.0-20211222120418-816f2192fec9
	sigs.k8s.io/controller-runtime v0.10.3
)
