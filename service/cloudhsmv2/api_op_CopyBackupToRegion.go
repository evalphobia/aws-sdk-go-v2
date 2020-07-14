// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudhsmv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CopyBackupToRegionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the backup that will be copied to the destination region.
	//
	// BackupId is a required field
	BackupId *string `type:"string" required:"true"`

	// The AWS region that will contain your copied CloudHSM cluster backup.
	//
	// DestinationRegion is a required field
	DestinationRegion *string `type:"string" required:"true"`

	TagList []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s CopyBackupToRegionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CopyBackupToRegionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CopyBackupToRegionInput"}

	if s.BackupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupId"))
	}

	if s.DestinationRegion == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationRegion"))
	}
	if s.TagList != nil && len(s.TagList) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TagList", 1))
	}
	if s.TagList != nil {
		for i, v := range s.TagList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TagList", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CopyBackupToRegionOutput struct {
	_ struct{} `type:"structure"`

	// Information on the backup that will be copied to the destination region,
	// including CreateTimestamp, SourceBackup, SourceCluster, and Source Region.
	// CreateTimestamp of the destination backup will be the same as that of the
	// source backup.
	//
	// You will need to use the sourceBackupID returned in this operation to use
	// the DescribeBackups operation on the backup that will be copied to the destination
	// region.
	DestinationBackup *DestinationBackup `type:"structure"`
}

// String returns the string representation
func (s CopyBackupToRegionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCopyBackupToRegion = "CopyBackupToRegion"

// CopyBackupToRegionRequest returns a request value for making API operation for
// AWS CloudHSM V2.
//
// Copy an AWS CloudHSM cluster backup to a different region.
//
//    // Example sending a request using CopyBackupToRegionRequest.
//    req := client.CopyBackupToRegionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsmv2-2017-04-28/CopyBackupToRegion
func (c *Client) CopyBackupToRegionRequest(input *CopyBackupToRegionInput) CopyBackupToRegionRequest {
	op := &aws.Operation{
		Name:       opCopyBackupToRegion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CopyBackupToRegionInput{}
	}

	req := c.newRequest(op, input, &CopyBackupToRegionOutput{})

	return CopyBackupToRegionRequest{Request: req, Input: input, Copy: c.CopyBackupToRegionRequest}
}

// CopyBackupToRegionRequest is the request type for the
// CopyBackupToRegion API operation.
type CopyBackupToRegionRequest struct {
	*aws.Request
	Input *CopyBackupToRegionInput
	Copy  func(*CopyBackupToRegionInput) CopyBackupToRegionRequest
}

// Send marshals and sends the CopyBackupToRegion API request.
func (r CopyBackupToRegionRequest) Send(ctx context.Context) (*CopyBackupToRegionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CopyBackupToRegionResponse{
		CopyBackupToRegionOutput: r.Request.Data.(*CopyBackupToRegionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CopyBackupToRegionResponse is the response type for the
// CopyBackupToRegion API operation.
type CopyBackupToRegionResponse struct {
	*CopyBackupToRegionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CopyBackupToRegion request.
func (r *CopyBackupToRegionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}