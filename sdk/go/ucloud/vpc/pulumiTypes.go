// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VPCNetworkInfo struct {
	// The CIDR block of the VPC.
	CidrBlock *string `pulumi:"cidrBlock"`
}

// VPCNetworkInfoInput is an input type that accepts VPCNetworkInfoArgs and VPCNetworkInfoOutput values.
// You can construct a concrete instance of `VPCNetworkInfoInput` via:
//
//          VPCNetworkInfoArgs{...}
type VPCNetworkInfoInput interface {
	pulumi.Input

	ToVPCNetworkInfoOutput() VPCNetworkInfoOutput
	ToVPCNetworkInfoOutputWithContext(context.Context) VPCNetworkInfoOutput
}

type VPCNetworkInfoArgs struct {
	// The CIDR block of the VPC.
	CidrBlock pulumi.StringPtrInput `pulumi:"cidrBlock"`
}

func (VPCNetworkInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VPCNetworkInfo)(nil)).Elem()
}

func (i VPCNetworkInfoArgs) ToVPCNetworkInfoOutput() VPCNetworkInfoOutput {
	return i.ToVPCNetworkInfoOutputWithContext(context.Background())
}

func (i VPCNetworkInfoArgs) ToVPCNetworkInfoOutputWithContext(ctx context.Context) VPCNetworkInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPCNetworkInfoOutput)
}

// VPCNetworkInfoArrayInput is an input type that accepts VPCNetworkInfoArray and VPCNetworkInfoArrayOutput values.
// You can construct a concrete instance of `VPCNetworkInfoArrayInput` via:
//
//          VPCNetworkInfoArray{ VPCNetworkInfoArgs{...} }
type VPCNetworkInfoArrayInput interface {
	pulumi.Input

	ToVPCNetworkInfoArrayOutput() VPCNetworkInfoArrayOutput
	ToVPCNetworkInfoArrayOutputWithContext(context.Context) VPCNetworkInfoArrayOutput
}

type VPCNetworkInfoArray []VPCNetworkInfoInput

func (VPCNetworkInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VPCNetworkInfo)(nil)).Elem()
}

func (i VPCNetworkInfoArray) ToVPCNetworkInfoArrayOutput() VPCNetworkInfoArrayOutput {
	return i.ToVPCNetworkInfoArrayOutputWithContext(context.Background())
}

func (i VPCNetworkInfoArray) ToVPCNetworkInfoArrayOutputWithContext(ctx context.Context) VPCNetworkInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPCNetworkInfoArrayOutput)
}

type VPCNetworkInfoOutput struct{ *pulumi.OutputState }

func (VPCNetworkInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VPCNetworkInfo)(nil)).Elem()
}

func (o VPCNetworkInfoOutput) ToVPCNetworkInfoOutput() VPCNetworkInfoOutput {
	return o
}

func (o VPCNetworkInfoOutput) ToVPCNetworkInfoOutputWithContext(ctx context.Context) VPCNetworkInfoOutput {
	return o
}

// The CIDR block of the VPC.
func (o VPCNetworkInfoOutput) CidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VPCNetworkInfo) *string { return v.CidrBlock }).(pulumi.StringPtrOutput)
}

type VPCNetworkInfoArrayOutput struct{ *pulumi.OutputState }

func (VPCNetworkInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VPCNetworkInfo)(nil)).Elem()
}

func (o VPCNetworkInfoArrayOutput) ToVPCNetworkInfoArrayOutput() VPCNetworkInfoArrayOutput {
	return o
}

func (o VPCNetworkInfoArrayOutput) ToVPCNetworkInfoArrayOutputWithContext(ctx context.Context) VPCNetworkInfoArrayOutput {
	return o
}

func (o VPCNetworkInfoArrayOutput) Index(i pulumi.IntInput) VPCNetworkInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VPCNetworkInfo {
		return vs[0].([]VPCNetworkInfo)[vs[1].(int)]
	}).(VPCNetworkInfoOutput)
}

