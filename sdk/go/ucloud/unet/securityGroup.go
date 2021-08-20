// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package unet

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Security Group resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-ucloud/sdk/go/ucloud/unet"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := unet.NewSecurityGroup(ctx, "example", &unet.SecurityGroupArgs{
// 			Rules: unet.SecurityGroupRuleArray{
// 				&unet.SecurityGroupRuleArgs{
// 					CidrBlock: pulumi.String("192.168.0.0/16"),
// 					Policy:    pulumi.String("accept"),
// 					PortRange: pulumi.String("80"),
// 					Protocol:  pulumi.String("tcp"),
// 				},
// 				&unet.SecurityGroupRuleArgs{
// 					CidrBlock: pulumi.String("192.168.0.0/16"),
// 					Policy:    pulumi.String("accept"),
// 					PortRange: pulumi.String("443"),
// 					Protocol:  pulumi.String("tcp"),
// 				},
// 			},
// 			Tag: pulumi.String("tf-example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Security Group can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import ucloud:unet/securityGroup:SecurityGroup example firewall-abc123456
// ```
type SecurityGroup struct {
	pulumi.CustomResourceState

	// The time of creation of security group, formatted in RFC3339 time string.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	Name       pulumi.StringOutput `pulumi:"name"`
	// The remarks of the security group. (Default: `""`).
	Remark pulumi.StringOutput `pulumi:"remark"`
	// A list of security group rules. Can be specified multiple times for each rules. See rules below for details on attributes.
	Rules SecurityGroupRuleArrayOutput `pulumi:"rules"`
	// A tag assigned to security group, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringPtrOutput `pulumi:"tag"`
}

// NewSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroup(ctx *pulumi.Context,
	name string, args *SecurityGroupArgs, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	var resource SecurityGroup
	err := ctx.RegisterResource("ucloud:unet/securityGroup:SecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityGroup gets an existing SecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityGroupState, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	var resource SecurityGroup
	err := ctx.ReadResource("ucloud:unet/securityGroup:SecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroup resources.
type securityGroupState struct {
	// The time of creation of security group, formatted in RFC3339 time string.
	CreateTime *string `pulumi:"createTime"`
	Name       *string `pulumi:"name"`
	// The remarks of the security group. (Default: `""`).
	Remark *string `pulumi:"remark"`
	// A list of security group rules. Can be specified multiple times for each rules. See rules below for details on attributes.
	Rules []SecurityGroupRule `pulumi:"rules"`
	// A tag assigned to security group, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag *string `pulumi:"tag"`
}

type SecurityGroupState struct {
	// The time of creation of security group, formatted in RFC3339 time string.
	CreateTime pulumi.StringPtrInput
	Name       pulumi.StringPtrInput
	// The remarks of the security group. (Default: `""`).
	Remark pulumi.StringPtrInput
	// A list of security group rules. Can be specified multiple times for each rules. See rules below for details on attributes.
	Rules SecurityGroupRuleArrayInput
	// A tag assigned to security group, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringPtrInput
}

func (SecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupState)(nil)).Elem()
}

type securityGroupArgs struct {
	Name *string `pulumi:"name"`
	// The remarks of the security group. (Default: `""`).
	Remark *string `pulumi:"remark"`
	// A list of security group rules. Can be specified multiple times for each rules. See rules below for details on attributes.
	Rules []SecurityGroupRule `pulumi:"rules"`
	// A tag assigned to security group, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag *string `pulumi:"tag"`
}

// The set of arguments for constructing a SecurityGroup resource.
type SecurityGroupArgs struct {
	Name pulumi.StringPtrInput
	// The remarks of the security group. (Default: `""`).
	Remark pulumi.StringPtrInput
	// A list of security group rules. Can be specified multiple times for each rules. See rules below for details on attributes.
	Rules SecurityGroupRuleArrayInput
	// A tag assigned to security group, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringPtrInput
}

func (SecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupArgs)(nil)).Elem()
}

type SecurityGroupInput interface {
	pulumi.Input

	ToSecurityGroupOutput() SecurityGroupOutput
	ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput
}

func (*SecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroup)(nil))
}

func (i *SecurityGroup) ToSecurityGroupOutput() SecurityGroupOutput {
	return i.ToSecurityGroupOutputWithContext(context.Background())
}

func (i *SecurityGroup) ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupOutput)
}

func (i *SecurityGroup) ToSecurityGroupPtrOutput() SecurityGroupPtrOutput {
	return i.ToSecurityGroupPtrOutputWithContext(context.Background())
}

func (i *SecurityGroup) ToSecurityGroupPtrOutputWithContext(ctx context.Context) SecurityGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupPtrOutput)
}

type SecurityGroupPtrInput interface {
	pulumi.Input

	ToSecurityGroupPtrOutput() SecurityGroupPtrOutput
	ToSecurityGroupPtrOutputWithContext(ctx context.Context) SecurityGroupPtrOutput
}

type securityGroupPtrType SecurityGroupArgs

