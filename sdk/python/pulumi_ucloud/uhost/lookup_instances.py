# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class LookupInstancesResult:
    """
    A collection of values returned by lookupInstances.
    """
    def __init__(__self__, availability_zone=None, ids=None, instances=None, name_regex=None, output_file=None, tag=None, total_count=None, id=None):
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        __self__.availability_zone = availability_zone
        """
        Availability zone where instances are located.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        __self__.instances = instances
        """
        It is a nested type. instances documented below.
        """
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
        A tag assigned to the instance.
        """
        if total_count and not isinstance(total_count, float):
            raise TypeError("Expected argument 'total_count' to be a float")
        __self__.total_count = total_count
        """
        Total number of instances that satisfy the condition.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableLookupInstancesResult(LookupInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return LookupInstancesResult(
            availability_zone=self.availability_zone,
            ids=self.ids,
            instances=self.instances,
            name_regex=self.name_regex,
            output_file=self.output_file,
            tag=self.tag,
            total_count=self.total_count,
            id=self.id)

def lookup_instances(availability_zone=None,ids=None,name_regex=None,output_file=None,tag=None,opts=None):
    """
    This data source providers a list of UHost instance resources according to their availability zone, instance ID and tag.
    
    :param str availability_zone: Availability zone where instances are located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
    :param list ids: A list of instance IDs, all the instances belongs to the defined region will be retrieved if this argument is "".
    :param str name_regex: A regex string to filter resulting instances by name.
    :param str tag: A tag assigned to instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).

    > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/d/instances.html.markdown.
    """
    __args__ = dict()

    __args__['availabilityZone'] = availability_zone
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['tag'] = tag
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ucloud:uhost/lookupInstances:lookupInstances', __args__, opts=opts).value

    return AwaitableLookupInstancesResult(
        availability_zone=__ret__.get('availabilityZone'),
        ids=__ret__.get('ids'),
        instances=__ret__.get('instances'),
        name_regex=__ret__.get('nameRegex'),
        output_file=__ret__.get('outputFile'),
        tag=__ret__.get('tag'),
        total_count=__ret__.get('totalCount'),
        id=__ret__.get('id'))