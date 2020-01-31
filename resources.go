// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ucloud

import (
	"strings"
	"unicode"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/resource"
	"github.com/pulumi/pulumi/pkg/tokens"
	"github.com/terraform-providers/terraform-provider-ucloud/ucloud"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "ucloud"
	// modules:
	mainMod     = "index" // the y module
	RootRes     = ""
	UHostRes    = "uhost"
	UNetRes     = "unet"
	ULBRes      = "ulb"
	VpcRes      = "vpc"
	UDBRes      = "udb"
	UMemRes     = "umem"
	IpSecVPNRes = "ipsecvpn"
	UDPNRes     = "udpn"
)

var dsModMaps = map[string]string{
	"ucloud_projects":              RootRes,
	"ucloud_zones":                 RootRes,
	"ucloud_images":                UHostRes,
	"ucloud_instances":             UHostRes,
	"ucloud_disks":                 UHostRes,
	"ucloud_eips":                  UNetRes,
	"ucloud_security_groups":       UNetRes,
	"ucloud_lb_ssls":               ULBRes,
	"ucloud_lbs":                   ULBRes,
	"ucloud_lb_listeners":          ULBRes,
	"ucloud_lb_rules":              ULBRes,
	"ucloud_lb_attachments":        ULBRes,
	"ucloud_db_instances":          UDBRes,
	"ucloud_subnets":               VpcRes,
	"ucloud_vpcs":                  VpcRes,
	"ucloud_nat_gateways":          VpcRes,
	"ucloud_vpn_gateways":          IpSecVPNRes,
	"ucloud_vpn_customer_gateways": IpSecVPNRes,
	"ucloud_vpn_connections":       IpSecVPNRes,
}

var resModMaps = map[string]string{
	"ucloud_instance":               UHostRes,
	"ucloud_disk":                   UHostRes,
	"ucloud_disk_attachment":        UHostRes,
	"ucloud_isolation_group":        UHostRes,
	"ucloud_security_group":         UNetRes,
	"ucloud_eip":                    UNetRes,
	"ucloud_eip_association":        UNetRes,
	"ucloud_lb":                     ULBRes,
	"ucloud_lb_attachment":          ULBRes,
	"ucloud_lb_listener":            ULBRes,
	"ucloud_lb_rule":                ULBRes,
	"ucloud_lb_ssl":                 ULBRes,
	"ucloud_lb_ssl_attachment":      ULBRes,
	"ucloud_vpc":                    VpcRes,
	"ucloud_vpc_peering_connection": VpcRes,
	"ucloud_subnet":                 VpcRes,
	"ucloud_nat_gateway":            VpcRes,
	"ucloud_nat_gateway_rule":       VpcRes,
	"ucloud_vip":                    VpcRes,
	"ucloud_db_instance":            UDBRes,
	"ucloud_redis_instance":         UMemRes,
	"ucloud_memcache_instance":      UMemRes,
	"ucloud_vpn_gateway":            IpSecVPNRes,
	"ucloud_vpn_customer_gateway":   IpSecVPNRes,
	"ucloud_vpn_connection":         IpSecVPNRes,
	"ucloud_udpn_connection":        UDPNRes,
}

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// boolRef returns a reference to the bool argument.
func boolRef(b bool) *bool {
	return &b
}

// stringValue gets a string value from a property map if present, else ""
func stringValue(vars resource.PropertyMap, prop resource.PropertyKey) string {
	val, ok := vars[prop]
	if ok && val.IsString() {
		return val.StringValue()
	}
	return ""
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	return nil
}

// managedByPulumi is a default used for some managed resources, in the absence of something more meaningful.
var managedByPulumi = &tfbridge.DefaultInfo{Value: "Managed by Pulumi"}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := ucloud.Provider().(*schema.Provider)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "ucloud",
		Description: "A Pulumi package for creating and managing ucloud cloud resources.",
		Keywords:    []string{"pulumi", "ucloud"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-ucloud",
		Config: map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			"region": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"UCLOUD_REGION", "UCLOUD_DEFAULT_REGION"},
				},
			},
			"public_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"UCLOUD_PUBLIC_KEY", "UCloud Public Key"},
				},
			},
			"private_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"UCLOUD_PRIVATE_KEY", "UCloud Private Key"},
				},
			},
			"project_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"UCLOUD_PROJECT_ID", "UCloud Project Id"},
				},
			},
			"base_url": {
				// Type: makeType("base_url", "BaseUrl"),
			},
			"max_retries": {},
			"insecure":    {},
			"profile": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"UCLOUD_PROFILE", "UCloud Profile Name"},
				},
			},
			"shared_credentials_file": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"UCLOUD_SHARED_CREDENTIAL_FILE", "Path To The Shared Credentials File"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources:            make(map[string]*tfbridge.ResourceInfo, len(p.ResourcesMap)),
		DataSources:          map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: makeDataSource(mainMod, "getAmi")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "latest",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=1.0.0,<2.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{},
	}

	for name, _ := range p.ResourcesMap {
		resStructName := ToStructName(name)
		prov.Resources[name] = &tfbridge.ResourceInfo{
			Tok: makeResource(resModMaps[name], resStructName),
		}
	}

	for name, _ := range p.DataSourcesMap {
		lookupFuncName := "lookup" + ToStructName(name)
		prov.DataSources[name] = &tfbridge.DataSourceInfo{
			Tok: makeDataSource(dsModMaps[name], lookupFuncName),
		}
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resName, res := range prov.Resources {
		if schema := p.ResourcesMap[resName]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}

func ToStructName(resName string) string {
	result := ""
	trimedResName := strings.TrimPrefix(resName, "ucloud_")
	for _, part := range strings.Split(trimedResName, "_") {
		result += strings.Title(part)
	}
	return result
}
