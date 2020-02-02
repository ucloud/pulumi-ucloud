// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ipsecvpn

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a VPN Customer Gateway resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/r/vpn_customer_gateway.html.markdown.
type VpnCustomerGateway struct {
	pulumi.CustomResourceState

	// The creation time for VPN Customer Gateway, formatted in RFC3339 time string.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The ip address of the VPN Customer Gateway.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	Name      pulumi.StringOutput `pulumi:"name"`
	// The remarks of the VPN Customer Gateway. (Default: `""`).
	Remark pulumi.StringOutput `pulumi:"remark"`
	// A tag assigned to VPN Customer Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	// * ``
	Tag pulumi.StringPtrOutput `pulumi:"tag"`
}

// NewVpnCustomerGateway registers a new resource with the given unique name, arguments, and options.
func NewVpnCustomerGateway(ctx *pulumi.Context,
	name string, args *VpnCustomerGatewayArgs, opts ...pulumi.ResourceOption) (*VpnCustomerGateway, error) {
	if args == nil || args.IpAddress == nil {
		return nil, errors.New("missing required argument 'IpAddress'")
	}
	if args == nil {
		args = &VpnCustomerGatewayArgs{}
	}
	var resource VpnCustomerGateway
	err := ctx.RegisterResource("ucloud:ipsecvpn/vpnCustomerGateway:VpnCustomerGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnCustomerGateway gets an existing VpnCustomerGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnCustomerGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnCustomerGatewayState, opts ...pulumi.ResourceOption) (*VpnCustomerGateway, error) {
	var resource VpnCustomerGateway
	err := ctx.ReadResource("ucloud:ipsecvpn/vpnCustomerGateway:VpnCustomerGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnCustomerGateway resources.
type vpnCustomerGatewayState struct {
	// The creation time for VPN Customer Gateway, formatted in RFC3339 time string.
	CreateTime *string `pulumi:"createTime"`
	// The ip address of the VPN Customer Gateway.
	IpAddress *string `pulumi:"ipAddress"`
	Name      *string `pulumi:"name"`
	// The remarks of the VPN Customer Gateway. (Default: `""`).
	Remark *string `pulumi:"remark"`
	// A tag assigned to VPN Customer Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	// * ``
	Tag *string `pulumi:"tag"`
}

type VpnCustomerGatewayState struct {
	// The creation time for VPN Customer Gateway, formatted in RFC3339 time string.
	CreateTime pulumi.StringPtrInput
	// The ip address of the VPN Customer Gateway.
	IpAddress pulumi.StringPtrInput
	Name      pulumi.StringPtrInput
	// The remarks of the VPN Customer Gateway. (Default: `""`).
	Remark pulumi.StringPtrInput
	// A tag assigned to VPN Customer Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	// * ``
	Tag pulumi.StringPtrInput
}

func (VpnCustomerGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnCustomerGatewayState)(nil)).Elem()
}

type vpnCustomerGatewayArgs struct {
	// The ip address of the VPN Customer Gateway.
	IpAddress string  `pulumi:"ipAddress"`
	Name      *string `pulumi:"name"`
	// The remarks of the VPN Customer Gateway. (Default: `""`).
	Remark *string `pulumi:"remark"`
	// A tag assigned to VPN Customer Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	// * ``
	Tag *string `pulumi:"tag"`
}

// The set of arguments for constructing a VpnCustomerGateway resource.
type VpnCustomerGatewayArgs struct {
	// The ip address of the VPN Customer Gateway.
	IpAddress pulumi.StringInput
	Name      pulumi.StringPtrInput
	// The remarks of the VPN Customer Gateway. (Default: `""`).
	Remark pulumi.StringPtrInput
	// A tag assigned to VPN Customer Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	// * ``
	Tag pulumi.StringPtrInput
}

func (VpnCustomerGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnCustomerGatewayArgs)(nil)).Elem()
}
