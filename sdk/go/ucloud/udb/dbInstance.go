// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package udb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Database instance resource.
//
// > **Note**  Please do confirm if any task pending submission before reset your password, since the password reset will take effect immediately.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/r/db_instance.html.markdown.
type DbInstance struct {
	pulumi.CustomResourceState

	// Availability zone where database instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Specifies when the backup starts, measured in hour, it starts at one o'clock of 1, 2, 3, 4 in the morning by default.
	BackupBeginTime pulumi.IntOutput `pulumi:"backupBeginTime"`
	// The backup for database such as "test.%" or table such as "city.address" specified in the black lists are not supported.
	BackupBlackLists pulumi.StringArrayOutput `pulumi:"backupBlackLists"`
	// Specifies the number of backup saved per week, it is 7 backups saved per week by default.
	BackupCount pulumi.IntPtrOutput `pulumi:"backupCount"`
	// Specifies whether the backup took place from Sunday to Saturday by displaying 7 digits. 0 stands for backup disabled and 1 stands for backup enabled. The rightmost digit specifies whether the backup took place on Sunday, and the digits from right to left specify whether the backup took place from Monday to Saturday, it's mandatory required to backup twice per week at least. such as: digits "1100000" stands for the backup took place on Saturday and Friday.
	BackupDate pulumi.StringOutput `pulumi:"backupDate"`
	// The charge type of db instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType pulumi.StringOutput `pulumi:"chargeType"`
	// The creation time of database, formatted by RFC3339 time string.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The duration that you will buy the db instance (Default: `1`). The value is `0` when pay by month and the instance will be vaild till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration pulumi.IntPtrOutput `pulumi:"duration"`
	// The type of database engine, possible values are: "mysql", "percona".
	Engine pulumi.StringOutput `pulumi:"engine"`
	// The database engine version, possible values are: "5.5", "5.6", "5.7".
	// - 5.5/5.6/5.7 for mysql and percona engine.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// The expiration time of database, formatted by RFC3339 time string.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// Specifies the allocated storage size in gigabytes (GB), range from 20 to 4500GB. The volume adjustment must be a multiple of 10 GB. The maximum disk volume for SSD type are：
	// - 500GB if the memory chosen is equal or less than 6GB;
	// - 1000GB if the memory chosen is from 8 to 16GB;
	// - 2000GB if the memory chosen is 24GB or 32GB;
	// - 3500GB if the memory chosen is 48GB or 64GB;
	// - 4500GB if the memory chosen is equal or more than 96GB;
	InstanceStorage pulumi.IntOutput    `pulumi:"instanceStorage"`
	InstanceType    pulumi.StringOutput `pulumi:"instanceType"`
	// The modification time of database, formatted by RFC3339 time string.
	ModifyTime pulumi.StringOutput `pulumi:"modifyTime"`
	Name       pulumi.StringOutput `pulumi:"name"`
	Password   pulumi.StringOutput `pulumi:"password"`
	// The port on which the database accepts connections, the default port is 3306 for mysql and percona.
	Port pulumi.IntOutput `pulumi:"port"`
	// The private IP address assigned to the database instance.
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// Availability zone where the standby database instance is located for the high availability database instance with multiple zone; The disaster recovery of data center can be activated by switching to the standby database instance for the high availability database instance.
	StandbyZone pulumi.StringPtrOutput `pulumi:"standbyZone"`
	// Specifies the status of database, possible values are: `Init`, `Fail`, `Starting`, `Running`, `Shutdown`, `Shutoff`, `Delete`, `Upgrading`, `Promoting`, `Recovering` and `Recover fail`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of subnet.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// A tag assigned to database instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringOutput `pulumi:"tag"`
	// The ID of VPC linked to the database instances.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewDbInstance registers a new resource with the given unique name, arguments, and options.
func NewDbInstance(ctx *pulumi.Context,
	name string, args *DbInstanceArgs, opts ...pulumi.ResourceOption) (*DbInstance, error) {
	if args == nil || args.AvailabilityZone == nil {
		return nil, errors.New("missing required argument 'AvailabilityZone'")
	}
	if args == nil || args.Engine == nil {
		return nil, errors.New("missing required argument 'Engine'")
	}
	if args == nil || args.EngineVersion == nil {
		return nil, errors.New("missing required argument 'EngineVersion'")
	}
	if args == nil || args.InstanceStorage == nil {
		return nil, errors.New("missing required argument 'InstanceStorage'")
	}
	if args == nil || args.InstanceType == nil {
		return nil, errors.New("missing required argument 'InstanceType'")
	}
	if args == nil {
		args = &DbInstanceArgs{}
	}
	var resource DbInstance
	err := ctx.RegisterResource("ucloud:udb/dbInstance:DbInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDbInstance gets an existing DbInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDbInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DbInstanceState, opts ...pulumi.ResourceOption) (*DbInstance, error) {
	var resource DbInstance
	err := ctx.ReadResource("ucloud:udb/dbInstance:DbInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DbInstance resources.
type dbInstanceState struct {
	// Availability zone where database instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Specifies when the backup starts, measured in hour, it starts at one o'clock of 1, 2, 3, 4 in the morning by default.
	BackupBeginTime *int `pulumi:"backupBeginTime"`
	// The backup for database such as "test.%" or table such as "city.address" specified in the black lists are not supported.
	BackupBlackLists []string `pulumi:"backupBlackLists"`
	// Specifies the number of backup saved per week, it is 7 backups saved per week by default.
	BackupCount *int `pulumi:"backupCount"`
	// Specifies whether the backup took place from Sunday to Saturday by displaying 7 digits. 0 stands for backup disabled and 1 stands for backup enabled. The rightmost digit specifies whether the backup took place on Sunday, and the digits from right to left specify whether the backup took place from Monday to Saturday, it's mandatory required to backup twice per week at least. such as: digits "1100000" stands for the backup took place on Saturday and Friday.
	BackupDate *string `pulumi:"backupDate"`
	// The charge type of db instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType *string `pulumi:"chargeType"`
	// The creation time of database, formatted by RFC3339 time string.
	CreateTime *string `pulumi:"createTime"`
	// The duration that you will buy the db instance (Default: `1`). The value is `0` when pay by month and the instance will be vaild till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration *int `pulumi:"duration"`
	// The type of database engine, possible values are: "mysql", "percona".
	Engine *string `pulumi:"engine"`
	// The database engine version, possible values are: "5.5", "5.6", "5.7".
	// - 5.5/5.6/5.7 for mysql and percona engine.
	EngineVersion *string `pulumi:"engineVersion"`
	// The expiration time of database, formatted by RFC3339 time string.
	ExpireTime *string `pulumi:"expireTime"`
	// Specifies the allocated storage size in gigabytes (GB), range from 20 to 4500GB. The volume adjustment must be a multiple of 10 GB. The maximum disk volume for SSD type are：
	// - 500GB if the memory chosen is equal or less than 6GB;
	// - 1000GB if the memory chosen is from 8 to 16GB;
	// - 2000GB if the memory chosen is 24GB or 32GB;
	// - 3500GB if the memory chosen is 48GB or 64GB;
	// - 4500GB if the memory chosen is equal or more than 96GB;
	InstanceStorage *int    `pulumi:"instanceStorage"`
	InstanceType    *string `pulumi:"instanceType"`
	// The modification time of database, formatted by RFC3339 time string.
	ModifyTime *string `pulumi:"modifyTime"`
	Name       *string `pulumi:"name"`
	Password   *string `pulumi:"password"`
	// The port on which the database accepts connections, the default port is 3306 for mysql and percona.
	Port *int `pulumi:"port"`
	// The private IP address assigned to the database instance.
	PrivateIp *string `pulumi:"privateIp"`
	// Availability zone where the standby database instance is located for the high availability database instance with multiple zone; The disaster recovery of data center can be activated by switching to the standby database instance for the high availability database instance.
	StandbyZone *string `pulumi:"standbyZone"`
	// Specifies the status of database, possible values are: `Init`, `Fail`, `Starting`, `Running`, `Shutdown`, `Shutoff`, `Delete`, `Upgrading`, `Promoting`, `Recovering` and `Recover fail`.
	Status *string `pulumi:"status"`
	// The ID of subnet.
	SubnetId *string `pulumi:"subnetId"`
	// A tag assigned to database instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag *string `pulumi:"tag"`
	// The ID of VPC linked to the database instances.
	VpcId *string `pulumi:"vpcId"`
}

type DbInstanceState struct {
	// Availability zone where database instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone pulumi.StringPtrInput
	// Specifies when the backup starts, measured in hour, it starts at one o'clock of 1, 2, 3, 4 in the morning by default.
	BackupBeginTime pulumi.IntPtrInput
	// The backup for database such as "test.%" or table such as "city.address" specified in the black lists are not supported.
	BackupBlackLists pulumi.StringArrayInput
	// Specifies the number of backup saved per week, it is 7 backups saved per week by default.
	BackupCount pulumi.IntPtrInput
	// Specifies whether the backup took place from Sunday to Saturday by displaying 7 digits. 0 stands for backup disabled and 1 stands for backup enabled. The rightmost digit specifies whether the backup took place on Sunday, and the digits from right to left specify whether the backup took place from Monday to Saturday, it's mandatory required to backup twice per week at least. such as: digits "1100000" stands for the backup took place on Saturday and Friday.
	BackupDate pulumi.StringPtrInput
	// The charge type of db instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType pulumi.StringPtrInput
	// The creation time of database, formatted by RFC3339 time string.
	CreateTime pulumi.StringPtrInput
	// The duration that you will buy the db instance (Default: `1`). The value is `0` when pay by month and the instance will be vaild till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration pulumi.IntPtrInput
	// The type of database engine, possible values are: "mysql", "percona".
	Engine pulumi.StringPtrInput
	// The database engine version, possible values are: "5.5", "5.6", "5.7".
	// - 5.5/5.6/5.7 for mysql and percona engine.
	EngineVersion pulumi.StringPtrInput
	// The expiration time of database, formatted by RFC3339 time string.
	ExpireTime pulumi.StringPtrInput
	// Specifies the allocated storage size in gigabytes (GB), range from 20 to 4500GB. The volume adjustment must be a multiple of 10 GB. The maximum disk volume for SSD type are：
	// - 500GB if the memory chosen is equal or less than 6GB;
	// - 1000GB if the memory chosen is from 8 to 16GB;
	// - 2000GB if the memory chosen is 24GB or 32GB;
	// - 3500GB if the memory chosen is 48GB or 64GB;
	// - 4500GB if the memory chosen is equal or more than 96GB;
	InstanceStorage pulumi.IntPtrInput
	InstanceType    pulumi.StringPtrInput
	// The modification time of database, formatted by RFC3339 time string.
	ModifyTime pulumi.StringPtrInput
	Name       pulumi.StringPtrInput
	Password   pulumi.StringPtrInput
	// The port on which the database accepts connections, the default port is 3306 for mysql and percona.
	Port pulumi.IntPtrInput
	// The private IP address assigned to the database instance.
	PrivateIp pulumi.StringPtrInput
	// Availability zone where the standby database instance is located for the high availability database instance with multiple zone; The disaster recovery of data center can be activated by switching to the standby database instance for the high availability database instance.
	StandbyZone pulumi.StringPtrInput
	// Specifies the status of database, possible values are: `Init`, `Fail`, `Starting`, `Running`, `Shutdown`, `Shutoff`, `Delete`, `Upgrading`, `Promoting`, `Recovering` and `Recover fail`.
	Status pulumi.StringPtrInput
	// The ID of subnet.
	SubnetId pulumi.StringPtrInput
	// A tag assigned to database instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringPtrInput
	// The ID of VPC linked to the database instances.
	VpcId pulumi.StringPtrInput
}

func (DbInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbInstanceState)(nil)).Elem()
}

type dbInstanceArgs struct {
	// Availability zone where database instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Specifies when the backup starts, measured in hour, it starts at one o'clock of 1, 2, 3, 4 in the morning by default.
	BackupBeginTime *int `pulumi:"backupBeginTime"`
	// The backup for database such as "test.%" or table such as "city.address" specified in the black lists are not supported.
	BackupBlackLists []string `pulumi:"backupBlackLists"`
	// Specifies the number of backup saved per week, it is 7 backups saved per week by default.
	BackupCount *int `pulumi:"backupCount"`
	// Specifies whether the backup took place from Sunday to Saturday by displaying 7 digits. 0 stands for backup disabled and 1 stands for backup enabled. The rightmost digit specifies whether the backup took place on Sunday, and the digits from right to left specify whether the backup took place from Monday to Saturday, it's mandatory required to backup twice per week at least. such as: digits "1100000" stands for the backup took place on Saturday and Friday.
	BackupDate *string `pulumi:"backupDate"`
	// The charge type of db instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType *string `pulumi:"chargeType"`
	// The duration that you will buy the db instance (Default: `1`). The value is `0` when pay by month and the instance will be vaild till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration *int `pulumi:"duration"`
	// The type of database engine, possible values are: "mysql", "percona".
	Engine string `pulumi:"engine"`
	// The database engine version, possible values are: "5.5", "5.6", "5.7".
	// - 5.5/5.6/5.7 for mysql and percona engine.
	EngineVersion string `pulumi:"engineVersion"`
	// Specifies the allocated storage size in gigabytes (GB), range from 20 to 4500GB. The volume adjustment must be a multiple of 10 GB. The maximum disk volume for SSD type are：
	// - 500GB if the memory chosen is equal or less than 6GB;
	// - 1000GB if the memory chosen is from 8 to 16GB;
	// - 2000GB if the memory chosen is 24GB or 32GB;
	// - 3500GB if the memory chosen is 48GB or 64GB;
	// - 4500GB if the memory chosen is equal or more than 96GB;
	InstanceStorage int     `pulumi:"instanceStorage"`
	InstanceType    string  `pulumi:"instanceType"`
	Name            *string `pulumi:"name"`
	Password        *string `pulumi:"password"`
	// The port on which the database accepts connections, the default port is 3306 for mysql and percona.
	Port *int `pulumi:"port"`
	// Availability zone where the standby database instance is located for the high availability database instance with multiple zone; The disaster recovery of data center can be activated by switching to the standby database instance for the high availability database instance.
	StandbyZone *string `pulumi:"standbyZone"`
	// The ID of subnet.
	SubnetId *string `pulumi:"subnetId"`
	// A tag assigned to database instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag *string `pulumi:"tag"`
	// The ID of VPC linked to the database instances.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a DbInstance resource.
type DbInstanceArgs struct {
	// Availability zone where database instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone pulumi.StringInput
	// Specifies when the backup starts, measured in hour, it starts at one o'clock of 1, 2, 3, 4 in the morning by default.
	BackupBeginTime pulumi.IntPtrInput
	// The backup for database such as "test.%" or table such as "city.address" specified in the black lists are not supported.
	BackupBlackLists pulumi.StringArrayInput
	// Specifies the number of backup saved per week, it is 7 backups saved per week by default.
	BackupCount pulumi.IntPtrInput
	// Specifies whether the backup took place from Sunday to Saturday by displaying 7 digits. 0 stands for backup disabled and 1 stands for backup enabled. The rightmost digit specifies whether the backup took place on Sunday, and the digits from right to left specify whether the backup took place from Monday to Saturday, it's mandatory required to backup twice per week at least. such as: digits "1100000" stands for the backup took place on Saturday and Friday.
	BackupDate pulumi.StringPtrInput
	// The charge type of db instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType pulumi.StringPtrInput
	// The duration that you will buy the db instance (Default: `1`). The value is `0` when pay by month and the instance will be vaild till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration pulumi.IntPtrInput
	// The type of database engine, possible values are: "mysql", "percona".
	Engine pulumi.StringInput
	// The database engine version, possible values are: "5.5", "5.6", "5.7".
	// - 5.5/5.6/5.7 for mysql and percona engine.
	EngineVersion pulumi.StringInput
	// Specifies the allocated storage size in gigabytes (GB), range from 20 to 4500GB. The volume adjustment must be a multiple of 10 GB. The maximum disk volume for SSD type are：
	// - 500GB if the memory chosen is equal or less than 6GB;
	// - 1000GB if the memory chosen is from 8 to 16GB;
	// - 2000GB if the memory chosen is 24GB or 32GB;
	// - 3500GB if the memory chosen is 48GB or 64GB;
	// - 4500GB if the memory chosen is equal or more than 96GB;
	InstanceStorage pulumi.IntInput
	InstanceType    pulumi.StringInput
	Name            pulumi.StringPtrInput
	Password        pulumi.StringPtrInput
	// The port on which the database accepts connections, the default port is 3306 for mysql and percona.
	Port pulumi.IntPtrInput
	// Availability zone where the standby database instance is located for the high availability database instance with multiple zone; The disaster recovery of data center can be activated by switching to the standby database instance for the high availability database instance.
	StandbyZone pulumi.StringPtrInput
	// The ID of subnet.
	SubnetId pulumi.StringPtrInput
	// A tag assigned to database instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringPtrInput
	// The ID of VPC linked to the database instances.
	VpcId pulumi.StringPtrInput
}

func (DbInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbInstanceArgs)(nil)).Elem()
}
