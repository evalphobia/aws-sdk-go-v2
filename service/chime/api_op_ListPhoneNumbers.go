// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListPhoneNumbersInput struct {
	_ struct{} `type:"structure"`

	// The filter to use to limit the number of results.
	FilterName PhoneNumberAssociationName `location:"querystring" locationName:"filter-name" type:"string" enum:"true"`

	// The value to use for the filter.
	FilterValue *string `location:"querystring" locationName:"filter-value" type:"string"`

	// The maximum number of results to return in a single call.
	MaxResults *int64 `location:"querystring" locationName:"max-results" min:"1" type:"integer"`

	// The token to use to retrieve the next page of results.
	NextToken *string `location:"querystring" locationName:"next-token" type:"string"`

	// The phone number product type.
	ProductType PhoneNumberProductType `location:"querystring" locationName:"product-type" type:"string" enum:"true"`

	// The phone number status.
	Status PhoneNumberStatus `location:"querystring" locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListPhoneNumbersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPhoneNumbersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPhoneNumbersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPhoneNumbersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if len(s.FilterName) > 0 {
		v := s.FilterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "filter-name", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.FilterValue != nil {
		v := *s.FilterValue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "filter-value", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max-results", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "next-token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ProductType) > 0 {
		v := s.ProductType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "product-type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type ListPhoneNumbersOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results.
	NextToken *string `type:"string"`

	// The phone number details.
	PhoneNumbers []PhoneNumber `type:"list"`
}

// String returns the string representation
func (s ListPhoneNumbersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPhoneNumbersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PhoneNumbers != nil {
		v := s.PhoneNumbers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "PhoneNumbers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListPhoneNumbers = "ListPhoneNumbers"

// ListPhoneNumbersRequest returns a request value for making API operation for
// Amazon Chime.
//
// Lists the phone numbers for the specified Amazon Chime account, Amazon Chime
// user, Amazon Chime Voice Connector, or Amazon Chime Voice Connector group.
//
//    // Example sending a request using ListPhoneNumbersRequest.
//    req := client.ListPhoneNumbersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/ListPhoneNumbers
func (c *Client) ListPhoneNumbersRequest(input *ListPhoneNumbersInput) ListPhoneNumbersRequest {
	op := &aws.Operation{
		Name:       opListPhoneNumbers,
		HTTPMethod: "GET",
		HTTPPath:   "/phone-numbers",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListPhoneNumbersInput{}
	}

	req := c.newRequest(op, input, &ListPhoneNumbersOutput{})

	return ListPhoneNumbersRequest{Request: req, Input: input, Copy: c.ListPhoneNumbersRequest}
}

// ListPhoneNumbersRequest is the request type for the
// ListPhoneNumbers API operation.
type ListPhoneNumbersRequest struct {
	*aws.Request
	Input *ListPhoneNumbersInput
	Copy  func(*ListPhoneNumbersInput) ListPhoneNumbersRequest
}

// Send marshals and sends the ListPhoneNumbers API request.
func (r ListPhoneNumbersRequest) Send(ctx context.Context) (*ListPhoneNumbersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPhoneNumbersResponse{
		ListPhoneNumbersOutput: r.Request.Data.(*ListPhoneNumbersOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPhoneNumbersRequestPaginator returns a paginator for ListPhoneNumbers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPhoneNumbersRequest(input)
//   p := chime.NewListPhoneNumbersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPhoneNumbersPaginator(req ListPhoneNumbersRequest) ListPhoneNumbersPaginator {
	return ListPhoneNumbersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListPhoneNumbersInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListPhoneNumbersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPhoneNumbersPaginator struct {
	aws.Pager
}

func (p *ListPhoneNumbersPaginator) CurrentPage() *ListPhoneNumbersOutput {
	return p.Pager.CurrentPage().(*ListPhoneNumbersOutput)
}

// ListPhoneNumbersResponse is the response type for the
// ListPhoneNumbers API operation.
type ListPhoneNumbersResponse struct {
	*ListPhoneNumbersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPhoneNumbers request.
func (r *ListPhoneNumbersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}