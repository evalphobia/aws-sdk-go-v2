// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Container for the parameters to the DescribeInboundCrossClusterSearchConnections
// operation.
type DescribeInboundCrossClusterSearchConnectionsInput struct {
	_ struct{} `type:"structure"`

	// A list of filters used to match properties for inbound cross-cluster search
	// connection. Available Filter names for this operation are:
	//    * cross-cluster-search-connection-id
	//
	//    * source-domain-info.domain-name
	//
	//    * source-domain-info.owner-id
	//
	//    * source-domain-info.region
	//
	//    * destination-domain-info.domain-name
	Filters []Filter `type:"list"`

	// Set this value to limit the number of results returned. If not specified,
	// defaults to 100.
	MaxResults *int64 `type:"integer"`

	// NextToken is sent in case the earlier API call results contain the NextToken.
	// It is used for pagination.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInboundCrossClusterSearchConnectionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInboundCrossClusterSearchConnectionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeInboundCrossClusterSearchConnectionsInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeInboundCrossClusterSearchConnectionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Filters != nil {
		v := s.Filters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Filters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result of a DescribeInboundCrossClusterSearchConnections request. Contains
// the list of connections matching the filter criteria.
type DescribeInboundCrossClusterSearchConnectionsOutput struct {
	_ struct{} `type:"structure"`

	// Consists of list of InboundCrossClusterSearchConnection matching the specified
	// filter criteria.
	CrossClusterSearchConnections []InboundCrossClusterSearchConnection `type:"list"`

	// If more results are available and NextToken is present, make the next request
	// to the same API with the received NextToken to paginate the remaining results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInboundCrossClusterSearchConnectionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeInboundCrossClusterSearchConnectionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CrossClusterSearchConnections != nil {
		v := s.CrossClusterSearchConnections

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "CrossClusterSearchConnections", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeInboundCrossClusterSearchConnections = "DescribeInboundCrossClusterSearchConnections"

// DescribeInboundCrossClusterSearchConnectionsRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Lists all the inbound cross-cluster search connections for a destination
// domain.
//
//    // Example sending a request using DescribeInboundCrossClusterSearchConnectionsRequest.
//    req := client.DescribeInboundCrossClusterSearchConnectionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeInboundCrossClusterSearchConnectionsRequest(input *DescribeInboundCrossClusterSearchConnectionsInput) DescribeInboundCrossClusterSearchConnectionsRequest {
	op := &aws.Operation{
		Name:       opDescribeInboundCrossClusterSearchConnections,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-01-01/es/ccs/inboundConnection/search",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeInboundCrossClusterSearchConnectionsInput{}
	}

	req := c.newRequest(op, input, &DescribeInboundCrossClusterSearchConnectionsOutput{})

	return DescribeInboundCrossClusterSearchConnectionsRequest{Request: req, Input: input, Copy: c.DescribeInboundCrossClusterSearchConnectionsRequest}
}

// DescribeInboundCrossClusterSearchConnectionsRequest is the request type for the
// DescribeInboundCrossClusterSearchConnections API operation.
type DescribeInboundCrossClusterSearchConnectionsRequest struct {
	*aws.Request
	Input *DescribeInboundCrossClusterSearchConnectionsInput
	Copy  func(*DescribeInboundCrossClusterSearchConnectionsInput) DescribeInboundCrossClusterSearchConnectionsRequest
}

// Send marshals and sends the DescribeInboundCrossClusterSearchConnections API request.
func (r DescribeInboundCrossClusterSearchConnectionsRequest) Send(ctx context.Context) (*DescribeInboundCrossClusterSearchConnectionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInboundCrossClusterSearchConnectionsResponse{
		DescribeInboundCrossClusterSearchConnectionsOutput: r.Request.Data.(*DescribeInboundCrossClusterSearchConnectionsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeInboundCrossClusterSearchConnectionsRequestPaginator returns a paginator for DescribeInboundCrossClusterSearchConnections.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeInboundCrossClusterSearchConnectionsRequest(input)
//   p := elasticsearchservice.NewDescribeInboundCrossClusterSearchConnectionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeInboundCrossClusterSearchConnectionsPaginator(req DescribeInboundCrossClusterSearchConnectionsRequest) DescribeInboundCrossClusterSearchConnectionsPaginator {
	return DescribeInboundCrossClusterSearchConnectionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeInboundCrossClusterSearchConnectionsInput
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

// DescribeInboundCrossClusterSearchConnectionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeInboundCrossClusterSearchConnectionsPaginator struct {
	aws.Pager
}

func (p *DescribeInboundCrossClusterSearchConnectionsPaginator) CurrentPage() *DescribeInboundCrossClusterSearchConnectionsOutput {
	return p.Pager.CurrentPage().(*DescribeInboundCrossClusterSearchConnectionsOutput)
}

// DescribeInboundCrossClusterSearchConnectionsResponse is the response type for the
// DescribeInboundCrossClusterSearchConnections API operation.
type DescribeInboundCrossClusterSearchConnectionsResponse struct {
	*DescribeInboundCrossClusterSearchConnectionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInboundCrossClusterSearchConnections request.
func (r *DescribeInboundCrossClusterSearchConnectionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
