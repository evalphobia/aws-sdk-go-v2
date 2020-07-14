// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecast

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDatasetImportJobInput struct {
	_ struct{} `type:"structure"`

	// The location of the training data to import and an AWS Identity and Access
	// Management (IAM) role that Amazon Forecast can assume to access the data.
	// The training data must be stored in an Amazon S3 bucket.
	//
	// If encryption is used, DataSource must include an AWS Key Management Service
	// (KMS) key and the IAM role must allow Amazon Forecast permission to access
	// the key. The KMS key and IAM role must match those specified in the EncryptionConfig
	// parameter of the CreateDataset operation.
	//
	// DataSource is a required field
	DataSource *DataSource `type:"structure" required:"true"`

	// The Amazon Resource Name (ARN) of the Amazon Forecast dataset that you want
	// to import data to.
	//
	// DatasetArn is a required field
	DatasetArn *string `type:"string" required:"true"`

	// The name for the dataset import job. We recommend including the current timestamp
	// in the name, for example, 20190721DatasetImport. This can help you avoid
	// getting a ResourceAlreadyExistsException exception.
	//
	// DatasetImportJobName is a required field
	DatasetImportJobName *string `min:"1" type:"string" required:"true"`

	// The format of timestamps in the dataset. The format that you specify depends
	// on the DataFrequency specified when the dataset was created. The following
	// formats are supported
	//
	//    * "yyyy-MM-dd" For the following data frequencies: Y, M, W, and D
	//
	//    * "yyyy-MM-dd HH:mm:ss" For the following data frequencies: H, 30min,
	//    15min, and 1min; and optionally, for: Y, M, W, and D
	//
	// If the format isn't specified, Amazon Forecast expects the format to be "yyyy-MM-dd
	// HH:mm:ss".
	TimestampFormat *string `type:"string"`
}

// String returns the string representation
func (s CreateDatasetImportJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDatasetImportJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDatasetImportJobInput"}

	if s.DataSource == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSource"))
	}

	if s.DatasetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetArn"))
	}

	if s.DatasetImportJobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetImportJobName"))
	}
	if s.DatasetImportJobName != nil && len(*s.DatasetImportJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatasetImportJobName", 1))
	}
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			invalidParams.AddNested("DataSource", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDatasetImportJobOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dataset import job.
	DatasetImportJobArn *string `type:"string"`
}

// String returns the string representation
func (s CreateDatasetImportJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDatasetImportJob = "CreateDatasetImportJob"

// CreateDatasetImportJobRequest returns a request value for making API operation for
// Amazon Forecast Service.
//
// Imports your training data to an Amazon Forecast dataset. You provide the
// location of your training data in an Amazon Simple Storage Service (Amazon
// S3) bucket and the Amazon Resource Name (ARN) of the dataset that you want
// to import the data to.
//
// You must specify a DataSource object that includes an AWS Identity and Access
// Management (IAM) role that Amazon Forecast can assume to access the data.
// For more information, see aws-forecast-iam-roles.
//
// The training data must be in CSV format. The delimiter must be a comma (,).
//
// You can specify the path to a specific CSV file, the S3 bucket, or to a folder
// in the S3 bucket. For the latter two cases, Amazon Forecast imports all files
// up to the limit of 10,000 files.
//
// To get a list of all your dataset import jobs, filtered by specified criteria,
// use the ListDatasetImportJobs operation.
//
//    // Example sending a request using CreateDatasetImportJobRequest.
//    req := client.CreateDatasetImportJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecast-2018-06-26/CreateDatasetImportJob
func (c *Client) CreateDatasetImportJobRequest(input *CreateDatasetImportJobInput) CreateDatasetImportJobRequest {
	op := &aws.Operation{
		Name:       opCreateDatasetImportJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDatasetImportJobInput{}
	}

	req := c.newRequest(op, input, &CreateDatasetImportJobOutput{})

	return CreateDatasetImportJobRequest{Request: req, Input: input, Copy: c.CreateDatasetImportJobRequest}
}

// CreateDatasetImportJobRequest is the request type for the
// CreateDatasetImportJob API operation.
type CreateDatasetImportJobRequest struct {
	*aws.Request
	Input *CreateDatasetImportJobInput
	Copy  func(*CreateDatasetImportJobInput) CreateDatasetImportJobRequest
}

// Send marshals and sends the CreateDatasetImportJob API request.
func (r CreateDatasetImportJobRequest) Send(ctx context.Context) (*CreateDatasetImportJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDatasetImportJobResponse{
		CreateDatasetImportJobOutput: r.Request.Data.(*CreateDatasetImportJobOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDatasetImportJobResponse is the response type for the
// CreateDatasetImportJob API operation.
type CreateDatasetImportJobResponse struct {
	*CreateDatasetImportJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDatasetImportJob request.
func (r *CreateDatasetImportJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}