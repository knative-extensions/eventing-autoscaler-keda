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
	knative.dev/eventing v0.22.1-0.20210423044837-a0a33025aee0
	knative.dev/eventing-awssqs v0.22.1-0.20210420222937-9c231d2460b6
	knative.dev/eventing-kafka v0.22.1-0.20210423113238-bd02a91f3250
	knative.dev/eventing-redis v0.22.1-0.20210423113539-7357f279f88e
	knative.dev/hack v0.0.0-20210423193138-b5f6e2587f6d
	knative.dev/pkg v0.0.0-20210423162638-78b8140ed19c
	knative.dev/reconciler-test v0.0.0-20210423124638-9a9c4dd3af2b
)

replace (
	k8s.io/api => k8s.io/api v0.19.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.7
	k8s.io/apiserver => k8s.io/apiserver v0.19.7
	k8s.io/client-go => k8s.io/client-go v0.19.7
	k8s.io/code-generator => k8s.io/code-generator v0.19.7
)
