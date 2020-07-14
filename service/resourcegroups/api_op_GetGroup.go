// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourcegroups

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the resource group.
	//
	// GroupName is a required field
	GroupName *string `location:"uri" locationName:"GroupName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetGroupInput"}

	if s.GroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupName"))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GroupName != nil {
		v := *s.GroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetGroupOutput struct {
	_ struct{} `type:"structure"`

	// A full description of the resource group.
	Group *Group `type:"structure"`
}

// String returns the string representation
func (s GetGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Group != nil {
		v := s.Group

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Group", v, metadata)
	}
	return nil
}

const opGetGroup = "GetGroup"

// GetGroupRequest returns a request value for making API operation for
// AWS Resource Groups.
//
// Returns information about a specified resource group.
//
//    // Example sending a request using GetGroupRequest.
//    req := client.GetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/resource-groups-2017-11-27/GetGroup
func (c *Client) GetGroupRequest(input *GetGroupInput) GetGroupRequest {
	op := &aws.Operation{
		Name:       opGetGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/groups/{GroupName}",
	}

	if input == nil {
		input = &GetGroupInput{}
	}

	req := c.newRequest(op, input, &GetGroupOutput{})

	return GetGroupRequest{Request: req, Input: input, Copy: c.GetGroupRequest}
}

// GetGroupRequest is the request type for the
// GetGroup API operation.
type GetGroupRequest struct {
	*aws.Request
	Input *GetGroupInput
	Copy  func(*GetGroupInput) GetGroupRequest
}

// Send marshals and sends the GetGroup API request.
func (r GetGroupRequest) Send(ctx context.Context) (*GetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetGroupResponse{
		GetGroupOutput: r.Request.Data.(*GetGroupOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetGroupResponse is the response type for the
// GetGroup API operation.
type GetGroupResponse struct {
	*GetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetGroup request.
func (r *GetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}