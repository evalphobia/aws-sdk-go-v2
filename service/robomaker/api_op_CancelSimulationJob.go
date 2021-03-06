// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CancelSimulationJobInput struct {
	_ struct{} `type:"structure"`

	// The simulation job ARN to cancel.
	//
	// Job is a required field
	Job *string `locationName:"job" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelSimulationJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelSimulationJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelSimulationJobInput"}

	if s.Job == nil {
		invalidParams.Add(aws.NewErrParamRequired("Job"))
	}
	if s.Job != nil && len(*s.Job) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Job", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelSimulationJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Job != nil {
		v := *s.Job

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "job", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CancelSimulationJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CancelSimulationJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelSimulationJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCancelSimulationJob = "CancelSimulationJob"

// CancelSimulationJobRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Cancels the specified simulation job.
//
//    // Example sending a request using CancelSimulationJobRequest.
//    req := client.CancelSimulationJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CancelSimulationJob
func (c *Client) CancelSimulationJobRequest(input *CancelSimulationJobInput) CancelSimulationJobRequest {
	op := &aws.Operation{
		Name:       opCancelSimulationJob,
		HTTPMethod: "POST",
		HTTPPath:   "/cancelSimulationJob",
	}

	if input == nil {
		input = &CancelSimulationJobInput{}
	}

	req := c.newRequest(op, input, &CancelSimulationJobOutput{})

	return CancelSimulationJobRequest{Request: req, Input: input, Copy: c.CancelSimulationJobRequest}
}

// CancelSimulationJobRequest is the request type for the
// CancelSimulationJob API operation.
type CancelSimulationJobRequest struct {
	*aws.Request
	Input *CancelSimulationJobInput
	Copy  func(*CancelSimulationJobInput) CancelSimulationJobRequest
}

// Send marshals and sends the CancelSimulationJob API request.
func (r CancelSimulationJobRequest) Send(ctx context.Context) (*CancelSimulationJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelSimulationJobResponse{
		CancelSimulationJobOutput: r.Request.Data.(*CancelSimulationJobOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelSimulationJobResponse is the response type for the
// CancelSimulationJob API operation.
type CancelSimulationJobResponse struct {
	*CancelSimulationJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelSimulationJob request.
func (r *CancelSimulationJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
