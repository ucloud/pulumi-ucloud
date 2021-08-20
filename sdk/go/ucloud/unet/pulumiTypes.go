// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package unet

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EIPIpSet struct {
	// Type of Elastic IP routes. Possible values are: `international` as international BGP IP and `bgp` as china mainland BGP IP.
	InternetType *string `pulumi:"internetType"`
	Ip           *string `pulumi:"ip"`
}

// EIPIpSetInput is an input type that accepts EIPIpSetArgs and EIPIpSetOutput values.
// You can construct a concrete instance of `EIPIpSetInput` via:
//
//          EIPIpSetArgs{...}
type EIPIpSetInput interface {
	pulumi.Input

	ToEIPIpSetOutput() EIPIpSetOutput
	ToEIPIpSetOutputWithContext(context.Context) EIPIpSetOutput
}

type EIPIpSetArgs struct {
	// Type of Elastic IP routes. Possible values are: `international` as international BGP IP and `bgp` as china mainland BGP IP.
	InternetType pulumi.StringPtrInput `pulumi:"internetType"`
	Ip           pulumi.StringPtrInput `pulumi:"ip"`
}

func (EIPIpSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EIPIpSet)(nil)).Elem()
}

func (i EIPIpSetArgs) ToEIPIpSetOutput() EIPIpSetOutput {
	return i.ToEIPIpSetOutputWithContext(context.Background())
}

func (i EIPIpSetArgs) ToEIPIpSetOutputWithContext(ctx context.Context) EIPIpSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EIPIpSetOutput)
}

// EIPIpSetArrayInput is an input type that accepts EIPIpSetArray and EIPIpSetArrayOutput values.
// You can construct a concrete instance of `EIPIpSetArrayInput` via:
//
//          EIPIpSetArray{ EIPIpSetArgs{...} }
type EIPIpSetArrayInput interface {
	pulumi.Input

	ToEIPIpSetArrayOutput() EIPIpSetArrayOutput
	ToEIPIpSetArrayOutputWithContext(context.Context) EIPIpSetArrayOutput
}

type EIPIpSetArray []EIPIpSetInput

func (EIPIpSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EIPIpSet)(nil)).Elem()
}

func (i EIPIpSetArray) ToEIPIpSetArrayOutput() EIPIpSetArrayOutput {
	return i.ToEIPIpSetArrayOutputWithContext(context.Background())
}

func (i EIPIpSetArray) ToEIPIpSetArrayOutputWithContext(ctx context.Context) EIPIpSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EIPIpSetArrayOutput)
}

type EIPIpSetOutput struct{ *pulumi.OutputState }

func (EIPIpSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EIPIpSet)(nil)).Elem()
}

func (o EIPIpSetOutput) ToEIPIpSetOutput() EIPIpSetOutput {
	return o
}

func (o EIPIpSetOutput) ToEIPIpSetOutputWithContext(ctx context.Context) EIPIpSetOutput {
	return o
}

// Type of Elastic IP routes. Possible values are: `international` as international BGP IP and `bgp` as china mainland BGP IP.
func (o EIPIpSetOutput) InternetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EIPIpSet) *string { return v.InternetType }).(pulumi.StringPtrOutput)
}

func (o EIPIpSetOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EIPIpSet) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

type EIPIpSetArrayOutput struct{ *pulumi.OutputState }

func (EIPIpSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EIPIpSet)(nil)).Elem()
}

func (o EIPIpSetArrayOutput) ToEIPIpSetArrayOutput() EIPIpSetArrayOutput {
	return o
}

func (o EIPIpSetArrayOutput) ToEIPIpSetArrayOutputWithContext(ctx context.Context) EIPIpSetArrayOutput {
	return o
}

func (o EIPIpSetArrayOutput) Index(i pulumi.IntInput) EIPIpSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EIPIpSet {
		return vs[0].([]EIPIpSet)[vs[1].(int)]
	}).(EIPIpSetOutput)
}

type EIPResource struct {
	// The ID of the resource with EIP attached.
	Id *string `pulumi:"id"`
	// The type of resource with EIP attached. Possible values are `instance` as instance, `lb` as load balancer.
	Type *string `pulumi:"type"`
}

// EIPResourceInput is an input type that accepts EIPResourceArgs and EIPResourceOutput values.
// You can construct a concrete instance of `EIPResourceInput` via:
//
//          EIPResourceArgs{...}
type EIPResourceInput interface {
	pulumi.Input

	ToEIPResourceOutput() EIPResourceOutput
	ToEIPResourceOutputWithContext(context.Context) EIPResourceOutput
}

type EIPResourceArgs struct {
	// The ID of the resource with EIP attached.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The type of resource with EIP attached. Possible values are `instance` as instance, `lb` as load balancer.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (EIPResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EIPResource)(nil)).Elem()
}

func (i EIPResourceArgs) ToEIPResourceOutput() EIPResourceOutput {
	return i.ToEIPResourceOutputWithContext(context.Background())
}

func (i EIPResourceArgs) ToEIPResourceOutputWithContext(ctx context.Context) EIPResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EIPResourceOutput)
}

func (i EIPResourceArgs) ToEIPResourcePtrOutput() EIPResourcePtrOutput {
	return i.ToEIPResourcePtrOutputWithContext(context.Background())
}

