// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeProductViewInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The product view identifier.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeProductViewInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeProductViewInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeProductViewInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeProductViewOutput struct {
	_ struct{} `type:"structure"`

	// Summary information about the product.
	ProductViewSummary *ProductViewSummary `type:"structure"`

	// Information about the provisioning artifacts for the product.
	ProvisioningArtifacts []ProvisioningArtifact `type:"list"`
}

// String returns the string representation
func (s DescribeProductViewOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeProductView = "DescribeProductView"

// DescribeProductViewRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Gets information about the specified product.
//
//    // Example sending a request using DescribeProductViewRequest.
//    req := client.DescribeProductViewRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeProductView
func (c *Client) DescribeProductViewRequest(input *DescribeProductViewInput) DescribeProductViewRequest {
	op := &aws.Operation{
		Name:       opDescribeProductView,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeProductViewInput{}
	}

	req := c.newRequest(op, input, &DescribeProductViewOutput{})

	return DescribeProductViewRequest{Request: req, Input: input, Copy: c.DescribeProductViewRequest}
}

// DescribeProductViewRequest is the request type for the
// DescribeProductView API operation.
type DescribeProductViewRequest struct {
	*aws.Request
	Input *DescribeProductViewInput
	Copy  func(*DescribeProductViewInput) DescribeProductViewRequest
}

// Send marshals and sends the DescribeProductView API request.
func (r DescribeProductViewRequest) Send(ctx context.Context) (*DescribeProductViewResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeProductViewResponse{
		DescribeProductViewOutput: r.Request.Data.(*DescribeProductViewOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeProductViewResponse is the response type for the
// DescribeProductView API operation.
type DescribeProductViewResponse struct {
	*DescribeProductViewOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeProductView request.
func (r *DescribeProductViewResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}