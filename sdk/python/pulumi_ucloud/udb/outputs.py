# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetDBInstanceDbInstanceResult',
]

@pulumi.output_type
class GetDBInstanceDbInstanceResult(dict):
    def __init__(__self__, *,
                 availability_zone: str,
                 backup_begin_time: int,
                 backup_black_lists: Sequence[str],
                 backup_count: int,
                 backup_date: str,
                 charge_type: str,
                 create_time: str,
                 engine: str,
                 engine_version: str,
                 expire_time: str,
                 id: str,
                 instance_storage: int,
                 instance_type: str,
                 modify_time: str,
                 name: str,
                 port: int,
                 private_ip: str,
                 standby_zone: str,
                 status: str,
                 subnet_id: str,
                 tag: str,
                 vpc_id: str):
        """
        :param str availability_zone: Availability zone where database instances are located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
        :param int backup_begin_time: Specifies when the backup starts, measured in hour.
        :param Sequence[str] backup_black_lists: The backup for database instance such as "test.%" or table such as "city.address" specified in the black lists are not supported.
        :param int backup_count: Specifies the number of backup saved per week.
        :param str backup_date: Specifies whether the backup took place from Sunday to Saturday by displaying 7 digits. 0 stands for backup disabled and 1 stands for backup enabled. The rightmost digit specifies whether the backup took place on Sunday, and the digits from right to left specify whether the backup took place from Monday to Saturday, it's mandatory required to backup twice per week at least. such as: digits "1100000" stands for the backup took place on Saturday and Friday.
        :param str charge_type: The charge type of db instance.
        :param str create_time: The creation time of database instance , formatted by RFC3339 time string.
        :param str engine: The type of database instance engine.
        :param str engine_version: The database instance engine version.
        :param str expire_time: The expiration time of database instance , formatted by RFC3339 time string.
        :param str id: The ID of database instance.
        :param int instance_storage: Specifies the allocated storage size in gigabytes (GB).
        :param str instance_type: Specifies the type of database instance.
        :param str modify_time: The modification time of database instance , formatted by RFC3339 time string.
        :param str name: The name of database instance.
        :param int port: The port on which the database instance accepts connections.
        :param str private_ip: The private IP address assigned to the database instance.
        :param str standby_zone: Availability zone where the standby database instance is located for the high availability database instance with multiple zone.
        :param str status: Specifies the status of database instance , possible values are: `Init`, `Fail`, `Starting`, `Running`, `Shutdown`, `Shutoff`, `Delete`, `Upgrading`, `Promoting`, `Recovering` and `Recover fail`.
        :param str subnet_id: The ID of subnet linked to the database instances.
        :param str tag: A tag assigned to database instance.
        :param str vpc_id: The ID of VPC linked to the database instances.
        """
        pulumi.set(__self__, "availability_zone", availability_zone)
        pulumi.set(__self__, "backup_begin_time", backup_begin_time)
        pulumi.set(__self__, "backup_black_lists", backup_black_lists)
        pulumi.set(__self__, "backup_count", backup_count)
        pulumi.set(__self__, "backup_date", backup_date)
        pulumi.set(__self__, "charge_type", charge_type)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "engine", engine)
        pulumi.set(__self__, "engine_version", engine_version)
        pulumi.set(__self__, "expire_time", expire_time)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "instance_storage", instance_storage)
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "modify_time", modify_time)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "private_ip", private_ip)
        pulumi.set(__self__, "standby_zone", standby_zone)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "tag", tag)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        Availability zone where database instances are located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="backupBeginTime")
    def backup_begin_time(self) -> int:
        """
        Specifies when the backup starts, measured in hour.
        """
        return pulumi.get(self, "backup_begin_time")

    @property
    @pulumi.getter(name="backupBlackLists")
    def backup_black_lists(self) -> Sequence[str]:
        """
        The backup for database instance such as "test.%" or table such as "city.address" specified in the black lists are not supported.
        """
        return pulumi.get(self, "backup_black_lists")

    @property
    @pulumi.getter(name="backupCount")
    def backup_count(self) -> int:
        """
        Specifies the number of backup saved per week.
        """
        return pulumi.get(self, "backup_count")

    @property
    @pulumi.getter(name="backupDate")
    def backup_date(self) -> str:
        """
        Specifies whether the backup took place from Sunday to Saturday by displaying 7 digits. 0 stands for backup disabled and 1 stands for backup enabled. The rightmost digit specifies whether the backup took place on Sunday, and the digits from right to left specify whether the backup took place from Monday to Saturday, it's mandatory required to backup twice per week at least. such as: digits "1100000" stands for the backup took place on Saturday and Friday.
        """
        return pulumi.get(self, "backup_date")

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> str:
        """
        The charge type of db instance.
        """
        return pulumi.get(self, "charge_type")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation time of database instance , formatted by RFC3339 time string.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def engine(self) -> str:
        """
        The type of database instance engine.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> str:
        """
        The database instance engine version.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> str:
        """
        The expiration time of database instance , formatted by RFC3339 time string.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of database instance.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceStorage")
    def instance_storage(self) -> int:
        """
        Specifies the allocated storage size in gigabytes (GB).
        """
        return pulumi.get(self, "instance_storage")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        """
        Specifies the type of database instance.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="modifyTime")
    def modify_time(self) -> str:
        """
        The modification time of database instance , formatted by RFC3339 time string.
        """
        return pulumi.get(self, "modify_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of database instance.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        The port on which the database instance accepts connections.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> str:
        """
        The private IP address assigned to the database instance.
        """
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter(name="standbyZone")
    def standby_zone(self) -> str:
        """
        Availability zone where the standby database instance is located for the high availability database instance with multiple zone.
        """
        return pulumi.get(self, "standby_zone")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Specifies the status of database instance , possible values are: `Init`, `Fail`, `Starting`, `Running`, `Shutdown`, `Shutoff`, `Delete`, `Upgrading`, `Promoting`, `Recovering` and `Recover fail`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        The ID of subnet linked to the database instances.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tag(self) -> str:
        """
        A tag assigned to database instance.
        """
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The ID of VPC linked to the database instances.
        """
        return pulumi.get(self, "vpc_id")