func (i EIPResourceArgs) ToEIPResourcePtrOutputWithContext(ctx context.Context) EIPResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EIPResourceOutput).ToEIPResourcePtrOutputWithContext(ctx)
}

// EIPResourcePtrInput is an input type that accepts EIPResourceArgs, EIPResourcePtr and EIPResourcePtrOutput values.
// You can construct a concrete instance of `EIPResourcePtrInput` via:
//
//          EIPResourceArgs{...}
//
//  or:
//
//          nil
type EIPResourcePtrInput interface {
	pulumi.Input

	ToEIPResourcePtrOutput() EIPResourcePtrOutput
	ToEIPResourcePtrOutputWithContext(context.Context) EIPResourcePtrOutput
}

type eipresourcePtrType EIPResourceArgs

func EIPResourcePtr(v *EIPResourceArgs) EIPResourcePtrInput {
	return (*eipresourcePtrType)(v)
}

func (*eipresourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EIPResource)(nil)).Elem()
}

func (i *eipresourcePtrType) ToEIPResourcePtrOutput() EIPResourcePtrOutput {
	return i.ToEIPResourcePtrOutputWithContext(context.Background())
}

func (i *eipresourcePtrType) ToEIPResourcePtrOutputWithContext(ctx context.Context) EIPResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EIPResourcePtrOutput)
}

type EIPResourceOutput struct{ *pulumi.OutputState }

func (EIPResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EIPResource)(nil)).Elem()
}

func (o EIPResourceOutput) ToEIPResourceOutput() EIPResourceOutput {
	return o
}

func (o EIPResourceOutput) ToEIPResourceOutputWithContext(ctx context.Context) EIPResourceOutput {
	return o
}

func (o EIPResourceOutput) ToEIPResourcePtrOutput() EIPResourcePtrOutput {
	return o.ToEIPResourcePtrOutputWithContext(context.Background())
}

func (o EIPResourceOutput) ToEIPResourcePtrOutputWithContext(ctx context.Context) EIPResourcePtrOutput {
	return o.ApplyT(func(v EIPResource) *EIPResource {
		return &v
	}).(EIPResourcePtrOutput)
}

// The ID of the resource with EIP attached.
func (o EIPResourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EIPResource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The type of resource with EIP attached. Possible values are `instance` as instance, `lb` as load balancer.
func (o EIPResourceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EIPResource) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type EIPResourcePtrOutput struct{ *pulumi.OutputState }

func (EIPResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EIPResource)(nil)).Elem()
}

func (o EIPResourcePtrOutput) ToEIPResourcePtrOutput() EIPResourcePtrOutput {
	return o
}

func (o EIPResourcePtrOutput) ToEIPResourcePtrOutputWithContext(ctx context.Context) EIPResourcePtrOutput {
	return o
}

func (o EIPResourcePtrOutput) Elem() EIPResourceOutput {
	return o.ApplyT(func(v *EIPResource) EIPResource { return *v }).(EIPResourceOutput)
}

// The ID of the resource with EIP attached.
func (o EIPResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EIPResource) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

// The type of resource with EIP attached. Possible values are `instance` as instance, `lb` as load balancer.
func (o EIPResourcePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EIPResource) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type SecurityGroupRule struct {
	// The cidr block of source.
	CidrBlock *string `pulumi:"cidrBlock"`
	// Authorization policy. Possible values are: `accept`, `drop`.
	Policy *string `pulumi:"policy"`
	// The range of port numbers, range: 1-65535. (eg: `port` or `port1-port2`).
	PortRange *string `pulumi:"portRange"`
	// Rule priority. Possible values are: `high`, `medium`, `low`.
	Priority *string `pulumi:"priority"`
	// The protocol. Possible values are: `tcp`, `udp`, `icmp`, `gre`.
	Protocol *string `pulumi:"protocol"`
}

// SecurityGroupRuleInput is an input type that accepts SecurityGroupRuleArgs and SecurityGroupRuleOutput values.
// You can construct a concrete instance of `SecurityGroupRuleInput` via:
//
//          SecurityGroupRuleArgs{...}
type SecurityGroupRuleInput interface {
	pulumi.Input

	ToSecurityGroupRuleOutput() SecurityGroupRuleOutput
	ToSecurityGroupRuleOutputWithContext(context.Context) SecurityGroupRuleOutput
}

type SecurityGroupRuleArgs struct {
	// The cidr block of source.
	CidrBlock pulumi.StringPtrInput `pulumi:"cidrBlock"`
	// Authorization policy. Possible values are: `accept`, `drop`.
	Policy pulumi.StringPtrInput `pulumi:"policy"`
	// The range of port numbers, range: 1-65535. (eg: `port` or `port1-port2`).
	PortRange pulumi.StringPtrInput `pulumi:"portRange"`
	// Rule priority. Possible values are: `high`, `medium`, `low`.
	Priority pulumi.StringPtrInput `pulumi:"priority"`
	// The protocol. Possible values are: `tcp`, `udp`, `icmp`, `gre`.
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
}

func (SecurityGroupRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroupRule)(nil)).Elem()
}

func (i SecurityGroupRuleArgs) ToSecurityGroupRuleOutput() SecurityGroupRuleOutput {
	return i.ToSecurityGroupRuleOutputWithContext(context.Background())
}

