// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateMitigationActionInput struct {
	_ struct{} `type:"structure"`

	// A friendly name for the action. Choose a friendly name that accurately describes
	// the action (for example, EnableLoggingAction).
	//
	// ActionName is a required field
	ActionName *string `location:"uri" locationName:"actionName" type:"string" required:"true"`

	// Defines the type of action and the parameters for that action.
	//
	// ActionParams is a required field
	ActionParams *MitigationActionParams `locationName:"actionParams" type:"structure" required:"true"`

	// The ARN of the IAM role that is used to apply the mitigation action.
	//
	// RoleArn is a required field
	RoleArn *string `locationName:"roleArn" min:"20" type:"string" required:"true"`

	// Metadata that can be used to manage the mitigation action.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreateMitigationActionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMitigationActionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMitigationActionInput"}

	if s.ActionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActionName"))
	}

	if s.ActionParams == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActionParams"))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 20))
	}
	if s.ActionParams != nil {
		if err := s.ActionParams.Validate(); err != nil {
			invalidParams.AddNested("ActionParams", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMitigationActionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ActionParams != nil {
		v := s.ActionParams

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "actionParams", v, metadata)
	}
	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ActionName != nil {
		v := *s.ActionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "actionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateMitigationActionOutput struct {
	_ struct{} `type:"structure"`

	// The ARN for the new mitigation action.
	ActionArn *string `locationName:"actionArn" type:"string"`

	// A unique identifier for the new mitigation action.
	ActionId *string `locationName:"actionId" type:"string"`
}

// String returns the string representation
func (s CreateMitigationActionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMitigationActionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ActionArn != nil {
		v := *s.ActionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "actionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ActionId != nil {
		v := *s.ActionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "actionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateMitigationAction = "CreateMitigationAction"

// CreateMitigationActionRequest returns a request value for making API operation for
// AWS IoT.
//
// Defines an action that can be applied to audit findings by using StartAuditMitigationActionsTask.
// Each mitigation action can apply only one type of change.
//
//    // Example sending a request using CreateMitigationActionRequest.
//    req := client.CreateMitigationActionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateMitigationActionRequest(input *CreateMitigationActionInput) CreateMitigationActionRequest {
	op := &aws.Operation{
		Name:       opCreateMitigationAction,
		HTTPMethod: "POST",
		HTTPPath:   "/mitigationactions/actions/{actionName}",
	}

	if input == nil {
		input = &CreateMitigationActionInput{}
	}

	req := c.newRequest(op, input, &CreateMitigationActionOutput{})

	return CreateMitigationActionRequest{Request: req, Input: input, Copy: c.CreateMitigationActionRequest}
}

// CreateMitigationActionRequest is the request type for the
// CreateMitigationAction API operation.
type CreateMitigationActionRequest struct {
	*aws.Request
	Input *CreateMitigationActionInput
	Copy  func(*CreateMitigationActionInput) CreateMitigationActionRequest
}

// Send marshals and sends the CreateMitigationAction API request.
func (r CreateMitigationActionRequest) Send(ctx context.Context) (*CreateMitigationActionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMitigationActionResponse{
		CreateMitigationActionOutput: r.Request.Data.(*CreateMitigationActionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMitigationActionResponse is the response type for the
// CreateMitigationAction API operation.
type CreateMitigationActionResponse struct {
	*CreateMitigationActionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMitigationAction request.
func (r *CreateMitigationActionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}