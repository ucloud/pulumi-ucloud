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
    'GetVPNGatewayResult',
    'AwaitableGetVPNGatewayResult',
    'get_vpn_gateway',
]

@pulumi.output_type
class GetVPNGatewayResult:
    """
    A collection of values returned by getVPNGateway.
    """
    def __init__(__self__, id=None, ids=None, name_regex=None, output_file=None, tag=None, total_count=None, vpc_id=None, vpn_gateways=None):
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
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        pulumi.set(__self__, "tag", tag)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)
        if vpn_gateways and not isinstance(vpn_gateways, list):
            raise TypeError("Expected argument 'vpn_gateways' to be a list")
        pulumi.set(__self__, "vpn_gateways", vpn_gateways)

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
    @pulumi.getter
    def tag(self) -> Optional[str]:
        """
        A tag assigned to the VPN Gateway.
        """
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        Total number of VPN Gateways that satisfy the condition.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        The ID of VPC linked to the VPN Gateway.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vpnGateways")
    def vpn_gateways(self) -> Sequence['outputs.GetVPNGatewayVpnGatewayResult']:
        """
        It is a nested type. VPN Gateways documented below.
        """
        return pulumi.get(self, "vpn_gateways")


class AwaitableGetVPNGatewayResult(GetVPNGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVPNGatewayResult(
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            output_file=self.output_file,
            tag=self.tag,
            total_count=self.total_count,
            vpc_id=self.vpc_id,
            vpn_gateways=self.vpn_gateways)


def get_vpn_gateway(ids: Optional[Sequence[str]] = None,
                    name_regex: Optional[str] = None,
                    output_file: Optional[str] = None,
                    tag: Optional[str] = None,
                    vpc_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVPNGatewayResult:
    """
    This data source providers a list of VPN Gateway resources according to their ID, name, vpc and tag.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ucloud as ucloud

    example = ucloud.ipsecvpn.get_vpn_gateway()
    pulumi.export("first", example.vpn_gateways[0].id)
    ```


    :param Sequence[str] ids: A list of VPN Gateway IDs, all the VPN Gateways belongs to the defined region will be retrieved if this argument is `[]`.
    :param str name_regex: A regex string to filter resulting VPN Gateways by name.
    :param str tag: A tag assigned to VPN Gateway.
    :param str vpc_id: The ID of VPC linked to the VPN Gateway.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['tag'] = tag
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ucloud:ipsecvpn/getVPNGateway:getVPNGateway', __args__, opts=opts, typ=GetVPNGatewayResult).value

    return AwaitableGetVPNGatewayResult(
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        tag=__ret__.tag,
        total_count=__ret__.total_count,
        vpc_id=__ret__.vpc_id,
        vpn_gateways=__ret__.vpn_gateways)
