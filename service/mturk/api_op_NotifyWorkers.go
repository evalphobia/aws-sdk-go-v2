// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type NotifyWorkersInput struct {
	_ struct{} `type:"structure"`

	// The text of the email message to send. Can include up to 4,096 characters
	//
	// MessageText is a required field
	MessageText *string `type:"string" required:"true"`

	// The subject line of the email message to send. Can include up to 200 characters.
	//
	// Subject is a required field
	Subject *string `type:"string" required:"true"`

	// A list of Worker IDs you wish to notify. You can notify upto 100 Workers
	// at a time.
	//
	// WorkerIds is a required field
	WorkerIds []string `type:"list" required:"true"`
}

// String returns the string representation
func (s NotifyWorkersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *NotifyWorkersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "NotifyWorkersInput"}

	if s.MessageText == nil {
		invalidParams.Add(aws.NewErrParamRequired("MessageText"))
	}

	if s.Subject == nil {
		invalidParams.Add(aws.NewErrParamRequired("Subject"))
	}

	if s.WorkerIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkerIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type NotifyWorkersOutput struct {
	_ struct{} `type:"structure"`

	// When MTurk sends notifications to the list of Workers, it returns back any
	// failures it encounters in this list of NotifyWorkersFailureStatus objects.
	NotifyWorkersFailureStatuses []NotifyWorkersFailureStatus `type:"list"`
}

// String returns the string representation
func (s NotifyWorkersOutput) String() string {
	return awsutil.Prettify(s)
}

const opNotifyWorkers = "NotifyWorkers"

// NotifyWorkersRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The NotifyWorkers operation sends an email to one or more Workers that you
// specify with the Worker ID. You can specify up to 100 Worker IDs to send
// the same message with a single call to the NotifyWorkers operation. The NotifyWorkers
// operation will send a notification email to a Worker only if you have previously
// approved or rejected work from the Worker.
//
//    // Example sending a request using NotifyWorkersRequest.
//    req := client.NotifyWorkersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/NotifyWorkers
func (c *Client) NotifyWorkersRequest(input *NotifyWorkersInput) NotifyWorkersRequest {
	op := &aws.Operation{
		Name:       opNotifyWorkers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &NotifyWorkersInput{}
	}

	req := c.newRequest(op, input, &NotifyWorkersOutput{})

	return NotifyWorkersRequest{Request: req, Input: input, Copy: c.NotifyWorkersRequest}
}

// NotifyWorkersRequest is the request type for the
// NotifyWorkers API operation.
type NotifyWorkersRequest struct {
	*aws.Request
	Input *NotifyWorkersInput
	Copy  func(*NotifyWorkersInput) NotifyWorkersRequest
}

// Send marshals and sends the NotifyWorkers API request.
func (r NotifyWorkersRequest) Send(ctx context.Context) (*NotifyWorkersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NotifyWorkersResponse{
		NotifyWorkersOutput: r.Request.Data.(*NotifyWorkersOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NotifyWorkersResponse is the response type for the
// NotifyWorkers API operation.
type NotifyWorkersResponse struct {
	*NotifyWorkersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NotifyWorkers request.
func (r *NotifyWorkersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}