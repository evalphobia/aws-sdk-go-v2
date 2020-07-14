// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type CreateFleetInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for a build to be deployed on the new fleet. You can
	// use either the build ID or ARN value. The custom game server build must have
	// been successfully uploaded to Amazon GameLift and be in a READY status. This
	// fleet setting cannot be changed once the fleet is created.
	BuildId *string `type:"string"`

	// Indicates whether to generate a TLS/SSL certificate for the new fleet. TLS
	// certificates are used for encrypting traffic between game clients and game
	// servers running on GameLift. If this parameter is not specified, the default
	// value, DISABLED, is used. This fleet setting cannot be changed once the fleet
	// is created. Learn more at Securing Client/Server Communication (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-howitworks.html#gamelift-howitworks-security).
	//
	// Note: This feature requires the AWS Certificate Manager (ACM) service, which
	// is available in the AWS global partition but not in all other partitions.
	// When working in a partition that does not support this feature, a request
	// for a new fleet with certificate generation results fails with a 4xx unsupported
	// Region error.
	//
	// Valid values include:
	//
	//    * GENERATED - Generate a TLS/SSL certificate for this fleet.
	//
	//    * DISABLED - (default) Do not generate a TLS/SSL certificate for this
	//    fleet.
	CertificateConfiguration *CertificateConfiguration `type:"structure"`

	// A human-readable description of a fleet.
	Description *string `min:"1" type:"string"`

	// Range of IP addresses and port settings that permit inbound traffic to access
	// game sessions that are running on the fleet. For fleets using a custom game
	// build, this parameter is required before game sessions running on the fleet
	// can accept connections. For Realtime Servers fleets, Amazon GameLift automatically
	// sets TCP and UDP ranges for use by the Realtime servers. You can specify
	// multiple permission settings or add more by updating the fleet.
	EC2InboundPermissions []IpPermission `type:"list"`

	// The name of an EC2 instance type that is supported in Amazon GameLift. A
	// fleet instance type determines the computing resources of each instance in
	// the fleet, including CPU, memory, storage, and networking capacity. Amazon
	// GameLift supports the following EC2 instance types. See Amazon EC2 Instance
	// Types (http://aws.amazon.com/ec2/instance-types/) for detailed descriptions.
	//
	// EC2InstanceType is a required field
	EC2InstanceType EC2InstanceType `type:"string" required:"true" enum:"true"`

	// Indicates whether to use On-Demand instances or Spot instances for this fleet.
	// If empty, the default is ON_DEMAND. Both categories of instances use identical
	// hardware and configurations based on the instance type selected for this
	// fleet. Learn more about On-Demand versus Spot Instances (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-ec2-instances.html#gamelift-ec2-instances-spot).
	FleetType FleetType `type:"string" enum:"true"`

	// A unique identifier for an AWS IAM role that manages access to your AWS services.
	// With an instance role ARN set, any application that runs on an instance in
	// this fleet can assume the role, including install scripts, server processes,
	// and daemons (background processes). Create a role or look up a role's ARN
	// from the IAM dashboard (https://console.aws.amazon.com/iam/) in the AWS Management
	// Console. Learn more about using on-box credentials for your game servers
	// at Access external resources from a game server (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-resources.html).
	InstanceRoleArn *string `min:"1" type:"string"`

	// This parameter is no longer used. Instead, to specify where Amazon GameLift
	// should store log files once a server process shuts down, use the Amazon GameLift
	// server API ProcessReady() and specify one or more directory paths in logParameters.
	// See more information in the Server API Reference (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api-ref.html#gamelift-sdk-server-api-ref-dataypes-process).
	LogPaths []string `type:"list"`

	// The name of an Amazon CloudWatch metric group to add this fleet to. A metric
	// group aggregates the metrics for all fleets in the group. Specify an existing
	// metric group name, or provide a new name to create a new metric group. A
	// fleet can only be included in one metric group at a time.
	MetricGroups []string `type:"list"`

	// A descriptive label that is associated with a fleet. Fleet names do not need
	// to be unique.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// A game session protection policy to apply to all instances in this fleet.
	// If this parameter is not set, instances in this fleet default to no protection.
	// You can change a fleet's protection policy using UpdateFleetAttributes, but
	// this change will only affect sessions created after the policy change. You
	// can also set protection for individual instances using UpdateGameSession.
	//
	//    * NoProtection - The game session can be terminated during a scale-down
	//    event.
	//
	//    * FullProtection - If the game session is in an ACTIVE status, it cannot
	//    be terminated during a scale-down event.
	NewGameSessionProtectionPolicy ProtectionPolicy `type:"string" enum:"true"`

	// A unique identifier for the AWS account with the VPC that you want to peer
	// your Amazon GameLift fleet with. You can find your account ID in the AWS
	// Management Console under account settings.
	PeerVpcAwsAccountId *string `min:"1" type:"string"`

	// A unique identifier for a VPC with resources to be accessed by your Amazon
	// GameLift fleet. The VPC must be in the same Region as your fleet. To look
	// up a VPC ID, use the VPC Dashboard (https://console.aws.amazon.com/vpc/)
	// in the AWS Management Console. Learn more about VPC peering in VPC Peering
	// with Amazon GameLift Fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html).
	PeerVpcId *string `min:"1" type:"string"`

	// A policy that limits the number of game sessions an individual player can
	// create over a span of time for this fleet.
	ResourceCreationLimitPolicy *ResourceCreationLimitPolicy `type:"structure"`

	// Instructions for launching server processes on each instance in the fleet.
	// Server processes run either a custom game build executable or a Realtime
	// script. The runtime configuration defines the server executables or launch
	// script file, launch parameters, and the number of processes to run concurrently
	// on each instance. When creating a fleet, the runtime configuration must have
	// at least one server process configuration; otherwise the request fails with
	// an invalid request exception. (This parameter replaces the parameters ServerLaunchPath
	// and ServerLaunchParameters, although requests that contain values for these
	// parameters instead of a runtime configuration will continue to work.) This
	// parameter is required unless the parameters ServerLaunchPath and ServerLaunchParameters
	// are defined. Runtime configuration replaced these parameters, but fleets
	// that use them will continue to work.
	RuntimeConfiguration *RuntimeConfiguration `type:"structure"`

	// A unique identifier for a Realtime script to be deployed on the new fleet.
	// You can use either the script ID or ARN value. The Realtime script must have
	// been successfully uploaded to Amazon GameLift. This fleet setting cannot
	// be changed once the fleet is created.
	ScriptId *string `type:"string"`

	// This parameter is no longer used. Instead, specify server launch parameters
	// in the RuntimeConfiguration parameter. (Requests that specify a server launch
	// path and launch parameters instead of a runtime configuration will continue
	// to work.)
	ServerLaunchParameters *string `min:"1" type:"string"`

	// This parameter is no longer used. Instead, specify a server launch path using
	// the RuntimeConfiguration parameter. Requests that specify a server launch
	// path and launch parameters instead of a runtime configuration will continue
	// to work.
	ServerLaunchPath *string `min:"1" type:"string"`

	// A list of labels to assign to the new fleet resource. Tags are developer-defined
	// key-value pairs. Tagging AWS resources are useful for resource management,
	// access management and cost allocation. For more information, see Tagging
	// AWS Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// in the AWS General Reference. Once the resource is created, you can use TagResource,
	// UntagResource, and ListTagsForResource to add, remove, and view tags. The
	// maximum tag limit may be lower than stated. See the AWS General Reference
	// for actual tagging limits.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateFleetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFleetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateFleetInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if len(s.EC2InstanceType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("EC2InstanceType"))
	}
	if s.InstanceRoleArn != nil && len(*s.InstanceRoleArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceRoleArn", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.PeerVpcAwsAccountId != nil && len(*s.PeerVpcAwsAccountId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PeerVpcAwsAccountId", 1))
	}
	if s.PeerVpcId != nil && len(*s.PeerVpcId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PeerVpcId", 1))
	}
	if s.ServerLaunchParameters != nil && len(*s.ServerLaunchParameters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerLaunchParameters", 1))
	}
	if s.ServerLaunchPath != nil && len(*s.ServerLaunchPath) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerLaunchPath", 1))
	}
	if s.CertificateConfiguration != nil {
		if err := s.CertificateConfiguration.Validate(); err != nil {
			invalidParams.AddNested("CertificateConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.EC2InboundPermissions != nil {
		for i, v := range s.EC2InboundPermissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "EC2InboundPermissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.RuntimeConfiguration != nil {
		if err := s.RuntimeConfiguration.Validate(); err != nil {
			invalidParams.AddNested("RuntimeConfiguration", err.(aws.ErrInvalidParams))
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

// Represents the returned data in response to a request action.
type CreateFleetOutput struct {
	_ struct{} `type:"structure"`

	// Properties for the newly created fleet.
	FleetAttributes *FleetAttributes `type:"structure"`
}

// String returns the string representation
func (s CreateFleetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateFleet = "CreateFleet"

// CreateFleetRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Creates a new fleet to run your game servers. whether they are custom game
// builds or Realtime Servers with game-specific script. A fleet is a set of
// Amazon Elastic Compute Cloud (Amazon EC2) instances, each of which can host
// multiple game sessions. When creating a fleet, you choose the hardware specifications,
// set some configuration options, and specify the game server to deploy on
// the new fleet.
//
// To create a new fleet, provide the following: (1) a fleet name, (2) an EC2
// instance type and fleet type (spot or on-demand), (3) the build ID for your
// game build or script ID if using Realtime Servers, and (4) a runtime configuration,
// which determines how game servers will run on each instance in the fleet.
//
// If the CreateFleet call is successful, Amazon GameLift performs the following
// tasks. You can track the process of a fleet by checking the fleet status
// or by monitoring fleet creation events:
//
//    * Creates a fleet resource. Status: NEW.
//
//    * Begins writing events to the fleet event log, which can be accessed
//    in the Amazon GameLift console.
//
//    * Sets the fleet's target capacity to 1 (desired instances), which triggers
//    Amazon GameLift to start one new EC2 instance.
//
//    * Downloads the game build or Realtime script to the new instance and
//    installs it. Statuses: DOWNLOADING, VALIDATING, BUILDING.
//
//    * Starts launching server processes on the instance. If the fleet is configured
//    to run multiple server processes per instance, Amazon GameLift staggers
//    each process launch by a few seconds. Status: ACTIVATING.
//
//    * Sets the fleet's status to ACTIVE as soon as one server process is ready
//    to host a game session.
//
// Learn more
//
// Setting Up Fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
//
// Debug Fleet Creation Issues (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-debug.html#fleets-creating-debug-creation)
//
// Related operations
//
//    * CreateFleet
//
//    * ListFleets
//
//    * DeleteFleet
//
//    * DescribeFleetAttributes
//
//    * UpdateFleetAttributes
//
//    * StartFleetActions or StopFleetActions
//
//    // Example sending a request using CreateFleetRequest.
//    req := client.CreateFleetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateFleet
func (c *Client) CreateFleetRequest(input *CreateFleetInput) CreateFleetRequest {
	op := &aws.Operation{
		Name:       opCreateFleet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateFleetInput{}
	}

	req := c.newRequest(op, input, &CreateFleetOutput{})

	return CreateFleetRequest{Request: req, Input: input, Copy: c.CreateFleetRequest}
}

// CreateFleetRequest is the request type for the
// CreateFleet API operation.
type CreateFleetRequest struct {
	*aws.Request
	Input *CreateFleetInput
	Copy  func(*CreateFleetInput) CreateFleetRequest
}

// Send marshals and sends the CreateFleet API request.
func (r CreateFleetRequest) Send(ctx context.Context) (*CreateFleetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFleetResponse{
		CreateFleetOutput: r.Request.Data.(*CreateFleetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFleetResponse is the response type for the
// CreateFleet API operation.
type CreateFleetResponse struct {
	*CreateFleetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFleet request.
func (r *CreateFleetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}