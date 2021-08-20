# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['UDPNConnectionArgs', 'UDPNConnection']

@pulumi.input_type
class UDPNConnectionArgs:
    def __init__(__self__, *,
                 peer_region: pulumi.Input[str],
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a UDPNConnection resource.
        :param pulumi.Input[str] peer_region: The correspondent region of dedicated connection, please refer to the region and [availability zone list](https://docs.ucloud.cn/api/summary/regionlist) and [UDPN price list](https://docs.ucloud.cn/network/udpn/udpn_price).
        :param pulumi.Input[int] bandwidth: Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). range from 2 - 1000M. The default value is "1"
        :param pulumi.Input[str] charge_type: Charge type. Possible values are: "year" as pay by year, "month" as pay by month, "dynamic" as pay by hour. The default value is "month".
        :param pulumi.Input[int] duration: The duration that you will buy the resource, the default value is "1". It is not required when "dynamic" (pay by hour), the value is "0" when pay by month and the instance will be valid till the last day of that month.
        """
        pulumi.set(__self__, "peer_region", peer_region)
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if charge_type is not None:
            pulumi.set(__self__, "charge_type", charge_type)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)

    @property
    @pulumi.getter(name="peerRegion")
    def peer_region(self) -> pulumi.Input[str]:
        """
        The correspondent region of dedicated connection, please refer to the region and [availability zone list](https://docs.ucloud.cn/api/summary/regionlist) and [UDPN price list](https://docs.ucloud.cn/network/udpn/udpn_price).
        """
        return pulumi.get(self, "peer_region")

    @peer_region.setter
    def peer_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "peer_region", value)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). range from 2 - 1000M. The default value is "1"
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        Charge type. Possible values are: "year" as pay by year, "month" as pay by month, "dynamic" as pay by hour. The default value is "month".
        """
        return pulumi.get(self, "charge_type")

    @charge_type.setter
    def charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charge_type", value)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[int]]:
        """
        The duration that you will buy the resource, the default value is "1". It is not required when "dynamic" (pay by hour), the value is "0" when pay by month and the instance will be valid till the last day of that month.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration", value)


@pulumi.input_type
class _UDPNConnectionState:
    def __init__(__self__, *,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[int]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 peer_region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UDPNConnection resources.
        :param pulumi.Input[int] bandwidth: Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). range from 2 - 1000M. The default value is "1"
        :param pulumi.Input[str] charge_type: Charge type. Possible values are: "year" as pay by year, "month" as pay by month, "dynamic" as pay by hour. The default value is "month".
        :param pulumi.Input[str] create_time: The time of creation for UDPN connection, formatted by RFC3339 time string.
        :param pulumi.Input[int] duration: The duration that you will buy the resource, the default value is "1". It is not required when "dynamic" (pay by hour), the value is "0" when pay by month and the instance will be valid till the last day of that month.
        :param pulumi.Input[str] expire_time: The expiration time for UDPN connection, formatted by RFC3339 time string.
        :param pulumi.Input[str] peer_region: The correspondent region of dedicated connection, please refer to the region and [availability zone list](https://docs.ucloud.cn/api/summary/regionlist) and [UDPN price list](https://docs.ucloud.cn/network/udpn/udpn_price).
        """
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if charge_type is not None:
            pulumi.set(__self__, "charge_type", charge_type)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if peer_region is not None:
            pulumi.set(__self__, "peer_region", peer_region)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). range from 2 - 1000M. The default value is "1"
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        Charge type. Possible values are: "year" as pay by year, "month" as pay by month, "dynamic" as pay by hour. The default value is "month".
        """
        return pulumi.get(self, "charge_type")

    @charge_type.setter
    def charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charge_type", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time of creation for UDPN connection, formatted by RFC3339 time string.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[int]]:
        """
        The duration that you will buy the resource, the default value is "1". It is not required when "dynamic" (pay by hour), the value is "0" when pay by month and the instance will be valid till the last day of that month.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[pulumi.Input[str]]:
        """
        The expiration time for UDPN connection, formatted by RFC3339 time string.
        """
        return pulumi.get(self, "expire_time")

    @expire_time.setter
    def expire_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_time", value)

    @property
    @pulumi.getter(name="peerRegion")
    def peer_region(self) -> Optional[pulumi.Input[str]]:
        """
        The correspondent region of dedicated connection, please refer to the region and [availability zone list](https://docs.ucloud.cn/api/summary/regionlist) and [UDPN price list](https://docs.ucloud.cn/network/udpn/udpn_price).
        """
        return pulumi.get(self, "peer_region")

    @peer_region.setter
    def peer_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_region", value)


class UDPNConnection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[int]] = None,
                 peer_region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        UDPN (UCloud Dedicated Private Network)，you can use Dedicated Private Network to achieve high-speed, stable, secure, and dedicated communications between different data centers. The most frequent scenario is to create network connection of clusters across regions.

        > **VPC Peering Connections with UDPN Connection** The cross-region Dedicated Private Network must be established if the two VPCs located in different regions are expected to be connected.

        > **Note** The additional packet head will be added and included in the overall length of packet due to the tunneling UDPN adopted. Since the number of the bytes of packet head is fixed, the bigger data packet is, the less usage will be taken for the packet head.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ucloud as ucloud

        # connect provider's region (cn-bj2) and peer region (cn-sh2)
        example = ucloud.udpn.UDPNConnection("example",
            bandwidth=2,
            peer_region="cn-sh2")
        ```

        ## Import

        UDPN connection can be imported using the `id`, e.g.

        ```sh
         $ pulumi import ucloud:udpn/uDPNConnection:UDPNConnection example udpn-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). range from 2 - 1000M. The default value is "1"
        :param pulumi.Input[str] charge_type: Charge type. Possible values are: "year" as pay by year, "month" as pay by month, "dynamic" as pay by hour. The default value is "month".
        :param pulumi.Input[int] duration: The duration that you will buy the resource, the default value is "1". It is not required when "dynamic" (pay by hour), the value is "0" when pay by month and the instance will be valid till the last day of that month.
        :param pulumi.Input[str] peer_region: The correspondent region of dedicated connection, please refer to the region and [availability zone list](https://docs.ucloud.cn/api/summary/regionlist) and [UDPN price list](https://docs.ucloud.cn/network/udpn/udpn_price).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UDPNConnectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        UDPN (UCloud Dedicated Private Network)，you can use Dedicated Private Network to achieve high-speed, stable, secure, and dedicated communications between different data centers. The most frequent scenario is to create network connection of clusters across regions.

        > **VPC Peering Connections with UDPN Connection** The cross-region Dedicated Private Network must be established if the two VPCs located in different regions are expected to be connected.

        > **Note** The additional packet head will be added and included in the overall length of packet due to the tunneling UDPN adopted. Since the number of the bytes of packet head is fixed, the bigger data packet is, the less usage will be taken for the packet head.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ucloud as ucloud

        # connect provider's region (cn-bj2) and peer region (cn-sh2)
        example = ucloud.udpn.UDPNConnection("example",
            bandwidth=2,
            peer_region="cn-sh2")
        ```

        ## Import

        UDPN connection can be imported using the `id`, e.g.

        ```sh
         $ pulumi import ucloud:udpn/uDPNConnection:UDPNConnection example udpn-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param UDPNConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UDPNConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[int]] = None,
                 peer_region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UDPNConnectionArgs.__new__(UDPNConnectionArgs)

            __props__.__dict__["bandwidth"] = bandwidth
            __props__.__dict__["charge_type"] = charge_type
            __props__.__dict__["duration"] = duration
            if peer_region is None and not opts.urn:
                raise TypeError("Missing required property 'peer_region'")
            __props__.__dict__["peer_region"] = peer_region
            __props__.__dict__["create_time"] = None
            __props__.__dict__["expire_time"] = None
        super(UDPNConnection, __self__).__init__(
            'ucloud:udpn/uDPNConnection:UDPNConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bandwidth: Optional[pulumi.Input[int]] = None,
            charge_type: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            duration: Optional[pulumi.Input[int]] = None,
            expire_time: Optional[pulumi.Input[str]] = None,
            peer_region: Optional[pulumi.Input[str]] = None) -> 'UDPNConnection':
        """
        Get an existing UDPNConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). range from 2 - 1000M. The default value is "1"
        :param pulumi.Input[str] charge_type: Charge type. Possible values are: "year" as pay by year, "month" as pay by month, "dynamic" as pay by hour. The default value is "month".
        :param pulumi.Input[str] create_time: The time of creation for UDPN connection, formatted by RFC3339 time string.
        :param pulumi.Input[int] duration: The duration that you will buy the resource, the default value is "1". It is not required when "dynamic" (pay by hour), the value is "0" when pay by month and the instance will be valid till the last day of that month.
        :param pulumi.Input[str] expire_time: The expiration time for UDPN connection, formatted by RFC3339 time string.
        :param pulumi.Input[str] peer_region: The correspondent region of dedicated connection, please refer to the region and [availability zone list](https://docs.ucloud.cn/api/summary/regionlist) and [UDPN price list](https://docs.ucloud.cn/network/udpn/udpn_price).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UDPNConnectionState.__new__(_UDPNConnectionState)

        __props__.__dict__["bandwidth"] = bandwidth
        __props__.__dict__["charge_type"] = charge_type
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["duration"] = duration
        __props__.__dict__["expire_time"] = expire_time
        __props__.__dict__["peer_region"] = peer_region
        return UDPNConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Output[Optional[int]]:
        """
        Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). range from 2 - 1000M. The default value is "1"
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> pulumi.Output[Optional[str]]:
        """
        Charge type. Possible values are: "year" as pay by year, "month" as pay by month, "dynamic" as pay by hour. The default value is "month".
        """
        return pulumi.get(self, "charge_type")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time of creation for UDPN connection, formatted by RFC3339 time string.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Output[Optional[int]]:
        """
        The duration that you will buy the resource, the default value is "1". It is not required when "dynamic" (pay by hour), the value is "0" when pay by month and the instance will be valid till the last day of that month.
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> pulumi.Output[str]:
        """
        The expiration time for UDPN connection, formatted by RFC3339 time string.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter(name="peerRegion")
    def peer_region(self) -> pulumi.Output[str]:
        """
        The correspondent region of dedicated connection, please refer to the region and [availability zone list](https://docs.ucloud.cn/api/summary/regionlist) and [UDPN price list](https://docs.ucloud.cn/network/udpn/udpn_price).
        """
        return pulumi.get(self, "peer_region")

