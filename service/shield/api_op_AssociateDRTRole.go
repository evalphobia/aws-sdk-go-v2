// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateDRTRoleInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the role the DRT will use to access your
	// AWS account.
	//
	// Prior to making the AssociateDRTRole request, you must attach the AWSShieldDRTAccessPolicy
	// (https://console.aws.amazon.com/iam/home?#/policies/arn:aws:iam::aws:policy/service-role/AWSShieldDRTAccessPolicy)
	// managed policy to this role. For more information see Attaching and Detaching
	// IAM Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_manage-attach-detach.html).
	//
	// RoleArn is a required field
	RoleArn *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateDRTRoleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateDRTRoleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateDRTRoleInput"}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateDRTRoleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateDRTRoleOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateDRTRole = "AssociateDRTRole"

// AssociateDRTRoleRequest returns a request value for making API operation for
// AWS Shield.
//
// Authorizes the DDoS Response team (DRT), using the specified role, to access
// your AWS account to assist with DDoS attack mitigation during potential attacks.
// This enables the DRT to inspect your AWS WAF configuration and create or
// update AWS WAF rules and web ACLs.
//
// You can associate only one RoleArn with your subscription. If you submit
// an AssociateDRTRole request for an account that already has an associated
// role, the new RoleArn will replace the existing RoleArn.
//
// Prior to making the AssociateDRTRole request, you must attach the AWSShieldDRTAccessPolicy
// (https://console.aws.amazon.com/iam/home?#/policies/arn:aws:iam::aws:policy/service-role/AWSShieldDRTAccessPolicy)
// managed policy to the role you will specify in the request. For more information
// see Attaching and Detaching IAM Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_manage-attach-detach.html).
// The role must also trust the service principal drt.shield.amazonaws.com.
// For more information, see IAM JSON Policy Elements: Principal (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html).
//
// The DRT will have access only to your AWS WAF and Shield resources. By submitting
// this request, you authorize the DRT to inspect your AWS WAF and Shield configuration
// and create and update AWS WAF rules and web ACLs on your behalf. The DRT
// takes these actions only if explicitly authorized by you.
//
// You must have the iam:PassRole permission to make an AssociateDRTRole request.
// For more information, see Granting a User Permissions to Pass a Role to an
// AWS Service (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html).
//
// To use the services of the DRT and make an AssociateDRTRole request, you
// must be subscribed to the Business Support plan (https://aws.amazon.com/premiumsupport/business-support/)
// or the Enterprise Support plan (https://aws.amazon.com/premiumsupport/enterprise-support/).
//
//    // Example sending a request using AssociateDRTRoleRequest.
//    req := client.AssociateDRTRoleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/AssociateDRTRole
func (c *Client) AssociateDRTRoleRequest(input *AssociateDRTRoleInput) AssociateDRTRoleRequest {
	op := &aws.Operation{
		Name:       opAssociateDRTRole,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateDRTRoleInput{}
	}

	req := c.newRequest(op, input, &AssociateDRTRoleOutput{})

	return AssociateDRTRoleRequest{Request: req, Input: input, Copy: c.AssociateDRTRoleRequest}
}

// AssociateDRTRoleRequest is the request type for the
// AssociateDRTRole API operation.
type AssociateDRTRoleRequest struct {
	*aws.Request
	Input *AssociateDRTRoleInput
	Copy  func(*AssociateDRTRoleInput) AssociateDRTRoleRequest
}

// Send marshals and sends the AssociateDRTRole API request.
func (r AssociateDRTRoleRequest) Send(ctx context.Context) (*AssociateDRTRoleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateDRTRoleResponse{
		AssociateDRTRoleOutput: r.Request.Data.(*AssociateDRTRoleOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateDRTRoleResponse is the response type for the
// AssociateDRTRole API operation.
type AssociateDRTRoleResponse struct {
	*AssociateDRTRoleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateDRTRole request.
func (r *AssociateDRTRoleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}