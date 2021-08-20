// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ucloud.Uhost.Outputs
{

    [OutputType]
    public sealed class GetInstanceInstanceDiskSetResult
    {
        /// <summary>
        /// The ID of disk.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specifies whether boot disk or not.
        /// </summary>
        public readonly bool IsBoot;
        /// <summary>
        /// The size of disk, measured in GB (Gigabyte).
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// The type of disk.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetInstanceInstanceDiskSetResult(
            string id,

            bool isBoot,

            int size,

            string type)
        {
            Id = id;
            IsBoot = isBoot;
            Size = size;
            Type = type;
        }
    }
}
