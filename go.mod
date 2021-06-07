module knative.dev/eventing-autoscaler-keda

go 1.16

require (
	github.com/kedacore/keda/v2 v2.3.0
	go.uber.org/zap v1.17.0
	k8s.io/api v0.20.7
	k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery v0.20.7
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.20.7
	knative.dev/eventing v0.23.1-0.20210604160145-ab3978c3656d
	knative.dev/eventing-awssqs v0.23.1-0.20210602221044-06d6790e808a
	knative.dev/eventing-kafka v0.23.1-0.20210604150745-b51681072a4c
	knative.dev/eventing-redis v0.23.1-0.20210604203745-4ad1af335f12
	knative.dev/hack v0.0.0-20210601210329-de04b70e00d0
	knative.dev/pkg v0.0.0-20210602095030-0e61d6763dd6
	knative.dev/reconciler-test v0.0.0-20210603210445-0071c48281c7
)

replace (
	k8s.io/api => k8s.io/api v0.19.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.7
	k8s.io/apiserver => k8s.io/apiserver v0.19.7
	k8s.io/client-go => k8s.io/client-go v0.19.7
	k8s.io/code-generator => k8s.io/code-generator v0.19.7
)
