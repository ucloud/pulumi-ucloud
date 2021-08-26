// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ucloud.Umem.Outputs
{

    [OutputType]
    public sealed class RedisInstanceIpSet
    {
        /// <summary>
        /// The virtual ip of Redis instance.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// The port on which Redis instance accepts connections, it is 6379 by default.
        /// </summary>
        public readonly int? Port;

        [OutputConstructor]
        private RedisInstanceIpSet(
            string? ip,

            int? port)
        {
            Ip = ip;
            Port = port;
        }
    }
}