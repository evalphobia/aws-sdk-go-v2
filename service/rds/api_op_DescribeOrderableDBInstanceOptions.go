// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeOrderableDBInstanceOptionsInput struct {
	_ struct{} `type:"structure"`

	// The Availability Zone group associated with a Local Zone. Specify this parameter
	// to retrieve available offerings for the Local Zones in the group.
	//
	// Omit this parameter to show the available offerings in the specified AWS
	// Region.
	AvailabilityZoneGroup *string `type:"string"`

	// The DB instance class filter value. Specify this parameter to show only the
	// available offerings matching the specified DB instance class.
	DBInstanceClass *string `type:"string"`

	// The name of the engine to retrieve DB instance options for.
	//
	// Engine is a required field
	Engine *string `type:"string" required:"true"`

	// The engine version filter value. Specify this parameter to show only the
	// available offerings matching the specified engine version.
	EngineVersion *string `type:"string"`

	// This parameter isn't currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// The license model filter value. Specify this parameter to show only the available
	// offerings matching the specified license model.
	LicenseModel *string `type:"string"`

	// An optional pagination token provided by a previous DescribeOrderableDBInstanceOptions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// A value that indicates whether to show only VPC or non-VPC offerings.
	Vpc *bool `type:"boolean"`
}

// String returns the string representation
func (s DescribeOrderableDBInstanceOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeOrderableDBInstanceOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeOrderableDBInstanceOptionsInput"}

	if s.Engine == nil {
		invalidParams.Add(aws.NewErrParamRequired("Engine"))
	}
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

// Contains the result of a successful invocation of the DescribeOrderableDBInstanceOptions
// action.
type DescribeOrderableDBInstanceOptionsOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous OrderableDBInstanceOptions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string `type:"string"`

	// An OrderableDBInstanceOption structure containing information about orderable
	// options for the DB instance.
	OrderableDBInstanceOptions []OrderableDBInstanceOption `locationNameList:"OrderableDBInstanceOption" type:"list"`
}

// String returns the string representation
func (s DescribeOrderableDBInstanceOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeOrderableDBInstanceOptions = "DescribeOrderableDBInstanceOptions"

// DescribeOrderableDBInstanceOptionsRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Returns a list of orderable DB instance options for the specified engine.
//
//    // Example sending a request using DescribeOrderableDBInstanceOptionsRequest.
//    req := client.DescribeOrderableDBInstanceOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeOrderableDBInstanceOptions
func (c *Client) DescribeOrderableDBInstanceOptionsRequest(input *DescribeOrderableDBInstanceOptionsInput) DescribeOrderableDBInstanceOptionsRequest {
	op := &aws.Operation{
		Name:       opDescribeOrderableDBInstanceOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeOrderableDBInstanceOptionsInput{}
	}

	req := c.newRequest(op, input, &DescribeOrderableDBInstanceOptionsOutput{})

	return DescribeOrderableDBInstanceOptionsRequest{Request: req, Input: input, Copy: c.DescribeOrderableDBInstanceOptionsRequest}
}

// DescribeOrderableDBInstanceOptionsRequest is the request type for the
// DescribeOrderableDBInstanceOptions API operation.
type DescribeOrderableDBInstanceOptionsRequest struct {
	*aws.Request
	Input *DescribeOrderableDBInstanceOptionsInput
	Copy  func(*DescribeOrderableDBInstanceOptionsInput) DescribeOrderableDBInstanceOptionsRequest
}

// Send marshals and sends the DescribeOrderableDBInstanceOptions API request.
func (r DescribeOrderableDBInstanceOptionsRequest) Send(ctx context.Context) (*DescribeOrderableDBInstanceOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeOrderableDBInstanceOptionsResponse{
		DescribeOrderableDBInstanceOptionsOutput: r.Request.Data.(*DescribeOrderableDBInstanceOptionsOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeOrderableDBInstanceOptionsRequestPaginator returns a paginator for DescribeOrderableDBInstanceOptions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeOrderableDBInstanceOptionsRequest(input)
//   p := rds.NewDescribeOrderableDBInstanceOptionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeOrderableDBInstanceOptionsPaginator(req DescribeOrderableDBInstanceOptionsRequest) DescribeOrderableDBInstanceOptionsPaginator {
	return DescribeOrderableDBInstanceOptionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeOrderableDBInstanceOptionsInput
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

// DescribeOrderableDBInstanceOptionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeOrderableDBInstanceOptionsPaginator struct {
	aws.Pager
}

func (p *DescribeOrderableDBInstanceOptionsPaginator) CurrentPage() *DescribeOrderableDBInstanceOptionsOutput {
	return p.Pager.CurrentPage().(*DescribeOrderableDBInstanceOptionsOutput)
}

// DescribeOrderableDBInstanceOptionsResponse is the response type for the
// DescribeOrderableDBInstanceOptions API operation.
type DescribeOrderableDBInstanceOptionsResponse struct {
	*DescribeOrderableDBInstanceOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeOrderableDBInstanceOptions request.
func (r *DescribeOrderableDBInstanceOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
