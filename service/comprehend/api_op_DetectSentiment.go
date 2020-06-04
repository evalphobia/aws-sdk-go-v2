// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DetectSentimentInput struct {
	_ struct{} `type:"structure"`

	// The language of the input documents. You can specify any of the primary languages
	// supported by Amazon Comprehend. All documents must be in the same language.
	//
	// LanguageCode is a required field
	LanguageCode LanguageCode `type:"string" required:"true" enum:"true"`

	// A UTF-8 text string. Each string must contain fewer that 5,000 bytes of UTF-8
	// encoded characters.
	//
	// Text is a required field
	Text *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s DetectSentimentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetectSentimentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetectSentimentInput"}
	if len(s.LanguageCode) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LanguageCode"))
	}

	if s.Text == nil {
		invalidParams.Add(aws.NewErrParamRequired("Text"))
	}
	if s.Text != nil && len(*s.Text) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Text", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DetectSentimentOutput struct {
	_ struct{} `type:"structure" sensitive:"true"`

	// The inferred sentiment that Amazon Comprehend has the highest level of confidence
	// in.
	Sentiment SentimentType `type:"string" enum:"true"`

	// An object that lists the sentiments, and their corresponding confidence levels.
	SentimentScore *SentimentScore `type:"structure"`
}

// String returns the string representation
func (s DetectSentimentOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetectSentiment = "DetectSentiment"

// DetectSentimentRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Inspects text and returns an inference of the prevailing sentiment (POSITIVE,
// NEUTRAL, MIXED, or NEGATIVE).
//
//    // Example sending a request using DetectSentimentRequest.
//    req := client.DetectSentimentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/DetectSentiment
func (c *Client) DetectSentimentRequest(input *DetectSentimentInput) DetectSentimentRequest {
	op := &aws.Operation{
		Name:       opDetectSentiment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetectSentimentInput{}
	}

	req := c.newRequest(op, input, &DetectSentimentOutput{})

	return DetectSentimentRequest{Request: req, Input: input, Copy: c.DetectSentimentRequest}
}

// DetectSentimentRequest is the request type for the
// DetectSentiment API operation.
type DetectSentimentRequest struct {
	*aws.Request
	Input *DetectSentimentInput
	Copy  func(*DetectSentimentInput) DetectSentimentRequest
}

// Send marshals and sends the DetectSentiment API request.
func (r DetectSentimentRequest) Send(ctx context.Context) (*DetectSentimentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetectSentimentResponse{
		DetectSentimentOutput: r.Request.Data.(*DetectSentimentOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetectSentimentResponse is the response type for the
// DetectSentiment API operation.
type DetectSentimentResponse struct {
	*DetectSentimentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetectSentiment request.
func (r *DetectSentimentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}