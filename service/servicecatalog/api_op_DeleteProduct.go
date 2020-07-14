// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteProductInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The product identifier.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteProductInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteProductInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteProductInput"}

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

type DeleteProductOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteProductOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteProduct = "DeleteProduct"

// DeleteProductRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Deletes the specified product.
//
// You cannot delete a product if it was shared with you or is associated with
// a portfolio.
//
//    // Example sending a request using DeleteProductRequest.
//    req := client.DeleteProductRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DeleteProduct
func (c *Client) DeleteProductRequest(input *DeleteProductInput) DeleteProductRequest {
	op := &aws.Operation{
		Name:       opDeleteProduct,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteProductInput{}
	}

	req := c.newRequest(op, input, &DeleteProductOutput{})

	return DeleteProductRequest{Request: req, Input: input, Copy: c.DeleteProductRequest}
}

// DeleteProductRequest is the request type for the
// DeleteProduct API operation.
type DeleteProductRequest struct {
	*aws.Request
	Input *DeleteProductInput
	Copy  func(*DeleteProductInput) DeleteProductRequest
}

// Send marshals and sends the DeleteProduct API request.
func (r DeleteProductRequest) Send(ctx context.Context) (*DeleteProductResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteProductResponse{
		DeleteProductOutput: r.Request.Data.(*DeleteProductOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteProductResponse is the response type for the
// DeleteProduct API operation.
type DeleteProductResponse struct {
	*DeleteProductOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteProduct request.
func (r *DeleteProductResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}