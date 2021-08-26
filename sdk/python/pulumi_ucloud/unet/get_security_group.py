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
    'GetSecurityGroupResult',
    'AwaitableGetSecurityGroupResult',
    'get_security_group',
]

@pulumi.output_type
class GetSecurityGroupResult:
    """
    A collection of values returned by getSecurityGroup.
    """
    def __init__(__self__, id=None, ids=None, name_regex=None, output_file=None, security_groups=None, total_count=None, type=None):
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
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        pulumi.set(__self__, "security_groups", security_groups)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

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
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Sequence['outputs.GetSecurityGroupSecurityGroupResult']:
        """
        It is a nested type which documented below.
        """
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        Total number of Security Group resources that satisfy the condition.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of Security Group.
        """
        return pulumi.get(self, "type")


class AwaitableGetSecurityGroupResult(GetSecurityGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecurityGroupResult(
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            output_file=self.output_file,
            security_groups=self.security_groups,
            total_count=self.total_count,
            type=self.type)


def get_security_group(ids: Optional[Sequence[str]] = None,
                       name_regex: Optional[str] = None,
                       output_file: Optional[str] = None,
                       type: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecurityGroupResult:
    """
    This data source provides a list of Security Group resources according to their Security Group ID, name and resource id.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ucloud as ucloud

    example = ucloud.unet.get_security_group()
    pulumi.export("first", example.security_groups[0].id)
    ```


    :param Sequence[str] ids: A list of Security Group IDs, all the Security Group resources belong to this region will be retrieved if the ID is `[]`.
    :param str name_regex: A regex string to filter resulting Security Group resources by name.
    :param str type: The type of Security Group. Possible values are: `recommend_web` as the default Web security group that UCloud recommend to users, default opened port include 80, 443, 22, 3389, `recommend_non_web` as the default non Web security group that UCloud recommend to users, default opened port include 22, 3389, `user_defined` as the security groups defined by users. You may refer to [security group](https://docs.ucloud.cn/network/firewall/firewall.html).
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['type'] = type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ucloud:unet/getSecurityGroup:getSecurityGroup', __args__, opts=opts, typ=GetSecurityGroupResult).value

    return AwaitableGetSecurityGroupResult(
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        security_groups=__ret__.security_groups,
        total_count=__ret__.total_count,
        type=__ret__.type)