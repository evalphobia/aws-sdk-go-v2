// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateAssociationInput struct {
	_ struct{} `type:"structure"`

	// Specify a descriptive name for the association.
	AssociationName *string `type:"string"`

	// Specify the target for the association. This target is required for associations
	// that use an Automation document and target resources by using rate controls.
	AutomationTargetParameterName *string `min:"1" type:"string"`

	// The severity level to assign to the association.
	ComplianceSeverity AssociationComplianceSeverity `type:"string" enum:"true"`

	// The document version you want to associate with the target(s). Can be a specific
	// version or the default version.
	DocumentVersion *string `type:"string"`

	// The instance ID.
	//
	// InstanceId has been deprecated. To specify an instance ID for an association,
	// use the Targets parameter. Requests that include the parameter InstanceID
	// with SSM documents that use schema version 2.0 or later will fail. In addition,
	// if you use the parameter InstanceId, you cannot use the parameters AssociationName,
	// DocumentVersion, MaxErrors, MaxConcurrency, OutputLocation, or ScheduleExpression.
	// To use these parameters, you must use the Targets parameter.
	InstanceId *string `type:"string"`

	// The maximum number of targets allowed to run the association at the same
	// time. You can specify a number, for example 10, or a percentage of the target
	// set, for example 10%. The default value is 100%, which means all targets
	// run the association at the same time.
	//
	// If a new instance starts and attempts to run an association while Systems
	// Manager is running MaxConcurrency associations, the association is allowed
	// to run. During the next association interval, the new instance will process
	// its association within the limit specified for MaxConcurrency.
	MaxConcurrency *string `min:"1" type:"string"`

	// The number of errors that are allowed before the system stops sending requests
	// to run the association on additional targets. You can specify either an absolute
	// number of errors, for example 10, or a percentage of the target set, for
	// example 10%. If you specify 3, for example, the system stops sending requests
	// when the fourth error is received. If you specify 0, then the system stops
	// sending requests after the first error is returned. If you run an association
	// on 50 instances and set MaxError to 10%, then the system stops sending the
	// request when the sixth error is received.
	//
	// Executions that are already running an association when MaxErrors is reached
	// are allowed to complete, but some of these executions may fail as well. If
	// you need to ensure that there won't be more than max-errors failed executions,
	// set MaxConcurrency to 1 so that executions proceed one at a time.
	MaxErrors *string `min:"1" type:"string"`

	// The name of the SSM document that contains the configuration information
	// for the instance. You can specify Command or Automation documents.
	//
	// You can specify AWS-predefined documents, documents you created, or a document
	// that is shared with you from another account.
	//
	// For SSM documents that are shared with you from other AWS accounts, you must
	// specify the complete SSM document ARN, in the following format:
	//
	// arn:partition:ssm:region:account-id:document/document-name
	//
	// For example:
	//
	// arn:aws:ssm:us-east-2:12345678912:document/My-Shared-Document
	//
	// For AWS-predefined documents and SSM documents you created in your account,
	// you only need to specify the document name. For example, AWS-ApplyPatchBaseline
	// or My-Document.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// An S3 bucket where you want to store the output details of the request.
	OutputLocation *InstanceAssociationOutputLocation `type:"structure"`

	// The parameters for the runtime configuration of the document.
	Parameters map[string][]string `type:"map"`

	// A cron expression when the association will be applied to the target(s).
	ScheduleExpression *string `min:"1" type:"string"`

	// The mode for generating association compliance. You can specify AUTO or MANUAL.
	// In AUTO mode, the system uses the status of the association execution to
	// determine the compliance status. If the association execution runs successfully,
	// then the association is COMPLIANT. If the association execution doesn't run
	// successfully, the association is NON-COMPLIANT.
	//
	// In MANUAL mode, you must specify the AssociationId as a parameter for the
	// PutComplianceItems API action. In this case, compliance data is not managed
	// by State Manager. It is managed by your direct call to the PutComplianceItems
	// API action.
	//
	// By default, all associations use AUTO mode.
	SyncCompliance AssociationSyncCompliance `type:"string" enum:"true"`

	// The targets for the association. You can target instances by using tags,
	// AWS Resource Groups, all instances in an AWS account, or individual instance
	// IDs. For more information about choosing targets for an association, see
	// Using targets and rate controls with State Manager associations (https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-state-manager-targets-and-rate-controls.html)
	// in the AWS Systems Manager User Guide.
	Targets []Target `type:"list"`
}

// String returns the string representation
func (s CreateAssociationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAssociationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAssociationInput"}
	if s.AutomationTargetParameterName != nil && len(*s.AutomationTargetParameterName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutomationTargetParameterName", 1))
	}
	if s.MaxConcurrency != nil && len(*s.MaxConcurrency) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MaxConcurrency", 1))
	}
	if s.MaxErrors != nil && len(*s.MaxErrors) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MaxErrors", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.ScheduleExpression != nil && len(*s.ScheduleExpression) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ScheduleExpression", 1))
	}
	if s.OutputLocation != nil {
		if err := s.OutputLocation.Validate(); err != nil {
			invalidParams.AddNested("OutputLocation", err.(aws.ErrInvalidParams))
		}
	}
	if s.Targets != nil {
		for i, v := range s.Targets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Targets", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateAssociationOutput struct {
	_ struct{} `type:"structure"`

	// Information about the association.
	AssociationDescription *AssociationDescription `type:"structure"`
}

// String returns the string representation
func (s CreateAssociationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateAssociation = "CreateAssociation"

// CreateAssociationRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Associates the specified Systems Manager document with the specified instances
// or targets.
//
// When you associate a document with one or more instances, SSM Agent running
// on the instance processes the document and configures the instance as specified.
// If you associate a document with an instance that already has an associated
// document, the system returns the AssociationAlreadyExists exception.
//
//    // Example sending a request using CreateAssociationRequest.
//    req := client.CreateAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/CreateAssociation
func (c *Client) CreateAssociationRequest(input *CreateAssociationInput) CreateAssociationRequest {
	op := &aws.Operation{
		Name:       opCreateAssociation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateAssociationInput{}
	}

	req := c.newRequest(op, input, &CreateAssociationOutput{})

	return CreateAssociationRequest{Request: req, Input: input, Copy: c.CreateAssociationRequest}
}

// CreateAssociationRequest is the request type for the
// CreateAssociation API operation.
type CreateAssociationRequest struct {
	*aws.Request
	Input *CreateAssociationInput
	Copy  func(*CreateAssociationInput) CreateAssociationRequest
}

// Send marshals and sends the CreateAssociation API request.
func (r CreateAssociationRequest) Send(ctx context.Context) (*CreateAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAssociationResponse{
		CreateAssociationOutput: r.Request.Data.(*CreateAssociationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAssociationResponse is the response type for the
// CreateAssociation API operation.
type CreateAssociationResponse struct {
	*CreateAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAssociation request.
func (r *CreateAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}