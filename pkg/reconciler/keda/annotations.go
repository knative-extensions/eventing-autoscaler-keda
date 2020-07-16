/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package keda

const (
	// Autoscaling refers to the autoscaling group.
	Autoscaling = "autoscaling.knative.dev"

	// AutoscalingClassAnnotation is the annotation for the explicit class of
	// scaler that a particular resource has opted into.
	AutoscalingClassAnnotation = Autoscaling + "/class"

	// AutoscalingMinScaleAnnotation is the annotation to specify the minimum number of pods to scale to.
	AutoscalingMinScaleAnnotation = Autoscaling + "/minScale"
	// AutoscalingMaxScaleAnnotation is the annotation to specify the maximum number of pods to scale to.
	AutoscalingMaxScaleAnnotation = Autoscaling + "/maxScale"

	// KEDA is Keda autoscaler.
	KEDA = "keda.autoscaling.knative.dev"

	// KedaAutoscalingPollingIntervalAnnotation is the annotation that refers to the interval in seconds Keda
	// uses to poll metrics in order to inform its scaling decisions.
	KedaAutoscalingPollingIntervalAnnotation = KEDA + "/pollingInterval"
	// KedaAutoscalingCooldownPeriodAnnotation is the annotation that refers to the period Keda waits until it
	// scales a Deployment down.
	KedaAutoscalingCooldownPeriodAnnotation = KEDA + "/cooldownPeriod"

	// KedaAutoscalingKafkaLagThreshold is the annotation that refers to the stream is lagging on the current consumer group
	KedaAutoscalingKafkaLagThreshold = KEDA + "/kafkaLagThreshold"

	// KedaAutoscalingAwsSqsQueueLength is the annotation that refers to the target value for ApproximateNumberOfMessages in the SQS Queue
	KedaAutoscalingAwsSqsQueueLength = KEDA + "/awsSqsQueueLength"
)
