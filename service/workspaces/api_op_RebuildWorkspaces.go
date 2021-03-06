// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RebuildWorkspacesInput struct {
	_ struct{} `type:"structure"`

	// The WorkSpace to rebuild. You can specify a single WorkSpace.
	//
	// RebuildWorkspaceRequests is a required field
	RebuildWorkspaceRequests []RebuildRequest `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s RebuildWorkspacesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RebuildWorkspacesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RebuildWorkspacesInput"}

	if s.RebuildWorkspaceRequests == nil {
		invalidParams.Add(aws.NewErrParamRequired("RebuildWorkspaceRequests"))
	}
	if s.RebuildWorkspaceRequests != nil && len(s.RebuildWorkspaceRequests) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RebuildWorkspaceRequests", 1))
	}
	if s.RebuildWorkspaceRequests != nil {
		for i, v := range s.RebuildWorkspaceRequests {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RebuildWorkspaceRequests", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RebuildWorkspacesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the WorkSpace that could not be rebuilt.
	FailedRequests []FailedWorkspaceChangeRequest `type:"list"`
}

// String returns the string representation
func (s RebuildWorkspacesOutput) String() string {
	return awsutil.Prettify(s)
}

const opRebuildWorkspaces = "RebuildWorkspaces"

// RebuildWorkspacesRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Rebuilds the specified WorkSpace.
//
// You cannot rebuild a WorkSpace unless its state is AVAILABLE, ERROR, UNHEALTHY,
// or STOPPED.
//
// Rebuilding a WorkSpace is a potentially destructive action that can result
// in the loss of data. For more information, see Rebuild a WorkSpace (https://docs.aws.amazon.com/workspaces/latest/adminguide/reset-workspace.html).
//
// This operation is asynchronous and returns before the WorkSpaces have been
// completely rebuilt.
//
//    // Example sending a request using RebuildWorkspacesRequest.
//    req := client.RebuildWorkspacesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/RebuildWorkspaces
func (c *Client) RebuildWorkspacesRequest(input *RebuildWorkspacesInput) RebuildWorkspacesRequest {
	op := &aws.Operation{
		Name:       opRebuildWorkspaces,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RebuildWorkspacesInput{}
	}

	req := c.newRequest(op, input, &RebuildWorkspacesOutput{})

	return RebuildWorkspacesRequest{Request: req, Input: input, Copy: c.RebuildWorkspacesRequest}
}

// RebuildWorkspacesRequest is the request type for the
// RebuildWorkspaces API operation.
type RebuildWorkspacesRequest struct {
	*aws.Request
	Input *RebuildWorkspacesInput
	Copy  func(*RebuildWorkspacesInput) RebuildWorkspacesRequest
}

// Send marshals and sends the RebuildWorkspaces API request.
func (r RebuildWorkspacesRequest) Send(ctx context.Context) (*RebuildWorkspacesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RebuildWorkspacesResponse{
		RebuildWorkspacesOutput: r.Request.Data.(*RebuildWorkspacesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RebuildWorkspacesResponse is the response type for the
// RebuildWorkspaces API operation.
type RebuildWorkspacesResponse struct {
	*RebuildWorkspacesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RebuildWorkspaces request.
func (r *RebuildWorkspacesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
