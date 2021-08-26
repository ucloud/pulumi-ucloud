// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package umem

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The UCloud Memcache instance is a key-value online storage service compatible with the Memcached protocol.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-ucloud/sdk/go/ucloud/uaccount"
// 	"github.com/pulumi/pulumi-ucloud/sdk/go/ucloud/umem"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_default, err := uaccount.GetZone(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = umem.NewMemcachedInstance(ctx, "master", &umem.MemcachedInstanceArgs{
// 			AvailabilityZone: pulumi.String(_default.Zones[0].Id),
// 			InstanceType:     pulumi.String("memcache-master-2"),
// 			Tag:              pulumi.String("tf-example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type MemcachedInstance struct {
	pulumi.CustomResourceState

	// Availability zone where Memcache instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The charge type of Memcache instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType pulumi.StringOutput `pulumi:"chargeType"`
	// The creation time of Memcache instance, formatted by RFC3339 time string.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The duration that you will buy the Memcache instance (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration pulumi.IntPtrOutput `pulumi:"duration"`
	// The expiration time of Memcache instance, formatted by RFC3339 time string.
	ExpireTime   pulumi.StringOutput `pulumi:"expireTime"`
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// ip_set is a nested type. ipSet documented below.
	IpSets MemcachedInstanceIpSetArrayOutput `pulumi:"ipSets"`
	Name   pulumi.StringOutput               `pulumi:"name"`
	// The status of KV Memcache instance.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of subnet linked to the Memcache instance.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// A tag assigned to Memcache instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringOutput `pulumi:"tag"`
	// The ID of VPC linked to the Memcache instance.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewMemcachedInstance registers a new resource with the given unique name, arguments, and options.
func NewMemcachedInstance(ctx *pulumi.Context,
	name string, args *MemcachedInstanceArgs, opts ...pulumi.ResourceOption) (*MemcachedInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	var resource MemcachedInstance
	err := ctx.RegisterResource("ucloud:umem/memcachedInstance:MemcachedInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMemcachedInstance gets an existing MemcachedInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMemcachedInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MemcachedInstanceState, opts ...pulumi.ResourceOption) (*MemcachedInstance, error) {
	var resource MemcachedInstance
	err := ctx.ReadResource("ucloud:umem/memcachedInstance:MemcachedInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MemcachedInstance resources.
type memcachedInstanceState struct {
	// Availability zone where Memcache instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The charge type of Memcache instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType *string `pulumi:"chargeType"`
	// The creation time of Memcache instance, formatted by RFC3339 time string.
	CreateTime *string `pulumi:"createTime"`
	// The duration that you will buy the Memcache instance (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration *int `pulumi:"duration"`
	// The expiration time of Memcache instance, formatted by RFC3339 time string.
	ExpireTime   *string `pulumi:"expireTime"`
	InstanceType *string `pulumi:"instanceType"`
	// ip_set is a nested type. ipSet documented below.
	IpSets []MemcachedInstanceIpSet `pulumi:"ipSets"`
	Name   *string                  `pulumi:"name"`
	// The status of KV Memcache instance.
	Status *string `pulumi:"status"`
	// The ID of subnet linked to the Memcache instance.
	SubnetId *string `pulumi:"subnetId"`
	// A tag assigned to Memcache instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag *string `pulumi:"tag"`
	// The ID of VPC linked to the Memcache instance.
	VpcId *string `pulumi:"vpcId"`
}

type MemcachedInstanceState struct {
	// Availability zone where Memcache instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone pulumi.StringPtrInput
	// The charge type of Memcache instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType pulumi.StringPtrInput
	// The creation time of Memcache instance, formatted by RFC3339 time string.
	CreateTime pulumi.StringPtrInput
	// The duration that you will buy the Memcache instance (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration pulumi.IntPtrInput
	// The expiration time of Memcache instance, formatted by RFC3339 time string.
	ExpireTime   pulumi.StringPtrInput
	InstanceType pulumi.StringPtrInput
	// ip_set is a nested type. ipSet documented below.
	IpSets MemcachedInstanceIpSetArrayInput
	Name   pulumi.StringPtrInput
	// The status of KV Memcache instance.
	Status pulumi.StringPtrInput
	// The ID of subnet linked to the Memcache instance.
	SubnetId pulumi.StringPtrInput
	// A tag assigned to Memcache instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringPtrInput
	// The ID of VPC linked to the Memcache instance.
	VpcId pulumi.StringPtrInput
}

func (MemcachedInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*memcachedInstanceState)(nil)).Elem()
}

type memcachedInstanceArgs struct {
	// Availability zone where Memcache instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone string `pulumi:"availabilityZone"`
	// The charge type of Memcache instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType *string `pulumi:"chargeType"`
	// The duration that you will buy the Memcache instance (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration     *int    `pulumi:"duration"`
	InstanceType string  `pulumi:"instanceType"`
	Name         *string `pulumi:"name"`
	// The ID of subnet linked to the Memcache instance.
	SubnetId *string `pulumi:"subnetId"`
	// A tag assigned to Memcache instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag *string `pulumi:"tag"`
	// The ID of VPC linked to the Memcache instance.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a MemcachedInstance resource.
type MemcachedInstanceArgs struct {
	// Availability zone where Memcache instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone pulumi.StringInput
	// The charge type of Memcache instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType pulumi.StringPtrInput
	// The duration that you will buy the Memcache instance (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration     pulumi.IntPtrInput
	InstanceType pulumi.StringInput
	Name         pulumi.StringPtrInput
	// The ID of subnet linked to the Memcache instance.
	SubnetId pulumi.StringPtrInput
	// A tag assigned to Memcache instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringPtrInput
	// The ID of VPC linked to the Memcache instance.
	VpcId pulumi.StringPtrInput
}

func (MemcachedInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*memcachedInstanceArgs)(nil)).Elem()
}