func (i SecurityGroupRuleArgs) ToSecurityGroupRuleOutputWithContext(ctx context.Context) SecurityGroupRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupRuleOutput)
}

// SecurityGroupRuleArrayInput is an input type that accepts SecurityGroupRuleArray and SecurityGroupRuleArrayOutput values.
// You can construct a concrete instance of `SecurityGroupRuleArrayInput` via:
//
//          SecurityGroupRuleArray{ SecurityGroupRuleArgs{...} }
type SecurityGroupRuleArrayInput interface {
	pulumi.Input

	ToSecurityGroupRuleArrayOutput() SecurityGroupRuleArrayOutput
	ToSecurityGroupRuleArrayOutputWithContext(context.Context) SecurityGroupRuleArrayOutput
}

type SecurityGroupRuleArray []SecurityGroupRuleInput

func (SecurityGroupRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityGroupRule)(nil)).Elem()
}

func (i SecurityGroupRuleArray) ToSecurityGroupRuleArrayOutput() SecurityGroupRuleArrayOutput {
	return i.ToSecurityGroupRuleArrayOutputWithContext(context.Background())
}

func (i SecurityGroupRuleArray) ToSecurityGroupRuleArrayOutputWithContext(ctx context.Context) SecurityGroupRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupRuleArrayOutput)
}

type SecurityGroupRuleOutput struct{ *pulumi.OutputState }

func (SecurityGroupRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroupRule)(nil)).Elem()
}

func (o SecurityGroupRuleOutput) ToSecurityGroupRuleOutput() SecurityGroupRuleOutput {
	return o
}

func (o SecurityGroupRuleOutput) ToSecurityGroupRuleOutputWithContext(ctx context.Context) SecurityGroupRuleOutput {
	return o
}

// The cidr block of source.
func (o SecurityGroupRuleOutput) CidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityGroupRule) *string { return v.CidrBlock }).(pulumi.StringPtrOutput)
}

// Authorization policy. Possible values are: `accept`, `drop`.
func (o SecurityGroupRuleOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityGroupRule) *string { return v.Policy }).(pulumi.StringPtrOutput)
}

// The range of port numbers, range: 1-65535. (eg: `port` or `port1-port2`).
func (o SecurityGroupRuleOutput) PortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityGroupRule) *string { return v.PortRange }).(pulumi.StringPtrOutput)
}

// Rule priority. Possible values are: `high`, `medium`, `low`.
func (o SecurityGroupRuleOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityGroupRule) *string { return v.Priority }).(pulumi.StringPtrOutput)
}

// The protocol. Possible values are: `tcp`, `udp`, `icmp`, `gre`.
func (o SecurityGroupRuleOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityGroupRule) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

type SecurityGroupRuleArrayOutput struct{ *pulumi.OutputState }

func (SecurityGroupRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityGroupRule)(nil)).Elem()
}

func (o SecurityGroupRuleArrayOutput) ToSecurityGroupRuleArrayOutput() SecurityGroupRuleArrayOutput {
	return o
}

func (o SecurityGroupRuleArrayOutput) ToSecurityGroupRuleArrayOutputWithContext(ctx context.Context) SecurityGroupRuleArrayOutput {
	return o
}

func (o SecurityGroupRuleArrayOutput) Index(i pulumi.IntInput) SecurityGroupRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityGroupRule {
		return vs[0].([]SecurityGroupRule)[vs[1].(int)]
	}).(SecurityGroupRuleOutput)
}

type GetEIPEip struct {
	// Maximum bandwidth to the elastic public network, measured in Mbps.
	Bandwidth int `pulumi:"bandwidth"`
	// The charge mode of Elastic IP. Possible values are: `traffic` as pay by traffic, `bandwidth` as pay by bandwidth.
	ChargeMode string `pulumi:"chargeMode"`
	// The charge type of Elastic IP. Possible values are: `year` as pay by year, `month` as pay by month, `dynamic` as pay by hour.
	ChargeType string `pulumi:"chargeType"`
	// The creation time of Elastic IP, formatted in RFC3339 time string.
	CreateTime string `pulumi:"createTime"`
	// The expiration time for Elastic IP, formatted in RFC3339 time string.
	ExpireTime string `pulumi:"expireTime"`
	// It is a nested type which documented below.
	IpSets []GetEIPEipIpSet `pulumi:"ipSets"`
	// The name of Elastic IP.
	Name string `pulumi:"name"`
	// The remarks of Elastic IP.
	Remark string `pulumi:"remark"`
	// Elastic IP status. Possible values are: `used` as in use, `free` as available and `freeze` as associating.
	Status string `pulumi:"status"`
	// A tag assigned to Elastic IP.
	Tag string `pulumi:"tag"`
}

// GetEIPEipInput is an input type that accepts GetEIPEipArgs and GetEIPEipOutput values.
// You can construct a concrete instance of `GetEIPEipInput` via:
//
//          GetEIPEipArgs{...}
type GetEIPEipInput interface {
	pulumi.Input

	ToGetEIPEipOutput() GetEIPEipOutput
	ToGetEIPEipOutputWithContext(context.Context) GetEIPEipOutput
}

