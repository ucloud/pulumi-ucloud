// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package udisk

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-ucloud/sdk/go/ucloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "ucloud:udisk/disk:Disk":
		r = &Disk{}
	case "ucloud:udisk/diskAttachment:DiskAttachment":
		r = &DiskAttachment{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := ucloud.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"ucloud",
		"udisk/disk",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ucloud",
		"udisk/diskAttachment",
		&module{version},
	)
}
