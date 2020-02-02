# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class LookupVpnCustomerGatewaysResult:
    """
    A collection of values returned by lookupVpnCustomerGateways.
    """
    def __init__(__self__, ids=None, name_regex=None, output_file=None, tag=None, total_count=None, vpn_customer_gateways=None, id=None):
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
        A tag assigned to the VPN Customer Gateway.
        """
        if total_count and not isinstance(total_count, float):
            raise TypeError("Expected argument 'total_count' to be a float")
        __self__.total_count = total_count
        """
        Total number of VPN Customer Gateways that satisfy the condition.
        """
        if vpn_customer_gateways and not isinstance(vpn_customer_gateways, list):
            raise TypeError("Expected argument 'vpn_customer_gateways' to be a list")
        __self__.vpn_customer_gateways = vpn_customer_gateways
        """
        It is a nested type. VPN Customer Gateways documented below.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableLookupVpnCustomerGatewaysResult(LookupVpnCustomerGatewaysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return LookupVpnCustomerGatewaysResult(
            ids=self.ids,
            name_regex=self.name_regex,
            output_file=self.output_file,
            tag=self.tag,
            total_count=self.total_count,
            vpn_customer_gateways=self.vpn_customer_gateways,
            id=self.id)

def lookup_vpn_customer_gateways(ids=None,name_regex=None,output_file=None,tag=None,opts=None):
    """
    This data source providers a list of VPN Customer Gateway resources according to their ID, name and tag.
    
    :param list ids: A list of VPN Customer Gateway IDs, all the VPN Customer Gateways belongs to the defined region will be retrieved if this argument is "".
    :param str name_regex: A regex string to filter resulting VPN Customer Gateways by name.
    :param str tag: A tag assigned to VPN Customer Gateway.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/d/vpn_customer_gateways.html.markdown.
    """
    __args__ = dict()

    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['tag'] = tag
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ucloud:ipsecvpn/lookupVpnCustomerGateways:lookupVpnCustomerGateways', __args__, opts=opts).value

    return AwaitableLookupVpnCustomerGatewaysResult(
        ids=__ret__.get('ids'),
        name_regex=__ret__.get('nameRegex'),
        output_file=__ret__.get('outputFile'),
        tag=__ret__.get('tag'),
        total_count=__ret__.get('totalCount'),
        vpn_customer_gateways=__ret__.get('vpnCustomerGateways'),
        id=__ret__.get('id'))
