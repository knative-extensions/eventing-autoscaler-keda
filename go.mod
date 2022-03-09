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
	knative.dev/eventing v0.30.0
	knative.dev/eventing-awssqs v0.29.1-0.20220302165045-502b3e543cc6
	knative.dev/eventing-kafka v0.28.1-0.20220308215949-a262ed153f01
	knative.dev/eventing-redis v0.30.0
	knative.dev/hack v0.0.0-20220224013837-e1785985d364
	knative.dev/pkg v0.0.0-20220301181942-2fdd5f232e77
	knative.dev/reconciler-test v0.0.0-20220303141206-84821d26ed1f
	sigs.k8s.io/controller-runtime v0.10.3
)
