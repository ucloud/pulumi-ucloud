# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'EIPIpSetArgs',
    'EIPResourceArgs',
    'SecurityGroupRuleArgs',
]

@pulumi.input_type
class EIPIpSetArgs:
    def __init__(__self__, *,
                 internet_type: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] internet_type: Type of Elastic IP routes. Possible values are: `international` as international BGP IP and `bgp` as china mainland BGP IP.
        """
        if internet_type is not None:
            pulumi.set(__self__, "internet_type", internet_type)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)

    @property
    @pulumi.getter(name="internetType")
    def internet_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of Elastic IP routes. Possible values are: `international` as international BGP IP and `bgp` as china mainland BGP IP.
        """
        return pulumi.get(self, "internet_type")

    @internet_type.setter
    def internet_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internet_type", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)


@pulumi.input_type
class EIPResourceArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: The ID of the resource with EIP attached.
        :param pulumi.Input[str] type: The type of resource with EIP attached. Possible values are `instance` as instance, `lb` as load balancer.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource with EIP attached.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of resource with EIP attached. Possible values are `instance` as instance, `lb` as load balancer.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class SecurityGroupRuleArgs:
    def __init__(__self__, *,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 port_range: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] cidr_block: The cidr block of source.
        :param pulumi.Input[str] policy: Authorization policy. Possible values are: `accept`, `drop`.
        :param pulumi.Input[str] port_range: The range of port numbers, range: 1-65535. (eg: `port` or `port1-port2`).
        :param pulumi.Input[str] priority: Rule priority. Possible values are: `high`, `medium`, `low`.
        :param pulumi.Input[str] protocol: The protocol. Possible values are: `tcp`, `udp`, `icmp`, `gre`.
        """
        if cidr_block is not None:
            pulumi.set(__self__, "cidr_block", cidr_block)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if port_range is not None:
            pulumi.set(__self__, "port_range", port_range)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The cidr block of source.
        """
        return pulumi.get(self, "cidr_block")

    @cidr_block.setter
    def cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr_block", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        Authorization policy. Possible values are: `accept`, `drop`.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter(name="portRange")
    def port_range(self) -> Optional[pulumi.Input[str]]:
        """
        The range of port numbers, range: 1-65535. (eg: `port` or `port1-port2`).
        """
        return pulumi.get(self, "port_range")

    @port_range.setter
    def port_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port_range", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[str]]:
        """
        Rule priority. Possible values are: `high`, `medium`, `low`.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The protocol. Possible values are: `tcp`, `udp`, `icmp`, `gre`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)


