// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for CreateLoadBalancer.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/CreateAccessPointInput
type CreateLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	// One or more Availability Zones from the same region as the load balancer.
	//
	// You must specify at least one Availability Zone.
	//
	// You can add more Availability Zones after you create the load balancer using
	// EnableAvailabilityZonesForLoadBalancer.
	AvailabilityZones []string `type:"list"`

	// The listeners.
	//
	// For more information, see Listeners for Your Classic Load Balancer (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-listener-config.html)
	// in the Classic Load Balancers Guide.
	//
	// Listeners is a required field
	Listeners []Listener `type:"list" required:"true"`

	// The name of the load balancer.
	//
	// This name must be unique within your set of load balancers for the region,
	// must have a maximum of 32 characters, must contain only alphanumeric characters
	// or hyphens, and cannot begin or end with a hyphen.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `type:"string" required:"true"`

	// The type of a load balancer. Valid only for load balancers in a VPC.
	//
	// By default, Elastic Load Balancing creates an Internet-facing load balancer
	// with a DNS name that resolves to public IP addresses. For more information
	// about Internet-facing and Internal load balancers, see Load Balancer Scheme
	// (http://docs.aws.amazon.com/elasticloadbalancing/latest/userguide/how-elastic-load-balancing-works.html#load-balancer-scheme)
	// in the Elastic Load Balancing User Guide.
	//
	// Specify internal to create a load balancer with a DNS name that resolves
	// to private IP addresses.
	Scheme *string `type:"string"`

	// The IDs of the security groups to assign to the load balancer.
	SecurityGroups []string `type:"list"`

	// The IDs of the subnets in your VPC to attach to the load balancer. Specify
	// one subnet per Availability Zone specified in AvailabilityZones.
	Subnets []string `type:"list"`

	// A list of tags to assign to the load balancer.
	//
	// For more information about tagging your load balancer, see Tag Your Classic
	// Load Balancer (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/add-remove-tags.html)
	// in the Classic Load Balancers Guide.
	Tags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateLoadBalancerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLoadBalancerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLoadBalancerInput"}

	if s.Listeners == nil {
		invalidParams.Add(aws.NewErrParamRequired("Listeners"))
	}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.Listeners != nil {
		for i, v := range s.Listeners {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Listeners", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output for CreateLoadBalancer.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/CreateAccessPointOutput
type CreateLoadBalancerOutput struct {
	_ struct{} `type:"structure"`

	// The DNS name of the load balancer.
	DNSName *string `type:"string"`
}

// String returns the string representation
func (s CreateLoadBalancerOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateLoadBalancer = "CreateLoadBalancer"

// CreateLoadBalancerRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Creates a Classic Load Balancer.
//
// You can add listeners, security groups, subnets, and tags when you create
// your load balancer, or you can add them later using CreateLoadBalancerListeners,
// ApplySecurityGroupsToLoadBalancer, AttachLoadBalancerToSubnets, and AddTags.
//
// To describe your current load balancers, see DescribeLoadBalancers. When
// you are finished with a load balancer, you can delete it using DeleteLoadBalancer.
//
// You can create up to 20 load balancers per region per account. You can request
// an increase for the number of load balancers for your account. For more information,
// see Limits for Your Classic Load Balancer (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-limits.html)
// in the Classic Load Balancers Guide.
//
//    // Example sending a request using CreateLoadBalancerRequest.
//    req := client.CreateLoadBalancerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/CreateLoadBalancer
func (c *Client) CreateLoadBalancerRequest(input *CreateLoadBalancerInput) CreateLoadBalancerRequest {
	op := &aws.Operation{
		Name:       opCreateLoadBalancer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateLoadBalancerInput{}
	}

	req := c.newRequest(op, input, &CreateLoadBalancerOutput{})
	return CreateLoadBalancerRequest{Request: req, Input: input, Copy: c.CreateLoadBalancerRequest}
}

// CreateLoadBalancerRequest is the request type for the
// CreateLoadBalancer API operation.
type CreateLoadBalancerRequest struct {
	*aws.Request
	Input *CreateLoadBalancerInput
	Copy  func(*CreateLoadBalancerInput) CreateLoadBalancerRequest
}

// Send marshals and sends the CreateLoadBalancer API request.
func (r CreateLoadBalancerRequest) Send(ctx context.Context) (*CreateLoadBalancerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLoadBalancerResponse{
		CreateLoadBalancerOutput: r.Request.Data.(*CreateLoadBalancerOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLoadBalancerResponse is the response type for the
// CreateLoadBalancer API operation.
type CreateLoadBalancerResponse struct {
	*CreateLoadBalancerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLoadBalancer request.
func (r *CreateLoadBalancerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}