type GetNATGatewayNatGateway struct {
	// The time of creation for Nat Gateway, formatted in RFC3339 time string.
	CreateTime string `pulumi:"createTime"`
	// The ID of Nat Gateway.
	Id string `pulumi:"id"`
	// It is a nested type which documented below.
	IpSets []GetNATGatewayNatGatewayIpSet `pulumi:"ipSets"`
	// The name of the Nat Gateway.
	Name string `pulumi:"name"`
	// The remarks of Nat Gateway.
	Remark        string `pulumi:"remark"`
	SecurityGroup string `pulumi:"securityGroup"`
	// The list of subnet ID under the VPC.
	// * `securityGroup` -The ID of the associated security group.
	SubnetIds []string `pulumi:"subnetIds"`
	// A tag assigned to the Nat Gateway.
	Tag string `pulumi:"tag"`
	// The ID of VPC linked to the Nat Gateway.
	VpcId string `pulumi:"vpcId"`
}

// GetNATGatewayNatGatewayInput is an input type that accepts GetNATGatewayNatGatewayArgs and GetNATGatewayNatGatewayOutput values.
// You can construct a concrete instance of `GetNATGatewayNatGatewayInput` via:
//
//          GetNATGatewayNatGatewayArgs{...}
type GetNATGatewayNatGatewayInput interface {
	pulumi.Input

	ToGetNATGatewayNatGatewayOutput() GetNATGatewayNatGatewayOutput
	ToGetNATGatewayNatGatewayOutputWithContext(context.Context) GetNATGatewayNatGatewayOutput
}

type GetNATGatewayNatGatewayArgs struct {
	// The time of creation for Nat Gateway, formatted in RFC3339 time string.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// The ID of Nat Gateway.
	Id pulumi.StringInput `pulumi:"id"`
	// It is a nested type which documented below.
	IpSets GetNATGatewayNatGatewayIpSetArrayInput `pulumi:"ipSets"`
	// The name of the Nat Gateway.
	Name pulumi.StringInput `pulumi:"name"`
	// The remarks of Nat Gateway.
	Remark        pulumi.StringInput `pulumi:"remark"`
	SecurityGroup pulumi.StringInput `pulumi:"securityGroup"`
	// The list of subnet ID under the VPC.
	// * `securityGroup` -The ID of the associated security group.
	SubnetIds pulumi.StringArrayInput `pulumi:"subnetIds"`
	// A tag assigned to the Nat Gateway.
	Tag pulumi.StringInput `pulumi:"tag"`
	// The ID of VPC linked to the Nat Gateway.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

func (GetNATGatewayNatGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNATGatewayNatGateway)(nil)).Elem()
}

func (i GetNATGatewayNatGatewayArgs) ToGetNATGatewayNatGatewayOutput() GetNATGatewayNatGatewayOutput {
	return i.ToGetNATGatewayNatGatewayOutputWithContext(context.Background())
}

func (i GetNATGatewayNatGatewayArgs) ToGetNATGatewayNatGatewayOutputWithContext(ctx context.Context) GetNATGatewayNatGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetNATGatewayNatGatewayOutput)
}

// GetNATGatewayNatGatewayArrayInput is an input type that accepts GetNATGatewayNatGatewayArray and GetNATGatewayNatGatewayArrayOutput values.
// You can construct a concrete instance of `GetNATGatewayNatGatewayArrayInput` via:
//
//          GetNATGatewayNatGatewayArray{ GetNATGatewayNatGatewayArgs{...} }
type GetNATGatewayNatGatewayArrayInput interface {
	pulumi.Input

	ToGetNATGatewayNatGatewayArrayOutput() GetNATGatewayNatGatewayArrayOutput
	ToGetNATGatewayNatGatewayArrayOutputWithContext(context.Context) GetNATGatewayNatGatewayArrayOutput
}

type GetNATGatewayNatGatewayArray []GetNATGatewayNatGatewayInput

func (GetNATGatewayNatGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetNATGatewayNatGateway)(nil)).Elem()
}

func (i GetNATGatewayNatGatewayArray) ToGetNATGatewayNatGatewayArrayOutput() GetNATGatewayNatGatewayArrayOutput {
	return i.ToGetNATGatewayNatGatewayArrayOutputWithContext(context.Background())
}

