// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ucloud.Uaccount.Outputs
{

    [OutputType]
    public sealed class GetZoneZoneResult
    {
        /// <summary>
        /// The ID of availability zone.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetZoneZoneResult(string id)
        {
            Id = id;
        }
    }
}
