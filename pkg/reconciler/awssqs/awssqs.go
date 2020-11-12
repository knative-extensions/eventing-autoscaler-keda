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

package awssqs

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	kedav1alpha1 "github.com/kedacore/keda/api/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"

	"knative.dev/eventing-autoscaler-keda/pkg/reconciler/keda"
	awssqsv1alpha1 "knative.dev/eventing-awssqs/pkg/apis/sources/v1alpha1"
)

const (
	defaultAwsSqsQueueLength = 5
)

func GenerateScaleTarget(ctx context.Context, kubeClient kubernetes.Interface, src *awssqsv1alpha1.AwsSqsSource) (*kedav1alpha1.ScaleTarget, error) {
	name, err := GenerateScaleTargetName(ctx, kubeClient, src)
	if err != nil {
		return nil, err
	}
	return &kedav1alpha1.ScaleTarget{Name: name}, nil
}

func GenerateScaleTargetName(ctx context.Context, kubeClient kubernetes.Interface, src *awssqsv1alpha1.AwsSqsSource) (string, error) {
	dl, err := kubeClient.AppsV1().Deployments(src.Namespace).List(ctx,
		metav1.ListOptions{
			LabelSelector: getLabelSelector(src).String(),
		})

	if err != nil {
		return "", err
	}
	for _, dep := range dl.Items {
		if metav1.IsControlledBy(&dep, src) {
			return dep.Name, nil
		}
	}
	return "", apierrors.NewNotFound(schema.GroupResource{}, "")
}

func GenerateScaleTriggers(src *awssqsv1alpha1.AwsSqsSource) ([]kedav1alpha1.ScaleTriggers, error) {
	queueLength, err := keda.GetInt32ValueFromMap(src.Annotations, keda.KedaAutoscalingAwsSqsQueueLength, defaultAwsSqsQueueLength)
	if err != nil {
		return nil, err
	}
	awsRegion, err := getRegion(src.Spec.QueueURL)
	if err != nil {
		return nil, err
	}

	triggerMetadata := map[string]string{
		"queueURL":    src.Spec.QueueURL,
		"queueLength": strconv.Itoa(int(*queueLength)),
		"awsRegion":   awsRegion,
	}

	return []kedav1alpha1.ScaleTriggers{
		{
			Type:     "aws-sqs-queue",
			Metadata: triggerMetadata,
		},
	}, nil
}

func getLabelSelector(src *awssqsv1alpha1.AwsSqsSource) labels.Selector {
	return labels.SelectorFromSet(getLabels(src))
}

func getLabels(src *awssqsv1alpha1.AwsSqsSource) map[string]string {
	return map[string]string{
		"knative-eventing-source":      "awssqssource.sources.eventing.knative.dev",
		"knative-eventing-source-name": src.Name,
	}
}

// getRegion takes an AWS SQS URL and extracts the region from it
// e.g. URLs have this form:
// https://sqs.<region>.amazonaws.com/<account_id>/<queue_name>
// See
// https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-general-identifiers.html
// for reference.  Note that AWS does not make any promises re. url
// structure although it feels reasonable to rely on it at this point
// rather than add an additional `region` parameter to the spec that
// will now be redundant most of the time.
func getRegion(url string) (string, error) {
	parts := strings.Split(url, ".")

	if len(parts) < 2 {
		err := fmt.Errorf("QueueURL does not look correct: %s", url)
		return "", err
	}
	return parts[1], nil
}
