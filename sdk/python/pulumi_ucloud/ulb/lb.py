# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['LBArgs', 'LB']

@pulumi.input_type
class LBArgs:
    def __init__(__self__, *,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 internal: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 security_group: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LB resource.
        :param pulumi.Input[str] charge_type: , argument `charge_type` is deprecated for optimizing parameters.
        :param pulumi.Input[bool] internal: Indicate whether the load balancer is intranet mode.(Default: `"false"`)
        :param pulumi.Input[str] remark: The remarks of the load balancer. (Default: `""`).
        :param pulumi.Input[str] security_group: The ID of the associated security group. The security_group only takes effect for ULB instances of request_proxy mode and extranet mode at present.
        :param pulumi.Input[str] subnet_id: The ID of subnet that intranet load balancer belongs to. This argument is not required if default subnet.
        :param pulumi.Input[str] tag: A tag assigned to load balancer, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        :param pulumi.Input[str] vpc_id: The ID of the VPC linked to the Load balancer, This argument is not required if default VPC.
        """
        if charge_type is not None:
            warnings.warn("""attribute `charge_type` is deprecated for optimizing parameters""", DeprecationWarning)
            pulumi.log.warn("""charge_type is deprecated: attribute `charge_type` is deprecated for optimizing parameters""")
        if charge_type is not None:
            pulumi.set(__self__, "charge_type", charge_type)
        if internal is not None:
            pulumi.set(__self__, "internal", internal)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if security_group is not None:
            pulumi.set(__self__, "security_group", security_group)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        , argument `charge_type` is deprecated for optimizing parameters.
        """
        return pulumi.get(self, "charge_type")

    @charge_type.setter
    def charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charge_type", value)

    @property
    @pulumi.getter
    def internal(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate whether the load balancer is intranet mode.(Default: `"false"`)
        """
        return pulumi.get(self, "internal")

    @internal.setter
    def internal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "internal", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        The remarks of the load balancer. (Default: `""`).
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="securityGroup")
    def security_group(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the associated security group. The security_group only takes effect for ULB instances of request_proxy mode and extranet mode at present.
        """
        return pulumi.get(self, "security_group")

    @security_group.setter
    def security_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of subnet that intranet load balancer belongs to. This argument is not required if default subnet.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[str]]:
        """
        A tag assigned to load balancer, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPC linked to the Load balancer, This argument is not required if default VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class _LBState:
    def __init__(__self__, *,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 internal: Optional[pulumi.Input[bool]] = None,
                 ip_sets: Optional[pulumi.Input[Sequence[pulumi.Input['LBIpSetArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_ip: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 security_group: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LB resources.
        :param pulumi.Input[str] charge_type: , argument `charge_type` is deprecated for optimizing parameters.
        :param pulumi.Input[str] create_time: The time of creation for load balancer, formatted in RFC3339 time string.
        :param pulumi.Input[str] expire_time: **Deprecated** attribute `expire_time` is deprecated for optimizing outputs.
        :param pulumi.Input[bool] internal: Indicate whether the load balancer is intranet mode.(Default: `"false"`)
        :param pulumi.Input[Sequence[pulumi.Input['LBIpSetArgs']]] ip_sets: It is a nested type which documented below.
        :param pulumi.Input[str] private_ip: The IP address of intranet IP. It is `""` if `internal` is `false`.
        :param pulumi.Input[str] remark: The remarks of the load balancer. (Default: `""`).
        :param pulumi.Input[str] security_group: The ID of the associated security group. The security_group only takes effect for ULB instances of request_proxy mode and extranet mode at present.
        :param pulumi.Input[str] subnet_id: The ID of subnet that intranet load balancer belongs to. This argument is not required if default subnet.
        :param pulumi.Input[str] tag: A tag assigned to load balancer, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        :param pulumi.Input[str] vpc_id: The ID of the VPC linked to the Load balancer, This argument is not required if default VPC.
        """
        if charge_type is not None:
            warnings.warn("""attribute `charge_type` is deprecated for optimizing parameters""", DeprecationWarning)
            pulumi.log.warn("""charge_type is deprecated: attribute `charge_type` is deprecated for optimizing parameters""")
        if charge_type is not None:
            pulumi.set(__self__, "charge_type", charge_type)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if expire_time is not None:
            warnings.warn("""attribute `expire_time` is deprecated for optimizing outputs""", DeprecationWarning)
            pulumi.log.warn("""expire_time is deprecated: attribute `expire_time` is deprecated for optimizing outputs""")
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if internal is not None:
            pulumi.set(__self__, "internal", internal)
        if ip_sets is not None:
            pulumi.set(__self__, "ip_sets", ip_sets)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if private_ip is not None:
            pulumi.set(__self__, "private_ip", private_ip)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if security_group is not None:
            pulumi.set(__self__, "security_group", security_group)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        , argument `charge_type` is deprecated for optimizing parameters.
        """
        return pulumi.get(self, "charge_type")

    @charge_type.setter
    def charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charge_type", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time of creation for load balancer, formatted in RFC3339 time string.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[pulumi.Input[str]]:
        """
        **Deprecated** attribute `expire_time` is deprecated for optimizing outputs.
        """
        return pulumi.get(self, "expire_time")

    @expire_time.setter
    def expire_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_time", value)

    @property
    @pulumi.getter
    def internal(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate whether the load balancer is intranet mode.(Default: `"false"`)
        """
        return pulumi.get(self, "internal")

    @internal.setter
    def internal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "internal", value)

    @property
    @pulumi.getter(name="ipSets")
    def ip_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LBIpSetArgs']]]]:
        """
        It is a nested type which documented below.
        """
        return pulumi.get(self, "ip_sets")

    @ip_sets.setter
    def ip_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LBIpSetArgs']]]]):
        pulumi.set(self, "ip_sets", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address of intranet IP. It is `""` if `internal` is `false`.
        """
        return pulumi.get(self, "private_ip")

    @private_ip.setter
    def private_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_ip", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        The remarks of the load balancer. (Default: `""`).
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="securityGroup")
    def security_group(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the associated security group. The security_group only takes effect for ULB instances of request_proxy mode and extranet mode at present.
        """
        return pulumi.get(self, "security_group")

    @security_group.setter
    def security_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of subnet that intranet load balancer belongs to. This argument is not required if default subnet.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[str]]:
        """
        A tag assigned to load balancer, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPC linked to the Load balancer, This argument is not required if default VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class LB(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 internal: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 security_group: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Load Balancer resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ucloud as ucloud

        web = ucloud.ulb.LB("web", tag="tf-example")
        ```

        ## Import

        LB can be imported using the `id`, e.g.

        ```sh
         $ pulumi import ucloud:ulb/lB:LB example ulb-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] charge_type: , argument `charge_type` is deprecated for optimizing parameters.
        :param pulumi.Input[bool] internal: Indicate whether the load balancer is intranet mode.(Default: `"false"`)
        :param pulumi.Input[str] remark: The remarks of the load balancer. (Default: `""`).
        :param pulumi.Input[str] security_group: The ID of the associated security group. The security_group only takes effect for ULB instances of request_proxy mode and extranet mode at present.
        :param pulumi.Input[str] subnet_id: The ID of subnet that intranet load balancer belongs to. This argument is not required if default subnet.
        :param pulumi.Input[str] tag: A tag assigned to load balancer, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        :param pulumi.Input[str] vpc_id: The ID of the VPC linked to the Load balancer, This argument is not required if default VPC.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[LBArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Load Balancer resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ucloud as ucloud

        web = ucloud.ulb.LB("web", tag="tf-example")
        ```

        ## Import

        LB can be imported using the `id`, e.g.

        ```sh
         $ pulumi import ucloud:ulb/lB:LB example ulb-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param LBArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LBArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 internal: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 security_group: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = LBArgs.__new__(LBArgs)

            if charge_type is not None and not opts.urn:
                warnings.warn("""attribute `charge_type` is deprecated for optimizing parameters""", DeprecationWarning)
                pulumi.log.warn("""charge_type is deprecated: attribute `charge_type` is deprecated for optimizing parameters""")
            __props__.__dict__["charge_type"] = charge_type
            __props__.__dict__["internal"] = internal
            __props__.__dict__["name"] = name
            __props__.__dict__["remark"] = remark
            __props__.__dict__["security_group"] = security_group
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["tag"] = tag
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["expire_time"] = None
            __props__.__dict__["ip_sets"] = None
            __props__.__dict__["private_ip"] = None
        super(LB, __self__).__init__(
            'ucloud:ulb/lB:LB',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            charge_type: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            expire_time: Optional[pulumi.Input[str]] = None,
            internal: Optional[pulumi.Input[bool]] = None,
            ip_sets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LBIpSetArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            private_ip: Optional[pulumi.Input[str]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            security_group: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tag: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'LB':
        """
        Get an existing LB resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] charge_type: , argument `charge_type` is deprecated for optimizing parameters.
        :param pulumi.Input[str] create_time: The time of creation for load balancer, formatted in RFC3339 time string.
        :param pulumi.Input[str] expire_time: **Deprecated** attribute `expire_time` is deprecated for optimizing outputs.
        :param pulumi.Input[bool] internal: Indicate whether the load balancer is intranet mode.(Default: `"false"`)
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LBIpSetArgs']]]] ip_sets: It is a nested type which documented below.
        :param pulumi.Input[str] private_ip: The IP address of intranet IP. It is `""` if `internal` is `false`.
        :param pulumi.Input[str] remark: The remarks of the load balancer. (Default: `""`).
        :param pulumi.Input[str] security_group: The ID of the associated security group. The security_group only takes effect for ULB instances of request_proxy mode and extranet mode at present.
        :param pulumi.Input[str] subnet_id: The ID of subnet that intranet load balancer belongs to. This argument is not required if default subnet.
        :param pulumi.Input[str] tag: A tag assigned to load balancer, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        :param pulumi.Input[str] vpc_id: The ID of the VPC linked to the Load balancer, This argument is not required if default VPC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LBState.__new__(_LBState)

        __props__.__dict__["charge_type"] = charge_type
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["expire_time"] = expire_time
        __props__.__dict__["internal"] = internal
        __props__.__dict__["ip_sets"] = ip_sets
        __props__.__dict__["name"] = name
        __props__.__dict__["private_ip"] = private_ip
        __props__.__dict__["remark"] = remark
        __props__.__dict__["security_group"] = security_group
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["tag"] = tag
        __props__.__dict__["vpc_id"] = vpc_id
        return LB(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> pulumi.Output[Optional[str]]:
        """
        , argument `charge_type` is deprecated for optimizing parameters.
        """
        return pulumi.get(self, "charge_type")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time of creation for load balancer, formatted in RFC3339 time string.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> pulumi.Output[str]:
        """
        **Deprecated** attribute `expire_time` is deprecated for optimizing outputs.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter
    def internal(self) -> pulumi.Output[bool]:
        """
        Indicate whether the load balancer is intranet mode.(Default: `"false"`)
        """
        return pulumi.get(self, "internal")

    @property
    @pulumi.getter(name="ipSets")
    def ip_sets(self) -> pulumi.Output[Sequence['outputs.LBIpSet']]:
        """
        It is a nested type which documented below.
        """
        return pulumi.get(self, "ip_sets")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> pulumi.Output[str]:
        """
        The IP address of intranet IP. It is `""` if `internal` is `false`.
        """
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[str]:
        """
        The remarks of the load balancer. (Default: `""`).
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="securityGroup")
    def security_group(self) -> pulumi.Output[str]:
        """
        The ID of the associated security group. The security_group only takes effect for ULB instances of request_proxy mode and extranet mode at present.
        """
        return pulumi.get(self, "security_group")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        The ID of subnet that intranet load balancer belongs to. This argument is not required if default subnet.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tag(self) -> pulumi.Output[Optional[str]]:
        """
        A tag assigned to load balancer, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        """
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of the VPC linked to the Load balancer, This argument is not required if default VPC.
        """
        return pulumi.get(self, "vpc_id")

