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
    'GetLBResult',
    'AwaitableGetLBResult',
    'get_lb',
]

@pulumi.output_type
class GetLBResult:
    """
    A collection of values returned by getLB.
    """
    def __init__(__self__, id=None, ids=None, lbs=None, name_regex=None, output_file=None, subnet_id=None, total_count=None, vpc_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if lbs and not isinstance(lbs, list):
            raise TypeError("Expected argument 'lbs' to be a list")
        pulumi.set(__self__, "lbs", lbs)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

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
    @pulumi.getter
    def lbs(self) -> Sequence['outputs.GetLBLbResult']:
        """
        It is a nested type which documented below.
        """
        return pulumi.get(self, "lbs")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        (Optional) The ID of subnet that intrant load balancer belongs to.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        Total number of Load Balancers that satisfy the condition.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The ID of the VPC linked to the Load Balancers.
        """
        return pulumi.get(self, "vpc_id")


class AwaitableGetLBResult(GetLBResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLBResult(
            id=self.id,
            ids=self.ids,
            lbs=self.lbs,
            name_regex=self.name_regex,
            output_file=self.output_file,
            subnet_id=self.subnet_id,
            total_count=self.total_count,
            vpc_id=self.vpc_id)


def get_lb(ids: Optional[Sequence[str]] = None,
           name_regex: Optional[str] = None,
           output_file: Optional[str] = None,
           subnet_id: Optional[str] = None,
           vpc_id: Optional[str] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLBResult:
    """
    This data source provides a list of Load Balancer resources according to their Load Balancer ID, VPC ID and Subnet ID.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ucloud as ucloud

    example = ucloud.ulb.get_lb()
    pulumi.export("first", example.lbs[0].id)
    ```


    :param Sequence[str] ids: A list of Load Balancer IDs, all the LBs belong to this region will be retrieved if the ID is `[]`.
    :param str name_regex: A regex string to filter resulting lbs by name.
    :param str subnet_id: The ID of subnet that intrant load balancer belongs to.
    :param str vpc_id: The ID of the VPC linked to the Load Balancers.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['subnetId'] = subnet_id
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ucloud:ulb/getLB:getLB', __args__, opts=opts, typ=GetLBResult).value

    return AwaitableGetLBResult(
        id=__ret__.id,
        ids=__ret__.ids,
        lbs=__ret__.lbs,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        subnet_id=__ret__.subnet_id,
        total_count=__ret__.total_count,
        vpc_id=__ret__.vpc_id)
