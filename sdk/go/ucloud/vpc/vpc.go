// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vpc

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a VPC resource.
//
// > **Note**  The network segment can only be created or deleted, can not perform both of them at the same time.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/r/vpc.html.markdown.
type Vpc struct {
	pulumi.CustomResourceState

	// The CIDR blocks of VPC.
	CidrBlocks pulumi.StringArrayOutput `pulumi:"cidrBlocks"`
	// The time of creation for VPC, formatted in RFC3339 time string.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	Name       pulumi.StringOutput `pulumi:"name"`
	// It is a nested type which documented below.
	NetworkInfos VpcNetworkInfoArrayOutput `pulumi:"networkInfos"`
	// The remarks of the VPC. (Default: `""`).
	Remark pulumi.StringOutput `pulumi:"remark"`
	// A tag assigned to VPC, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringPtrOutput `pulumi:"tag"`
	// The time whenever there is a change made to VPC, formatted in RFC3339 time string.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewVpc registers a new resource with the given unique name, arguments, and options.
func NewVpc(ctx *pulumi.Context,
	name string, args *VpcArgs, opts ...pulumi.ResourceOption) (*Vpc, error) {
	if args == nil || args.CidrBlocks == nil {
		return nil, errors.New("missing required argument 'CidrBlocks'")
	}
	if args == nil {
		args = &VpcArgs{}
	}
	var resource Vpc
	err := ctx.RegisterResource("ucloud:vpc/vpc:Vpc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpc gets an existing Vpc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcState, opts ...pulumi.ResourceOption) (*Vpc, error) {
	var resource Vpc
	err := ctx.ReadResource("ucloud:vpc/vpc:Vpc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vpc resources.
type vpcState struct {
	// The CIDR blocks of VPC.
	CidrBlocks []string `pulumi:"cidrBlocks"`
	// The time of creation for VPC, formatted in RFC3339 time string.
	CreateTime *string `pulumi:"createTime"`
	Name       *string `pulumi:"name"`
	// It is a nested type which documented below.
	NetworkInfos []VpcNetworkInfo `pulumi:"networkInfos"`
	// The remarks of the VPC. (Default: `""`).
	Remark *string `pulumi:"remark"`
	// A tag assigned to VPC, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag *string `pulumi:"tag"`
	// The time whenever there is a change made to VPC, formatted in RFC3339 time string.
	UpdateTime *string `pulumi:"updateTime"`
}

type VpcState struct {
	// The CIDR blocks of VPC.
	CidrBlocks pulumi.StringArrayInput
	// The time of creation for VPC, formatted in RFC3339 time string.
	CreateTime pulumi.StringPtrInput
	Name       pulumi.StringPtrInput
	// It is a nested type which documented below.
	NetworkInfos VpcNetworkInfoArrayInput
	// The remarks of the VPC. (Default: `""`).
	Remark pulumi.StringPtrInput
	// A tag assigned to VPC, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringPtrInput
	// The time whenever there is a change made to VPC, formatted in RFC3339 time string.
	UpdateTime pulumi.StringPtrInput
}

func (VpcState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcState)(nil)).Elem()
}

type vpcArgs struct {
	// The CIDR blocks of VPC.
	CidrBlocks []string `pulumi:"cidrBlocks"`
	Name       *string  `pulumi:"name"`
	// The remarks of the VPC. (Default: `""`).
	Remark *string `pulumi:"remark"`
	// A tag assigned to VPC, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag *string `pulumi:"tag"`
}

// The set of arguments for constructing a Vpc resource.
type VpcArgs struct {
	// The CIDR blocks of VPC.
	CidrBlocks pulumi.StringArrayInput
	Name       pulumi.StringPtrInput
	// The remarks of the VPC. (Default: `""`).
	Remark pulumi.StringPtrInput
	// A tag assigned to VPC, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringPtrInput
}

func (VpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcArgs)(nil)).Elem()
}
