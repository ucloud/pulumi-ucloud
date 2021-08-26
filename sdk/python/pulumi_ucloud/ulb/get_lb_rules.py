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
    'GetLBRulesResult',
    'AwaitableGetLBRulesResult',
    'get_lb_rules',
]

@pulumi.output_type
class GetLBRulesResult:
    """
    A collection of values returned by getLBRules.
    """
    def __init__(__self__, id=None, ids=None, lb_rules=None, listener_id=None, load_balancer_id=None, output_file=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if lb_rules and not isinstance(lb_rules, list):
            raise TypeError("Expected argument 'lb_rules' to be a list")
        pulumi.set(__self__, "lb_rules", lb_rules)
        if listener_id and not isinstance(listener_id, str):
            raise TypeError("Expected argument 'listener_id' to be a str")
        pulumi.set(__self__, "listener_id", listener_id)
        if load_balancer_id and not isinstance(load_balancer_id, str):
            raise TypeError("Expected argument 'load_balancer_id' to be a str")
        pulumi.set(__self__, "load_balancer_id", load_balancer_id)
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
    @pulumi.getter(name="lbRules")
    def lb_rules(self) -> Sequence['outputs.GetLBRulesLbRuleResult']:
        """
        It is a nested type which documented below.
        """
        return pulumi.get(self, "lb_rules")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> str:
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> str:
        return pulumi.get(self, "load_balancer_id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        Total number of LB Rules that satisfy the condition.
        """
        return pulumi.get(self, "total_count")


class AwaitableGetLBRulesResult(GetLBRulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLBRulesResult(
            id=self.id,
            ids=self.ids,
            lb_rules=self.lb_rules,
            listener_id=self.listener_id,
            load_balancer_id=self.load_balancer_id,
            output_file=self.output_file,
            total_count=self.total_count)


def get_lb_rules(ids: Optional[Sequence[str]] = None,
                 listener_id: Optional[str] = None,
                 load_balancer_id: Optional[str] = None,
                 output_file: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLBRulesResult:
    """
    This data source provides a list of Load Balancer Rule resources according to their Load Balancer Rule ID.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ucloud as ucloud

    example = ucloud.ulb.get_lb_rules(load_balancer_id="ulb-xxx",
        listener_id="vserver-xxx")
    pulumi.export("first", example.lb_rules[0].id)
    ```


    :param Sequence[str] ids: A list of LB Rule IDs, all the LB Rules belong to the Load Balancer listener will be retrieved if the ID is `[]`.
    :param str listener_id: The ID of a listener server.
    :param str load_balancer_id: The ID of a load balancer.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['listenerId'] = listener_id
    __args__['loadBalancerId'] = load_balancer_id
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ucloud:ulb/getLBRules:getLBRules', __args__, opts=opts, typ=GetLBRulesResult).value

    return AwaitableGetLBRulesResult(
        id=__ret__.id,
        ids=__ret__.ids,
        lb_rules=__ret__.lb_rules,
        listener_id=__ret__.listener_id,
        load_balancer_id=__ret__.load_balancer_id,
        output_file=__ret__.output_file,
        total_count=__ret__.total_count)