// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// DescribeTaskExecutionRequest
type DescribeTaskExecutionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the task that is being executed.
	//
	// TaskExecutionArn is a required field
	TaskExecutionArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTaskExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTaskExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTaskExecutionInput"}

	if s.TaskExecutionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskExecutionArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// DescribeTaskExecutionResponse
type DescribeTaskExecutionOutput struct {
	_ struct{} `type:"structure"`

	// The physical number of bytes transferred over the network.
	BytesTransferred *int64 `type:"long"`

	// The number of logical bytes written to the destination AWS storage resource.
	BytesWritten *int64 `type:"long"`

	// The estimated physical number of bytes that is to be transferred over the
	// network.
	EstimatedBytesToTransfer *int64 `type:"long"`

	// The expected number of files that is to be transferred over the network.
	// This value is calculated during the PREPARING phase, before the TRANSFERRING
	// phase. This value is the expected number of files to be transferred. It's
	// calculated based on comparing the content of the source and destination locations
	// and finding the delta that needs to be transferred.
	EstimatedFilesToTransfer *int64 `type:"long"`

	// A list of filter rules that determines which files to exclude from a task.
	// The list should contain a single filter string that consists of the patterns
	// to exclude. The patterns are delimited by "|" (that is, a pipe), for example:
	// "/folder1|/folder2"
	Excludes []FilterRule `type:"list"`

	// The actual number of files that was transferred over the network. This value
	// is calculated and updated on an ongoing basis during the TRANSFERRING phase.
	// It's updated periodically when each file is read from the source and sent
	// over the network.
	//
	// If failures occur during a transfer, this value can be less than EstimatedFilesToTransfer.
	// This value can also be greater than EstimatedFilesTransferred in some cases.
	// This element is implementation-specific for some location types, so don't
	// use it as an indicator for a correct file number or to monitor your task
	// execution.
	FilesTransferred *int64 `type:"long"`

	// A list of filter rules that determines which files to include when running
	// a task. The list should contain a single filter string that consists of the
	// patterns to include. The patterns are delimited by "|" (that is, a pipe),
	// for example: "/folder1|/folder2"
	Includes []FilterRule `type:"list"`

	// Represents the options that are available to control the behavior of a StartTaskExecution
	// operation. Behavior includes preserving metadata such as user ID (UID), group
	// ID (GID), and file permissions, and also overwriting files in the destination,
	// data integrity verification, and so on.
	//
	// A task has a set of default options associated with it. If you don't specify
	// an option in StartTaskExecution, the default value is used. You can override
	// the defaults options on each task execution by specifying an overriding Options
	// value to StartTaskExecution.
	Options *Options `type:"structure"`

	// The result of the task execution.
	Result *TaskExecutionResultDetail `type:"structure"`

	// The time that the task execution was started.
	StartTime *time.Time `type:"timestamp"`

	// The status of the task execution.
	//
	// For detailed information about task execution statuses, see Understanding
	// Task Statuses in the AWS DataSync User Guide.
	Status TaskExecutionStatus `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the task execution that was described.
	// TaskExecutionArn is hierarchical and includes TaskArn for the task that was
	// executed.
	//
	// For example, a TaskExecution value with the ARN arn:aws:datasync:us-east-1:111222333444:task/task-0208075f79cedf4a2/execution/exec-08ef1e88ec491019b
	// executed the task with the ARN arn:aws:datasync:us-east-1:111222333444:task/task-0208075f79cedf4a2.
	TaskExecutionArn *string `type:"string"`
}

// String returns the string representation
func (s DescribeTaskExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTaskExecution = "DescribeTaskExecution"

// DescribeTaskExecutionRequest returns a request value for making API operation for
// AWS DataSync.
//
// Returns detailed metadata about a task that is being executed.
//
//    // Example sending a request using DescribeTaskExecutionRequest.
//    req := client.DescribeTaskExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/DescribeTaskExecution
func (c *Client) DescribeTaskExecutionRequest(input *DescribeTaskExecutionInput) DescribeTaskExecutionRequest {
	op := &aws.Operation{
		Name:       opDescribeTaskExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTaskExecutionInput{}
	}

	req := c.newRequest(op, input, &DescribeTaskExecutionOutput{})

	return DescribeTaskExecutionRequest{Request: req, Input: input, Copy: c.DescribeTaskExecutionRequest}
}

// DescribeTaskExecutionRequest is the request type for the
// DescribeTaskExecution API operation.
type DescribeTaskExecutionRequest struct {
	*aws.Request
	Input *DescribeTaskExecutionInput
	Copy  func(*DescribeTaskExecutionInput) DescribeTaskExecutionRequest
}

// Send marshals and sends the DescribeTaskExecution API request.
func (r DescribeTaskExecutionRequest) Send(ctx context.Context) (*DescribeTaskExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTaskExecutionResponse{
		DescribeTaskExecutionOutput: r.Request.Data.(*DescribeTaskExecutionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTaskExecutionResponse is the response type for the
// DescribeTaskExecution API operation.
type DescribeTaskExecutionResponse struct {
	*DescribeTaskExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTaskExecution request.
func (r *DescribeTaskExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}