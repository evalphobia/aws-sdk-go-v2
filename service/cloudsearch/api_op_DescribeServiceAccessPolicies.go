// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Container for the parameters to the DescribeServiceAccessPolicies operation.
// Specifies the name of the domain you want to describe. To show the active
// configuration and exclude any pending changes, set the Deployed option to
// true.
type DescribeServiceAccessPoliciesInput struct {
	_ struct{} `type:"structure"`

	// Whether to display the deployed configuration (true) or include any pending
	// changes (false). Defaults to false.
	Deployed *bool `type:"boolean"`

	// The name of the domain you want to describe.
	//
	// DomainName is a required field
	DomainName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeServiceAccessPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeServiceAccessPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeServiceAccessPoliciesInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a DescribeServiceAccessPolicies request.
type DescribeServiceAccessPoliciesOutput struct {
	_ struct{} `type:"structure"`

	// The access rules configured for the domain specified in the request.
	//
	// AccessPolicies is a required field
	AccessPolicies *AccessPoliciesStatus `type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeServiceAccessPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeServiceAccessPolicies = "DescribeServiceAccessPolicies"

// DescribeServiceAccessPoliciesRequest returns a request value for making API operation for
// Amazon CloudSearch.
//
// Gets information about the access policies that control access to the domain's
// document and search endpoints. By default, shows the configuration with any
// pending changes. Set the Deployed option to true to show the active configuration
// and exclude pending changes. For more information, see Configuring Access
// for a Search Domain (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-access.html)
// in the Amazon CloudSearch Developer Guide.
//
//    // Example sending a request using DescribeServiceAccessPoliciesRequest.
//    req := client.DescribeServiceAccessPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeServiceAccessPoliciesRequest(input *DescribeServiceAccessPoliciesInput) DescribeServiceAccessPoliciesRequest {
	op := &aws.Operation{
		Name:       opDescribeServiceAccessPolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeServiceAccessPoliciesInput{}
	}

	req := c.newRequest(op, input, &DescribeServiceAccessPoliciesOutput{})

	return DescribeServiceAccessPoliciesRequest{Request: req, Input: input, Copy: c.DescribeServiceAccessPoliciesRequest}
}

// DescribeServiceAccessPoliciesRequest is the request type for the
// DescribeServiceAccessPolicies API operation.
type DescribeServiceAccessPoliciesRequest struct {
	*aws.Request
	Input *DescribeServiceAccessPoliciesInput
	Copy  func(*DescribeServiceAccessPoliciesInput) DescribeServiceAccessPoliciesRequest
}

// Send marshals and sends the DescribeServiceAccessPolicies API request.
func (r DescribeServiceAccessPoliciesRequest) Send(ctx context.Context) (*DescribeServiceAccessPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeServiceAccessPoliciesResponse{
		DescribeServiceAccessPoliciesOutput: r.Request.Data.(*DescribeServiceAccessPoliciesOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeServiceAccessPoliciesResponse is the response type for the
// DescribeServiceAccessPolicies API operation.
type DescribeServiceAccessPoliciesResponse struct {
	*DescribeServiceAccessPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeServiceAccessPolicies request.
func (r *DescribeServiceAccessPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}