// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type ExecutePolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	AutoScalingGroupName *string `min:"1" type:"string"`

	// The breach threshold for the alarm.
	//
	// Conditional: This parameter is required if the policy type is StepScaling
	// and not supported otherwise.
	BreachThreshold *float64 `type:"double"`

	// Indicates whether Amazon EC2 Auto Scaling waits for the cooldown period to
	// complete before executing the policy.
	//
	// This parameter is not supported if the policy type is StepScaling or TargetTrackingScaling.
	//
	// For more information, see Scaling Cooldowns (https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	HonorCooldown *bool `type:"boolean"`

	// The metric value to compare to BreachThreshold. This enables you to execute
	// a policy of type StepScaling and determine which step adjustment to use.
	// For example, if the breach threshold is 50 and you want to use a step adjustment
	// with a lower bound of 0 and an upper bound of 10, you can set the metric
	// value to 59.
	//
	// If you specify a metric value that doesn't correspond to a step adjustment
	// for the policy, the call returns an error.
	//
	// Conditional: This parameter is required if the policy type is StepScaling
	// and not supported otherwise.
	MetricValue *float64 `type:"double"`

	// The name or ARN of the policy.
	//
	// PolicyName is a required field
	PolicyName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ExecutePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExecutePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ExecutePolicyInput"}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}

	if s.PolicyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyName != nil && len(*s.PolicyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ExecutePolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ExecutePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opExecutePolicy = "ExecutePolicy"

// ExecutePolicyRequest returns a request value for making API operation for
// Auto Scaling.
//
// Executes the specified policy.
//
//    // Example sending a request using ExecutePolicyRequest.
//    req := client.ExecutePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/ExecutePolicy
func (c *Client) ExecutePolicyRequest(input *ExecutePolicyInput) ExecutePolicyRequest {
	op := &aws.Operation{
		Name:       opExecutePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ExecutePolicyInput{}
	}

	req := c.newRequest(op, input, &ExecutePolicyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return ExecutePolicyRequest{Request: req, Input: input, Copy: c.ExecutePolicyRequest}
}

// ExecutePolicyRequest is the request type for the
// ExecutePolicy API operation.
type ExecutePolicyRequest struct {
	*aws.Request
	Input *ExecutePolicyInput
	Copy  func(*ExecutePolicyInput) ExecutePolicyRequest
}

// Send marshals and sends the ExecutePolicy API request.
func (r ExecutePolicyRequest) Send(ctx context.Context) (*ExecutePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ExecutePolicyResponse{
		ExecutePolicyOutput: r.Request.Data.(*ExecutePolicyOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ExecutePolicyResponse is the response type for the
// ExecutePolicy API operation.
type ExecutePolicyResponse struct {
	*ExecutePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ExecutePolicy request.
func (r *ExecutePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}