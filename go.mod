module knative.dev/eventing-autoscaler-keda

go 1.14

require (
	github.com/google/licenseclassifier v0.0.0-20200708223521-3d09a0ea2f39
	github.com/kedacore/keda v1.5.1-0.20201012074936-071d5a0732c5
	go.uber.org/zap v1.15.0
	k8s.io/api v0.18.8
	k8s.io/apiextensions-apiserver v0.18.8
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/code-generator v0.18.8
	k8s.io/kube-openapi v0.0.0-20200805222855-6aeccd4b50c6
	knative.dev/eventing v0.18.1-0.20201027155533-17e1562518ef
	knative.dev/eventing-awssqs v0.0.0-20201021084318-8fd20cd38fdd
	knative.dev/eventing-kafka v0.0.0-20201027135533-5e99f5d2f7b4
	knative.dev/pkg v0.0.0-20201027121533-273ba59a1132
	knative.dev/test-infra v0.0.0-20201026182042-46291de4ab66
)

replace (
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6
)
