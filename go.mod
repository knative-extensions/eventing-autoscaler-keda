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
	knative.dev/eventing v0.30.1-0.20220316055858-c02e63aea351
	knative.dev/eventing-kafka v0.30.1-0.20220314063517-bc254c796e2f
	knative.dev/eventing-redis v0.30.2-0.20220314154919-b09f9af34286
	knative.dev/hack v0.0.0-20220314052818-c9c3ea17a2e9
	knative.dev/pkg v0.0.0-20220316002959-3a4cc56708b9
	knative.dev/reconciler-test v0.0.0-20220314160418-3b7a0d7f7b4b
	sigs.k8s.io/controller-runtime v0.10.3
)
