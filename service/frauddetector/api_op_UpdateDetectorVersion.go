// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateDetectorVersionInput struct {
	_ struct{} `type:"structure"`

	// The detector version description.
	Description *string `locationName:"description" min:"1" type:"string"`

	// The parent detector ID for the detector version you want to update.
	//
	// DetectorId is a required field
	DetectorId *string `locationName:"detectorId" min:"1" type:"string" required:"true"`

	// The detector version ID.
	//
	// DetectorVersionId is a required field
	DetectorVersionId *string `locationName:"detectorVersionId" min:"1" type:"string" required:"true"`

	// The Amazon SageMaker model endpoints to include in the detector version.
	//
	// ExternalModelEndpoints is a required field
	ExternalModelEndpoints []string `locationName:"externalModelEndpoints" type:"list" required:"true"`

	// The model versions to include in the detector version.
	ModelVersions []ModelVersion `locationName:"modelVersions" type:"list"`

	// The rule execution mode to add to the detector.
	//
	// If you specify FIRST_MATCHED, Amazon Fraud Detector evaluates rules sequentially,
	// first to last, stopping at the first matched rule. Amazon Fraud dectector
	// then provides the outcomes for that single rule.
	//
	// If you specifiy ALL_MATCHED, Amazon Fraud Detector evaluates all rules and
	// returns the outcomes for all matched rules. You can define and edit the rule
	// mode at the detector version level, when it is in draft status.
	//
	// The default behavior is FIRST_MATCHED.
	RuleExecutionMode RuleExecutionMode `locationName:"ruleExecutionMode" type:"string" enum:"true"`

	// The rules to include in the detector version.
	//
	// Rules is a required field
	Rules []Rule `locationName:"rules" type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateDetectorVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDetectorVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDetectorVersionInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if s.DetectorVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorVersionId"))
	}
	if s.DetectorVersionId != nil && len(*s.DetectorVersionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorVersionId", 1))
	}

	if s.ExternalModelEndpoints == nil {
		invalidParams.Add(aws.NewErrParamRequired("ExternalModelEndpoints"))
	}

	if s.Rules == nil {
		invalidParams.Add(aws.NewErrParamRequired("Rules"))
	}
	if s.ModelVersions != nil {
		for i, v := range s.ModelVersions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ModelVersions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Rules != nil {
		for i, v := range s.Rules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Rules", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateDetectorVersionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateDetectorVersionOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateDetectorVersion = "UpdateDetectorVersion"

// UpdateDetectorVersionRequest returns a request value for making API operation for
// Amazon Fraud Detector.
//
// Updates a detector version. The detector version attributes that you can
// update include models, external model endpoints, rules, and description.
// You can only update a DRAFT detector version.
//
//    // Example sending a request using UpdateDetectorVersionRequest.
//    req := client.UpdateDetectorVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/frauddetector-2019-11-15/UpdateDetectorVersion
func (c *Client) UpdateDetectorVersionRequest(input *UpdateDetectorVersionInput) UpdateDetectorVersionRequest {
	op := &aws.Operation{
		Name:       opUpdateDetectorVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateDetectorVersionInput{}
	}

	req := c.newRequest(op, input, &UpdateDetectorVersionOutput{})

	return UpdateDetectorVersionRequest{Request: req, Input: input, Copy: c.UpdateDetectorVersionRequest}
}

// UpdateDetectorVersionRequest is the request type for the
// UpdateDetectorVersion API operation.
type UpdateDetectorVersionRequest struct {
	*aws.Request
	Input *UpdateDetectorVersionInput
	Copy  func(*UpdateDetectorVersionInput) UpdateDetectorVersionRequest
}

// Send marshals and sends the UpdateDetectorVersion API request.
func (r UpdateDetectorVersionRequest) Send(ctx context.Context) (*UpdateDetectorVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDetectorVersionResponse{
		UpdateDetectorVersionOutput: r.Request.Data.(*UpdateDetectorVersionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDetectorVersionResponse is the response type for the
// UpdateDetectorVersion API operation.
type UpdateDetectorVersionResponse struct {
	*UpdateDetectorVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDetectorVersion request.
func (r *UpdateDetectorVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}