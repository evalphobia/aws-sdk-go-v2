// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build go1.10,integration

package kinesis_test

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/integration"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

var _ aws.Config
var _ awserr.Error

func TestInteg_00_ListStreams(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := integration.ConfigWithDefaultRegion("us-west-2")
	svc := kinesis.New(cfg)
	params := &kinesis.ListStreamsInput{}

	req := svc.ListStreamsRequest(params)

	_, err := req.Send(ctx)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
func TestInteg_01_DescribeStream(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := integration.ConfigWithDefaultRegion("us-west-2")
	svc := kinesis.New(cfg)
	params := &kinesis.DescribeStreamInput{
		StreamName: aws.String("bogus-stream-name"),
	}

	req := svc.DescribeStreamRequest(params)

	_, err := req.Send(ctx)
	if err == nil {
		t.Fatalf("expect request to fail")
	}
	aerr, ok := err.(awserr.RequestFailure)
	if !ok {
		t.Fatalf("expect awserr, was %T", err)
	}
	if len(aerr.Code()) == 0 {
		t.Errorf("expect non-empty error code")
	}
	if v := aerr.Code(); v == aws.ErrCodeSerialization {
		t.Errorf("expect API error code got serialization failure")
	}
}