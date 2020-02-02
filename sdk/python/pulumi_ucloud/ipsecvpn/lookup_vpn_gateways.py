# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class LookupVpnGatewaysResult:
    """
    A collection of values returned by lookupVpnGateways.
    """
    def __init__(__self__, ids=None, name_regex=None, output_file=None, tag=None, total_count=None, vpc_id=None, vpn_gateways=None, id=None):
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        __self__.tag = tag
        """
        A tag assigned to the VPN Gateway.
        """
        if total_count and not isinstance(total_count, float):
            raise TypeError("Expected argument 'total_count' to be a float")
        __self__.total_count = total_count
        """
        Total number of VPN Gateways that satisfy the condition.
        """
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        __self__.vpc_id = vpc_id
        """
        The ID of VPC linked to the VPN Gateway.
        """
        if vpn_gateways and not isinstance(vpn_gateways, list):
            raise TypeError("Expected argument 'vpn_gateways' to be a list")
        __self__.vpn_gateways = vpn_gateways
        """
        It is a nested type. VPN Gateways documented below.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableLookupVpnGatewaysResult(LookupVpnGatewaysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return LookupVpnGatewaysResult(
            ids=self.ids,
            name_regex=self.name_regex,
            output_file=self.output_file,
            tag=self.tag,
            total_count=self.total_count,
            vpc_id=self.vpc_id,
            vpn_gateways=self.vpn_gateways,
            id=self.id)

def lookup_vpn_gateways(ids=None,name_regex=None,output_file=None,tag=None,vpc_id=None,opts=None):
    """
    This data source providers a list of VPN Gateway resources according to their ID, name, vpc and tag.
    
    :param list ids: A list of VPN Gateway IDs, all the VPN Gateways belongs to the defined region will be retrieved if this argument is "".
    :param str name_regex: A regex string to filter resulting VPN Gateways by name.
    :param str tag: A tag assigned to VPN Gateway.
    :param str vpc_id: The ID of VPC linked to the VPN Gateway.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/d/vpn_gateways.html.markdown.
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
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ucloud:ipsecvpn/lookupVpnGateways:lookupVpnGateways', __args__, opts=opts).value

    return AwaitableLookupVpnGatewaysResult(
        ids=__ret__.get('ids'),
        name_regex=__ret__.get('nameRegex'),
        output_file=__ret__.get('outputFile'),
        tag=__ret__.get('tag'),
        total_count=__ret__.get('totalCount'),
        vpc_id=__ret__.get('vpcId'),
        vpn_gateways=__ret__.get('vpnGateways'),
        id=__ret__.get('id'))
