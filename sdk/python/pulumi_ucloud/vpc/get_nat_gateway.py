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
    'GetNATGatewayResult',
    'AwaitableGetNATGatewayResult',
    'get_nat_gateway',
]

@pulumi.output_type
class GetNATGatewayResult:
    """
    A collection of values returned by getNATGateway.
    """
    def __init__(__self__, id=None, ids=None, name_regex=None, nat_gateways=None, output_file=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if nat_gateways and not isinstance(nat_gateways, list):
            raise TypeError("Expected argument 'nat_gateways' to be a list")
        pulumi.set(__self__, "nat_gateways", nat_gateways)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

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
    @pulumi.getter(name="natGateways")
    def nat_gateways(self) -> Sequence['outputs.GetNATGatewayNatGatewayResult']:
        """
        It is a nested type. Nat Gateways documented below.
        """
        return pulumi.get(self, "nat_gateways")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        Total number of Nat Gateways that satisfy the condition.
        """
        return pulumi.get(self, "total_count")


class AwaitableGetNATGatewayResult(GetNATGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNATGatewayResult(
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            nat_gateways=self.nat_gateways,
            output_file=self.output_file,
            total_count=self.total_count)


def get_nat_gateway(ids: Optional[Sequence[str]] = None,
                    name_regex: Optional[str] = None,
                    output_file: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNATGatewayResult:
    """
    This data source providers a list of Nat Gateway resources according to their ID and name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ucloud as ucloud

    example = ucloud.vpc.get_nat_gateway()
    pulumi.export("first", example.nat_gateways[0].id)
    ```


    :param Sequence[str] ids: A list of Nat Gateway IDs, all the Nat Gateways belongs to the defined region will be retrieved if this argument is `[]`.
    :param str name_regex: A regex string to filter resulting Nat Gateways by name.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ucloud:vpc/getNATGateway:getNATGateway', __args__, opts=opts, typ=GetNATGatewayResult).value

    return AwaitableGetNATGatewayResult(
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        nat_gateways=__ret__.nat_gateways,
        output_file=__ret__.output_file,
        total_count=__ret__.total_count)