type MemcachedInstanceInput interface {
	pulumi.Input

	ToMemcachedInstanceOutput() MemcachedInstanceOutput
	ToMemcachedInstanceOutputWithContext(ctx context.Context) MemcachedInstanceOutput
}

func (*MemcachedInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*MemcachedInstance)(nil))
}

func (i *MemcachedInstance) ToMemcachedInstanceOutput() MemcachedInstanceOutput {
	return i.ToMemcachedInstanceOutputWithContext(context.Background())
}

func (i *MemcachedInstance) ToMemcachedInstanceOutputWithContext(ctx context.Context) MemcachedInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemcachedInstanceOutput)
}

func (i *MemcachedInstance) ToMemcachedInstancePtrOutput() MemcachedInstancePtrOutput {
	return i.ToMemcachedInstancePtrOutputWithContext(context.Background())
}

func (i *MemcachedInstance) ToMemcachedInstancePtrOutputWithContext(ctx context.Context) MemcachedInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemcachedInstancePtrOutput)
}

type MemcachedInstancePtrInput interface {
	pulumi.Input

	ToMemcachedInstancePtrOutput() MemcachedInstancePtrOutput
	ToMemcachedInstancePtrOutputWithContext(ctx context.Context) MemcachedInstancePtrOutput
}

type memcachedInstancePtrType MemcachedInstanceArgs

func (*memcachedInstancePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MemcachedInstance)(nil))
}

func (i *memcachedInstancePtrType) ToMemcachedInstancePtrOutput() MemcachedInstancePtrOutput {
	return i.ToMemcachedInstancePtrOutputWithContext(context.Background())
}

func (i *memcachedInstancePtrType) ToMemcachedInstancePtrOutputWithContext(ctx context.Context) MemcachedInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemcachedInstancePtrOutput)
}

// MemcachedInstanceArrayInput is an input type that accepts MemcachedInstanceArray and MemcachedInstanceArrayOutput values.
// You can construct a concrete instance of `MemcachedInstanceArrayInput` via:
//
//          MemcachedInstanceArray{ MemcachedInstanceArgs{...} }
type MemcachedInstanceArrayInput interface {
	pulumi.Input

	ToMemcachedInstanceArrayOutput() MemcachedInstanceArrayOutput
	ToMemcachedInstanceArrayOutputWithContext(context.Context) MemcachedInstanceArrayOutput
}

type MemcachedInstanceArray []MemcachedInstanceInput

func (MemcachedInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*MemcachedInstance)(nil))
}

func (i MemcachedInstanceArray) ToMemcachedInstanceArrayOutput() MemcachedInstanceArrayOutput {
	return i.ToMemcachedInstanceArrayOutputWithContext(context.Background())
}

