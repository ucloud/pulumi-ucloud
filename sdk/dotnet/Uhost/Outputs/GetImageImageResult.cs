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
    public sealed class GetImageImageResult
    {
        /// <summary>
        /// Availability zone where images are located. such as: `cn-bj2-02`. You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist).
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// The time of creation for image, formatted in RFC3339 time string.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of image if any.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// To identify if any particular feature belongs to the instance, the value is `NetEnhnced` as I/O enhanced instance for now.
        /// </summary>
        public readonly ImmutableArray<string> Features;
        /// <summary>
        /// The ID of image.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of image.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of OS.
        /// </summary>
        public readonly string OsName;
        /// <summary>
        /// The type of OS. Possible values are: `linux` and `windows`, all the OS types will be retrieved by default.
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// The size of image.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// The status of image. Possible values are `Available`, `Making` and `Unavailable`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The type of image.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetImageImageResult(
            string availabilityZone,

            string createTime,

            string description,

            ImmutableArray<string> features,

            string id,

            string name,

            string osName,

            string osType,

            int size,

            string status,

            string type)
        {
            AvailabilityZone = availabilityZone;
            CreateTime = createTime;
            Description = description;
            Features = features;
            Id = id;
            Name = name;
            OsName = osName;
            OsType = osType;
            Size = size;
            Status = status;
            Type = type;
        }
    }
}