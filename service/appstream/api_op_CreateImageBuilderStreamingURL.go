// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateImageBuilderStreamingURLInput struct {
	_ struct{} `type:"structure"`

	// The name of the image builder.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The time that the streaming URL will be valid, in seconds. Specify a value
	// between 1 and 604800 seconds. The default is 3600 seconds.
	Validity *int64 `type:"long"`
}

// String returns the string representation
func (s CreateImageBuilderStreamingURLInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateImageBuilderStreamingURLInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateImageBuilderStreamingURLInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateImageBuilderStreamingURLOutput struct {
	_ struct{} `type:"structure"`

	// The elapsed time, in seconds after the Unix epoch, when this URL expires.
	Expires *time.Time `type:"timestamp"`

	// The URL to start the AppStream 2.0 streaming session.
	StreamingURL *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateImageBuilderStreamingURLOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateImageBuilderStreamingURL = "CreateImageBuilderStreamingURL"

// CreateImageBuilderStreamingURLRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Creates a URL to start an image builder streaming session.
//
//    // Example sending a request using CreateImageBuilderStreamingURLRequest.
//    req := client.CreateImageBuilderStreamingURLRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/CreateImageBuilderStreamingURL
func (c *Client) CreateImageBuilderStreamingURLRequest(input *CreateImageBuilderStreamingURLInput) CreateImageBuilderStreamingURLRequest {
	op := &aws.Operation{
		Name:       opCreateImageBuilderStreamingURL,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateImageBuilderStreamingURLInput{}
	}

	req := c.newRequest(op, input, &CreateImageBuilderStreamingURLOutput{})

	return CreateImageBuilderStreamingURLRequest{Request: req, Input: input, Copy: c.CreateImageBuilderStreamingURLRequest}
}

// CreateImageBuilderStreamingURLRequest is the request type for the
// CreateImageBuilderStreamingURL API operation.
type CreateImageBuilderStreamingURLRequest struct {
	*aws.Request
	Input *CreateImageBuilderStreamingURLInput
	Copy  func(*CreateImageBuilderStreamingURLInput) CreateImageBuilderStreamingURLRequest
}

// Send marshals and sends the CreateImageBuilderStreamingURL API request.
func (r CreateImageBuilderStreamingURLRequest) Send(ctx context.Context) (*CreateImageBuilderStreamingURLResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateImageBuilderStreamingURLResponse{
		CreateImageBuilderStreamingURLOutput: r.Request.Data.(*CreateImageBuilderStreamingURLOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateImageBuilderStreamingURLResponse is the response type for the
// CreateImageBuilderStreamingURL API operation.
type CreateImageBuilderStreamingURLResponse struct {
	*CreateImageBuilderStreamingURLOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateImageBuilderStreamingURL request.
func (r *CreateImageBuilderStreamingURLResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
