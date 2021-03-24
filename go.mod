module knative.dev/eventing-autoscaler-keda

go 1.15

require (
	github.com/kedacore/keda/v2 v2.2.0
	go.uber.org/zap v1.16.0
	k8s.io/api v0.20.4
	k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.20.4
	knative.dev/eventing v0.21.1-0.20210324094318-d3e856a2a2bd
	knative.dev/eventing-awssqs v0.21.1-0.20210316190343-db02cadaac30
	knative.dev/eventing-kafka v0.21.1-0.20210324134818-d78321896b22
	knative.dev/eventing-redis v0.21.1-0.20210322200754-5a1e2c4a8fed
	knative.dev/hack v0.0.0-20210317214554-58edbdc42966
	knative.dev/pkg v0.0.0-20210323202917-b558677ab034
)

replace (
	k8s.io/api => k8s.io/api v0.19.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.7
	k8s.io/apiserver => k8s.io/apiserver v0.19.7
	k8s.io/client-go => k8s.io/client-go v0.19.7
	k8s.io/code-generator => k8s.io/code-generator v0.19.7
)
