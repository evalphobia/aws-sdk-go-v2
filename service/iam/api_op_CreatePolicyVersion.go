// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreatePolicyVersionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the IAM policy to which you want to add
	// a new version.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// PolicyArn is a required field
	PolicyArn *string `min:"20" type:"string" required:"true"`

	// The JSON policy document that you want to use as the content for this new
	// version of the policy.
	//
	// You must provide policies in JSON format in IAM. However, for AWS CloudFormation
	// templates formatted in YAML, you can provide the policy in JSON or YAML format.
	// AWS CloudFormation always converts a YAML policy to JSON format before submitting
	// it to IAM.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) used to validate this
	// parameter is a string of characters consisting of the following:
	//
	//    * Any printable ASCII character ranging from the space character (\u0020)
	//    through the end of the ASCII character range
	//
	//    * The printable characters in the Basic Latin and Latin-1 Supplement character
	//    set (through \u00FF)
	//
	//    * The special characters tab (\u0009), line feed (\u000A), and carriage
	//    return (\u000D)
	//
	// PolicyDocument is a required field
	PolicyDocument *string `min:"1" type:"string" required:"true"`

	// Specifies whether to set this version as the policy's default version.
	//
	// When this parameter is true, the new policy version becomes the operative
	// version. That is, it becomes the version that is in effect for the IAM users,
	// groups, and roles that the policy is attached to.
	//
	// For more information about managed policy versions, see Versioning for Managed
	// Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html)
	// in the IAM User Guide.
	SetAsDefault *bool `type:"boolean"`
}

// String returns the string representation
func (s CreatePolicyVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePolicyVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePolicyVersionInput"}

	if s.PolicyArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyArn"))
	}
	if s.PolicyArn != nil && len(*s.PolicyArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyArn", 20))
	}

	if s.PolicyDocument == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyDocument"))
	}
	if s.PolicyDocument != nil && len(*s.PolicyDocument) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyDocument", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful CreatePolicyVersion request.
type CreatePolicyVersionOutput struct {
	_ struct{} `type:"structure"`

	// A structure containing details about the new policy version.
	PolicyVersion *PolicyVersion `type:"structure"`
}

// String returns the string representation
func (s CreatePolicyVersionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreatePolicyVersion = "CreatePolicyVersion"

// CreatePolicyVersionRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Creates a new version of the specified managed policy. To update a managed
// policy, you create a new policy version. A managed policy can have up to
// five versions. If the policy has five versions, you must delete an existing
// version using DeletePolicyVersion before you create a new version.
//
// Optionally, you can set the new version as the policy's default version.
// The default version is the version that is in effect for the IAM users, groups,
// and roles to which the policy is attached.
//
// For more information about managed policy versions, see Versioning for Managed
// Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html)
// in the IAM User Guide.
//
//    // Example sending a request using CreatePolicyVersionRequest.
//    req := client.CreatePolicyVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreatePolicyVersion
func (c *Client) CreatePolicyVersionRequest(input *CreatePolicyVersionInput) CreatePolicyVersionRequest {
	op := &aws.Operation{
		Name:       opCreatePolicyVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreatePolicyVersionInput{}
	}

	req := c.newRequest(op, input, &CreatePolicyVersionOutput{})

	return CreatePolicyVersionRequest{Request: req, Input: input, Copy: c.CreatePolicyVersionRequest}
}

// CreatePolicyVersionRequest is the request type for the
// CreatePolicyVersion API operation.
type CreatePolicyVersionRequest struct {
	*aws.Request
	Input *CreatePolicyVersionInput
	Copy  func(*CreatePolicyVersionInput) CreatePolicyVersionRequest
}

// Send marshals and sends the CreatePolicyVersion API request.
func (r CreatePolicyVersionRequest) Send(ctx context.Context) (*CreatePolicyVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePolicyVersionResponse{
		CreatePolicyVersionOutput: r.Request.Data.(*CreatePolicyVersionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePolicyVersionResponse is the response type for the
// CreatePolicyVersion API operation.
type CreatePolicyVersionResponse struct {
	*CreatePolicyVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePolicyVersion request.
func (r *CreatePolicyVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}