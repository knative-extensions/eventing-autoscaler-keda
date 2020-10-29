# Experimental KEDA support for Knative Event Sources Autoscaling

[![Build status](https://github.com/knative-sandbox/eventing-autoscaler-keda//workflows/master%20build/badge.svg)](https://github.com/knative-sandbox/eventing-autoscaler-keda/actions)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0)

|                   |                                                                                                                        |
| ----------------- | ---------------------------------------------------------------------------------------------------------------------- |
| **STATUS**        | Experimental                                                                                                           |
| **Sponsoring WG** | [Eventing Sources](https://github.com/knative/community/blob/master/working-groups/WORKING-GROUPS.md#eventing-sources) |

> Warning: Still under development. Not meant for production deployment.

## Design

To enable KEDA Autoscaling of Knative Event Sources (and other components in the
future) there is a separate controller implemented, ie. no hard dependency in
Knative. This contoller si watching for `CustomResourcesDefinitions` resources
in the cluster, if there is installed a new CRD which is supported by this
controller a new dynamic controller watching these resources is created.

Currently there is support for **Kafka Source** and **AWS SQS Source**. We also
have experimental support for **RabbitMQ Broker**.

## Annotations

User can enable and configure autoscaling on a particular Source or Broker by a
set of annotations.

```yaml
metadata:
  annotations:
    autoscaling.knative.dev/class: keda.autoscaling.knative.dev
    autoscaling.knative.dev/minScale: "0"
    autoscaling.knative.dev/maxScale: "5"
    keda.autoscaling.knative.dev/pollingInterval: "30"
    keda.autoscaling.knative.dev/cooldownPeriod: "30"

    # Kafka Source
    keda.autoscaling.knative.dev/kafkaLagThreshold: "10"

    # AWS SQS Source
    keda.autoscaling.knative.dev/awsSqsQueueLength: "5"
```

- `autoscaling.knative.dev/class: keda.autoscaling.knative.dev` - needs to be
  specified on a Source to enable KEDA autoscaling
- `autoscaling.knative.dev/minScale` - minimum number of replicas to scale down
  to. Default: `0`
- `autoscaling.knative.dev/maxScale` - maximum number of replicas to scale out
  to. Default: `50`
- `keda.autoscaling.knative.dev/pollingInterval` - interval in seconds KEDA uses
  to poll metrics. Default: `30`
- `keda.autoscaling.knative.dev/cooldownPeriod` - period of time in seconds KEDA
  waits until it scales down. Default: `300`
- `keda.autoscaling.knative.dev/kafkaLagThreshold` - only for Kafka Source,
  refers to the stream is lagging on the current consumer group. Default: `10`
- `keda.autoscaling.knative.dev/awsSqsQueueLength` - only for AWS SQS Source,
  refers to the target value for ApproximateNumberOfMessages in the SQS Queue.
  Default: `5`
- `keda.autoscaling.knative.dev/rabbitMQQueueLength` - only for AWS SQS Source,
  refers to the target value for number of messages in a RabbitMQ brokers
  trigger queue: `1`

## HOW TO

### Install KEDA v2

It is needed to install KEDA v2, which is using different namespace for it's
CRDs (`keda.k8s.io` -> `keda.sh`).

Currently there is development (Alpha) version of KEDA v2, to install it follow
instructions on:
https://github.com/kedacore/keda#how-can-i-try-keda-v2-beta-version

Confirm there are 2 pods running in `keda` namespace:

```bash
$ kubectl get pods -n keda
NAME                                      READY   STATUS    RESTARTS   AGE
keda-metrics-apiserver-7cf7765dc8-k9lnc   1/1     Running   0          5m2s
keda-operator-55658855fc-rc9rb            1/1     Running   0          5m3s
```

### Install `eventing-autoscaler-keda` Controller

```bash
export KO_DOCKER_REPO=...
ko apply -f /config
```

Confirm there is 1 pod running in `eventing-autoscaler-keda` namespace:

```bash
$ kubectl get pods -n eventing-autoscaler-keda
NAME                          READY   STATUS    RESTARTS   AGE
controller-76fb8d6756-5f4vm   1/1     Running   0          21m
```

## Example of Kafka Source autoscaled by KEDA

1. Set up Kafka Cluster, eg. use [Strimzi operator](https://strimzi.io/)

2. Install Knative Serving and Eventing

3. Install Knative Eventing
   [Kafka Source](https://github.com/knative/eventing-contrib/tree/master/kafka/source)

4. Create `KafkaSource` resource, with annotation
   `autoscaling.knative.dev/class: keda.autoscaling.knative.dev`. There are
   other KEDA related annotations, see the example:

```yaml
apiVersion: sources.knative.dev/v1alpha1
kind: KafkaSource
metadata:
  name: kafka-source
  namespace: default
  annotations:
    autoscaling.knative.dev/class: keda.autoscaling.knative.dev
    autoscaling.knative.dev/minScale: "0"
    autoscaling.knative.dev/maxScale: "5"
    keda.autoscaling.knative.dev/pollingInterval: "30"
    keda.autoscaling.knative.dev/cooldownPeriod: "30"
    keda.autoscaling.knative.dev/kafkaLagThreshold: "10"
spec:
  consumerGroup: knative-group
  bootstrapServers:
    - my-cluster-kafka-bootstrap.openshift-operators:9092
  topics:
    - test
  sink:
    ref:
      apiVersion: serving.knative.dev/v1
      kind: Service
      name: event-display
```

5. Check that `ScaledObject` was created for this `KafkaSource`:

```bash
$ kubectl get scaledobjects
NAME                                      SCALETARGETKIND      SCALETARGETNAME                                                 TRIGGERS   AUTHENTICATION   READY   ACTIVE   AGE
so-f87369e5-c320-4f44-b23a-8c535a523e3a   apps/v1.Deployment   kafkasource-kafka-source-f87369e5-c320-4f44-b23a-8c535a523e3a   kafka                       True    False     6m5s
```

## Example of RabbitMQ Broker autoscaled by KEDA

1. Install Knative Serving and Eventing

2. Install
   [RabbitMQ Broker](https://github.com/knative-sandbox/eventing-rabbitmq/tree/master/broker)

3. Install a Broker / Trigger and sources as directed in the above guide.

4. Enable the autoscaler by applying the KEDA patch:

```shell
kubectl patch broker default --type merge --patch '{"metadata": {"annotations": {"autoscaling.knative.dev/class": "keda.autoscaling.knative.dev"}}}'

```

5. Check that the scaled resources were created and are ready

```shell
vaikas-a01:eventing-autoscaler-keda vaikas$ kubectl get triggerauthentications
NAME                   PODIDENTITY   SECRET                  ENV
default-trigger-auth                 default-broker-rabbit
vaikas-a01:eventing-autoscaler-keda vaikas$ kubectl get scaledobjects
NAME           SCALETARGETKIND      SCALETARGETNAME           TRIGGERS   AUTHENTICATION         READY   ACTIVE   AGE
ping-trigger   apps/v1.Deployment   ping-trigger-dispatcher   rabbitmq   default-trigger-auth   True    True     14m
```