type GetEIPEipArgs struct {
	// Maximum bandwidth to the elastic public network, measured in Mbps.
	Bandwidth pulumi.IntInput `pulumi:"bandwidth"`
	// The charge mode of Elastic IP. Possible values are: `traffic` as pay by traffic, `bandwidth` as pay by bandwidth.
	ChargeMode pulumi.StringInput `pulumi:"chargeMode"`
	// The charge type of Elastic IP. Possible values are: `year` as pay by year, `month` as pay by month, `dynamic` as pay by hour.
	ChargeType pulumi.StringInput `pulumi:"chargeType"`
	// The creation time of Elastic IP, formatted in RFC3339 time string.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// The expiration time for Elastic IP, formatted in RFC3339 time string.
	ExpireTime pulumi.StringInput `pulumi:"expireTime"`
	// It is a nested type which documented below.
	IpSets GetEIPEipIpSetArrayInput `pulumi:"ipSets"`
	// The name of Elastic IP.
	Name pulumi.StringInput `pulumi:"name"`
	// The remarks of Elastic IP.
	Remark pulumi.StringInput `pulumi:"remark"`
	// Elastic IP status. Possible values are: `used` as in use, `free` as available and `freeze` as associating.
	Status pulumi.StringInput `pulumi:"status"`
	// A tag assigned to Elastic IP.
	Tag pulumi.StringInput `pulumi:"tag"`
}

func (GetEIPEipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEIPEip)(nil)).Elem()
}

func (i GetEIPEipArgs) ToGetEIPEipOutput() GetEIPEipOutput {
	return i.ToGetEIPEipOutputWithContext(context.Background())
}

func (i GetEIPEipArgs) ToGetEIPEipOutputWithContext(ctx context.Context) GetEIPEipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetEIPEipOutput)
}

// GetEIPEipArrayInput is an input type that accepts GetEIPEipArray and GetEIPEipArrayOutput values.
// You can construct a concrete instance of `GetEIPEipArrayInput` via:
//
//          GetEIPEipArray{ GetEIPEipArgs{...} }
type GetEIPEipArrayInput interface {
	pulumi.Input

	ToGetEIPEipArrayOutput() GetEIPEipArrayOutput
	ToGetEIPEipArrayOutputWithContext(context.Context) GetEIPEipArrayOutput
}

type GetEIPEipArray []GetEIPEipInput

func (GetEIPEipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetEIPEip)(nil)).Elem()
}

func (i GetEIPEipArray) ToGetEIPEipArrayOutput() GetEIPEipArrayOutput {
	return i.ToGetEIPEipArrayOutputWithContext(context.Background())
}

func (i GetEIPEipArray) ToGetEIPEipArrayOutputWithContext(ctx context.Context) GetEIPEipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetEIPEipArrayOutput)
}

type GetEIPEipOutput struct{ *pulumi.OutputState }

func (GetEIPEipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEIPEip)(nil)).Elem()
}

func (o GetEIPEipOutput) ToGetEIPEipOutput() GetEIPEipOutput {
	return o
}

func (o GetEIPEipOutput) ToGetEIPEipOutputWithContext(ctx context.Context) GetEIPEipOutput {
	return o
}

// Maximum bandwidth to the elastic public network, measured in Mbps.
func (o GetEIPEipOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v GetEIPEip) int { return v.Bandwidth }).(pulumi.IntOutput)
}

// The charge mode of Elastic IP. Possible values are: `traffic` as pay by traffic, `bandwidth` as pay by bandwidth.
func (o GetEIPEipOutput) ChargeMode() pulumi.StringOutput {
	return o.ApplyT(func(v GetEIPEip) string { return v.ChargeMode }).(pulumi.StringOutput)
}

// The charge type of Elastic IP. Possible values are: `year` as pay by year, `month` as pay by month, `dynamic` as pay by hour.
func (o GetEIPEipOutput) ChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v GetEIPEip) string { return v.ChargeType }).(pulumi.StringOutput)
}

// The creation time of Elastic IP, formatted in RFC3339 time string.
func (o GetEIPEipOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetEIPEip) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The expiration time for Elastic IP, formatted in RFC3339 time string.
func (o GetEIPEipOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetEIPEip) string { return v.ExpireTime }).(pulumi.StringOutput)
}

// It is a nested type which documented below.
func (o GetEIPEipOutput) IpSets() GetEIPEipIpSetArrayOutput {
	return o.ApplyT(func(v GetEIPEip) []GetEIPEipIpSet { return v.IpSets }).(GetEIPEipIpSetArrayOutput)
}

// The name of Elastic IP.
func (o GetEIPEipOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetEIPEip) string { return v.Name }).(pulumi.StringOutput)
}

// The remarks of Elastic IP.
func (o GetEIPEipOutput) Remark() pulumi.StringOutput {
	return o.ApplyT(func(v GetEIPEip) string { return v.Remark }).(pulumi.StringOutput)
}

// Elastic IP status. Possible values are: `used` as in use, `free` as available and `freeze` as associating.
func (o GetEIPEipOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetEIPEip) string { return v.Status }).(pulumi.StringOutput)
}

// A tag assigned to Elastic IP.
func (o GetEIPEipOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v GetEIPEip) string { return v.Tag }).(pulumi.StringOutput)
}

type GetEIPEipArrayOutput struct{ *pulumi.OutputState }

func (GetEIPEipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetEIPEip)(nil)).Elem()
}

func (o GetEIPEipArrayOutput) ToGetEIPEipArrayOutput() GetEIPEipArrayOutput {
	return o
}

