// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type RegisterCertificateWithoutCAInput struct {
	_ struct{} `type:"structure"`

	// The certificate data, in PEM format.
	//
	// CertificatePem is a required field
	CertificatePem *string `locationName:"certificatePem" min:"1" type:"string" required:"true"`

	// The status of the register certificate request.
	Status CertificateStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s RegisterCertificateWithoutCAInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterCertificateWithoutCAInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterCertificateWithoutCAInput"}

	if s.CertificatePem == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificatePem"))
	}
	if s.CertificatePem != nil && len(*s.CertificatePem) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificatePem", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RegisterCertificateWithoutCAInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CertificatePem != nil {
		v := *s.CertificatePem

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificatePem", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type RegisterCertificateWithoutCAOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the registered certificate.
	CertificateArn *string `locationName:"certificateArn" type:"string"`

	// The ID of the registered certificate. (The last part of the certificate ARN
	// contains the certificate ID.
	CertificateId *string `locationName:"certificateId" min:"64" type:"string"`
}

// String returns the string representation
func (s RegisterCertificateWithoutCAOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RegisterCertificateWithoutCAOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CertificateArn != nil {
		v := *s.CertificateArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificateArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CertificateId != nil {
		v := *s.CertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opRegisterCertificateWithoutCA = "RegisterCertificateWithoutCA"

// RegisterCertificateWithoutCARequest returns a request value for making API operation for
// AWS IoT.
//
// Register a certificate that does not have a certificate authority (CA).
//
//    // Example sending a request using RegisterCertificateWithoutCARequest.
//    req := client.RegisterCertificateWithoutCARequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) RegisterCertificateWithoutCARequest(input *RegisterCertificateWithoutCAInput) RegisterCertificateWithoutCARequest {
	op := &aws.Operation{
		Name:       opRegisterCertificateWithoutCA,
		HTTPMethod: "POST",
		HTTPPath:   "/certificate/register-no-ca",
	}

	if input == nil {
		input = &RegisterCertificateWithoutCAInput{}
	}

	req := c.newRequest(op, input, &RegisterCertificateWithoutCAOutput{})

	return RegisterCertificateWithoutCARequest{Request: req, Input: input, Copy: c.RegisterCertificateWithoutCARequest}
}

// RegisterCertificateWithoutCARequest is the request type for the
// RegisterCertificateWithoutCA API operation.
type RegisterCertificateWithoutCARequest struct {
	*aws.Request
	Input *RegisterCertificateWithoutCAInput
	Copy  func(*RegisterCertificateWithoutCAInput) RegisterCertificateWithoutCARequest
}

// Send marshals and sends the RegisterCertificateWithoutCA API request.
func (r RegisterCertificateWithoutCARequest) Send(ctx context.Context) (*RegisterCertificateWithoutCAResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterCertificateWithoutCAResponse{
		RegisterCertificateWithoutCAOutput: r.Request.Data.(*RegisterCertificateWithoutCAOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterCertificateWithoutCAResponse is the response type for the
// RegisterCertificateWithoutCA API operation.
type RegisterCertificateWithoutCAResponse struct {
	*RegisterCertificateWithoutCAOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterCertificateWithoutCA request.
func (r *RegisterCertificateWithoutCAResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
