// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemakera2iruntime

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Attributes of the data specified by the customer. Use these to describe the
// data to be labeled.
type HumanLoopDataAttributes struct {
	_ struct{} `type:"structure"`

	// Declares that your content is free of personally identifiable information
	// or adult content.
	//
	// Amazon SageMaker can restrict the Amazon Mechanical Turk workers who can
	// view your task based on this information.
	//
	// ContentClassifiers is a required field
	ContentClassifiers []ContentClassifier `type:"list" required:"true"`
}

// String returns the string representation
func (s HumanLoopDataAttributes) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *HumanLoopDataAttributes) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "HumanLoopDataAttributes"}

	if s.ContentClassifiers == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContentClassifiers"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s HumanLoopDataAttributes) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContentClassifiers != nil {
		v := s.ContentClassifiers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ContentClassifiers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// An object containing the human loop input in JSON format.
type HumanLoopInput struct {
	_ struct{} `type:"structure"`

	// Serialized input from the human loop. The input must be a string representation
	// of a file in JSON format.
	//
	// InputContent is a required field
	InputContent *string `type:"string" required:"true"`
}

// String returns the string representation
func (s HumanLoopInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *HumanLoopInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "HumanLoopInput"}

	if s.InputContent == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputContent"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s HumanLoopInput) MarshalFields(e protocol.FieldEncoder) error {
	if s.InputContent != nil {
		v := *s.InputContent

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "InputContent", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about where the human output will be stored.
type HumanLoopOutput struct {
	_ struct{} `type:"structure"`

	// The location of the Amazon S3 object where Amazon Augmented AI stores your
	// human loop output.
	//
	// OutputS3Uri is a required field
	OutputS3Uri *string `type:"string" required:"true"`
}

// String returns the string representation
func (s HumanLoopOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s HumanLoopOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.OutputS3Uri != nil {
		v := *s.OutputS3Uri

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OutputS3Uri", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Summary information about the human loop.
type HumanLoopSummary struct {
	_ struct{} `type:"structure"`

	// When Amazon Augmented AI created the human loop.
	CreationTime *time.Time `type:"timestamp"`

	// The reason why the human loop failed. A failure reason is returned when the
	// status of the human loop is Failed.
	FailureReason *string `type:"string"`

	// The Amazon Resource Name (ARN) of the flow definition used to configure the
	// human loop.
	FlowDefinitionArn *string `type:"string"`

	// The name of the human loop.
	HumanLoopName *string `min:"1" type:"string"`

	// The status of the human loop.
	HumanLoopStatus HumanLoopStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s HumanLoopSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s HumanLoopSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.FailureReason != nil {
		v := *s.FailureReason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FailureReason", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FlowDefinitionArn != nil {
		v := *s.FlowDefinitionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FlowDefinitionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HumanLoopName != nil {
		v := *s.HumanLoopName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "HumanLoopName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.HumanLoopStatus) > 0 {
		v := s.HumanLoopStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "HumanLoopStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}