func (i GetNATGatewayNatGatewayArray) ToGetNATGatewayNatGatewayArrayOutputWithContext(ctx context.Context) GetNATGatewayNatGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetNATGatewayNatGatewayArrayOutput)
}

type GetNATGatewayNatGatewayOutput struct{ *pulumi.OutputState }

func (GetNATGatewayNatGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNATGatewayNatGateway)(nil)).Elem()
}

func (o GetNATGatewayNatGatewayOutput) ToGetNATGatewayNatGatewayOutput() GetNATGatewayNatGatewayOutput {
	return o
}

func (o GetNATGatewayNatGatewayOutput) ToGetNATGatewayNatGatewayOutputWithContext(ctx context.Context) GetNATGatewayNatGatewayOutput {
	return o
}

// The time of creation for Nat Gateway, formatted in RFC3339 time string.
func (o GetNATGatewayNatGatewayOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetNATGatewayNatGateway) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The ID of Nat Gateway.
func (o GetNATGatewayNatGatewayOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNATGatewayNatGateway) string { return v.Id }).(pulumi.StringOutput)
}

// It is a nested type which documented below.
func (o GetNATGatewayNatGatewayOutput) IpSets() GetNATGatewayNatGatewayIpSetArrayOutput {
	return o.ApplyT(func(v GetNATGatewayNatGateway) []GetNATGatewayNatGatewayIpSet { return v.IpSets }).(GetNATGatewayNatGatewayIpSetArrayOutput)
}

// The name of the Nat Gateway.
func (o GetNATGatewayNatGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetNATGatewayNatGateway) string { return v.Name }).(pulumi.StringOutput)
}

// The remarks of Nat Gateway.
func (o GetNATGatewayNatGatewayOutput) Remark() pulumi.StringOutput {
	return o.ApplyT(func(v GetNATGatewayNatGateway) string { return v.Remark }).(pulumi.StringOutput)
}

func (o GetNATGatewayNatGatewayOutput) SecurityGroup() pulumi.StringOutput {
	return o.ApplyT(func(v GetNATGatewayNatGateway) string { return v.SecurityGroup }).(pulumi.StringOutput)
}

// The list of subnet ID under the VPC.
// * `securityGroup` -The ID of the associated security group.
func (o GetNATGatewayNatGatewayOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNATGatewayNatGateway) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// A tag assigned to the Nat Gateway.
func (o GetNATGatewayNatGatewayOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v GetNATGatewayNatGateway) string { return v.Tag }).(pulumi.StringOutput)
}

// The ID of VPC linked to the Nat Gateway.
func (o GetNATGatewayNatGatewayOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNATGatewayNatGateway) string { return v.VpcId }).(pulumi.StringOutput)
}

type GetNATGatewayNatGatewayArrayOutput struct{ *pulumi.OutputState }

func (GetNATGatewayNatGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetNATGatewayNatGateway)(nil)).Elem()
}

func (o GetNATGatewayNatGatewayArrayOutput) ToGetNATGatewayNatGatewayArrayOutput() GetNATGatewayNatGatewayArrayOutput {
	return o
}

func (o GetNATGatewayNatGatewayArrayOutput) ToGetNATGatewayNatGatewayArrayOutputWithContext(ctx context.Context) GetNATGatewayNatGatewayArrayOutput {
	return o
}

func (o GetNATGatewayNatGatewayArrayOutput) Index(i pulumi.IntInput) GetNATGatewayNatGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetNATGatewayNatGateway {
		return vs[0].([]GetNATGatewayNatGateway)[vs[1].(int)]
	}).(GetNATGatewayNatGatewayOutput)
}

type GetNATGatewayNatGatewayIpSet struct {
	// Type of Elastic IP routes.
	InternetType string `pulumi:"internetType"`
	// Elastic IP address.
	Ip string `pulumi:"ip"`
}

// GetNATGatewayNatGatewayIpSetInput is an input type that accepts GetNATGatewayNatGatewayIpSetArgs and GetNATGatewayNatGatewayIpSetOutput values.
// You can construct a concrete instance of `GetNATGatewayNatGatewayIpSetInput` via:
//
//          GetNATGatewayNatGatewayIpSetArgs{...}
type GetNATGatewayNatGatewayIpSetInput interface {
	pulumi.Input

	ToGetNATGatewayNatGatewayIpSetOutput() GetNATGatewayNatGatewayIpSetOutput
	ToGetNATGatewayNatGatewayIpSetOutputWithContext(context.Context) GetNATGatewayNatGatewayIpSetOutput
}

