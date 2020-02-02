// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ipsecvpn

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source providers a list of VPN Connection resources according to their ID, name and tag.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/d/vpn_connections.html.markdown.
func LookupVpnConnections(ctx *pulumi.Context, args *LookupVpnConnectionsArgs, opts ...pulumi.InvokeOption) (*LookupVpnConnectionsResult, error) {
	var rv LookupVpnConnectionsResult
	err := ctx.Invoke("ucloud:ipsecvpn/lookupVpnConnections:lookupVpnConnections", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking lookupVpnConnections.
type LookupVpnConnectionsArgs struct {
	// A list of VPN Connection IDs, all the VPN Connections belongs to the defined region will be retrieved if this argument is "".
	Ids []string `pulumi:"ids"`
	// A regex string to filter resulting VPN Connections by name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// A tag assigned to VPN Connection.
	Tag *string `pulumi:"tag"`
}

// A collection of values returned by lookupVpnConnections.
type LookupVpnConnectionsResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	OutputFile *string  `pulumi:"outputFile"`
	// A tag assigned to the VPN Connection.
	Tag *string `pulumi:"tag"`
	// Total number of VPN Connections that satisfy the condition.
	TotalCount int `pulumi:"totalCount"`
	// It is a nested type. VPN Connections documented below.
	VpnConnections []LookupVpnConnectionsVpnConnection `pulumi:"vpnConnections"`
}
