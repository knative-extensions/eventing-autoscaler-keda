module knative.dev/eventing-autoscaler-keda

go 1.14

require (
	github.com/kedacore/keda/v2 v2.0.1-0.20201118092520-5c1257d8c726
	go.uber.org/zap v1.16.0
	k8s.io/api v0.18.12
	k8s.io/apiextensions-apiserver v0.18.12
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/code-generator v0.18.12
	k8s.io/kube-openapi v0.0.0-20200805222855-6aeccd4b50c6 // indirect
	knative.dev/eventing v0.20.0
	knative.dev/eventing-awssqs v0.20.0
	knative.dev/eventing-kafka v0.20.0
	knative.dev/eventing-redis v0.20.0
	knative.dev/hack v0.0.0-20210108203236-ea9c9a0cac5c
	knative.dev/pkg v0.0.0-20210107022335-51c72e24c179
)

replace (
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6
)
