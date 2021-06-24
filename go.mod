module knative.dev/eventing-autoscaler-keda

go 1.16

require (
	github.com/kedacore/keda/v2 v2.3.0
	go.uber.org/zap v1.17.0
	k8s.io/api v0.20.7
	k8s.io/apiextensions-apiserver v0.20.7
	k8s.io/apimachinery v0.20.7
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.20.7
	knative.dev/eventing v0.23.1-0.20210624052545-d69382b45e27
	knative.dev/eventing-awssqs v0.23.1-0.20210623025002-3ca4a9c94f80
	knative.dev/eventing-kafka v0.23.1-0.20210624052644-542772d58179
	knative.dev/eventing-redis v0.23.1-0.20210623024946-788c74d39e1f
	knative.dev/hack v0.0.0-20210622141627-e28525d8d260
	knative.dev/pkg v0.0.0-20210622173328-dd0db4b05c80
	knative.dev/reconciler-test v0.0.0-20210623134345-88c84739abd9
)

replace (
	k8s.io/api => k8s.io/api v0.19.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.7
	k8s.io/apiserver => k8s.io/apiserver v0.19.7
	k8s.io/client-go => k8s.io/client-go v0.19.7
	k8s.io/code-generator => k8s.io/code-generator v0.19.7
)