type GetNATGatewayNatGatewayIpSetArgs struct {
	// Type of Elastic IP routes.
	InternetType pulumi.StringInput `pulumi:"internetType"`
	// Elastic IP address.
	Ip pulumi.StringInput `pulumi:"ip"`
}

func (GetNATGatewayNatGatewayIpSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNATGatewayNatGatewayIpSet)(nil)).Elem()
}

func (i GetNATGatewayNatGatewayIpSetArgs) ToGetNATGatewayNatGatewayIpSetOutput() GetNATGatewayNatGatewayIpSetOutput {
	return i.ToGetNATGatewayNatGatewayIpSetOutputWithContext(context.Background())
}

func (i GetNATGatewayNatGatewayIpSetArgs) ToGetNATGatewayNatGatewayIpSetOutputWithContext(ctx context.Context) GetNATGatewayNatGatewayIpSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetNATGatewayNatGatewayIpSetOutput)
}

// GetNATGatewayNatGatewayIpSetArrayInput is an input type that accepts GetNATGatewayNatGatewayIpSetArray and GetNATGatewayNatGatewayIpSetArrayOutput values.
// You can construct a concrete instance of `GetNATGatewayNatGatewayIpSetArrayInput` via:
//
//          GetNATGatewayNatGatewayIpSetArray{ GetNATGatewayNatGatewayIpSetArgs{...} }
type GetNATGatewayNatGatewayIpSetArrayInput interface {
	pulumi.Input

	ToGetNATGatewayNatGatewayIpSetArrayOutput() GetNATGatewayNatGatewayIpSetArrayOutput
	ToGetNATGatewayNatGatewayIpSetArrayOutputWithContext(context.Context) GetNATGatewayNatGatewayIpSetArrayOutput
}

type GetNATGatewayNatGatewayIpSetArray []GetNATGatewayNatGatewayIpSetInput

func (GetNATGatewayNatGatewayIpSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetNATGatewayNatGatewayIpSet)(nil)).Elem()
}

func (i GetNATGatewayNatGatewayIpSetArray) ToGetNATGatewayNatGatewayIpSetArrayOutput() GetNATGatewayNatGatewayIpSetArrayOutput {
	return i.ToGetNATGatewayNatGatewayIpSetArrayOutputWithContext(context.Background())
}

func (i GetNATGatewayNatGatewayIpSetArray) ToGetNATGatewayNatGatewayIpSetArrayOutputWithContext(ctx context.Context) GetNATGatewayNatGatewayIpSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetNATGatewayNatGatewayIpSetArrayOutput)
}

type GetNATGatewayNatGatewayIpSetOutput struct{ *pulumi.OutputState }

func (GetNATGatewayNatGatewayIpSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNATGatewayNatGatewayIpSet)(nil)).Elem()
}

func (o GetNATGatewayNatGatewayIpSetOutput) ToGetNATGatewayNatGatewayIpSetOutput() GetNATGatewayNatGatewayIpSetOutput {
	return o
}

func (o GetNATGatewayNatGatewayIpSetOutput) ToGetNATGatewayNatGatewayIpSetOutputWithContext(ctx context.Context) GetNATGatewayNatGatewayIpSetOutput {
	return o
}

// Type of Elastic IP routes.
func (o GetNATGatewayNatGatewayIpSetOutput) InternetType() pulumi.StringOutput {
	return o.ApplyT(func(v GetNATGatewayNatGatewayIpSet) string { return v.InternetType }).(pulumi.StringOutput)
}

// Elastic IP address.
func (o GetNATGatewayNatGatewayIpSetOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v GetNATGatewayNatGatewayIpSet) string { return v.Ip }).(pulumi.StringOutput)
}

type GetNATGatewayNatGatewayIpSetArrayOutput struct{ *pulumi.OutputState }

func (GetNATGatewayNatGatewayIpSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetNATGatewayNatGatewayIpSet)(nil)).Elem()
}

func (o GetNATGatewayNatGatewayIpSetArrayOutput) ToGetNATGatewayNatGatewayIpSetArrayOutput() GetNATGatewayNatGatewayIpSetArrayOutput {
	return o
}

func (o GetNATGatewayNatGatewayIpSetArrayOutput) ToGetNATGatewayNatGatewayIpSetArrayOutputWithContext(ctx context.Context) GetNATGatewayNatGatewayIpSetArrayOutput {
	return o
}

func (o GetNATGatewayNatGatewayIpSetArrayOutput) Index(i pulumi.IntInput) GetNATGatewayNatGatewayIpSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetNATGatewayNatGatewayIpSet {
		return vs[0].([]GetNATGatewayNatGatewayIpSet)[vs[1].(int)]
	}).(GetNATGatewayNatGatewayIpSetOutput)
}

type GetSubnetSubnet struct {
	// The cidr block of the desired Subnet.
	CidrBlock string `pulumi:"cidrBlock"`
	// The time of creation of Subnet, formatted in RFC3339 time string.
	CreateTime string `pulumi:"createTime"`
	// The ID of Subnet.
	Id string `pulumi:"id"`
	// The name of Subnet.
	Name string `pulumi:"name"`
	// The remark of the Subnet.
	Remark string `pulumi:"remark"`
	// A tag assigned to Subnet.
	Tag string `pulumi:"tag"`
}

// GetSubnetSubnetInput is an input type that accepts GetSubnetSubnetArgs and GetSubnetSubnetOutput values.
// You can construct a concrete instance of `GetSubnetSubnetInput` via:
//
//          GetSubnetSubnetArgs{...}
type GetSubnetSubnetInput interface {
	pulumi.Input

	ToGetSubnetSubnetOutput() GetSubnetSubnetOutput
	ToGetSubnetSubnetOutputWithContext(context.Context) GetSubnetSubnetOutput
}

type GetSubnetSubnetArgs struct {
	// The cidr block of the desired Subnet.
	CidrBlock pulumi.StringInput `pulumi:"cidrBlock"`
	// The time of creation of Subnet, formatted in RFC3339 time string.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// The ID of Subnet.
	Id pulumi.StringInput `pulumi:"id"`
	// The name of Subnet.
	Name pulumi.StringInput `pulumi:"name"`
	// The remark of the Subnet.
	Remark pulumi.StringInput `pulumi:"remark"`
	// A tag assigned to Subnet.
	Tag pulumi.StringInput `pulumi:"tag"`
}

func (GetSubnetSubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetSubnet)(nil)).Elem()
}

func (i GetSubnetSubnetArgs) ToGetSubnetSubnetOutput() GetSubnetSubnetOutput {
	return i.ToGetSubnetSubnetOutputWithContext(context.Background())
}

func (i GetSubnetSubnetArgs) ToGetSubnetSubnetOutputWithContext(ctx context.Context) GetSubnetSubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSubnetSubnetOutput)
}

// GetSubnetSubnetArrayInput is an input type that accepts GetSubnetSubnetArray and GetSubnetSubnetArrayOutput values.
// You can construct a concrete instance of `GetSubnetSubnetArrayInput` via:
//
//          GetSubnetSubnetArray{ GetSubnetSubnetArgs{...} }
type GetSubnetSubnetArrayInput interface {
	pulumi.Input

	ToGetSubnetSubnetArrayOutput() GetSubnetSubnetArrayOutput
	ToGetSubnetSubnetArrayOutputWithContext(context.Context) GetSubnetSubnetArrayOutput
}

type GetSubnetSubnetArray []GetSubnetSubnetInput

func (GetSubnetSubnetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSubnetSubnet)(nil)).Elem()
}

func (i GetSubnetSubnetArray) ToGetSubnetSubnetArrayOutput() GetSubnetSubnetArrayOutput {
	return i.ToGetSubnetSubnetArrayOutputWithContext(context.Background())
}

func (i GetSubnetSubnetArray) ToGetSubnetSubnetArrayOutputWithContext(ctx context.Context) GetSubnetSubnetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSubnetSubnetArrayOutput)
}

