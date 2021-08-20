// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ucloud.Umem.Inputs
{

    public sealed class RedisInstanceIpSetGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The virtual ip of Redis instance.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The port on which Redis instance accepts connections, it is 6379 by default.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        public RedisInstanceIpSetGetArgs()
        {
        }
    }
}
