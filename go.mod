module github.com/zroubalik/autoscaler-keda

go 1.14

require (
	github.com/google/licenseclassifier v0.0.0-20200708223521-3d09a0ea2f39
	github.com/kedacore/keda v1.4.2-0.20200617120630-97df7e08e24b
	go.uber.org/zap v1.15.0
	k8s.io/api v0.18.6
	k8s.io/apiextensions-apiserver v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/code-generator v0.18.6
	k8s.io/kube-openapi v0.0.0-20200410145947-bcb3869e6f29
	knative.dev/eventing v0.16.1
	knative.dev/eventing-contrib v0.16.0
	knative.dev/pkg v0.0.0-20200702222342-ea4d6e985ba0
	knative.dev/test-infra v0.0.0-20200713185018-6b52776d44a4
)

replace knative.dev/eventing => github.com/zroubalik/eventing v0.15.1-0.20200714134632-e7996ac5d4c3

replace knative.dev/pkg => github.com/zroubalik/pkg v0.0.0-20200714090639-88ee0a9b8a22

replace github.com/kedacore/keda => github.com/kedacore/keda v1.4.2-0.20200617120630-97df7e08e24b

replace k8s.io/client-go => k8s.io/client-go v0.18.6
