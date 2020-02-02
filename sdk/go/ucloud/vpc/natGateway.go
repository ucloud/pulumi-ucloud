// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vpc

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Nat Gateway resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/r/nat_gateway.html.markdown.
type NatGateway struct {
	pulumi.CustomResourceState

	// The time of creation of Nat Gateway, formatted in RFC3339 time string.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The ID of eip associate to the Nat Gateway.
	EipId pulumi.StringOutput `pulumi:"eipId"`
	// The boolean value to Controls whether or not start the whitelist mode.
	EnableWhiteList pulumi.BoolOutput   `pulumi:"enableWhiteList"`
	Name            pulumi.StringOutput `pulumi:"name"`
	// The remarks of the Nat Gateway. (Default: `""`).
	Remark pulumi.StringOutput `pulumi:"remark"`
	// The ID of the associated security group.
	SecurityGroup pulumi.StringOutput `pulumi:"securityGroup"`
	// The list of subnet ID under the VPC.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// A tag assigned to Nat Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	// * ``
	Tag pulumi.StringPtrOutput `pulumi:"tag"`
	// The ID of VPC linked to the Nat Gateway.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The white list of instance under the Nat Gateway.
	WhiteLists pulumi.StringArrayOutput `pulumi:"whiteLists"`
}

// NewNatGateway registers a new resource with the given unique name, arguments, and options.
func NewNatGateway(ctx *pulumi.Context,
	name string, args *NatGatewayArgs, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	if args == nil || args.EipId == nil {
		return nil, errors.New("missing required argument 'EipId'")
	}
	if args == nil || args.EnableWhiteList == nil {
		return nil, errors.New("missing required argument 'EnableWhiteList'")
	}
	if args == nil || args.SecurityGroup == nil {
		return nil, errors.New("missing required argument 'SecurityGroup'")
	}
	if args == nil || args.SubnetIds == nil {
		return nil, errors.New("missing required argument 'SubnetIds'")
	}
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	if args == nil {
		args = &NatGatewayArgs{}
	}
	var resource NatGateway
	err := ctx.RegisterResource("ucloud:vpc/natGateway:NatGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNatGateway gets an existing NatGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNatGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NatGatewayState, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	var resource NatGateway
	err := ctx.ReadResource("ucloud:vpc/natGateway:NatGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NatGateway resources.
type natGatewayState struct {
	// The time of creation of Nat Gateway, formatted in RFC3339 time string.
	CreateTime *string `pulumi:"createTime"`
	// The ID of eip associate to the Nat Gateway.
	EipId *string `pulumi:"eipId"`
	// The boolean value to Controls whether or not start the whitelist mode.
	EnableWhiteList *bool   `pulumi:"enableWhiteList"`
	Name            *string `pulumi:"name"`
	// The remarks of the Nat Gateway. (Default: `""`).
	Remark *string `pulumi:"remark"`
	// The ID of the associated security group.
	SecurityGroup *string `pulumi:"securityGroup"`
	// The list of subnet ID under the VPC.
	SubnetIds []string `pulumi:"subnetIds"`
	// A tag assigned to Nat Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	// * ``
	Tag *string `pulumi:"tag"`
	// The ID of VPC linked to the Nat Gateway.
	VpcId *string `pulumi:"vpcId"`
	// The white list of instance under the Nat Gateway.
	WhiteLists []string `pulumi:"whiteLists"`
}

type NatGatewayState struct {
	// The time of creation of Nat Gateway, formatted in RFC3339 time string.
	CreateTime pulumi.StringPtrInput
	// The ID of eip associate to the Nat Gateway.
	EipId pulumi.StringPtrInput
	// The boolean value to Controls whether or not start the whitelist mode.
	EnableWhiteList pulumi.BoolPtrInput
	Name            pulumi.StringPtrInput
	// The remarks of the Nat Gateway. (Default: `""`).
	Remark pulumi.StringPtrInput
	// The ID of the associated security group.
	SecurityGroup pulumi.StringPtrInput
	// The list of subnet ID under the VPC.
	SubnetIds pulumi.StringArrayInput
	// A tag assigned to Nat Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	// * ``
	Tag pulumi.StringPtrInput
	// The ID of VPC linked to the Nat Gateway.
	VpcId pulumi.StringPtrInput
	// The white list of instance under the Nat Gateway.
	WhiteLists pulumi.StringArrayInput
}

func (NatGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayState)(nil)).Elem()
}

type natGatewayArgs struct {
	// The ID of eip associate to the Nat Gateway.
	EipId string `pulumi:"eipId"`
	// The boolean value to Controls whether or not start the whitelist mode.
	EnableWhiteList bool    `pulumi:"enableWhiteList"`
	Name            *string `pulumi:"name"`
	// The remarks of the Nat Gateway. (Default: `""`).
	Remark *string `pulumi:"remark"`
	// The ID of the associated security group.
	SecurityGroup string `pulumi:"securityGroup"`
	// The list of subnet ID under the VPC.
	SubnetIds []string `pulumi:"subnetIds"`
	// A tag assigned to Nat Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	// * ``
	Tag *string `pulumi:"tag"`
	// The ID of VPC linked to the Nat Gateway.
	VpcId string `pulumi:"vpcId"`
	// The white list of instance under the Nat Gateway.
	WhiteLists []string `pulumi:"whiteLists"`
}

// The set of arguments for constructing a NatGateway resource.
type NatGatewayArgs struct {
	// The ID of eip associate to the Nat Gateway.
	EipId pulumi.StringInput
	// The boolean value to Controls whether or not start the whitelist mode.
	EnableWhiteList pulumi.BoolInput
	Name            pulumi.StringPtrInput
	// The remarks of the Nat Gateway. (Default: `""`).
	Remark pulumi.StringPtrInput
	// The ID of the associated security group.
	SecurityGroup pulumi.StringInput
	// The list of subnet ID under the VPC.
	SubnetIds pulumi.StringArrayInput
	// A tag assigned to Nat Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	// * ``
	Tag pulumi.StringPtrInput
	// The ID of VPC linked to the Nat Gateway.
	VpcId pulumi.StringInput
	// The white list of instance under the Nat Gateway.
	WhiteLists pulumi.StringArrayInput
}

func (NatGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayArgs)(nil)).Elem()
}
