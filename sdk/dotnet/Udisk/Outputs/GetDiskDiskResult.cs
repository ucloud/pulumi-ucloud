// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ucloud.Udisk.Outputs
{

    [OutputType]
    public sealed class GetDiskDiskResult
    {
        /// <summary>
        /// Availability zone where Disk are located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// The charge type of disk. Possible values are: `year` as pay by year, `month` as pay by month, `dynamic` as pay by hour.
        /// </summary>
        public readonly string ChargeType;
        /// <summary>
        /// The creation time of Disk, formatted in RFC3339 time string.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The size of disk. Purchase the size of disk in GB.
        /// </summary>
        public readonly int DiskSize;
        /// <summary>
        /// The type of disk. Possible values are: `data_disk`as cloud disk, `ssd_data_disk` as SSD cloud disk, `system_disk`as system disk, `ssd_system_disk` as SSD system disk, `rssd_data_disk` as RDMA-SSD cloud disk.
        /// </summary>
        public readonly string DiskType;
        /// <summary>
        /// The expiration time of disk, formatted in RFC3339 time string.
        /// </summary>
        public readonly string ExpireTime;
        /// <summary>
        /// The ID of Disk.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of Disk.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The status of disk. Possible values are: `Available`, `InUse`, `Detaching`, `Initializating`, `Failed`, `Cloning`, `Restoring`, `RestoreFailed`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A tag assigned to Disk.
        /// </summary>
        public readonly string Tag;

        [OutputConstructor]
        private GetDiskDiskResult(
            string availabilityZone,

            string chargeType,

            string createTime,

            int diskSize,

            string diskType,

            string expireTime,

            string id,

            string name,

            string status,

            string tag)
        {
            AvailabilityZone = availabilityZone;
            ChargeType = chargeType;
            CreateTime = createTime;
            DiskSize = diskSize;
            DiskType = diskType;
            ExpireTime = expireTime;
            Id = id;
            Name = name;
            Status = status;
            Tag = tag;
        }
    }
}
