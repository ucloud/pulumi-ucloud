// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ipsecvpn

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source providers a list of VPN Gateway resources according to their ID, name, vpc and tag.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/d/vpn_gateways.html.markdown.
func LookupVpnGateways(ctx *pulumi.Context, args *LookupVpnGatewaysArgs, opts ...pulumi.InvokeOption) (*LookupVpnGatewaysResult, error) {
	var rv LookupVpnGatewaysResult
	err := ctx.Invoke("ucloud:ipsecvpn/lookupVpnGateways:lookupVpnGateways", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking lookupVpnGateways.
type LookupVpnGatewaysArgs struct {
	// A list of VPN Gateway IDs, all the VPN Gateways belongs to the defined region will be retrieved if this argument is "".
	Ids []string `pulumi:"ids"`
	// A regex string to filter resulting VPN Gateways by name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// A tag assigned to VPN Gateway.
	Tag *string `pulumi:"tag"`
	// The ID of VPC linked to the VPN Gateway.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by lookupVpnGateways.
type LookupVpnGatewaysResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	OutputFile *string  `pulumi:"outputFile"`
	// A tag assigned to the VPN Gateway.
	Tag *string `pulumi:"tag"`
	// Total number of VPN Gateways that satisfy the condition.
	TotalCount int `pulumi:"totalCount"`
	// The ID of VPC linked to the VPN Gateway.
	VpcId *string `pulumi:"vpcId"`
	// It is a nested type. VPN Gateways documented below.
	VpnGateways []LookupVpnGatewaysVpnGateway `pulumi:"vpnGateways"`
}