type GetSubnetSubnetOutput struct{ *pulumi.OutputState }

func (GetSubnetSubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetSubnet)(nil)).Elem()
}

func (o GetSubnetSubnetOutput) ToGetSubnetSubnetOutput() GetSubnetSubnetOutput {
	return o
}

func (o GetSubnetSubnetOutput) ToGetSubnetSubnetOutputWithContext(ctx context.Context) GetSubnetSubnetOutput {
	return o
}

// The cidr block of the desired Subnet.
func (o GetSubnetSubnetOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetSubnet) string { return v.CidrBlock }).(pulumi.StringOutput)
}

// The time of creation of Subnet, formatted in RFC3339 time string.
func (o GetSubnetSubnetOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetSubnet) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The ID of Subnet.
func (o GetSubnetSubnetOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetSubnet) string { return v.Id }).(pulumi.StringOutput)
}

// The name of Subnet.
func (o GetSubnetSubnetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetSubnet) string { return v.Name }).(pulumi.StringOutput)
}

// The remark of the Subnet.
func (o GetSubnetSubnetOutput) Remark() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetSubnet) string { return v.Remark }).(pulumi.StringOutput)
}

// A tag assigned to Subnet.
func (o GetSubnetSubnetOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetSubnet) string { return v.Tag }).(pulumi.StringOutput)
}

type GetSubnetSubnetArrayOutput struct{ *pulumi.OutputState }

func (GetSubnetSubnetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSubnetSubnet)(nil)).Elem()
}

func (o GetSubnetSubnetArrayOutput) ToGetSubnetSubnetArrayOutput() GetSubnetSubnetArrayOutput {
	return o
}

func (o GetSubnetSubnetArrayOutput) ToGetSubnetSubnetArrayOutputWithContext(ctx context.Context) GetSubnetSubnetArrayOutput {
	return o
}

func (o GetSubnetSubnetArrayOutput) Index(i pulumi.IntInput) GetSubnetSubnetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSubnetSubnet {
		return vs[0].([]GetSubnetSubnet)[vs[1].(int)]
	}).(GetSubnetSubnetOutput)
}

type GetVPCVpc struct {
	// The CIDR blocks of VPC.
	CidrBlocks []string `pulumi:"cidrBlocks"`
	// The time of creation for VPC, formatted in RFC3339 time string.
	CreateTime string `pulumi:"createTime"`
	// The ID of VPC.
	Id string `pulumi:"id"`
	// The name of VPC.
	Name string `pulumi:"name"`
	// A tag assigned to VPC.
	Tag string `pulumi:"tag"`
	// The time whenever there is a change made to VPC, formatted in RFC3339 time string.
	UpdateTime string `pulumi:"updateTime"`
}

// GetVPCVpcInput is an input type that accepts GetVPCVpcArgs and GetVPCVpcOutput values.
// You can construct a concrete instance of `GetVPCVpcInput` via:
//
//          GetVPCVpcArgs{...}
type GetVPCVpcInput interface {
	pulumi.Input

	ToGetVPCVpcOutput() GetVPCVpcOutput
	ToGetVPCVpcOutputWithContext(context.Context) GetVPCVpcOutput
}

type GetVPCVpcArgs struct {
	// The CIDR blocks of VPC.
	CidrBlocks pulumi.StringArrayInput `pulumi:"cidrBlocks"`
	// The time of creation for VPC, formatted in RFC3339 time string.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// The ID of VPC.
	Id pulumi.StringInput `pulumi:"id"`
	// The name of VPC.
	Name pulumi.StringInput `pulumi:"name"`
	// A tag assigned to VPC.
	Tag pulumi.StringInput `pulumi:"tag"`
	// The time whenever there is a change made to VPC, formatted in RFC3339 time string.
	UpdateTime pulumi.StringInput `pulumi:"updateTime"`
}

func (GetVPCVpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVPCVpc)(nil)).Elem()
}

func (i GetVPCVpcArgs) ToGetVPCVpcOutput() GetVPCVpcOutput {
	return i.ToGetVPCVpcOutputWithContext(context.Background())
}

