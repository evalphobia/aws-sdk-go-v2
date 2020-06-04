// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type TagLogGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the log group.
	//
	// LogGroupName is a required field
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string" required:"true"`

	// The key-value pairs to use for the tags.
	//
	// Tags is a required field
	Tags map[string]string `locationName:"tags" min:"1" type:"map" required:"true"`
}

// String returns the string representation
func (s TagLogGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagLogGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagLogGroupInput"}

	if s.LogGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogGroupName"))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TagLogGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TagLogGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opTagLogGroup = "TagLogGroup"

// TagLogGroupRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Adds or updates the specified tags for the specified log group.
//
// To list the tags for a log group, use ListTagsLogGroup (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ListTagsLogGroup.html).
// To remove tags, use UntagLogGroup (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_UntagLogGroup.html).
//
// For more information about tags, see Tag Log Groups in Amazon CloudWatch
// Logs (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Working-with-log-groups-and-streams.html#log-group-tagging)
// in the Amazon CloudWatch Logs User Guide.
//
//    // Example sending a request using TagLogGroupRequest.
//    req := client.TagLogGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/TagLogGroup
func (c *Client) TagLogGroupRequest(input *TagLogGroupInput) TagLogGroupRequest {
	op := &aws.Operation{
		Name:       opTagLogGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TagLogGroupInput{}
	}

	req := c.newRequest(op, input, &TagLogGroupOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return TagLogGroupRequest{Request: req, Input: input, Copy: c.TagLogGroupRequest}
}

// TagLogGroupRequest is the request type for the
// TagLogGroup API operation.
type TagLogGroupRequest struct {
	*aws.Request
	Input *TagLogGroupInput
	Copy  func(*TagLogGroupInput) TagLogGroupRequest
}

// Send marshals and sends the TagLogGroup API request.
func (r TagLogGroupRequest) Send(ctx context.Context) (*TagLogGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagLogGroupResponse{
		TagLogGroupOutput: r.Request.Data.(*TagLogGroupOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagLogGroupResponse is the response type for the
// TagLogGroup API operation.
type TagLogGroupResponse struct {
	*TagLogGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TagLogGroup request.
func (r *TagLogGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}