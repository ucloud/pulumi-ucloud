# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 base_url: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 max_retries: Optional[pulumi.Input[int]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 shared_credentials_file: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] base_url: ...
        :param pulumi.Input[bool] insecure: ...
        :param pulumi.Input[int] max_retries: ...
        :param pulumi.Input[str] private_key: ...
        :param pulumi.Input[str] profile: ...
        :param pulumi.Input[str] project_id: ...
        :param pulumi.Input[str] public_key: ...
        :param pulumi.Input[str] region: ...
        :param pulumi.Input[str] shared_credentials_file: ...
        """
        if base_url is not None:
            pulumi.set(__self__, "base_url", base_url)
        if insecure is not None:
            pulumi.set(__self__, "insecure", insecure)
        if max_retries is not None:
            pulumi.set(__self__, "max_retries", max_retries)
        if private_key is None:
            private_key = _utilities.get_env('UCLOUD_PRIVATE_KEY')
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if profile is None:
            profile = _utilities.get_env('UCLOUD_PROFILE')
        if profile is not None:
            pulumi.set(__self__, "profile", profile)
        if project_id is None:
            project_id = _utilities.get_env('UCLOUD_PROJECT_ID')
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if public_key is None:
            public_key = _utilities.get_env('UCLOUD_PUBLIC_KEY')
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)
        if region is None:
            region = _utilities.get_env('UCLOUD_REGION')
        if region is not None:
            pulumi.set(__self__, "region", region)
        if shared_credentials_file is None:
            shared_credentials_file = _utilities.get_env('UCLOUD_SHARED_CREDENTIAL_FILE')
        if shared_credentials_file is not None:
            pulumi.set(__self__, "shared_credentials_file", shared_credentials_file)

    @property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> Optional[pulumi.Input[str]]:
        """
        ...
        """
        return pulumi.get(self, "base_url")

    @base_url.setter
    def base_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base_url", value)

    @property
    @pulumi.getter
    def insecure(self) -> Optional[pulumi.Input[bool]]:
        """
        ...
        """
        return pulumi.get(self, "insecure")

    @insecure.setter
    def insecure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "insecure", value)

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> Optional[pulumi.Input[int]]:
        """
        ...
        """
        return pulumi.get(self, "max_retries")

    @max_retries.setter
    def max_retries(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_retries", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        ...
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter
    def profile(self) -> Optional[pulumi.Input[str]]:
        """
        ...
        """
        return pulumi.get(self, "profile")

    @profile.setter
    def profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "profile", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        ...
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        ...
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        ...
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="sharedCredentialsFile")
    def shared_credentials_file(self) -> Optional[pulumi.Input[str]]:
        """
        ...
        """
        return pulumi.get(self, "shared_credentials_file")

    @shared_credentials_file.setter
    def shared_credentials_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "shared_credentials_file", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base_url: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 max_retries: Optional[pulumi.Input[int]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 shared_credentials_file: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the ucloud package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] base_url: ...
        :param pulumi.Input[bool] insecure: ...
        :param pulumi.Input[int] max_retries: ...
        :param pulumi.Input[str] private_key: ...
        :param pulumi.Input[str] profile: ...
        :param pulumi.Input[str] project_id: ...
        :param pulumi.Input[str] public_key: ...
        :param pulumi.Input[str] region: ...
        :param pulumi.Input[str] shared_credentials_file: ...
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the ucloud package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base_url: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 max_retries: Optional[pulumi.Input[int]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 shared_credentials_file: Optional[pulumi.Input[str]] = None,
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
            __props__ = ProviderArgs.__new__(ProviderArgs)

            __props__.__dict__["base_url"] = base_url
            __props__.__dict__["insecure"] = pulumi.Output.from_input(insecure).apply(pulumi.runtime.to_json) if insecure is not None else None
            __props__.__dict__["max_retries"] = pulumi.Output.from_input(max_retries).apply(pulumi.runtime.to_json) if max_retries is not None else None
            if private_key is None:
                private_key = _utilities.get_env('UCLOUD_PRIVATE_KEY')
            __props__.__dict__["private_key"] = private_key
            if profile is None:
                profile = _utilities.get_env('UCLOUD_PROFILE')
            __props__.__dict__["profile"] = profile
            if project_id is None:
                project_id = _utilities.get_env('UCLOUD_PROJECT_ID')
            __props__.__dict__["project_id"] = project_id
            if public_key is None:
                public_key = _utilities.get_env('UCLOUD_PUBLIC_KEY')
            __props__.__dict__["public_key"] = public_key
            if region is None:
                region = _utilities.get_env('UCLOUD_REGION')
            __props__.__dict__["region"] = region
            if shared_credentials_file is None:
                shared_credentials_file = _utilities.get_env('UCLOUD_SHARED_CREDENTIAL_FILE')
            __props__.__dict__["shared_credentials_file"] = shared_credentials_file
        super(Provider, __self__).__init__(
            'ucloud',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> pulumi.Output[Optional[str]]:
        """
        ...
        """
        return pulumi.get(self, "base_url")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[Optional[str]]:
        """
        ...
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter
    def profile(self) -> pulumi.Output[Optional[str]]:
        """
        ...
        """
        return pulumi.get(self, "profile")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[Optional[str]]:
        """
        ...
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[Optional[str]]:
        """
        ...
        """
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional[str]]:
        """
        ...
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sharedCredentialsFile")
    def shared_credentials_file(self) -> pulumi.Output[Optional[str]]:
        """
        ...
        """
        return pulumi.get(self, "shared_credentials_file")

