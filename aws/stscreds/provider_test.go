package stscreds

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type stubSTS struct {
	TestInput func(*sts.AssumeRoleInput)
}

func (s *stubSTS) AssumeRoleRequest(input *sts.AssumeRoleInput) sts.AssumeRoleRequest {
	if s.TestInput != nil {
		s.TestInput(input)
	}
	expiry := time.Now().Add(60 * time.Minute)

	req := sts.AssumeRoleRequest{
		Input: input,
		Request: &aws.Request{
			HTTPRequest: func() *http.Request {
				r, _ := http.NewRequest("POST", "https://sts.amazonaws.com", nil)
				return r
			}(),
			Retryer: retry.NewStandard(),
			Handlers: func() aws.Handlers {
				h := aws.Handlers{}

				h.Send.PushBack(func(r *aws.Request) {
					r.HTTPResponse = &http.Response{StatusCode: 200}
					r.Data = &sts.AssumeRoleOutput{
						Credentials: &sts.Credentials{
							// Just reflect the role arn to the provider.
							AccessKeyId:     input.RoleArn,
							SecretAccessKey: aws.String("assumedSecretAccessKey"),
							SessionToken:    aws.String("assumedSessionToken"),
							Expiration:      &expiry,
						},
					}
				})
				return h
			}(),
		},
	}

	return req
}

const roleARN = "00000000000000000000000000000000000"
const tokenCode = "00000000000000000000"

func TestAssumeRoleProvider(t *testing.T) {
	stub := &stubSTS{}
	p := NewAssumeRoleProvider(stub, roleARN)

	creds, err := p.Retrieve(context.Background())
	if err != nil {
		t.Fatalf("Expect no error, %v", err)
	}

	if e, a := roleARN, creds.AccessKeyID; e != a {
		t.Errorf("Expect access key ID to be reflected role ARN")
	}
	if e, a := "assumedSecretAccessKey", creds.SecretAccessKey; e != a {
		t.Errorf("Expect secret access key to match")
	}
	if e, a := "assumedSessionToken", creds.SessionToken; e != a {
		t.Errorf("Expect session token to match")
	}
}

func TestAssumeRoleProvider_WithTokenProvider(t *testing.T) {
	stub := &stubSTS{
		TestInput: func(in *sts.AssumeRoleInput) {
			if e, a := "0123456789", *in.SerialNumber; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
			if e, a := tokenCode, *in.TokenCode; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
		},
	}
	p := NewAssumeRoleProvider(stub, roleARN, func(options *AssumeRoleProviderOptions) {
		options.SerialNumber = aws.String("0123456789")
		options.TokenProvider = func() (string, error) {
			return tokenCode, nil
		}
	})

	creds, err := p.Retrieve(context.Background())
	if err != nil {
		t.Fatalf("Expect no error, %v", err)
	}

	if e, a := roleARN, creds.AccessKeyID; e != a {
		t.Errorf("Expect access key ID to be reflected role ARN")
	}
	if e, a := "assumedSecretAccessKey", creds.SecretAccessKey; e != a {
		t.Errorf("Expect secret access key to match")
	}
	if e, a := "assumedSessionToken", creds.SessionToken; e != a {
		t.Errorf("Expect session token to match")
	}
}

func TestAssumeRoleProvider_WithTokenProviderError(t *testing.T) {
	stub := &stubSTS{
		TestInput: func(in *sts.AssumeRoleInput) {
			t.Fatalf("API request should not of been called")
		},
	}
	p := NewAssumeRoleProvider(stub, roleARN, func(options *AssumeRoleProviderOptions) {
		options.SerialNumber = aws.String("0123456789")
		options.TokenProvider = func() (string, error) {
			return "", fmt.Errorf("error occurred")
		}
	})

	creds, err := p.Retrieve(context.Background())
	if err == nil {
		t.Fatalf("expect error, got none")
	}

	if v := creds.AccessKeyID; len(v) != 0 {
		t.Errorf("expect zero, got %v", v)
	}
	if v := creds.SecretAccessKey; len(v) != 0 {
		t.Errorf("expect zero, got %v", v)
	}
	if v := creds.SessionToken; len(v) != 0 {
		t.Errorf("expect zero, got %v", v)
	}
}

func TestAssumeRoleProvider_MFAWithNoToken(t *testing.T) {
	stub := &stubSTS{
		TestInput: func(in *sts.AssumeRoleInput) {
			t.Fatalf("API request should not of been called")
		},
	}
	p := NewAssumeRoleProvider(stub, roleARN, func(options *AssumeRoleProviderOptions) {
		options.SerialNumber = aws.String("0123456789")
	})

	creds, err := p.Retrieve(context.Background())
	if err == nil {
		t.Fatalf("expect error, got none")
	}

	if v := creds.AccessKeyID; len(v) != 0 {
		t.Errorf("expect zero, got %v", v)
	}
	if v := creds.SecretAccessKey; len(v) != 0 {
		t.Errorf("expect zero, got %v", v)
	}
	if v := creds.SessionToken; len(v) != 0 {
		t.Errorf("expect zero, got %v", v)
	}
}

func BenchmarkAssumeRoleProvider(b *testing.B) {
	stub := &stubSTS{}
	p := NewAssumeRoleProvider(stub, roleARN)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := p.Retrieve(context.Background()); err != nil {
			b.Fatal(err)
		}
	}
}
