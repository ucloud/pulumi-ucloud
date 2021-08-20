// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ucloud.Ipsecvpn.Outputs
{

    [OutputType]
    public sealed class GetVPNConnectionVpnConnectionResult
    {
        /// <summary>
        /// The time of creation for VPN Connection, formatted in RFC3339 time string.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The ID of VPN Customer Gateway.
        /// </summary>
        public readonly string CustomerGatewayId;
        /// <summary>
        /// The ID of VPN Connection.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// It is a nested type which documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVPNConnectionVpnConnectionIkeConfigResult> IkeConfigs;
        /// <summary>
        /// It is a nested type which documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVPNConnectionVpnConnectionIpsecConfigResult> IpsecConfigs;
        /// <summary>
        /// The name of the VPN Connection.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The remarks of VPN Connection.
        /// </summary>
        public readonly string Remark;
        /// <summary>
        /// A tag assigned to VPN Connection.
        /// </summary>
        public readonly string Tag;
        /// <summary>
        /// The ID of VPC linked to the VPN Connection.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// The ID of VPN Gateway.
        /// </summary>
        public readonly string VpnGatewayId;

        [OutputConstructor]
        private GetVPNConnectionVpnConnectionResult(
            string createTime,

            string customerGatewayId,

            string id,

            ImmutableArray<Outputs.GetVPNConnectionVpnConnectionIkeConfigResult> ikeConfigs,

            ImmutableArray<Outputs.GetVPNConnectionVpnConnectionIpsecConfigResult> ipsecConfigs,

            string name,

            string remark,

            string tag,

            string vpcId,

            string vpnGatewayId)
        {
            CreateTime = createTime;
            CustomerGatewayId = customerGatewayId;
            Id = id;
            IkeConfigs = ikeConfigs;
            IpsecConfigs = ipsecConfigs;
            Name = name;
            Remark = remark;
            Tag = tag;
            VpcId = vpcId;
            VpnGatewayId = vpnGatewayId;
        }
    }
}
