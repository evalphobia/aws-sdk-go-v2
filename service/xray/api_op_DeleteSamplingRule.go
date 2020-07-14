// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package xray

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteSamplingRuleInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the sampling rule. Specify a rule by either name or ARN, but not
	// both.
	RuleARN *string `type:"string"`

	// The name of the sampling rule. Specify a rule by either name or ARN, but
	// not both.
	RuleName *string `type:"string"`
}

// String returns the string representation
func (s DeleteSamplingRuleInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSamplingRuleInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RuleARN != nil {
		v := *s.RuleARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RuleARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RuleName != nil {
		v := *s.RuleName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RuleName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteSamplingRuleOutput struct {
	_ struct{} `type:"structure"`

	// The deleted rule definition and metadata.
	SamplingRuleRecord *SamplingRuleRecord `type:"structure"`
}

// String returns the string representation
func (s DeleteSamplingRuleOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSamplingRuleOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SamplingRuleRecord != nil {
		v := s.SamplingRuleRecord

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SamplingRuleRecord", v, metadata)
	}
	return nil
}

const opDeleteSamplingRule = "DeleteSamplingRule"

// DeleteSamplingRuleRequest returns a request value for making API operation for
// AWS X-Ray.
//
// Deletes a sampling rule.
//
//    // Example sending a request using DeleteSamplingRuleRequest.
//    req := client.DeleteSamplingRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/xray-2016-04-12/DeleteSamplingRule
func (c *Client) DeleteSamplingRuleRequest(input *DeleteSamplingRuleInput) DeleteSamplingRuleRequest {
	op := &aws.Operation{
		Name:       opDeleteSamplingRule,
		HTTPMethod: "POST",
		HTTPPath:   "/DeleteSamplingRule",
	}

	if input == nil {
		input = &DeleteSamplingRuleInput{}
	}

	req := c.newRequest(op, input, &DeleteSamplingRuleOutput{})

	return DeleteSamplingRuleRequest{Request: req, Input: input, Copy: c.DeleteSamplingRuleRequest}
}

// DeleteSamplingRuleRequest is the request type for the
// DeleteSamplingRule API operation.
type DeleteSamplingRuleRequest struct {
	*aws.Request
	Input *DeleteSamplingRuleInput
	Copy  func(*DeleteSamplingRuleInput) DeleteSamplingRuleRequest
}

// Send marshals and sends the DeleteSamplingRule API request.
func (r DeleteSamplingRuleRequest) Send(ctx context.Context) (*DeleteSamplingRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSamplingRuleResponse{
		DeleteSamplingRuleOutput: r.Request.Data.(*DeleteSamplingRuleOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSamplingRuleResponse is the response type for the
// DeleteSamplingRule API operation.
type DeleteSamplingRuleResponse struct {
	*DeleteSamplingRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSamplingRule request.
func (r *DeleteSamplingRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}