func (o GetEIPEipArrayOutput) ToGetEIPEipArrayOutputWithContext(ctx context.Context) GetEIPEipArrayOutput {
	return o
}

func (o GetEIPEipArrayOutput) Index(i pulumi.IntInput) GetEIPEipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetEIPEip {
		return vs[0].([]GetEIPEip)[vs[1].(int)]
	}).(GetEIPEipOutput)
}

type GetEIPEipIpSet struct {
	// Type of Elastic IP routes.
	InternetType string `pulumi:"internetType"`
	// Elastic IP address.
	Ip string `pulumi:"ip"`
}

// GetEIPEipIpSetInput is an input type that accepts GetEIPEipIpSetArgs and GetEIPEipIpSetOutput values.
// You can construct a concrete instance of `GetEIPEipIpSetInput` via:
//
//          GetEIPEipIpSetArgs{...}
type GetEIPEipIpSetInput interface {
	pulumi.Input

	ToGetEIPEipIpSetOutput() GetEIPEipIpSetOutput
	ToGetEIPEipIpSetOutputWithContext(context.Context) GetEIPEipIpSetOutput
}

type GetEIPEipIpSetArgs struct {
	// Type of Elastic IP routes.
	InternetType pulumi.StringInput `pulumi:"internetType"`
	// Elastic IP address.
	Ip pulumi.StringInput `pulumi:"ip"`
}

func (GetEIPEipIpSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEIPEipIpSet)(nil)).Elem()
}

func (i GetEIPEipIpSetArgs) ToGetEIPEipIpSetOutput() GetEIPEipIpSetOutput {
	return i.ToGetEIPEipIpSetOutputWithContext(context.Background())
}

func (i GetEIPEipIpSetArgs) ToGetEIPEipIpSetOutputWithContext(ctx context.Context) GetEIPEipIpSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetEIPEipIpSetOutput)
}

// GetEIPEipIpSetArrayInput is an input type that accepts GetEIPEipIpSetArray and GetEIPEipIpSetArrayOutput values.
// You can construct a concrete instance of `GetEIPEipIpSetArrayInput` via:
//
//          GetEIPEipIpSetArray{ GetEIPEipIpSetArgs{...} }
type GetEIPEipIpSetArrayInput interface {
	pulumi.Input

	ToGetEIPEipIpSetArrayOutput() GetEIPEipIpSetArrayOutput
	ToGetEIPEipIpSetArrayOutputWithContext(context.Context) GetEIPEipIpSetArrayOutput
}

type GetEIPEipIpSetArray []GetEIPEipIpSetInput

func (GetEIPEipIpSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetEIPEipIpSet)(nil)).Elem()
}

func (i GetEIPEipIpSetArray) ToGetEIPEipIpSetArrayOutput() GetEIPEipIpSetArrayOutput {
	return i.ToGetEIPEipIpSetArrayOutputWithContext(context.Background())
}

func (i GetEIPEipIpSetArray) ToGetEIPEipIpSetArrayOutputWithContext(ctx context.Context) GetEIPEipIpSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetEIPEipIpSetArrayOutput)
}

type GetEIPEipIpSetOutput struct{ *pulumi.OutputState }

func (GetEIPEipIpSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEIPEipIpSet)(nil)).Elem()
}

func (o GetEIPEipIpSetOutput) ToGetEIPEipIpSetOutput() GetEIPEipIpSetOutput {
	return o
}

func (o GetEIPEipIpSetOutput) ToGetEIPEipIpSetOutputWithContext(ctx context.Context) GetEIPEipIpSetOutput {
	return o
}

// Type of Elastic IP routes.
func (o GetEIPEipIpSetOutput) InternetType() pulumi.StringOutput {
	return o.ApplyT(func(v GetEIPEipIpSet) string { return v.InternetType }).(pulumi.StringOutput)
}

// Elastic IP address.
func (o GetEIPEipIpSetOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v GetEIPEipIpSet) string { return v.Ip }).(pulumi.StringOutput)
}

type GetEIPEipIpSetArrayOutput struct{ *pulumi.OutputState }

func (GetEIPEipIpSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetEIPEipIpSet)(nil)).Elem()
}

func (o GetEIPEipIpSetArrayOutput) ToGetEIPEipIpSetArrayOutput() GetEIPEipIpSetArrayOutput {
	return o
}

func (o GetEIPEipIpSetArrayOutput) ToGetEIPEipIpSetArrayOutputWithContext(ctx context.Context) GetEIPEipIpSetArrayOutput {
	return o
}

func (o GetEIPEipIpSetArrayOutput) Index(i pulumi.IntInput) GetEIPEipIpSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetEIPEipIpSet {
		return vs[0].([]GetEIPEipIpSet)[vs[1].(int)]
	}).(GetEIPEipIpSetOutput)
}

