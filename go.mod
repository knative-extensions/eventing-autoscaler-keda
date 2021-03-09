module knative.dev/eventing-autoscaler-keda

go 1.15

require (
	github.com/kedacore/keda/v2 v2.1.0
	go.uber.org/zap v1.16.0
	k8s.io/api v0.20.2
	k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.20.0
	knative.dev/eventing v0.21.1-0.20210306204021-6136d09fb405
	knative.dev/eventing-awssqs v0.21.1-0.20210304153115-152a90f5a00a
	knative.dev/eventing-kafka v0.21.1-0.20210308090421-c1e9dbb621e3
	knative.dev/eventing-redis v0.21.1-0.20210304151515-d2dcaa9a496f
	knative.dev/hack v0.0.0-20210305150220-f99a25560134
	knative.dev/pkg v0.0.0-20210308052421-737401c38b22
)

replace (
	k8s.io/api => k8s.io/api v0.19.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.7
	k8s.io/apiserver => k8s.io/apiserver v0.19.7
	k8s.io/client-go => k8s.io/client-go v0.19.7
	k8s.io/code-generator => k8s.io/code-generator v0.19.7
)
