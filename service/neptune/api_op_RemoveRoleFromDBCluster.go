// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type RemoveRoleFromDBClusterInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB cluster to disassociate the IAM role from.
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the IAM role to disassociate from the DB
	// cluster, for example arn:aws:iam::123456789012:role/NeptuneAccessRole.
	//
	// RoleArn is a required field
	RoleArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RemoveRoleFromDBClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveRoleFromDBClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveRoleFromDBClusterInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RemoveRoleFromDBClusterOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveRoleFromDBClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opRemoveRoleFromDBCluster = "RemoveRoleFromDBCluster"

// RemoveRoleFromDBClusterRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Disassociates an Identity and Access Management (IAM) role from a DB cluster.
//
//    // Example sending a request using RemoveRoleFromDBClusterRequest.
//    req := client.RemoveRoleFromDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/RemoveRoleFromDBCluster
func (c *Client) RemoveRoleFromDBClusterRequest(input *RemoveRoleFromDBClusterInput) RemoveRoleFromDBClusterRequest {
	op := &aws.Operation{
		Name:       opRemoveRoleFromDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveRoleFromDBClusterInput{}
	}

	req := c.newRequest(op, input, &RemoveRoleFromDBClusterOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return RemoveRoleFromDBClusterRequest{Request: req, Input: input, Copy: c.RemoveRoleFromDBClusterRequest}
}

// RemoveRoleFromDBClusterRequest is the request type for the
// RemoveRoleFromDBCluster API operation.
type RemoveRoleFromDBClusterRequest struct {
	*aws.Request
	Input *RemoveRoleFromDBClusterInput
	Copy  func(*RemoveRoleFromDBClusterInput) RemoveRoleFromDBClusterRequest
}

// Send marshals and sends the RemoveRoleFromDBCluster API request.
func (r RemoveRoleFromDBClusterRequest) Send(ctx context.Context) (*RemoveRoleFromDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveRoleFromDBClusterResponse{
		RemoveRoleFromDBClusterOutput: r.Request.Data.(*RemoveRoleFromDBClusterOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveRoleFromDBClusterResponse is the response type for the
// RemoveRoleFromDBCluster API operation.
type RemoveRoleFromDBClusterResponse struct {
	*RemoveRoleFromDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveRoleFromDBCluster request.
func (r *RemoveRoleFromDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}