type GetSecurityGroupSecurityGroup struct {
	// The time of creation for the security group, formatted in RFC3339 time string.
	CreateTime string `pulumi:"createTime"`
	// The ID of Security Group.
	Id string `pulumi:"id"`
	// The name of Security Group.
	Name string `pulumi:"name"`
	// The remarks of the security group.
	Remark string `pulumi:"remark"`
	// It is a nested type which documented below.
	Rules []GetSecurityGroupSecurityGroupRule `pulumi:"rules"`
	// A tag assigned to the security group.
	Tag string `pulumi:"tag"`
	// The type of Security Group. Possible values are: `recommendWeb` as the default Web security group that UCloud recommend to users, default opened port include 80, 443, 22, 3389, `recommendNonWeb` as the default non Web security group that UCloud recommend to users, default opened port include 22, 3389, `userDefined` as the security groups defined by users. You may refer to [security group](https://docs.ucloud.cn/network/firewall/firewall.html).
	Type string `pulumi:"type"`
}

// GetSecurityGroupSecurityGroupInput is an input type that accepts GetSecurityGroupSecurityGroupArgs and GetSecurityGroupSecurityGroupOutput values.
// You can construct a concrete instance of `GetSecurityGroupSecurityGroupInput` via:
//
//          GetSecurityGroupSecurityGroupArgs{...}
type GetSecurityGroupSecurityGroupInput interface {
	pulumi.Input

	ToGetSecurityGroupSecurityGroupOutput() GetSecurityGroupSecurityGroupOutput
	ToGetSecurityGroupSecurityGroupOutputWithContext(context.Context) GetSecurityGroupSecurityGroupOutput
}

type GetSecurityGroupSecurityGroupArgs struct {
	// The time of creation for the security group, formatted in RFC3339 time string.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// The ID of Security Group.
	Id pulumi.StringInput `pulumi:"id"`
	// The name of Security Group.
	Name pulumi.StringInput `pulumi:"name"`
	// The remarks of the security group.
	Remark pulumi.StringInput `pulumi:"remark"`
	// It is a nested type which documented below.
	Rules GetSecurityGroupSecurityGroupRuleArrayInput `pulumi:"rules"`
	// A tag assigned to the security group.
	Tag pulumi.StringInput `pulumi:"tag"`
	// The type of Security Group. Possible values are: `recommendWeb` as the default Web security group that UCloud recommend to users, default opened port include 80, 443, 22, 3389, `recommendNonWeb` as the default non Web security group that UCloud recommend to users, default opened port include 22, 3389, `userDefined` as the security groups defined by users. You may refer to [security group](https://docs.ucloud.cn/network/firewall/firewall.html).
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetSecurityGroupSecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecurityGroupSecurityGroup)(nil)).Elem()
}

func (i GetSecurityGroupSecurityGroupArgs) ToGetSecurityGroupSecurityGroupOutput() GetSecurityGroupSecurityGroupOutput {
	return i.ToGetSecurityGroupSecurityGroupOutputWithContext(context.Background())
}

func (i GetSecurityGroupSecurityGroupArgs) ToGetSecurityGroupSecurityGroupOutputWithContext(ctx context.Context) GetSecurityGroupSecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSecurityGroupSecurityGroupOutput)
}

// GetSecurityGroupSecurityGroupArrayInput is an input type that accepts GetSecurityGroupSecurityGroupArray and GetSecurityGroupSecurityGroupArrayOutput values.
// You can construct a concrete instance of `GetSecurityGroupSecurityGroupArrayInput` via:
//
//          GetSecurityGroupSecurityGroupArray{ GetSecurityGroupSecurityGroupArgs{...} }
type GetSecurityGroupSecurityGroupArrayInput interface {
	pulumi.Input

	ToGetSecurityGroupSecurityGroupArrayOutput() GetSecurityGroupSecurityGroupArrayOutput
	ToGetSecurityGroupSecurityGroupArrayOutputWithContext(context.Context) GetSecurityGroupSecurityGroupArrayOutput
}

type GetSecurityGroupSecurityGroupArray []GetSecurityGroupSecurityGroupInput

func (GetSecurityGroupSecurityGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSecurityGroupSecurityGroup)(nil)).Elem()
}

func (i GetSecurityGroupSecurityGroupArray) ToGetSecurityGroupSecurityGroupArrayOutput() GetSecurityGroupSecurityGroupArrayOutput {
	return i.ToGetSecurityGroupSecurityGroupArrayOutputWithContext(context.Background())
}

func (i GetSecurityGroupSecurityGroupArray) ToGetSecurityGroupSecurityGroupArrayOutputWithContext(ctx context.Context) GetSecurityGroupSecurityGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSecurityGroupSecurityGroupArrayOutput)
}

type GetSecurityGroupSecurityGroupOutput struct{ *pulumi.OutputState }

func (GetSecurityGroupSecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecurityGroupSecurityGroup)(nil)).Elem()
}

func (o GetSecurityGroupSecurityGroupOutput) ToGetSecurityGroupSecurityGroupOutput() GetSecurityGroupSecurityGroupOutput {
	return o
}

func (o GetSecurityGroupSecurityGroupOutput) ToGetSecurityGroupSecurityGroupOutputWithContext(ctx context.Context) GetSecurityGroupSecurityGroupOutput {
	return o
}

// The time of creation for the security group, formatted in RFC3339 time string.
func (o GetSecurityGroupSecurityGroupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupSecurityGroup) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The ID of Security Group.
func (o GetSecurityGroupSecurityGroupOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupSecurityGroup) string { return v.Id }).(pulumi.StringOutput)
}

// The name of Security Group.
func (o GetSecurityGroupSecurityGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupSecurityGroup) string { return v.Name }).(pulumi.StringOutput)
}