func (i GetVPCVpcArgs) ToGetVPCVpcOutputWithContext(ctx context.Context) GetVPCVpcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVPCVpcOutput)
}

// GetVPCVpcArrayInput is an input type that accepts GetVPCVpcArray and GetVPCVpcArrayOutput values.
// You can construct a concrete instance of `GetVPCVpcArrayInput` via:
//
//          GetVPCVpcArray{ GetVPCVpcArgs{...} }
type GetVPCVpcArrayInput interface {
	pulumi.Input

	ToGetVPCVpcArrayOutput() GetVPCVpcArrayOutput
	ToGetVPCVpcArrayOutputWithContext(context.Context) GetVPCVpcArrayOutput
}

type GetVPCVpcArray []GetVPCVpcInput

func (GetVPCVpcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVPCVpc)(nil)).Elem()
}

func (i GetVPCVpcArray) ToGetVPCVpcArrayOutput() GetVPCVpcArrayOutput {
	return i.ToGetVPCVpcArrayOutputWithContext(context.Background())
}

func (i GetVPCVpcArray) ToGetVPCVpcArrayOutputWithContext(ctx context.Context) GetVPCVpcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVPCVpcArrayOutput)
}

type GetVPCVpcOutput struct{ *pulumi.OutputState }

func (GetVPCVpcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVPCVpc)(nil)).Elem()
}

func (o GetVPCVpcOutput) ToGetVPCVpcOutput() GetVPCVpcOutput {
	return o
}

func (o GetVPCVpcOutput) ToGetVPCVpcOutputWithContext(ctx context.Context) GetVPCVpcOutput {
	return o
}

// The CIDR blocks of VPC.
func (o GetVPCVpcOutput) CidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVPCVpc) []string { return v.CidrBlocks }).(pulumi.StringArrayOutput)
}

// The time of creation for VPC, formatted in RFC3339 time string.
func (o GetVPCVpcOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetVPCVpc) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The ID of VPC.
func (o GetVPCVpcOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVPCVpc) string { return v.Id }).(pulumi.StringOutput)
}

// The name of VPC.
func (o GetVPCVpcOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetVPCVpc) string { return v.Name }).(pulumi.StringOutput)
}

// A tag assigned to VPC.
func (o GetVPCVpcOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v GetVPCVpc) string { return v.Tag }).(pulumi.StringOutput)
}

// The time whenever there is a change made to VPC, formatted in RFC3339 time string.
func (o GetVPCVpcOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetVPCVpc) string { return v.UpdateTime }).(pulumi.StringOutput)
}

type GetVPCVpcArrayOutput struct{ *pulumi.OutputState }

func (GetVPCVpcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVPCVpc)(nil)).Elem()
}

func (o GetVPCVpcArrayOutput) ToGetVPCVpcArrayOutput() GetVPCVpcArrayOutput {
	return o
}

func (o GetVPCVpcArrayOutput) ToGetVPCVpcArrayOutputWithContext(ctx context.Context) GetVPCVpcArrayOutput {
	return o
}

func (o GetVPCVpcArrayOutput) Index(i pulumi.IntInput) GetVPCVpcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetVPCVpc {
		return vs[0].([]GetVPCVpc)[vs[1].(int)]
	}).(GetVPCVpcOutput)
}

func init() {
	pulumi.RegisterOutputType(VPCNetworkInfoOutput{})
	pulumi.RegisterOutputType(VPCNetworkInfoArrayOutput{})
	pulumi.RegisterOutputType(GetNATGatewayNatGatewayOutput{})
	pulumi.RegisterOutputType(GetNATGatewayNatGatewayArrayOutput{})
	pulumi.RegisterOutputType(GetNATGatewayNatGatewayIpSetOutput{})
	pulumi.RegisterOutputType(GetNATGatewayNatGatewayIpSetArrayOutput{})
	pulumi.RegisterOutputType(GetSubnetSubnetOutput{})
	pulumi.RegisterOutputType(GetSubnetSubnetArrayOutput{})
	pulumi.RegisterOutputType(GetVPCVpcOutput{})
	pulumi.RegisterOutputType(GetVPCVpcArrayOutput{})
}