func (i MemcachedInstanceArray) ToMemcachedInstanceArrayOutputWithContext(ctx context.Context) MemcachedInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemcachedInstanceArrayOutput)
}

// MemcachedInstanceMapInput is an input type that accepts MemcachedInstanceMap and MemcachedInstanceMapOutput values.
// You can construct a concrete instance of `MemcachedInstanceMapInput` via:
//
//          MemcachedInstanceMap{ "key": MemcachedInstanceArgs{...} }
type MemcachedInstanceMapInput interface {
	pulumi.Input

	ToMemcachedInstanceMapOutput() MemcachedInstanceMapOutput
	ToMemcachedInstanceMapOutputWithContext(context.Context) MemcachedInstanceMapOutput
}

type MemcachedInstanceMap map[string]MemcachedInstanceInput

func (MemcachedInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*MemcachedInstance)(nil))
}

func (i MemcachedInstanceMap) ToMemcachedInstanceMapOutput() MemcachedInstanceMapOutput {
	return i.ToMemcachedInstanceMapOutputWithContext(context.Background())
}

func (i MemcachedInstanceMap) ToMemcachedInstanceMapOutputWithContext(ctx context.Context) MemcachedInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemcachedInstanceMapOutput)
}

type MemcachedInstanceOutput struct {
	*pulumi.OutputState
}

func (MemcachedInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MemcachedInstance)(nil))
}

func (o MemcachedInstanceOutput) ToMemcachedInstanceOutput() MemcachedInstanceOutput {
	return o
}

func (o MemcachedInstanceOutput) ToMemcachedInstanceOutputWithContext(ctx context.Context) MemcachedInstanceOutput {
	return o
}

func (o MemcachedInstanceOutput) ToMemcachedInstancePtrOutput() MemcachedInstancePtrOutput {
	return o.ToMemcachedInstancePtrOutputWithContext(context.Background())
}

func (o MemcachedInstanceOutput) ToMemcachedInstancePtrOutputWithContext(ctx context.Context) MemcachedInstancePtrOutput {
	return o.ApplyT(func(v MemcachedInstance) *MemcachedInstance {
		return &v
	}).(MemcachedInstancePtrOutput)
}

type MemcachedInstancePtrOutput struct {
	*pulumi.OutputState
}

func (MemcachedInstancePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MemcachedInstance)(nil))
}

func (o MemcachedInstancePtrOutput) ToMemcachedInstancePtrOutput() MemcachedInstancePtrOutput {
	return o
}

func (o MemcachedInstancePtrOutput) ToMemcachedInstancePtrOutputWithContext(ctx context.Context) MemcachedInstancePtrOutput {
	return o
}

type MemcachedInstanceArrayOutput struct{ *pulumi.OutputState }

func (MemcachedInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MemcachedInstance)(nil))
}

func (o MemcachedInstanceArrayOutput) ToMemcachedInstanceArrayOutput() MemcachedInstanceArrayOutput {
	return o
}

func (o MemcachedInstanceArrayOutput) ToMemcachedInstanceArrayOutputWithContext(ctx context.Context) MemcachedInstanceArrayOutput {
	return o
}

func (o MemcachedInstanceArrayOutput) Index(i pulumi.IntInput) MemcachedInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MemcachedInstance {
		return vs[0].([]MemcachedInstance)[vs[1].(int)]
	}).(MemcachedInstanceOutput)
}

type MemcachedInstanceMapOutput struct{ *pulumi.OutputState }

func (MemcachedInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MemcachedInstance)(nil))
}

func (o MemcachedInstanceMapOutput) ToMemcachedInstanceMapOutput() MemcachedInstanceMapOutput {
	return o
}

func (o MemcachedInstanceMapOutput) ToMemcachedInstanceMapOutputWithContext(ctx context.Context) MemcachedInstanceMapOutput {
	return o
}

func (o MemcachedInstanceMapOutput) MapIndex(k pulumi.StringInput) MemcachedInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MemcachedInstance {
		return vs[0].(map[string]MemcachedInstance)[vs[1].(string)]
	}).(MemcachedInstanceOutput)
}

func init() {
	pulumi.RegisterOutputType(MemcachedInstanceOutput{})
	pulumi.RegisterOutputType(MemcachedInstancePtrOutput{})
	pulumi.RegisterOutputType(MemcachedInstanceArrayOutput{})
	pulumi.RegisterOutputType(MemcachedInstanceMapOutput{})
}