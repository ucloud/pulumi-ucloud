// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package uhost

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/r/instance.html.markdown.
type Instance struct {
	pulumi.CustomResourceState

	AllowStoppingForUpdate pulumi.BoolPtrOutput `pulumi:"allowStoppingForUpdate"`
	// Whether to renew an instance automatically or not.
	AutoRenew pulumi.BoolOutput `pulumi:"autoRenew"`
	// Availability zone where instance is located. such as: `cn-bj2-02`. You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The size of the boot disk, measured in GB (GigaByte). Range: 20-100. The value set of disk size must be larger or equal to `20`(default: `20`) for Linux and `40` (default: `40`) for Windows. The responsive time is a bit longer if the value set is larger than default for local boot disk, and further settings may be required on host instance if the value set is larger than default for cloud boot disk. The disk volume adjustment must be a multiple of 10 GB. In addition, any reduction of boot disk size is not supported.
	BootDiskSize pulumi.IntOutput `pulumi:"bootDiskSize"`
	// The type of boot disk. Possible values are: `localNormal` and `localSsd` for local boot disk, `cloudSsd` for cloud SSD boot disk. (Default: `localNormal`). The `localSsd` and `cloudSsd` are not fully support by all regions as boot disk type, please proceed to UCloud console for more details.
	BootDiskType pulumi.StringOutput `pulumi:"bootDiskType"`
	// The charge type of instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType pulumi.StringOutput `pulumi:"chargeType"`
	// The number of cores of virtual CPU, measured in core.
	Cpu pulumi.IntOutput `pulumi:"cpu"`
	// The time of creation for instance, formatted in RFC3339 time string.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The size of local data disk, measured in GB (GigaByte), range: 0-8000 (Default: `20`), 0-8000 for cloud disk, 0-2000 for local sata disk and 100-1000 for local ssd disk (all the GPU type instances are included). The volume adjustment must be a multiple of 10 GB. In addition, any reduction of data disk size is not supported.
	DataDiskSize pulumi.IntOutput `pulumi:"dataDiskSize"`
	// The type of local data disk. Possible values are: `localNormal` and `localSsd` for local data disk. (Default: `localNormal`). The `localSsd` is not fully support by all regions as data disk type, please proceed to UCloud console for more details. In addition, the `dataDiskType` must be same as `bootDiskType` if specified.
	DataDiskType pulumi.StringOutput `pulumi:"dataDiskType"`
	// It is a nested type which documented below.
	DiskSets InstanceDiskSetArrayOutput `pulumi:"diskSets"`
	// The duration that you will buy the instance (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration pulumi.IntPtrOutput `pulumi:"duration"`
	// The expiration time for instance, formatted in RFC3339 time string.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// The ID for the image to use for the instance.
	ImageId      pulumi.StringOutput `pulumi:"imageId"`
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// It is a nested type which documented below.
	IpSets InstanceIpSetArrayOutput `pulumi:"ipSets"`
	// The ID of the associated isolation group.
	IsolationGroup pulumi.StringOutput `pulumi:"isolationGroup"`
	// The size of memory, measured in GB(Gigabyte).
	Memory pulumi.IntOutput    `pulumi:"memory"`
	Name   pulumi.StringOutput `pulumi:"name"`
	// The private IP address assigned to the instance.
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// The remarks of instance. (Default: `""`).
	Remark       pulumi.StringOutput `pulumi:"remark"`
	RootPassword pulumi.StringOutput `pulumi:"rootPassword"`
	// The ID of the associated security group.
	SecurityGroup pulumi.StringOutput `pulumi:"securityGroup"`
	// Instance current status. Possible values are `Initializing`, `Starting`, `Running`, `Stopping`, `Stopped`, `Install Fail`, `ResizeFail` and `Rebooting`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of subnet. If defined `vpcId`, the `subnetId` is Required. If not defined `vpcId` and `subnetId`, the instance will use the default subnet in the current region.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// A tag assigned to instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringPtrOutput `pulumi:"tag"`
	// The ID of VPC linked to the instance. If not defined `vpcId`, the instance will use the default VPC in the current region.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil || args.AvailabilityZone == nil {
		return nil, errors.New("missing required argument 'AvailabilityZone'")
	}
	if args == nil || args.ImageId == nil {
		return nil, errors.New("missing required argument 'ImageId'")
	}
	if args == nil || args.InstanceType == nil {
		return nil, errors.New("missing required argument 'InstanceType'")
	}
	if args == nil {
		args = &InstanceArgs{}
	}
	var resource Instance
	err := ctx.RegisterResource("ucloud:uhost/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("ucloud:uhost/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	AllowStoppingForUpdate *bool `pulumi:"allowStoppingForUpdate"`
	// Whether to renew an instance automatically or not.
	AutoRenew *bool `pulumi:"autoRenew"`
	// Availability zone where instance is located. such as: `cn-bj2-02`. You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The size of the boot disk, measured in GB (GigaByte). Range: 20-100. The value set of disk size must be larger or equal to `20`(default: `20`) for Linux and `40` (default: `40`) for Windows. The responsive time is a bit longer if the value set is larger than default for local boot disk, and further settings may be required on host instance if the value set is larger than default for cloud boot disk. The disk volume adjustment must be a multiple of 10 GB. In addition, any reduction of boot disk size is not supported.
	BootDiskSize *int `pulumi:"bootDiskSize"`
	// The type of boot disk. Possible values are: `localNormal` and `localSsd` for local boot disk, `cloudSsd` for cloud SSD boot disk. (Default: `localNormal`). The `localSsd` and `cloudSsd` are not fully support by all regions as boot disk type, please proceed to UCloud console for more details.
	BootDiskType *string `pulumi:"bootDiskType"`
	// The charge type of instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType *string `pulumi:"chargeType"`
	// The number of cores of virtual CPU, measured in core.
	Cpu *int `pulumi:"cpu"`
	// The time of creation for instance, formatted in RFC3339 time string.
	CreateTime *string `pulumi:"createTime"`
	// The size of local data disk, measured in GB (GigaByte), range: 0-8000 (Default: `20`), 0-8000 for cloud disk, 0-2000 for local sata disk and 100-1000 for local ssd disk (all the GPU type instances are included). The volume adjustment must be a multiple of 10 GB. In addition, any reduction of data disk size is not supported.
	DataDiskSize *int `pulumi:"dataDiskSize"`
	// The type of local data disk. Possible values are: `localNormal` and `localSsd` for local data disk. (Default: `localNormal`). The `localSsd` is not fully support by all regions as data disk type, please proceed to UCloud console for more details. In addition, the `dataDiskType` must be same as `bootDiskType` if specified.
	DataDiskType *string `pulumi:"dataDiskType"`
	// It is a nested type which documented below.
	DiskSets []InstanceDiskSet `pulumi:"diskSets"`
	// The duration that you will buy the instance (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration *int `pulumi:"duration"`
	// The expiration time for instance, formatted in RFC3339 time string.
	ExpireTime *string `pulumi:"expireTime"`
	// The ID for the image to use for the instance.
	ImageId      *string `pulumi:"imageId"`
	InstanceType *string `pulumi:"instanceType"`
	// It is a nested type which documented below.
	IpSets []InstanceIpSet `pulumi:"ipSets"`
	// The ID of the associated isolation group.
	IsolationGroup *string `pulumi:"isolationGroup"`
	// The size of memory, measured in GB(Gigabyte).
	Memory *int    `pulumi:"memory"`
	Name   *string `pulumi:"name"`
	// The private IP address assigned to the instance.
	PrivateIp *string `pulumi:"privateIp"`
	// The remarks of instance. (Default: `""`).
	Remark       *string `pulumi:"remark"`
	RootPassword *string `pulumi:"rootPassword"`
	// The ID of the associated security group.
	SecurityGroup *string `pulumi:"securityGroup"`
	// Instance current status. Possible values are `Initializing`, `Starting`, `Running`, `Stopping`, `Stopped`, `Install Fail`, `ResizeFail` and `Rebooting`.
	Status *string `pulumi:"status"`
	// The ID of subnet. If defined `vpcId`, the `subnetId` is Required. If not defined `vpcId` and `subnetId`, the instance will use the default subnet in the current region.
	SubnetId *string `pulumi:"subnetId"`
	// A tag assigned to instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag *string `pulumi:"tag"`
	// The ID of VPC linked to the instance. If not defined `vpcId`, the instance will use the default VPC in the current region.
	VpcId *string `pulumi:"vpcId"`
}

