# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'MemcachedInstanceIpSet',
    'RedisInstanceIpSet',
]

@pulumi.output_type
class MemcachedInstanceIpSet(dict):
    def __init__(__self__, *,
                 ip: Optional[str] = None,
                 port: Optional[int] = None):
        """
        :param str ip: The virtual ip of Memcache instance.
        :param int port: The port on which Memcache instance accepts connections, it is 6379 by default.
        """
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def ip(self) -> Optional[str]:
        """
        The virtual ip of Memcache instance.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def port(self) -> Optional[int]:
        """
        The port on which Memcache instance accepts connections, it is 6379 by default.
        """
        return pulumi.get(self, "port")


@pulumi.output_type
class RedisInstanceIpSet(dict):
    def __init__(__self__, *,
                 ip: Optional[str] = None,
                 port: Optional[int] = None):
        """
        :param str ip: The virtual ip of Redis instance.
        :param int port: The port on which Redis instance accepts connections, it is 6379 by default.
        """
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def ip(self) -> Optional[str]:
        """
        The virtual ip of Redis instance.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def port(self) -> Optional[int]:
        """
        The port on which Redis instance accepts connections, it is 6379 by default.
        """
        return pulumi.get(self, "port")


