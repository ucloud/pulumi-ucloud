// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ucloud.Uhost
{
    /// <summary>
    /// ## Import
    /// 
    /// Instance can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import ucloud:uhost/instance:Instance example uhost-abcdefg
    /// ```
    /// </summary>
    [UcloudResourceType("ucloud:uhost/instance:Instance")]
    public partial class Instance : Pulumi.CustomResource
    {
        [Output("allowStoppingForUpdate")]
        public Output<bool?> AllowStoppingForUpdate { get; private set; } = null!;

        /// <summary>
        /// Whether to renew an instance automatically or not.
        /// </summary>
        [Output("autoRenew")]
        public Output<bool> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// Availability zone where instance is located. such as: `cn-bj2-02`. You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// The size of the boot disk, measured in GB (GigaByte). Range: 20-500. The value set of disk size must be larger or equal to `20`(default: `20`) for Linux and `40` (default: `40`) for Windows. The responsive time is a bit longer if the value set is larger than default for local boot disk, and further settings may be required on host instance if the value set is larger than default for cloud boot disk. The disk volume adjustment must be a multiple of 10 GB. In addition, any reduction of boot disk size is not supported.
        /// </summary>
        [Output("bootDiskSize")]
        public Output<int> BootDiskSize { get; private set; } = null!;

        /// <summary>
        /// The type of boot disk. Possible values are: `local_normal` and `local_ssd` for local boot disk, `cloud_ssd` for cloud SSD boot disk. (Default: `local_normal`). The `local_ssd` and `cloud_ssd` are not fully support by all regions as boot disk type, please proceed to UCloud console for more details.
        /// </summary>
        [Output("bootDiskType")]
        public Output<string> BootDiskType { get; private set; } = null!;

        /// <summary>
        /// The charge type of instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
        /// </summary>
        [Output("chargeType")]
        public Output<string> ChargeType { get; private set; } = null!;

        /// <summary>
        /// The number of cores of virtual CPU, measured in core.
        /// </summary>
        [Output("cpu")]
        public Output<int> Cpu { get; private set; } = null!;

        [Output("cpuPlatform")]
        public Output<string> CpuPlatform { get; private set; } = null!;

        /// <summary>
        /// The time of creation for instance, formatted in RFC3339 time string.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The size of local data disk, measured in GB (GigaByte), 20-2000 for local sata disk and 20-1000 for local ssd disk (all the GPU type instances are included). The volume adjustment must be a multiple of 10 GB. In addition, any reduction of data disk size is not supported.
        /// </summary>
        [Output("dataDiskSize")]
        public Output<int> DataDiskSize { get; private set; } = null!;

        /// <summary>
        /// The type of local data disk. Possible values are: `local_normal` and `local_ssd` for local data disk. (Default: `local_normal`). The `local_ssd` is not fully support by all regions as data disk type, please proceed to UCloud console for more details. In addition, the `data_disk_type` must be same as `boot_disk_type` if specified.
        /// </summary>
        [Output("dataDiskType")]
        public Output<string> DataDiskType { get; private set; } = null!;

        /// <summary>
        /// Additional cloud data disks to attach to the instance. `data_disks` configurations only apply on resource creation. The count of `data_disks` can only be one. See data_disks below for details on attributes. When set `data_disks`, the argument `delete_disks_with_instance` must bet set.
        /// </summary>
        [Output("dataDisks")]
        public Output<Outputs.InstanceDataDisks?> DataDisks { get; private set; } = null!;

        /// <summary>
        /// Whether the cloud data disks attached instance should be destroyed on instance termination.
        /// </summary>
        [Output("deleteDisksWithInstance")]
        public Output<bool?> DeleteDisksWithInstance { get; private set; } = null!;

        /// <summary>
        /// It is a nested type which documented below.
        /// </summary>
        [Output("diskSets")]
        public Output<ImmutableArray<Outputs.InstanceDiskSet>> DiskSets { get; private set; } = null!;

        /// <summary>
        /// The duration that you will buy the instance (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
        /// </summary>
        [Output("duration")]
        public Output<int?> Duration { get; private set; } = null!;

        /// <summary>
        /// The expiration time for instance, formatted in RFC3339 time string.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// The ID for the image to use for the instance.
        /// </summary>
        [Output("imageId")]
        public Output<string> ImageId { get; private set; } = null!;

        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// It is a nested type which documented below.
        /// </summary>
        [Output("ipSets")]
        public Output<ImmutableArray<Outputs.InstanceIpSet>> IpSets { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated isolation group.
        /// </summary>
        [Output("isolationGroup")]
        public Output<string> IsolationGroup { get; private set; } = null!;

        /// <summary>
        /// The size of memory, measured in GB(Gigabyte).
        /// </summary>
        [Output("memory")]
        public Output<int> Memory { get; private set; } = null!;

        /// <summary>
        /// Specifies a minimum CPU platform for the the VM instance. (Default: `Intel/Auto`). You may refer to [min_cpu_platform](https://docs.ucloud.cn/uhost/introduction/uhost/type_new)
        /// - The Intel CPU platform:
        /// - `Intel/Auto` as the Intel CPU platform version will be selected randomly by system;
        /// - `Intel/IvyBridge` as Intel V2, the version of Intel CPU platform selected by system will be `Intel/IvyBridge` and above;
        /// - `Intel/Haswell` as Intel V3,  the version of Intel CPU platform selected by system will be `Intel/Haswell` and above;
        /// - `Intel/Broadwell` as Intel V4, the version of Intel CPU platform selected by system will be `Intel/Broadwell` and above;
        /// - `Intel/Skylake` as Intel V5, the version of Intel CPU platform selected by system will be `Intel/Skylake` and above;
        /// - `Intel/Cascadelake` as Intel V6, the version of Intel CPU platform selected by system will be `Intel/Cascadelake`;
        /// - The AMD CPU platform:
        /// - `Amd/Auto` as the Amd CPU platform version will be selected randomly by system;
        /// - `Amd/Epyc2` as the version of Amd CPU platform selected by system will be `Amd/Epyc2` and above;
        /// </summary>
        [Output("minCpuPlatform")]
        public Output<string?> MinCpuPlatform { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The private IP address assigned to the instance.
        /// </summary>
        [Output("privateIp")]
        public Output<string> PrivateIp { get; private set; } = null!;

        /// <summary>
        /// The remarks of instance. (Default: `""`).
        /// </summary>
        [Output("remark")]
        public Output<string> Remark { get; private set; } = null!;

        [Output("rootPassword")]
        public Output<string> RootPassword { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated security group.
        /// </summary>
        [Output("securityGroup")]
        public Output<string> SecurityGroup { get; private set; } = null!;

        /// <summary>
        /// Instance current status. Possible values are `Initializing`, `Starting`, `Running`, `Stopping`, `Stopped`, `Install Fail`, `ResizeFail` and `Rebooting`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The ID of subnet. If defined `vpc_id`, the `subnet_id` is Required. If not defined `vpc_id` and `subnet_id`, the instance will use the default subnet in the current region.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// A tag assigned to instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        /// </summary>
        [Output("tag")]
        public Output<string?> Tag { get; private set; } = null!;

        /// <summary>
        /// The user data to customize the startup behaviors when launching the instance. You may refer to [user_data_document](https://docs.ucloud.cn/uhost/guide/metadata/userdata)
        /// </summary>
        [Output("userData")]
        public Output<string?> UserData { get; private set; } = null!;

        /// <summary>
        /// The ID of VPC linked to the instance. If not defined `vpc_id`, the instance will use the default VPC in the current region.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("ucloud:uhost/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("ucloud:uhost/instance:Instance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        [Input("allowStoppingForUpdate")]
        public Input<bool>? AllowStoppingForUpdate { get; set; }

        /// <summary>
        /// Availability zone where instance is located. such as: `cn-bj2-02`. You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
        /// </summary>
        [Input("availabilityZone", required: true)]
        public Input<string> AvailabilityZone { get; set; } = null!;

        /// <summary>
        /// The size of the boot disk, measured in GB (GigaByte). Range: 20-500. The value set of disk size must be larger or equal to `20`(default: `20`) for Linux and `40` (default: `40`) for Windows. The responsive time is a bit longer if the value set is larger than default for local boot disk, and further settings may be required on host instance if the value set is larger than default for cloud boot disk. The disk volume adjustment must be a multiple of 10 GB. In addition, any reduction of boot disk size is not supported.
        /// </summary>
        [Input("bootDiskSize")]
        public Input<int>? BootDiskSize { get; set; }

        /// <summary>
        /// The type of boot disk. Possible values are: `local_normal` and `local_ssd` for local boot disk, `cloud_ssd` for cloud SSD boot disk. (Default: `local_normal`). The `local_ssd` and `cloud_ssd` are not fully support by all regions as boot disk type, please proceed to UCloud console for more details.
        /// </summary>
        [Input("bootDiskType")]
        public Input<string>? BootDiskType { get; set; }

        /// <summary>
        /// The charge type of instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// The size of local data disk, measured in GB (GigaByte), 20-2000 for local sata disk and 20-1000 for local ssd disk (all the GPU type instances are included). The volume adjustment must be a multiple of 10 GB. In addition, any reduction of data disk size is not supported.
        /// </summary>
        [Input("dataDiskSize")]
        public Input<int>? DataDiskSize { get; set; }

        /// <summary>
        /// The type of local data disk. Possible values are: `local_normal` and `local_ssd` for local data disk. (Default: `local_normal`). The `local_ssd` is not fully support by all regions as data disk type, please proceed to UCloud console for more details. In addition, the `data_disk_type` must be same as `boot_disk_type` if specified.
        /// </summary>
        [Input("dataDiskType")]
        public Input<string>? DataDiskType { get; set; }

        /// <summary>
        /// Additional cloud data disks to attach to the instance. `data_disks` configurations only apply on resource creation. The count of `data_disks` can only be one. See data_disks below for details on attributes. When set `data_disks`, the argument `delete_disks_with_instance` must bet set.
        /// </summary>
        [Input("dataDisks")]
        public Input<Inputs.InstanceDataDisksArgs>? DataDisks { get; set; }

        /// <summary>
        /// Whether the cloud data disks attached instance should be destroyed on instance termination.
        /// </summary>
        [Input("deleteDisksWithInstance")]
        public Input<bool>? DeleteDisksWithInstance { get; set; }

        /// <summary>
        /// The duration that you will buy the instance (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// The ID for the image to use for the instance.
        /// </summary>
        [Input("imageId", required: true)]
        public Input<string> ImageId { get; set; } = null!;

        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// The ID of the associated isolation group.
        /// </summary>
        [Input("isolationGroup")]
        public Input<string>? IsolationGroup { get; set; }

        /// <summary>
        /// Specifies a minimum CPU platform for the the VM instance. (Default: `Intel/Auto`). You may refer to [min_cpu_platform](https://docs.ucloud.cn/uhost/introduction/uhost/type_new)
        /// - The Intel CPU platform:
        /// - `Intel/Auto` as the Intel CPU platform version will be selected randomly by system;
        /// - `Intel/IvyBridge` as Intel V2, the version of Intel CPU platform selected by system will be `Intel/IvyBridge` and above;
        /// - `Intel/Haswell` as Intel V3,  the version of Intel CPU platform selected by system will be `Intel/Haswell` and above;
        /// - `Intel/Broadwell` as Intel V4, the version of Intel CPU platform selected by system will be `Intel/Broadwell` and above;
        /// - `Intel/Skylake` as Intel V5, the version of Intel CPU platform selected by system will be `Intel/Skylake` and above;
        /// - `Intel/Cascadelake` as Intel V6, the version of Intel CPU platform selected by system will be `Intel/Cascadelake`;
        /// - The AMD CPU platform:
        /// - `Amd/Auto` as the Amd CPU platform version will be selected randomly by system;
        /// - `Amd/Epyc2` as the version of Amd CPU platform selected by system will be `Amd/Epyc2` and above;
        /// </summary>
        [Input("minCpuPlatform")]
        public Input<string>? MinCpuPlatform { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The private IP address assigned to the instance.
        /// </summary>
        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        /// <summary>
        /// The remarks of instance. (Default: `""`).
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        [Input("rootPassword")]
        public Input<string>? RootPassword { get; set; }

        /// <summary>
        /// The ID of the associated security group.
        /// </summary>
        [Input("securityGroup")]
        public Input<string>? SecurityGroup { get; set; }

        /// <summary>
        /// The ID of subnet. If defined `vpc_id`, the `subnet_id` is Required. If not defined `vpc_id` and `subnet_id`, the instance will use the default subnet in the current region.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// A tag assigned to instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// The user data to customize the startup behaviors when launching the instance. You may refer to [user_data_document](https://docs.ucloud.cn/uhost/guide/metadata/userdata)
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        /// <summary>
        /// The ID of VPC linked to the instance. If not defined `vpc_id`, the instance will use the default VPC in the current region.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        [Input("allowStoppingForUpdate")]
        public Input<bool>? AllowStoppingForUpdate { get; set; }

        /// <summary>
        /// Whether to renew an instance automatically or not.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// Availability zone where instance is located. such as: `cn-bj2-02`. You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The size of the boot disk, measured in GB (GigaByte). Range: 20-500. The value set of disk size must be larger or equal to `20`(default: `20`) for Linux and `40` (default: `40`) for Windows. The responsive time is a bit longer if the value set is larger than default for local boot disk, and further settings may be required on host instance if the value set is larger than default for cloud boot disk. The disk volume adjustment must be a multiple of 10 GB. In addition, any reduction of boot disk size is not supported.
        /// </summary>
        [Input("bootDiskSize")]
        public Input<int>? BootDiskSize { get; set; }

        /// <summary>
        /// The type of boot disk. Possible values are: `local_normal` and `local_ssd` for local boot disk, `cloud_ssd` for cloud SSD boot disk. (Default: `local_normal`). The `local_ssd` and `cloud_ssd` are not fully support by all regions as boot disk type, please proceed to UCloud console for more details.
        /// </summary>
        [Input("bootDiskType")]
        public Input<string>? BootDiskType { get; set; }

        /// <summary>
        /// The charge type of instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// The number of cores of virtual CPU, measured in core.
        /// </summary>
        [Input("cpu")]
        public Input<int>? Cpu { get; set; }

        [Input("cpuPlatform")]
        public Input<string>? CpuPlatform { get; set; }

        /// <summary>
        /// The time of creation for instance, formatted in RFC3339 time string.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The size of local data disk, measured in GB (GigaByte), 20-2000 for local sata disk and 20-1000 for local ssd disk (all the GPU type instances are included). The volume adjustment must be a multiple of 10 GB. In addition, any reduction of data disk size is not supported.
        /// </summary>
        [Input("dataDiskSize")]
        public Input<int>? DataDiskSize { get; set; }

        /// <summary>
        /// The type of local data disk. Possible values are: `local_normal` and `local_ssd` for local data disk. (Default: `local_normal`). The `local_ssd` is not fully support by all regions as data disk type, please proceed to UCloud console for more details. In addition, the `data_disk_type` must be same as `boot_disk_type` if specified.
        /// </summary>
        [Input("dataDiskType")]
        public Input<string>? DataDiskType { get; set; }

        /// <summary>
        /// Additional cloud data disks to attach to the instance. `data_disks` configurations only apply on resource creation. The count of `data_disks` can only be one. See data_disks below for details on attributes. When set `data_disks`, the argument `delete_disks_with_instance` must bet set.
        /// </summary>
        [Input("dataDisks")]
        public Input<Inputs.InstanceDataDisksGetArgs>? DataDisks { get; set; }

        /// <summary>
        /// Whether the cloud data disks attached instance should be destroyed on instance termination.
        /// </summary>
        [Input("deleteDisksWithInstance")]
        public Input<bool>? DeleteDisksWithInstance { get; set; }

        [Input("diskSets")]
        private InputList<Inputs.InstanceDiskSetGetArgs>? _diskSets;

        /// <summary>
        /// It is a nested type which documented below.
        /// </summary>
        public InputList<Inputs.InstanceDiskSetGetArgs> DiskSets
        {
            get => _diskSets ?? (_diskSets = new InputList<Inputs.InstanceDiskSetGetArgs>());
            set => _diskSets = value;
        }

        /// <summary>
        /// The duration that you will buy the instance (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// The expiration time for instance, formatted in RFC3339 time string.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        /// <summary>
        /// The ID for the image to use for the instance.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("ipSets")]
        private InputList<Inputs.InstanceIpSetGetArgs>? _ipSets;

        /// <summary>
        /// It is a nested type which documented below.
        /// </summary>
        public InputList<Inputs.InstanceIpSetGetArgs> IpSets
        {
            get => _ipSets ?? (_ipSets = new InputList<Inputs.InstanceIpSetGetArgs>());
            set => _ipSets = value;
        }

        /// <summary>
        /// The ID of the associated isolation group.
        /// </summary>
        [Input("isolationGroup")]
        public Input<string>? IsolationGroup { get; set; }

        /// <summary>
        /// The size of memory, measured in GB(Gigabyte).
        /// </summary>
        [Input("memory")]
        public Input<int>? Memory { get; set; }

        /// <summary>
        /// Specifies a minimum CPU platform for the the VM instance. (Default: `Intel/Auto`). You may refer to [min_cpu_platform](https://docs.ucloud.cn/uhost/introduction/uhost/type_new)
        /// - The Intel CPU platform:
        /// - `Intel/Auto` as the Intel CPU platform version will be selected randomly by system;
        /// - `Intel/IvyBridge` as Intel V2, the version of Intel CPU platform selected by system will be `Intel/IvyBridge` and above;
        /// - `Intel/Haswell` as Intel V3,  the version of Intel CPU platform selected by system will be `Intel/Haswell` and above;
        /// - `Intel/Broadwell` as Intel V4, the version of Intel CPU platform selected by system will be `Intel/Broadwell` and above;
        /// - `Intel/Skylake` as Intel V5, the version of Intel CPU platform selected by system will be `Intel/Skylake` and above;
        /// - `Intel/Cascadelake` as Intel V6, the version of Intel CPU platform selected by system will be `Intel/Cascadelake`;
        /// - The AMD CPU platform:
        /// - `Amd/Auto` as the Amd CPU platform version will be selected randomly by system;
        /// - `Amd/Epyc2` as the version of Amd CPU platform selected by system will be `Amd/Epyc2` and above;
        /// </summary>
        [Input("minCpuPlatform")]
        public Input<string>? MinCpuPlatform { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The private IP address assigned to the instance.
        /// </summary>
        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        /// <summary>
        /// The remarks of instance. (Default: `""`).
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        [Input("rootPassword")]
        public Input<string>? RootPassword { get; set; }

        /// <summary>
        /// The ID of the associated security group.
        /// </summary>
        [Input("securityGroup")]
        public Input<string>? SecurityGroup { get; set; }

        /// <summary>
        /// Instance current status. Possible values are `Initializing`, `Starting`, `Running`, `Stopping`, `Stopped`, `Install Fail`, `ResizeFail` and `Rebooting`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of subnet. If defined `vpc_id`, the `subnet_id` is Required. If not defined `vpc_id` and `subnet_id`, the instance will use the default subnet in the current region.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// A tag assigned to instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// The user data to customize the startup behaviors when launching the instance. You may refer to [user_data_document](https://docs.ucloud.cn/uhost/guide/metadata/userdata)
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        /// <summary>
        /// The ID of VPC linked to the instance. If not defined `vpc_id`, the instance will use the default VPC in the current region.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public InstanceState()
        {
        }
    }
}