type InstanceState struct {
	AllowStoppingForUpdate pulumi.BoolPtrInput
	// Whether to renew an instance automatically or not.
	AutoRenew pulumi.BoolPtrInput
	// Availability zone where instance is located. such as: `cn-bj2-02`. You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone pulumi.StringPtrInput
	// The size of the boot disk, measured in GB (GigaByte). Range: 20-100. The value set of disk size must be larger or equal to `20`(default: `20`) for Linux and `40` (default: `40`) for Windows. The responsive time is a bit longer if the value set is larger than default for local boot disk, and further settings may be required on host instance if the value set is larger than default for cloud boot disk. The disk volume adjustment must be a multiple of 10 GB. In addition, any reduction of boot disk size is not supported.
	BootDiskSize pulumi.IntPtrInput
	// The type of boot disk. Possible values are: `localNormal` and `localSsd` for local boot disk, `cloudSsd` for cloud SSD boot disk. (Default: `localNormal`). The `localSsd` and `cloudSsd` are not fully support by all regions as boot disk type, please proceed to UCloud console for more details.
	BootDiskType pulumi.StringPtrInput
	// The charge type of instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType pulumi.StringPtrInput
	// The number of cores of virtual CPU, measured in core.
	Cpu pulumi.IntPtrInput
	// The time of creation for instance, formatted in RFC3339 time string.
	CreateTime pulumi.StringPtrInput
	// The size of local data disk, measured in GB (GigaByte), range: 0-8000 (Default: `20`), 0-8000 for cloud disk, 0-2000 for local sata disk and 100-1000 for local ssd disk (all the GPU type instances are included). The volume adjustment must be a multiple of 10 GB. In addition, any reduction of data disk size is not supported.
	DataDiskSize pulumi.IntPtrInput
	// The type of local data disk. Possible values are: `localNormal` and `localSsd` for local data disk. (Default: `localNormal`). The `localSsd` is not fully support by all regions as data disk type, please proceed to UCloud console for more details. In addition, the `dataDiskType` must be same as `bootDiskType` if specified.
	DataDiskType pulumi.StringPtrInput
	// It is a nested type which documented below.
	DiskSets InstanceDiskSetArrayInput
	// The duration that you will buy the instance (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration pulumi.IntPtrInput
	// The expiration time for instance, formatted in RFC3339 time string.
	ExpireTime pulumi.StringPtrInput
	// The ID for the image to use for the instance.
	ImageId      pulumi.StringPtrInput
	InstanceType pulumi.StringPtrInput
	// It is a nested type which documented below.
	IpSets InstanceIpSetArrayInput
	// The ID of the associated isolation group.
	IsolationGroup pulumi.StringPtrInput
	// The size of memory, measured in GB(Gigabyte).
	Memory pulumi.IntPtrInput
	Name   pulumi.StringPtrInput
	// The private IP address assigned to the instance.
	PrivateIp pulumi.StringPtrInput
	// The remarks of instance. (Default: `""`).
	Remark       pulumi.StringPtrInput
	RootPassword pulumi.StringPtrInput
	// The ID of the associated security group.
	SecurityGroup pulumi.StringPtrInput
	// Instance current status. Possible values are `Initializing`, `Starting`, `Running`, `Stopping`, `Stopped`, `Install Fail`, `ResizeFail` and `Rebooting`.
	Status pulumi.StringPtrInput
	// The ID of subnet. If defined `vpcId`, the `subnetId` is Required. If not defined `vpcId` and `subnetId`, the instance will use the default subnet in the current region.
	SubnetId pulumi.StringPtrInput
	// A tag assigned to instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringPtrInput
	// The ID of VPC linked to the instance. If not defined `vpcId`, the instance will use the default VPC in the current region.
	VpcId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	AllowStoppingForUpdate *bool `pulumi:"allowStoppingForUpdate"`
	// Availability zone where instance is located. such as: `cn-bj2-02`. You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone string `pulumi:"availabilityZone"`
	// The size of the boot disk, measured in GB (GigaByte). Range: 20-100. The value set of disk size must be larger or equal to `20`(default: `20`) for Linux and `40` (default: `40`) for Windows. The responsive time is a bit longer if the value set is larger than default for local boot disk, and further settings may be required on host instance if the value set is larger than default for cloud boot disk. The disk volume adjustment must be a multiple of 10 GB. In addition, any reduction of boot disk size is not supported.
	BootDiskSize *int `pulumi:"bootDiskSize"`
	// The type of boot disk. Possible values are: `localNormal` and `localSsd` for local boot disk, `cloudSsd` for cloud SSD boot disk. (Default: `localNormal`). The `localSsd` and `cloudSsd` are not fully support by all regions as boot disk type, please proceed to UCloud console for more details.
	BootDiskType *string `pulumi:"bootDiskType"`
	// The charge type of instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType *string `pulumi:"chargeType"`
	// The size of local data disk, measured in GB (GigaByte), range: 0-8000 (Default: `20`), 0-8000 for cloud disk, 0-2000 for local sata disk and 100-1000 for local ssd disk (all the GPU type instances are included). The volume adjustment must be a multiple of 10 GB. In addition, any reduction of data disk size is not supported.
	DataDiskSize *int `pulumi:"dataDiskSize"`
	// The type of local data disk. Possible values are: `localNormal` and `localSsd` for local data disk. (Default: `localNormal`). The `localSsd` is not fully support by all regions as data disk type, please proceed to UCloud console for more details. In addition, the `dataDiskType` must be same as `bootDiskType` if specified.
	DataDiskType *string `pulumi:"dataDiskType"`
	// The duration that you will buy the instance (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration *int `pulumi:"duration"`
	// The ID for the image to use for the instance.
	ImageId      string `pulumi:"imageId"`
	InstanceType string `pulumi:"instanceType"`
	// The ID of the associated isolation group.
	IsolationGroup *string `pulumi:"isolationGroup"`
	Name           *string `pulumi:"name"`
	// The private IP address assigned to the instance.
	PrivateIp *string `pulumi:"privateIp"`
	// The remarks of instance. (Default: `""`).
	Remark       *string `pulumi:"remark"`
	RootPassword *string `pulumi:"rootPassword"`
	// The ID of the associated security group.
	SecurityGroup *string `pulumi:"securityGroup"`
	// The ID of subnet. If defined `vpcId`, the `subnetId` is Required. If not defined `vpcId` and `subnetId`, the instance will use the default subnet in the current region.
	SubnetId *string `pulumi:"subnetId"`
	// A tag assigned to instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag *string `pulumi:"tag"`
	// The ID of VPC linked to the instance. If not defined `vpcId`, the instance will use the default VPC in the current region.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	AllowStoppingForUpdate pulumi.BoolPtrInput
	// Availability zone where instance is located. such as: `cn-bj2-02`. You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone pulumi.StringInput
	// The size of the boot disk, measured in GB (GigaByte). Range: 20-100. The value set of disk size must be larger or equal to `20`(default: `20`) for Linux and `40` (default: `40`) for Windows. The responsive time is a bit longer if the value set is larger than default for local boot disk, and further settings may be required on host instance if the value set is larger than default for cloud boot disk. The disk volume adjustment must be a multiple of 10 GB. In addition, any reduction of boot disk size is not supported.
	BootDiskSize pulumi.IntPtrInput
	// The type of boot disk. Possible values are: `localNormal` and `localSsd` for local boot disk, `cloudSsd` for cloud SSD boot disk. (Default: `localNormal`). The `localSsd` and `cloudSsd` are not fully support by all regions as boot disk type, please proceed to UCloud console for more details.
	BootDiskType pulumi.StringPtrInput
	// The charge type of instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType pulumi.StringPtrInput
	// The size of local data disk, measured in GB (GigaByte), range: 0-8000 (Default: `20`), 0-8000 for cloud disk, 0-2000 for local sata disk and 100-1000 for local ssd disk (all the GPU type instances are included). The volume adjustment must be a multiple of 10 GB. In addition, any reduction of data disk size is not supported.
	DataDiskSize pulumi.IntPtrInput
	// The type of local data disk. Possible values are: `localNormal` and `localSsd` for local data disk. (Default: `localNormal`). The `localSsd` is not fully support by all regions as data disk type, please proceed to UCloud console for more details. In addition, the `dataDiskType` must be same as `bootDiskType` if specified.
	DataDiskType pulumi.StringPtrInput
	// The duration that you will buy the instance (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration pulumi.IntPtrInput
	// The ID for the image to use for the instance.
	ImageId      pulumi.StringInput
	InstanceType pulumi.StringInput
	// The ID of the associated isolation group.
	IsolationGroup pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	// The private IP address assigned to the instance.
	PrivateIp pulumi.StringPtrInput
	// The remarks of instance. (Default: `""`).
	Remark       pulumi.StringPtrInput
	RootPassword pulumi.StringPtrInput
	// The ID of the associated security group.
	SecurityGroup pulumi.StringPtrInput
	// The ID of subnet. If defined `vpcId`, the `subnetId` is Required. If not defined `vpcId` and `subnetId`, the instance will use the default subnet in the current region.
	SubnetId pulumi.StringPtrInput
	// A tag assigned to instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringPtrInput
	// The ID of VPC linked to the instance. If not defined `vpcId`, the instance will use the default VPC in the current region.
	VpcId pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}
