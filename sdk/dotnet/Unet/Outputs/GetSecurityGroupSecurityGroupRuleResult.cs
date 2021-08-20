// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ucloud.Unet.Outputs
{

    [OutputType]
    public sealed class GetSecurityGroupSecurityGroupRuleResult
    {
        /// <summary>
        /// The cidr block of source.
        /// </summary>
        public readonly string CidrBlock;
        /// <summary>
        /// Authorization policy. Can be either `accept` or `drop`.
        /// </summary>
        public readonly string Policy;
        /// <summary>
        /// The range of port numbers, range: 1-65535. (eg: `port` or `port1-port2`).
        /// </summary>
        public readonly string PortRange;
        /// <summary>
        /// Rule priority. Can be `high`, `medium`, `low`.
        /// </summary>
        public readonly string Priority;
        /// <summary>
        /// The protocol. Can be `tcp`, `udp`, `icmp`, `gre`.
        /// </summary>
        public readonly string Protocol;

        [OutputConstructor]
        private GetSecurityGroupSecurityGroupRuleResult(
            string cidrBlock,

            string policy,

            string portRange,

            string priority,

            string protocol)
        {
            CidrBlock = cidrBlock;
            Policy = policy;
            PortRange = portRange;
            Priority = priority;
            Protocol = protocol;
        }
    }
}
