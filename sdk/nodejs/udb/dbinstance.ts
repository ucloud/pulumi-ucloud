// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * DB Instance can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import ucloud:udb/dBInstance:DBInstance example udbha-abc123456
 * ```
 */
export class DBInstance extends pulumi.CustomResource {
    /**
     * Get an existing DBInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DBInstanceState, opts?: pulumi.CustomResourceOptions): DBInstance {
        return new DBInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ucloud:udb/dBInstance:DBInstance';

    /**
     * Returns true if the given object is an instance of DBInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DBInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DBInstance.__pulumiType;
    }

    public readonly allowStoppingForUpdate!: pulumi.Output<boolean | undefined>;
    /**
     * Availability zone where database instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * Specifies when the backup starts, measured in hour, it starts at one o'clock of 1, 2, 3, 4 in the morning by default.
     */
    public readonly backupBeginTime!: pulumi.Output<number>;
    /**
     * The backup for database such as "test.%" or table such as "city.address" specified in the black lists are not supported.
     */
    public readonly backupBlackLists!: pulumi.Output<string[]>;
    /**
     * Specifies the number of backup saved per week, it is 7 backups saved per week by default.
     */
    public readonly backupCount!: pulumi.Output<number | undefined>;
    /**
     * Specifies whether the backup took place from Sunday to Saturday by displaying 7 digits. 0 stands for backup disabled and 1 stands for backup enabled. The rightmost digit specifies whether the backup took place on Sunday, and the digits from right to left specify whether the backup took place from Monday to Saturday, it's mandatory required to backup twice per week at least. such as: digits "1100000" stands for the backup took place on Saturday and Friday.
     */
    public readonly backupDate!: pulumi.Output<string>;
    /**
     * The charge type of db instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
     */
    public readonly chargeType!: pulumi.Output<string>;
    /**
     * The creation time of database, formatted by RFC3339 time string.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The duration that you will buy the db instance (Default: `1`). The value is `0` when pay by month and the instance will be vaild till the last day of that month. It is not required when `dynamic` (pay by hour).
     */
    public readonly duration!: pulumi.Output<number | undefined>;
    /**
     * The type of database engine, possible values are: "mysql", "percona".
     */
    public readonly engine!: pulumi.Output<string>;
    /**
     * The database engine version, possible values are: "5.5", "5.6", "5.7".
     * - 5.5/5.6/5.7 for mysql and percona engine.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * The expiration time of database, formatted by RFC3339 time string.
     */
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    /**
     * Specifies the allocated storage size in gigabytes (GB), range from 20 to 4500GB. The volume adjustment must be a multiple of 10 GB. The maximum disk volume for SSD type are：
     * - 500GB if the memory chosen is equal or less than 6GB;
     * - 1000GB if the memory chosen is from 8 to 16GB;
     * - 2000GB if the memory chosen is 24GB or 32GB;
     * - 3500GB if the memory chosen is 48GB or 64GB;
     * - 4500GB if the memory chosen is equal or more than 96GB;
     */
    public readonly instanceStorage!: pulumi.Output<number>;
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * The modification time of database, formatted by RFC3339 time string.
     */
    public /*out*/ readonly modifyTime!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly parameterGroup!: pulumi.Output<string>;
    public readonly password!: pulumi.Output<string>;
    /**
     * The port on which the database accepts connections, the default port is 3306 for mysql and percona.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * The private IP address assigned to the database instance.
     */
    public /*out*/ readonly privateIp!: pulumi.Output<string>;
    /**
     * Availability zone where the standby database instance is located for the high availability database instance with multiple zone; The disaster recovery of data center can be activated by switching to the standby database instance for the high availability database instance.
     */
    public readonly standbyZone!: pulumi.Output<string | undefined>;
    /**
     * Specifies the status of database, possible values are: `Init`, `Fail`, `Starting`, `Running`, `Shutdown`, `Shutoff`, `Delete`, `Upgrading`, `Promoting`, `Recovering` and `Recover fail`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The ID of subnet.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * A tag assigned to database instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
     */
    public readonly tag!: pulumi.Output<string>;
    /**
     * The ID of VPC linked to the database instances.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a DBInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DBInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DBInstanceArgs | DBInstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DBInstanceState | undefined;
            inputs["allowStoppingForUpdate"] = state ? state.allowStoppingForUpdate : undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["backupBeginTime"] = state ? state.backupBeginTime : undefined;
            inputs["backupBlackLists"] = state ? state.backupBlackLists : undefined;
            inputs["backupCount"] = state ? state.backupCount : undefined;
            inputs["backupDate"] = state ? state.backupDate : undefined;
            inputs["chargeType"] = state ? state.chargeType : undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["duration"] = state ? state.duration : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["engineVersion"] = state ? state.engineVersion : undefined;
            inputs["expireTime"] = state ? state.expireTime : undefined;
            inputs["instanceStorage"] = state ? state.instanceStorage : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["modifyTime"] = state ? state.modifyTime : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parameterGroup"] = state ? state.parameterGroup : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["privateIp"] = state ? state.privateIp : undefined;
            inputs["standbyZone"] = state ? state.standbyZone : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["tag"] = state ? state.tag : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as DBInstanceArgs | undefined;
            if ((!args || args.availabilityZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            if ((!args || args.engine === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engine'");
            }
            if ((!args || args.engineVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engineVersion'");
            }
            if ((!args || args.instanceStorage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceStorage'");
            }
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            inputs["allowStoppingForUpdate"] = args ? args.allowStoppingForUpdate : undefined;
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["backupBeginTime"] = args ? args.backupBeginTime : undefined;
            inputs["backupBlackLists"] = args ? args.backupBlackLists : undefined;
            inputs["backupCount"] = args ? args.backupCount : undefined;
            inputs["backupDate"] = args ? args.backupDate : undefined;
            inputs["chargeType"] = args ? args.chargeType : undefined;
            inputs["duration"] = args ? args.duration : undefined;
            inputs["engine"] = args ? args.engine : undefined;
            inputs["engineVersion"] = args ? args.engineVersion : undefined;
            inputs["instanceStorage"] = args ? args.instanceStorage : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameterGroup"] = args ? args.parameterGroup : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["standbyZone"] = args ? args.standbyZone : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tag"] = args ? args.tag : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["expireTime"] = undefined /*out*/;
            inputs["modifyTime"] = undefined /*out*/;
            inputs["privateIp"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DBInstance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DBInstance resources.
 */
export interface DBInstanceState {
    allowStoppingForUpdate?: pulumi.Input<boolean>;
    /**
     * Availability zone where database instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Specifies when the backup starts, measured in hour, it starts at one o'clock of 1, 2, 3, 4 in the morning by default.
     */
    backupBeginTime?: pulumi.Input<number>;
    /**
     * The backup for database such as "test.%" or table such as "city.address" specified in the black lists are not supported.
     */
    backupBlackLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the number of backup saved per week, it is 7 backups saved per week by default.
     */
    backupCount?: pulumi.Input<number>;
    /**
     * Specifies whether the backup took place from Sunday to Saturday by displaying 7 digits. 0 stands for backup disabled and 1 stands for backup enabled. The rightmost digit specifies whether the backup took place on Sunday, and the digits from right to left specify whether the backup took place from Monday to Saturday, it's mandatory required to backup twice per week at least. such as: digits "1100000" stands for the backup took place on Saturday and Friday.
     */
    backupDate?: pulumi.Input<string>;
    /**
     * The charge type of db instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
     */
    chargeType?: pulumi.Input<string>;
    /**
     * The creation time of database, formatted by RFC3339 time string.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The duration that you will buy the db instance (Default: `1`). The value is `0` when pay by month and the instance will be vaild till the last day of that month. It is not required when `dynamic` (pay by hour).
     */
    duration?: pulumi.Input<number>;
    /**
     * The type of database engine, possible values are: "mysql", "percona".
     */
    engine?: pulumi.Input<string>;
    /**
     * The database engine version, possible values are: "5.5", "5.6", "5.7".
     * - 5.5/5.6/5.7 for mysql and percona engine.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * The expiration time of database, formatted by RFC3339 time string.
     */
    expireTime?: pulumi.Input<string>;
    /**
     * Specifies the allocated storage size in gigabytes (GB), range from 20 to 4500GB. The volume adjustment must be a multiple of 10 GB. The maximum disk volume for SSD type are：
     * - 500GB if the memory chosen is equal or less than 6GB;
     * - 1000GB if the memory chosen is from 8 to 16GB;
     * - 2000GB if the memory chosen is 24GB or 32GB;
     * - 3500GB if the memory chosen is 48GB or 64GB;
     * - 4500GB if the memory chosen is equal or more than 96GB;
     */
    instanceStorage?: pulumi.Input<number>;
    instanceType?: pulumi.Input<string>;
    /**
     * The modification time of database, formatted by RFC3339 time string.
     */
    modifyTime?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    parameterGroup?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    /**
     * The port on which the database accepts connections, the default port is 3306 for mysql and percona.
     */
    port?: pulumi.Input<number>;
    /**
     * The private IP address assigned to the database instance.
     */
    privateIp?: pulumi.Input<string>;
    /**
     * Availability zone where the standby database instance is located for the high availability database instance with multiple zone; The disaster recovery of data center can be activated by switching to the standby database instance for the high availability database instance.
     */
    standbyZone?: pulumi.Input<string>;
    /**
     * Specifies the status of database, possible values are: `Init`, `Fail`, `Starting`, `Running`, `Shutdown`, `Shutoff`, `Delete`, `Upgrading`, `Promoting`, `Recovering` and `Recover fail`.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of subnet.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * A tag assigned to database instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
     */
    tag?: pulumi.Input<string>;
    /**
     * The ID of VPC linked to the database instances.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DBInstance resource.
 */
export interface DBInstanceArgs {
    allowStoppingForUpdate?: pulumi.Input<boolean>;
    /**
     * Availability zone where database instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
     */
    availabilityZone: pulumi.Input<string>;
    /**
     * Specifies when the backup starts, measured in hour, it starts at one o'clock of 1, 2, 3, 4 in the morning by default.
     */
    backupBeginTime?: pulumi.Input<number>;
    /**
     * The backup for database such as "test.%" or table such as "city.address" specified in the black lists are not supported.
     */
    backupBlackLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the number of backup saved per week, it is 7 backups saved per week by default.
     */
    backupCount?: pulumi.Input<number>;
    /**
     * Specifies whether the backup took place from Sunday to Saturday by displaying 7 digits. 0 stands for backup disabled and 1 stands for backup enabled. The rightmost digit specifies whether the backup took place on Sunday, and the digits from right to left specify whether the backup took place from Monday to Saturday, it's mandatory required to backup twice per week at least. such as: digits "1100000" stands for the backup took place on Saturday and Friday.
     */
    backupDate?: pulumi.Input<string>;
    /**
     * The charge type of db instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
     */
    chargeType?: pulumi.Input<string>;
    /**
     * The duration that you will buy the db instance (Default: `1`). The value is `0` when pay by month and the instance will be vaild till the last day of that month. It is not required when `dynamic` (pay by hour).
     */
    duration?: pulumi.Input<number>;
    /**
     * The type of database engine, possible values are: "mysql", "percona".
     */
    engine: pulumi.Input<string>;
    /**
     * The database engine version, possible values are: "5.5", "5.6", "5.7".
     * - 5.5/5.6/5.7 for mysql and percona engine.
     */
    engineVersion: pulumi.Input<string>;
    /**
     * Specifies the allocated storage size in gigabytes (GB), range from 20 to 4500GB. The volume adjustment must be a multiple of 10 GB. The maximum disk volume for SSD type are：
     * - 500GB if the memory chosen is equal or less than 6GB;
     * - 1000GB if the memory chosen is from 8 to 16GB;
     * - 2000GB if the memory chosen is 24GB or 32GB;
     * - 3500GB if the memory chosen is 48GB or 64GB;
     * - 4500GB if the memory chosen is equal or more than 96GB;
     */
    instanceStorage: pulumi.Input<number>;
    instanceType: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    parameterGroup?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    /**
     * The port on which the database accepts connections, the default port is 3306 for mysql and percona.
     */
    port?: pulumi.Input<number>;
    /**
     * Availability zone where the standby database instance is located for the high availability database instance with multiple zone; The disaster recovery of data center can be activated by switching to the standby database instance for the high availability database instance.
     */
    standbyZone?: pulumi.Input<string>;
    /**
     * The ID of subnet.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * A tag assigned to database instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
     */
    tag?: pulumi.Input<string>;
    /**
     * The ID of VPC linked to the database instances.
     */
    vpcId?: pulumi.Input<string>;
}