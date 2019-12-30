// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sso

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Provides information about your AWS account.
type AccountInfo struct {
	_ struct{} `type:"structure"`

	// The identifier of the AWS account that is assigned to the user.
	AccountId *string `locationName:"accountId" type:"string"`

	// The display name of the AWS account that is assigned to the user.
	AccountName *string `locationName:"accountName" type:"string"`

	// The email address of the AWS account that is assigned to the user.
	EmailAddress *string `locationName:"emailAddress" min:"1" type:"string"`
}

// String returns the string representation
func (s AccountInfo) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AccountInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AccountName != nil {
		v := *s.AccountName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "accountName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EmailAddress != nil {
		v := *s.EmailAddress

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "emailAddress", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Provides information about the role credentials that are assigned to the
// user.
type RoleCredentials struct {
	_ struct{} `type:"structure"`

	// The identifier used for the temporary security credentials. For more information,
	// see Using Temporary Security Credentials to Request Access to AWS Resources
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html)
	// in the AWS IAM User Guide.
	AccessKeyId *string `locationName:"accessKeyId" type:"string"`

	// The date on which temporary security credentials expire.
	Expiration *int64 `locationName:"expiration" type:"long"`

	// The key that is used to sign the request. For more information, see Using
	// Temporary Security Credentials to Request Access to AWS Resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html)
	// in the AWS IAM User Guide.
	SecretAccessKey *string `locationName:"secretAccessKey" type:"string" sensitive:"true"`

	// The token used for temporary credentials. For more information, see Using
	// Temporary Security Credentials to Request Access to AWS Resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html)
	// in the AWS IAM User Guide.
	SessionToken *string `locationName:"sessionToken" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s RoleCredentials) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RoleCredentials) MarshalFields(e protocol.FieldEncoder) error {
	if s.AccessKeyId != nil {
		v := *s.AccessKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "accessKeyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Expiration != nil {
		v := *s.Expiration

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "expiration", protocol.Int64Value(v), metadata)
	}
	if s.SecretAccessKey != nil {
		v := *s.SecretAccessKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "secretAccessKey", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SessionToken != nil {
		v := *s.SessionToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sessionToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Provides information about the role that is assigned to the user.
type RoleInfo struct {
	_ struct{} `type:"structure"`

	// The identifier of the AWS account assigned to the user.
	AccountId *string `locationName:"accountId" type:"string"`

	// The friendly name of the role that is assigned to the user.
	RoleName *string `locationName:"roleName" type:"string"`
}

// String returns the string representation
func (s RoleInfo) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RoleInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RoleName != nil {
		v := *s.RoleName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}