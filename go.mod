module knative.dev/eventing-autoscaler-keda

go 1.14

require (
	github.com/google/licenseclassifier v0.0.0-20200708223521-3d09a0ea2f39
	github.com/kedacore/keda/v2 v2.0.1-0.20201118092520-5c1257d8c726
	go.uber.org/zap v1.16.0
	k8s.io/api v0.18.12
	k8s.io/apiextensions-apiserver v0.18.12
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/code-generator v0.18.12
	k8s.io/kube-openapi v0.0.0-20200805222855-6aeccd4b50c6
	knative.dev/eventing v0.19.1-0.20201127155535-ad755bdfccc6
	knative.dev/eventing-awssqs v0.19.1-0.20201123083558-14e79dc9ff5b
	knative.dev/eventing-kafka v0.19.1-0.20201127112435-e00890307a57
	knative.dev/eventing-redis v0.19.1-0.20201127145036-2c4f1b10dfcf
	knative.dev/hack v0.0.0-20201125230335-c46a6498e9ed
	knative.dev/pkg v0.0.0-20201127013335-0d896b5c87b8
)

replace (
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6
)