// The remarks of the security group.
func (o GetSecurityGroupSecurityGroupOutput) Remark() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupSecurityGroup) string { return v.Remark }).(pulumi.StringOutput)
}

// It is a nested type which documented below.
func (o GetSecurityGroupSecurityGroupOutput) Rules() GetSecurityGroupSecurityGroupRuleArrayOutput {
	return o.ApplyT(func(v GetSecurityGroupSecurityGroup) []GetSecurityGroupSecurityGroupRule { return v.Rules }).(GetSecurityGroupSecurityGroupRuleArrayOutput)
}

// A tag assigned to the security group.
func (o GetSecurityGroupSecurityGroupOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupSecurityGroup) string { return v.Tag }).(pulumi.StringOutput)
}

// The type of Security Group. Possible values are: `recommendWeb` as the default Web security group that UCloud recommend to users, default opened port include 80, 443, 22, 3389, `recommendNonWeb` as the default non Web security group that UCloud recommend to users, default opened port include 22, 3389, `userDefined` as the security groups defined by users. You may refer to [security group](https://docs.ucloud.cn/network/firewall/firewall.html).
func (o GetSecurityGroupSecurityGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupSecurityGroup) string { return v.Type }).(pulumi.StringOutput)
}

type GetSecurityGroupSecurityGroupArrayOutput struct{ *pulumi.OutputState }

func (GetSecurityGroupSecurityGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSecurityGroupSecurityGroup)(nil)).Elem()
}

func (o GetSecurityGroupSecurityGroupArrayOutput) ToGetSecurityGroupSecurityGroupArrayOutput() GetSecurityGroupSecurityGroupArrayOutput {
	return o
}

func (o GetSecurityGroupSecurityGroupArrayOutput) ToGetSecurityGroupSecurityGroupArrayOutputWithContext(ctx context.Context) GetSecurityGroupSecurityGroupArrayOutput {
	return o
}

func (o GetSecurityGroupSecurityGroupArrayOutput) Index(i pulumi.IntInput) GetSecurityGroupSecurityGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSecurityGroupSecurityGroup {
		return vs[0].([]GetSecurityGroupSecurityGroup)[vs[1].(int)]
	}).(GetSecurityGroupSecurityGroupOutput)
}

type GetSecurityGroupSecurityGroupRule struct {
	// The cidr block of source.
	CidrBlock string `pulumi:"cidrBlock"`
	// Authorization policy. Can be either `accept` or `drop`.
	Policy string `pulumi:"policy"`
	// The range of port numbers, range: 1-65535. (eg: `port` or `port1-port2`).
	PortRange string `pulumi:"portRange"`
	// Rule priority. Can be `high`, `medium`, `low`.
	Priority string `pulumi:"priority"`
	// The protocol. Can be `tcp`, `udp`, `icmp`, `gre`.
	Protocol string `pulumi:"protocol"`
}

// GetSecurityGroupSecurityGroupRuleInput is an input type that accepts GetSecurityGroupSecurityGroupRuleArgs and GetSecurityGroupSecurityGroupRuleOutput values.
// You can construct a concrete instance of `GetSecurityGroupSecurityGroupRuleInput` via:
//
//          GetSecurityGroupSecurityGroupRuleArgs{...}
type GetSecurityGroupSecurityGroupRuleInput interface {
	pulumi.Input

	ToGetSecurityGroupSecurityGroupRuleOutput() GetSecurityGroupSecurityGroupRuleOutput
	ToGetSecurityGroupSecurityGroupRuleOutputWithContext(context.Context) GetSecurityGroupSecurityGroupRuleOutput
}

type GetSecurityGroupSecurityGroupRuleArgs struct {
	// The cidr block of source.
	CidrBlock pulumi.StringInput `pulumi:"cidrBlock"`
	// Authorization policy. Can be either `accept` or `drop`.
	Policy pulumi.StringInput `pulumi:"policy"`
	// The range of port numbers, range: 1-65535. (eg: `port` or `port1-port2`).
	PortRange pulumi.StringInput `pulumi:"portRange"`
	// Rule priority. Can be `high`, `medium`, `low`.
	Priority pulumi.StringInput `pulumi:"priority"`
	// The protocol. Can be `tcp`, `udp`, `icmp`, `gre`.
	Protocol pulumi.StringInput `pulumi:"protocol"`
}

func (GetSecurityGroupSecurityGroupRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecurityGroupSecurityGroupRule)(nil)).Elem()
}

func (i GetSecurityGroupSecurityGroupRuleArgs) ToGetSecurityGroupSecurityGroupRuleOutput() GetSecurityGroupSecurityGroupRuleOutput {
	return i.ToGetSecurityGroupSecurityGroupRuleOutputWithContext(context.Background())
}

func (i GetSecurityGroupSecurityGroupRuleArgs) ToGetSecurityGroupSecurityGroupRuleOutputWithContext(ctx context.Context) GetSecurityGroupSecurityGroupRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSecurityGroupSecurityGroupRuleOutput)
}

