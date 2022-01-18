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
	knative.dev/eventing v0.28.1-0.20220118155632-76366bce7449
	knative.dev/eventing-awssqs v0.28.1-0.20220105135633-cb44f263b413
	knative.dev/eventing-kafka v0.28.1-0.20220117141029-46249ce2a5f4
	knative.dev/eventing-redis v0.28.1-0.20220113062511-8a5ddfc4e7bb
	knative.dev/hack v0.0.0-20220118141833-9b2ed8471e30
	knative.dev/pkg v0.0.0-20220118151132-768f44f3fce2
	knative.dev/reconciler-test v0.0.0-20220117082429-6a9b91eef10c
	sigs.k8s.io/controller-runtime v0.10.3
)
