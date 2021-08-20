# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetDBInstanceResult',
    'AwaitableGetDBInstanceResult',
    'get_db_instance',
]

@pulumi.output_type
class GetDBInstanceResult:
    """
    A collection of values returned by getDBInstance.
    """
    def __init__(__self__, availability_zone=None, db_instances=None, id=None, ids=None, name_regex=None, output_file=None, total_count=None):
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if db_instances and not isinstance(db_instances, list):
            raise TypeError("Expected argument 'db_instances' to be a list")
        pulumi.set(__self__, "db_instances", db_instances)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[str]:
        """
        Availability zone where database instance is located.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="dbInstances")
    def db_instances(self) -> Sequence['outputs.GetDBInstanceDbInstanceResult']:
        """
        It is a nested type which documented below.
        """
        return pulumi.get(self, "db_instances")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        Total number of database instances that satisfy the condition.
        """
        return pulumi.get(self, "total_count")


class AwaitableGetDBInstanceResult(GetDBInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDBInstanceResult(
            availability_zone=self.availability_zone,
            db_instances=self.db_instances,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            output_file=self.output_file,
            total_count=self.total_count)


def get_db_instance(availability_zone: Optional[str] = None,
                    ids: Optional[Sequence[str]] = None,
                    name_regex: Optional[str] = None,
                    output_file: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDBInstanceResult:
    """
    This data source provides a list of database instance resources according to their database instance ID and name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ucloud as ucloud

    example = ucloud.udb.get_db_instance()
    pulumi.export("first", example.db_instances[0].id)
    ```


    :param str availability_zone: Availability zone where database instances are located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
    :param Sequence[str] ids: A list of database instance IDs, all the database instances belong to this region will be retrieved if the ID is `[]`.
    :param str name_regex: A regex string to filter resulting database instances by name.
    """
    __args__ = dict()
    __args__['availabilityZone'] = availability_zone
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ucloud:udb/getDBInstance:getDBInstance', __args__, opts=opts, typ=GetDBInstanceResult).value

    return AwaitableGetDBInstanceResult(
        availability_zone=__ret__.availability_zone,
        db_instances=__ret__.db_instances,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        total_count=__ret__.total_count)