// GetSecurityGroupSecurityGroupRuleArrayInput is an input type that accepts GetSecurityGroupSecurityGroupRuleArray and GetSecurityGroupSecurityGroupRuleArrayOutput values.
// You can construct a concrete instance of `GetSecurityGroupSecurityGroupRuleArrayInput` via:
//
//          GetSecurityGroupSecurityGroupRuleArray{ GetSecurityGroupSecurityGroupRuleArgs{...} }
type GetSecurityGroupSecurityGroupRuleArrayInput interface {
	pulumi.Input

	ToGetSecurityGroupSecurityGroupRuleArrayOutput() GetSecurityGroupSecurityGroupRuleArrayOutput
	ToGetSecurityGroupSecurityGroupRuleArrayOutputWithContext(context.Context) GetSecurityGroupSecurityGroupRuleArrayOutput
}

type GetSecurityGroupSecurityGroupRuleArray []GetSecurityGroupSecurityGroupRuleInput

func (GetSecurityGroupSecurityGroupRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSecurityGroupSecurityGroupRule)(nil)).Elem()
}

func (i GetSecurityGroupSecurityGroupRuleArray) ToGetSecurityGroupSecurityGroupRuleArrayOutput() GetSecurityGroupSecurityGroupRuleArrayOutput {
	return i.ToGetSecurityGroupSecurityGroupRuleArrayOutputWithContext(context.Background())
}

func (i GetSecurityGroupSecurityGroupRuleArray) ToGetSecurityGroupSecurityGroupRuleArrayOutputWithContext(ctx context.Context) GetSecurityGroupSecurityGroupRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSecurityGroupSecurityGroupRuleArrayOutput)
}

type GetSecurityGroupSecurityGroupRuleOutput struct{ *pulumi.OutputState }

func (GetSecurityGroupSecurityGroupRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecurityGroupSecurityGroupRule)(nil)).Elem()
}

func (o GetSecurityGroupSecurityGroupRuleOutput) ToGetSecurityGroupSecurityGroupRuleOutput() GetSecurityGroupSecurityGroupRuleOutput {
	return o
}

func (o GetSecurityGroupSecurityGroupRuleOutput) ToGetSecurityGroupSecurityGroupRuleOutputWithContext(ctx context.Context) GetSecurityGroupSecurityGroupRuleOutput {
	return o
}

// The cidr block of source.
func (o GetSecurityGroupSecurityGroupRuleOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupSecurityGroupRule) string { return v.CidrBlock }).(pulumi.StringOutput)
}

// Authorization policy. Can be either `accept` or `drop`.
func (o GetSecurityGroupSecurityGroupRuleOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupSecurityGroupRule) string { return v.Policy }).(pulumi.StringOutput)
}

// The range of port numbers, range: 1-65535. (eg: `port` or `port1-port2`).
func (o GetSecurityGroupSecurityGroupRuleOutput) PortRange() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupSecurityGroupRule) string { return v.PortRange }).(pulumi.StringOutput)
}

// Rule priority. Can be `high`, `medium`, `low`.
func (o GetSecurityGroupSecurityGroupRuleOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupSecurityGroupRule) string { return v.Priority }).(pulumi.StringOutput)
}

// The protocol. Can be `tcp`, `udp`, `icmp`, `gre`.
func (o GetSecurityGroupSecurityGroupRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupSecurityGroupRule) string { return v.Protocol }).(pulumi.StringOutput)
}

type GetSecurityGroupSecurityGroupRuleArrayOutput struct{ *pulumi.OutputState }

func (GetSecurityGroupSecurityGroupRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSecurityGroupSecurityGroupRule)(nil)).Elem()
}

func (o GetSecurityGroupSecurityGroupRuleArrayOutput) ToGetSecurityGroupSecurityGroupRuleArrayOutput() GetSecurityGroupSecurityGroupRuleArrayOutput {
	return o
}

func (o GetSecurityGroupSecurityGroupRuleArrayOutput) ToGetSecurityGroupSecurityGroupRuleArrayOutputWithContext(ctx context.Context) GetSecurityGroupSecurityGroupRuleArrayOutput {
	return o
}

func (o GetSecurityGroupSecurityGroupRuleArrayOutput) Index(i pulumi.IntInput) GetSecurityGroupSecurityGroupRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSecurityGroupSecurityGroupRule {
		return vs[0].([]GetSecurityGroupSecurityGroupRule)[vs[1].(int)]
	}).(GetSecurityGroupSecurityGroupRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(EIPIpSetOutput{})
	pulumi.RegisterOutputType(EIPIpSetArrayOutput{})
	pulumi.RegisterOutputType(EIPResourceOutput{})
	pulumi.RegisterOutputType(EIPResourcePtrOutput{})
	pulumi.RegisterOutputType(SecurityGroupRuleOutput{})
	pulumi.RegisterOutputType(SecurityGroupRuleArrayOutput{})
	pulumi.RegisterOutputType(GetEIPEipOutput{})
	pulumi.RegisterOutputType(GetEIPEipArrayOutput{})
	pulumi.RegisterOutputType(GetEIPEipIpSetOutput{})
	pulumi.RegisterOutputType(GetEIPEipIpSetArrayOutput{})
	pulumi.RegisterOutputType(GetSecurityGroupSecurityGroupOutput{})
	pulumi.RegisterOutputType(GetSecurityGroupSecurityGroupArrayOutput{})
	pulumi.RegisterOutputType(GetSecurityGroupSecurityGroupRuleOutput{})
	pulumi.RegisterOutputType(GetSecurityGroupSecurityGroupRuleArrayOutput{})
}
