// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ucloud.Vpc.Outputs
{

    [OutputType]
    public sealed class GetNATGatewayNatGatewayResult
    {
        /// <summary>
        /// The time of creation for Nat Gateway, formatted in RFC3339 time string.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The ID of Nat Gateway.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// It is a nested type which documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNATGatewayNatGatewayIpSetResult> IpSets;
        /// <summary>
        /// The name of the Nat Gateway.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The remarks of Nat Gateway.
        /// </summary>
        public readonly string Remark;
        public readonly string SecurityGroup;
        /// <summary>
        /// The list of subnet ID under the VPC.
        /// * `security_group` -The ID of the associated security group.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// A tag assigned to the Nat Gateway.
        /// </summary>
        public readonly string Tag;
        /// <summary>
        /// The ID of VPC linked to the Nat Gateway.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private GetNATGatewayNatGatewayResult(
            string createTime,

            string id,

            ImmutableArray<Outputs.GetNATGatewayNatGatewayIpSetResult> ipSets,

            string name,

            string remark,

            string securityGroup,

            ImmutableArray<string> subnetIds,

            string tag,

            string vpcId)
        {
            CreateTime = createTime;
            Id = id;
            IpSets = ipSets;
            Name = name;
            Remark = remark;
            SecurityGroup = securityGroup;
            SubnetIds = subnetIds;
            Tag = tag;
            VpcId = vpcId;
        }
    }
}