func (*securityGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroup)(nil))
}

func (i *securityGroupPtrType) ToSecurityGroupPtrOutput() SecurityGroupPtrOutput {
	return i.ToSecurityGroupPtrOutputWithContext(context.Background())
}

func (i *securityGroupPtrType) ToSecurityGroupPtrOutputWithContext(ctx context.Context) SecurityGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupPtrOutput)
}

// SecurityGroupArrayInput is an input type that accepts SecurityGroupArray and SecurityGroupArrayOutput values.
// You can construct a concrete instance of `SecurityGroupArrayInput` via:
//
//          SecurityGroupArray{ SecurityGroupArgs{...} }
type SecurityGroupArrayInput interface {
	pulumi.Input

	ToSecurityGroupArrayOutput() SecurityGroupArrayOutput
	ToSecurityGroupArrayOutputWithContext(context.Context) SecurityGroupArrayOutput
}

type SecurityGroupArray []SecurityGroupInput

func (SecurityGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SecurityGroup)(nil))
}

func (i SecurityGroupArray) ToSecurityGroupArrayOutput() SecurityGroupArrayOutput {
	return i.ToSecurityGroupArrayOutputWithContext(context.Background())
}

func (i SecurityGroupArray) ToSecurityGroupArrayOutputWithContext(ctx context.Context) SecurityGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupArrayOutput)
}

// SecurityGroupMapInput is an input type that accepts SecurityGroupMap and SecurityGroupMapOutput values.
// You can construct a concrete instance of `SecurityGroupMapInput` via:
//
//          SecurityGroupMap{ "key": SecurityGroupArgs{...} }
type SecurityGroupMapInput interface {
	pulumi.Input

	ToSecurityGroupMapOutput() SecurityGroupMapOutput
	ToSecurityGroupMapOutputWithContext(context.Context) SecurityGroupMapOutput
}

type SecurityGroupMap map[string]SecurityGroupInput

func (SecurityGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SecurityGroup)(nil))
}

func (i SecurityGroupMap) ToSecurityGroupMapOutput() SecurityGroupMapOutput {
	return i.ToSecurityGroupMapOutputWithContext(context.Background())
}

func (i SecurityGroupMap) ToSecurityGroupMapOutputWithContext(ctx context.Context) SecurityGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupMapOutput)
}

type SecurityGroupOutput struct {
	*pulumi.OutputState
}

func (SecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroup)(nil))
}

func (o SecurityGroupOutput) ToSecurityGroupOutput() SecurityGroupOutput {
	return o
}

func (o SecurityGroupOutput) ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput {
	return o
}

func (o SecurityGroupOutput) ToSecurityGroupPtrOutput() SecurityGroupPtrOutput {
	return o.ToSecurityGroupPtrOutputWithContext(context.Background())
}

func (o SecurityGroupOutput) ToSecurityGroupPtrOutputWithContext(ctx context.Context) SecurityGroupPtrOutput {
	return o.ApplyT(func(v SecurityGroup) *SecurityGroup {
		return &v
	}).(SecurityGroupPtrOutput)
}

type SecurityGroupPtrOutput struct {
	*pulumi.OutputState
}

func (SecurityGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroup)(nil))
}

func (o SecurityGroupPtrOutput) ToSecurityGroupPtrOutput() SecurityGroupPtrOutput {
	return o
}

func (o SecurityGroupPtrOutput) ToSecurityGroupPtrOutputWithContext(ctx context.Context) SecurityGroupPtrOutput {
	return o
}

type SecurityGroupArrayOutput struct{ *pulumi.OutputState }

func (SecurityGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityGroup)(nil))
}

func (o SecurityGroupArrayOutput) ToSecurityGroupArrayOutput() SecurityGroupArrayOutput {
	return o
}

func (o SecurityGroupArrayOutput) ToSecurityGroupArrayOutputWithContext(ctx context.Context) SecurityGroupArrayOutput {
	return o
}

func (o SecurityGroupArrayOutput) Index(i pulumi.IntInput) SecurityGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityGroup {
		return vs[0].([]SecurityGroup)[vs[1].(int)]
	}).(SecurityGroupOutput)
}

type SecurityGroupMapOutput struct{ *pulumi.OutputState }

func (SecurityGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SecurityGroup)(nil))
}

func (o SecurityGroupMapOutput) ToSecurityGroupMapOutput() SecurityGroupMapOutput {
	return o
}

func (o SecurityGroupMapOutput) ToSecurityGroupMapOutputWithContext(ctx context.Context) SecurityGroupMapOutput {
	return o
}

func (o SecurityGroupMapOutput) MapIndex(k pulumi.StringInput) SecurityGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SecurityGroup {
		return vs[0].(map[string]SecurityGroup)[vs[1].(string)]
	}).(SecurityGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityGroupOutput{})
	pulumi.RegisterOutputType(SecurityGroupPtrOutput{})
	pulumi.RegisterOutputType(SecurityGroupArrayOutput{})
	pulumi.RegisterOutputType(SecurityGroupMapOutput{})
}
