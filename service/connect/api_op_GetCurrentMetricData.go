// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetCurrentMetricDataInput struct {
	_ struct{} `type:"structure"`

	// The metrics to retrieve. Specify the name and unit for each metric. The following
	// metrics are available:
	//
	// AGENTS_AFTER_CONTACT_WORK
	//
	// Unit: COUNT
	//
	// AGENTS_AVAILABLE
	//
	// Unit: COUNT
	//
	// AGENTS_ERROR
	//
	// Unit: COUNT
	//
	// AGENTS_NON_PRODUCTIVE
	//
	// Unit: COUNT
	//
	// AGENTS_ON_CALL
	//
	// Unit: COUNT
	//
	// AGENTS_ON_CONTACT
	//
	// Unit: COUNT
	//
	// AGENTS_ONLINE
	//
	// Unit: COUNT
	//
	// AGENTS_STAFFED
	//
	// Unit: COUNT
	//
	// CONTACTS_IN_QUEUE
	//
	// Unit: COUNT
	//
	// CONTACTS_SCHEDULED
	//
	// Unit: COUNT
	//
	// OLDEST_CONTACT_AGE
	//
	// Unit: SECONDS
	//
	// SLOTS_ACTIVE
	//
	// Unit: COUNT
	//
	// SLOTS_AVAILABLE
	//
	// Unit: COUNT
	//
	// CurrentMetrics is a required field
	CurrentMetrics []CurrentMetric `type:"list" required:"true"`

	// The queues, up to 100, or channels, to use to filter the metrics returned.
	// Metric data is retrieved only for the resources associated with the queues
	// or channels included in the filter. You can include both queue IDs and queue
	// ARNs in the same request. The only supported channel is VOICE.
	//
	// Filters is a required field
	Filters *Filters `type:"structure" required:"true"`

	// The grouping applied to the metrics returned. For example, when grouped by
	// QUEUE, the metrics returned apply to each queue rather than aggregated for
	// all queues. If you group by CHANNEL, you should include a Channels filter.
	// The only supported channel is VOICE.
	//
	// If no Grouping is included in the request, a summary of metrics is returned.
	Groupings []Grouping `type:"list"`

	// The identifier of the Amazon Connect instance.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`

	// The maximimum number of results to return per page.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	//
	// The token expires after 5 minutes from the time it is created. Subsequent
	// requests that use the token must use the same request parameters as the request
	// that generated the token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetCurrentMetricDataInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCurrentMetricDataInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCurrentMetricDataInput"}

	if s.CurrentMetrics == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentMetrics"))
	}

	if s.Filters == nil {
		invalidParams.Add(aws.NewErrParamRequired("Filters"))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Filters != nil {
		if err := s.Filters.Validate(); err != nil {
			invalidParams.AddNested("Filters", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCurrentMetricDataInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CurrentMetrics != nil {
		v := s.CurrentMetrics

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "CurrentMetrics", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Filters != nil {
		v := s.Filters

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Filters", v, metadata)
	}
	if s.Groupings != nil {
		v := s.Groupings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Groupings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
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
	if s.InstanceId != nil {
		v := *s.InstanceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "InstanceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetCurrentMetricDataOutput struct {
	_ struct{} `type:"structure"`

	// The time at which the metrics were retrieved and cached for pagination.
	DataSnapshotTime *time.Time `type:"timestamp"`

	// Information about the real-time metrics.
	MetricResults []CurrentMetricResult `type:"list"`

	// If there are additional results, this is the token for the next set of results.
	//
	// The token expires after 5 minutes from the time it is created. Subsequent
	// requests that use the token must use the same request parameters as the request
	// that generated the token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetCurrentMetricDataOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCurrentMetricDataOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DataSnapshotTime != nil {
		v := *s.DataSnapshotTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DataSnapshotTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.MetricResults != nil {
		v := s.MetricResults

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "MetricResults", metadata)
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

const opGetCurrentMetricData = "GetCurrentMetricData"

// GetCurrentMetricDataRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// Gets the real-time metric data from the specified Amazon Connect instance.
//
// For more information, see Real-time Metrics Reports (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-reports.html)
// in the Amazon Connect Administrator Guide.
//
//    // Example sending a request using GetCurrentMetricDataRequest.
//    req := client.GetCurrentMetricDataRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/GetCurrentMetricData
func (c *Client) GetCurrentMetricDataRequest(input *GetCurrentMetricDataInput) GetCurrentMetricDataRequest {
	op := &aws.Operation{
		Name:       opGetCurrentMetricData,
		HTTPMethod: "POST",
		HTTPPath:   "/metrics/current/{InstanceId}",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetCurrentMetricDataInput{}
	}

	req := c.newRequest(op, input, &GetCurrentMetricDataOutput{})

	return GetCurrentMetricDataRequest{Request: req, Input: input, Copy: c.GetCurrentMetricDataRequest}
}

// GetCurrentMetricDataRequest is the request type for the
// GetCurrentMetricData API operation.
type GetCurrentMetricDataRequest struct {
	*aws.Request
	Input *GetCurrentMetricDataInput
	Copy  func(*GetCurrentMetricDataInput) GetCurrentMetricDataRequest
}

// Send marshals and sends the GetCurrentMetricData API request.
func (r GetCurrentMetricDataRequest) Send(ctx context.Context) (*GetCurrentMetricDataResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCurrentMetricDataResponse{
		GetCurrentMetricDataOutput: r.Request.Data.(*GetCurrentMetricDataOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetCurrentMetricDataRequestPaginator returns a paginator for GetCurrentMetricData.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetCurrentMetricDataRequest(input)
//   p := connect.NewGetCurrentMetricDataRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetCurrentMetricDataPaginator(req GetCurrentMetricDataRequest) GetCurrentMetricDataPaginator {
	return GetCurrentMetricDataPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetCurrentMetricDataInput
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

// GetCurrentMetricDataPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetCurrentMetricDataPaginator struct {
	aws.Pager
}

func (p *GetCurrentMetricDataPaginator) CurrentPage() *GetCurrentMetricDataOutput {
	return p.Pager.CurrentPage().(*GetCurrentMetricDataOutput)
}

// GetCurrentMetricDataResponse is the response type for the
// GetCurrentMetricData API operation.
type GetCurrentMetricDataResponse struct {
	*GetCurrentMetricDataOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCurrentMetricData request.
func (r *GetCurrentMetricDataResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}