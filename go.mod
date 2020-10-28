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
	knative.dev/eventing v0.18.1-0.20201028151034-7c24cda5fd71
	knative.dev/eventing-awssqs v0.0.0-20201028020434-65245b947a6a
	knative.dev/eventing-kafka v0.0.0-20201028124334-8325f1b9852d
	knative.dev/hack v0.0.0-20201027221733-0d7f2f064b7b
	knative.dev/pkg v0.0.0-20201028142834-e135a1737847
)

replace (
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6
)
