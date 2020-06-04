// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package polly

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetSpeechSynthesisTaskInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Polly generated identifier for a speech synthesis task.
	//
	// TaskId is a required field
	TaskId *string `location:"uri" locationName:"TaskId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSpeechSynthesisTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSpeechSynthesisTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSpeechSynthesisTaskInput"}

	if s.TaskId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSpeechSynthesisTaskInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TaskId != nil {
		v := *s.TaskId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "TaskId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetSpeechSynthesisTaskOutput struct {
	_ struct{} `type:"structure"`

	// SynthesisTask object that provides information from the requested task, including
	// output format, creation time, task status, and so on.
	SynthesisTask *SynthesisTask `type:"structure"`
}

// String returns the string representation
func (s GetSpeechSynthesisTaskOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSpeechSynthesisTaskOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SynthesisTask != nil {
		v := s.SynthesisTask

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SynthesisTask", v, metadata)
	}
	return nil
}

const opGetSpeechSynthesisTask = "GetSpeechSynthesisTask"

// GetSpeechSynthesisTaskRequest returns a request value for making API operation for
// Amazon Polly.
//
// Retrieves a specific SpeechSynthesisTask object based on its TaskID. This
// object contains information about the given speech synthesis task, including
// the status of the task, and a link to the S3 bucket containing the output
// of the task.
//
//    // Example sending a request using GetSpeechSynthesisTaskRequest.
//    req := client.GetSpeechSynthesisTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/GetSpeechSynthesisTask
func (c *Client) GetSpeechSynthesisTaskRequest(input *GetSpeechSynthesisTaskInput) GetSpeechSynthesisTaskRequest {
	op := &aws.Operation{
		Name:       opGetSpeechSynthesisTask,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/synthesisTasks/{TaskId}",
	}

	if input == nil {
		input = &GetSpeechSynthesisTaskInput{}
	}

	req := c.newRequest(op, input, &GetSpeechSynthesisTaskOutput{})

	return GetSpeechSynthesisTaskRequest{Request: req, Input: input, Copy: c.GetSpeechSynthesisTaskRequest}
}

// GetSpeechSynthesisTaskRequest is the request type for the
// GetSpeechSynthesisTask API operation.
type GetSpeechSynthesisTaskRequest struct {
	*aws.Request
	Input *GetSpeechSynthesisTaskInput
	Copy  func(*GetSpeechSynthesisTaskInput) GetSpeechSynthesisTaskRequest
}

// Send marshals and sends the GetSpeechSynthesisTask API request.
func (r GetSpeechSynthesisTaskRequest) Send(ctx context.Context) (*GetSpeechSynthesisTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSpeechSynthesisTaskResponse{
		GetSpeechSynthesisTaskOutput: r.Request.Data.(*GetSpeechSynthesisTaskOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSpeechSynthesisTaskResponse is the response type for the
// GetSpeechSynthesisTask API operation.
type GetSpeechSynthesisTaskResponse struct {
	*GetSpeechSynthesisTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSpeechSynthesisTask request.
func (r *GetSpeechSynthesisTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}