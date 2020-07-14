// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to modify the delivery options for a configuration set.
type PutConfigurationSetDeliveryOptionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set that you want to specify the delivery options
	// for.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `type:"string" required:"true"`

	// Specifies whether messages that use the configuration set are required to
	// use Transport Layer Security (TLS).
	DeliveryOptions *DeliveryOptions `type:"structure"`
}

// String returns the string representation
func (s PutConfigurationSetDeliveryOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutConfigurationSetDeliveryOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutConfigurationSetDeliveryOptionsInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type PutConfigurationSetDeliveryOptionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutConfigurationSetDeliveryOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutConfigurationSetDeliveryOptions = "PutConfigurationSetDeliveryOptions"

// PutConfigurationSetDeliveryOptionsRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Adds or updates the delivery options for a configuration set.
//
//    // Example sending a request using PutConfigurationSetDeliveryOptionsRequest.
//    req := client.PutConfigurationSetDeliveryOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/PutConfigurationSetDeliveryOptions
func (c *Client) PutConfigurationSetDeliveryOptionsRequest(input *PutConfigurationSetDeliveryOptionsInput) PutConfigurationSetDeliveryOptionsRequest {
	op := &aws.Operation{
		Name:       opPutConfigurationSetDeliveryOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutConfigurationSetDeliveryOptionsInput{}
	}

	req := c.newRequest(op, input, &PutConfigurationSetDeliveryOptionsOutput{})

	return PutConfigurationSetDeliveryOptionsRequest{Request: req, Input: input, Copy: c.PutConfigurationSetDeliveryOptionsRequest}
}

// PutConfigurationSetDeliveryOptionsRequest is the request type for the
// PutConfigurationSetDeliveryOptions API operation.
type PutConfigurationSetDeliveryOptionsRequest struct {
	*aws.Request
	Input *PutConfigurationSetDeliveryOptionsInput
	Copy  func(*PutConfigurationSetDeliveryOptionsInput) PutConfigurationSetDeliveryOptionsRequest
}

// Send marshals and sends the PutConfigurationSetDeliveryOptions API request.
func (r PutConfigurationSetDeliveryOptionsRequest) Send(ctx context.Context) (*PutConfigurationSetDeliveryOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutConfigurationSetDeliveryOptionsResponse{
		PutConfigurationSetDeliveryOptionsOutput: r.Request.Data.(*PutConfigurationSetDeliveryOptionsOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutConfigurationSetDeliveryOptionsResponse is the response type for the
// PutConfigurationSetDeliveryOptions API operation.
type PutConfigurationSetDeliveryOptionsResponse struct {
	*PutConfigurationSetDeliveryOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutConfigurationSetDeliveryOptions request.
func (r *PutConfigurationSetDeliveryOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}