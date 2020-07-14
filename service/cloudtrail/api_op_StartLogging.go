// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudtrail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The request to CloudTrail to start logging AWS API calls for an account.
type StartLoggingInput struct {
	_ struct{} `type:"structure"`

	// Specifies the name or the CloudTrail ARN of the trail for which CloudTrail
	// logs AWS API calls. The format of a trail ARN is:
	//
	// arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StartLoggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartLoggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartLoggingInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Returns the objects or data listed below if successful. Otherwise, returns
// an error.
type StartLoggingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartLoggingOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartLogging = "StartLogging"

// StartLoggingRequest returns a request value for making API operation for
// AWS CloudTrail.
//
// Starts the recording of AWS API calls and log file delivery for a trail.
// For a trail that is enabled in all regions, this operation must be called
// from the region in which the trail was created. This operation cannot be
// called on the shadow trails (replicated trails in other regions) of a trail
// that is enabled in all regions.
//
//    // Example sending a request using StartLoggingRequest.
//    req := client.StartLoggingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudtrail-2013-11-01/StartLogging
func (c *Client) StartLoggingRequest(input *StartLoggingInput) StartLoggingRequest {
	op := &aws.Operation{
		Name:       opStartLogging,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartLoggingInput{}
	}

	req := c.newRequest(op, input, &StartLoggingOutput{})

	return StartLoggingRequest{Request: req, Input: input, Copy: c.StartLoggingRequest}
}

// StartLoggingRequest is the request type for the
// StartLogging API operation.
type StartLoggingRequest struct {
	*aws.Request
	Input *StartLoggingInput
	Copy  func(*StartLoggingInput) StartLoggingRequest
}

// Send marshals and sends the StartLogging API request.
func (r StartLoggingRequest) Send(ctx context.Context) (*StartLoggingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartLoggingResponse{
		StartLoggingOutput: r.Request.Data.(*StartLoggingOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartLoggingResponse is the response type for the
// StartLogging API operation.
type StartLoggingResponse struct {
	*StartLoggingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartLogging request.
func (r *StartLoggingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}