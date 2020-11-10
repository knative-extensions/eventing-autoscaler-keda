module knative.dev/eventing-autoscaler-keda

go 1.14

require (
	github.com/google/licenseclassifier v0.0.0-20200708223521-3d09a0ea2f39
	github.com/kedacore/keda v1.5.1-0.20201012074936-071d5a0732c5
	go.uber.org/zap v1.16.0
	k8s.io/api v0.18.8
	k8s.io/apiextensions-apiserver v0.18.8
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/code-generator v0.18.8
	k8s.io/kube-openapi v0.0.0-20200805222855-6aeccd4b50c6
	knative.dev/eventing v0.18.1-0.20201105172507-a6fc5408cb44
	knative.dev/eventing-awssqs v0.0.0-20201030071135-487e66030fae
	knative.dev/eventing-kafka v0.0.0-20201103101804-55d0e83c4bff
	knative.dev/eventing-redis v0.0.0-20201109225209-c82e1c132654
	knative.dev/hack v0.0.0-20201103151104-3d5abc3a0075
	knative.dev/pkg v0.0.0-20201103163404-5514ab0c1fdf
)

replace (
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6
)
