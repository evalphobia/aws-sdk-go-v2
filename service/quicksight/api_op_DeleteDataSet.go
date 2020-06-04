// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteDataSetInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID for the dataset that you want to create. This ID is unique per AWS
	// Region for each AWS account.
	//
	// DataSetId is a required field
	DataSetId *string `location:"uri" locationName:"DataSetId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDataSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDataSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDataSetInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.DataSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDataSetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DataSetId != nil {
		v := *s.DataSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DataSetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteDataSetOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dataset.
	Arn *string `type:"string"`

	// The ID for the dataset that you want to create. This ID is unique per AWS
	// Region for each AWS account.
	DataSetId *string `type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s DeleteDataSetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDataSetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DataSetId != nil {
		v := *s.DataSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DataSetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opDeleteDataSet = "DeleteDataSet"

// DeleteDataSetRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Deletes a dataset.
//
//    // Example sending a request using DeleteDataSetRequest.
//    req := client.DeleteDataSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DeleteDataSet
func (c *Client) DeleteDataSetRequest(input *DeleteDataSetInput) DeleteDataSetRequest {
	op := &aws.Operation{
		Name:       opDeleteDataSet,
		HTTPMethod: "DELETE",
		HTTPPath:   "/accounts/{AwsAccountId}/data-sets/{DataSetId}",
	}

	if input == nil {
		input = &DeleteDataSetInput{}
	}

	req := c.newRequest(op, input, &DeleteDataSetOutput{})

	return DeleteDataSetRequest{Request: req, Input: input, Copy: c.DeleteDataSetRequest}
}

// DeleteDataSetRequest is the request type for the
// DeleteDataSet API operation.
type DeleteDataSetRequest struct {
	*aws.Request
	Input *DeleteDataSetInput
	Copy  func(*DeleteDataSetInput) DeleteDataSetRequest
}

// Send marshals and sends the DeleteDataSet API request.
func (r DeleteDataSetRequest) Send(ctx context.Context) (*DeleteDataSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDataSetResponse{
		DeleteDataSetOutput: r.Request.Data.(*DeleteDataSetOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDataSetResponse is the response type for the
// DeleteDataSet API operation.
type DeleteDataSetResponse struct {
	*DeleteDataSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDataSet request.
func (r *DeleteDataSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}