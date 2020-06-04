// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the DeleteRegistrationCode operation.
type DeleteRegistrationCodeInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRegistrationCodeInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRegistrationCodeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	return nil
}

// The output for the DeleteRegistrationCode operation.
type DeleteRegistrationCodeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRegistrationCodeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRegistrationCodeOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteRegistrationCode = "DeleteRegistrationCode"

// DeleteRegistrationCodeRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes a CA certificate registration code.
//
//    // Example sending a request using DeleteRegistrationCodeRequest.
//    req := client.DeleteRegistrationCodeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteRegistrationCodeRequest(input *DeleteRegistrationCodeInput) DeleteRegistrationCodeRequest {
	op := &aws.Operation{
		Name:       opDeleteRegistrationCode,
		HTTPMethod: "DELETE",
		HTTPPath:   "/registrationcode",
	}

	if input == nil {
		input = &DeleteRegistrationCodeInput{}
	}

	req := c.newRequest(op, input, &DeleteRegistrationCodeOutput{})

	return DeleteRegistrationCodeRequest{Request: req, Input: input, Copy: c.DeleteRegistrationCodeRequest}
}

// DeleteRegistrationCodeRequest is the request type for the
// DeleteRegistrationCode API operation.
type DeleteRegistrationCodeRequest struct {
	*aws.Request
	Input *DeleteRegistrationCodeInput
	Copy  func(*DeleteRegistrationCodeInput) DeleteRegistrationCodeRequest
}

// Send marshals and sends the DeleteRegistrationCode API request.
func (r DeleteRegistrationCodeRequest) Send(ctx context.Context) (*DeleteRegistrationCodeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRegistrationCodeResponse{
		DeleteRegistrationCodeOutput: r.Request.Data.(*DeleteRegistrationCodeOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRegistrationCodeResponse is the response type for the
// DeleteRegistrationCode API operation.
type DeleteRegistrationCodeResponse struct {
	*DeleteRegistrationCodeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRegistrationCode request.
func (r *DeleteRegistrationCodeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}