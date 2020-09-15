// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package platform_application

import (
	"context"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/sns"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws/aws-controllers-k8s/services/sns/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = &svcsdk.SNS{}
	_ = &svcapitypes.PlatformApplication{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newGetAttributesRequestPayload(r)
	if err != nil {
		return nil, err
	}

	_, respErr := rm.sdkapi.GetPlatformApplicationAttributesWithContext(ctx, input)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "NotFound" {
			return nil, ackerr.NotFound
		}
		return nil, respErr
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}

	return &resource{ko}, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.ListPlatformApplicationsInput, error) {
	res := &svcsdk.ListPlatformApplicationsInput{}

	return res, nil
}

// newGetAttributesRequestPayload returns SDK-specific struct for the HTTP
// request payload of the GetAttributes API call for the resource
func (rm *resourceManager) newGetAttributesRequestPayload(
	r *resource,
) (*svcsdk.GetPlatformApplicationAttributesInput, error) {
	res := &svcsdk.GetPlatformApplicationAttributesInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetPlatformApplicationArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	} else {
		res.SetPlatformApplicationArn(rm.ARNFromName(*r.ko.Spec.Name))
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	_, respErr := rm.sdkapi.CreatePlatformApplicationWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{OwnerAccountID: &rm.awsAccountID}
	ko.Status.Conditions = []*ackv1alpha1.Condition{}
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreatePlatformApplicationInput, error) {
	res := &svcsdk.CreatePlatformApplicationInput{}

	attrMap := map[string]*string{}
	if r.ko.Spec.EventDeliveryFailure != nil {
		attrMap["EventDeliveryFailure"] = r.ko.Spec.EventDeliveryFailure
	}
	if r.ko.Spec.EventEndpointCreated != nil {
		attrMap["EventEndpointCreated"] = r.ko.Spec.EventEndpointCreated
	}
	if r.ko.Spec.EventEndpointDeleted != nil {
		attrMap["EventEndpointDeleted"] = r.ko.Spec.EventEndpointDeleted
	}
	if r.ko.Spec.EventEndpointUpdated != nil {
		attrMap["EventEndpointUpdated"] = r.ko.Spec.EventEndpointUpdated
	}
	if r.ko.Spec.FailureFeedbackRoleARN != nil {
		attrMap["FailureFeedbackRoleArn"] = r.ko.Spec.FailureFeedbackRoleARN
	}
	if r.ko.Spec.PlatformCredential != nil {
		attrMap["PlatformCredential"] = r.ko.Spec.PlatformCredential
	}
	if r.ko.Spec.PlatformPrincipal != nil {
		attrMap["PlatformPrincipal"] = r.ko.Spec.PlatformPrincipal
	}
	if r.ko.Spec.SuccessFeedbackRoleARN != nil {
		attrMap["SuccessFeedbackRoleArn"] = r.ko.Spec.SuccessFeedbackRoleARN
	}
	if r.ko.Spec.SuccessFeedbackSampleRate != nil {
		attrMap["SuccessFeedbackSampleRate"] = r.ko.Spec.SuccessFeedbackSampleRate
	}
	res.SetAttributes(attrMap)
	if r.ko.Spec.Name != nil {
		res.SetName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.Platform != nil {
		res.SetPlatform(*r.ko.Spec.Platform)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	r *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeletePlatformApplicationWithContext(ctx, input)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeletePlatformApplicationInput, error) {
	res := &svcsdk.DeletePlatformApplicationInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetPlatformApplicationArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	} else {
		res.SetPlatformApplicationArn(rm.ARNFromName(*r.ko.Spec.Name))
	}

	return res, nil
}
