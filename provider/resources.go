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
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v1"
	"github.com/pulumi/pulumi-ucloud/provider/pkg/version"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/terraform-providers/terraform-provider-ucloud/ucloud"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "ucloud"

	ucloudUHost    = "uhost"
	ucloudUNet     = "unet"
	ucloudVPC      = "vpc"
	ucloudUDPN     = "udpn"
	ucloudULB      = "ulb"
	ucloudUDisk    = "udisk"
	ucloudUDB      = "udb"
	ucloudUMem     = "umem"
	ucloudIPSecVPN = "ipsecvpn"
	ucloudUFS      = "ufs"
	ucloudUS3      = "us3"
	ucloudCube     = "cube"
	ucloudUK8S     = "uk8s"
	ucloudUAccount = "uaccount"
)

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
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// managedByPulumi is a default used for some managed resources, in the absence of something more meaningful.
var managedByPulumi = &tfbridge.DefaultInfo{Value: "Managed by Pulumi"}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv1.NewProvider(ucloud.Provider().(*schema.Provider))

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
			"region": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"UCLOUD_REGION"},
				},
			},
			"public_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"UCLOUD_PUBLIC_KEY"},
				},
			},
			"private_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"UCLOUD_PRIVATE_KEY"},
				},
			},
			"project_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"UCLOUD_PROJECT_ID"},
				},
			},
			"profile": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"UCLOUD_PROFILE"},
				},
			},
			"base_url":    {},
			"max_retries": {},
			"insecure":    {},
			"shared_credentials_file": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"UCLOUD_SHARED_CREDENTIAL_FILE"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"ucloud_instance":               {Tok: makeResource(ucloudUHost, "Instance")},
			"ucloud_eip":                    {Tok: makeResource(ucloudUNet, "EIP")},
			"ucloud_eip_association":        {Tok: makeResource(ucloudUNet, "EIPAssociation")},
			"ucloud_vpc":                    {Tok: makeResource(ucloudVPC, "VPC")},
			"ucloud_subnet":                 {Tok: makeResource(ucloudVPC, "Subnet")},
			"ucloud_vpc_peering_connection": {Tok: makeResource(ucloudVPC, "VPCPeeringConnection")},
			"ucloud_udpn_connection":        {Tok: makeResource(ucloudUDPN, "UDPNConnection")},
			"ucloud_lb":                     {Tok: makeResource(ucloudULB, "LB")},
			"ucloud_lb_attachment":          {Tok: makeResource(ucloudULB, "LBAttachment")},
			"ucloud_lb_listener":            {Tok: makeResource(ucloudULB, "LBListener")},
			"ucloud_lb_rule":                {Tok: makeResource(ucloudULB, "LBRule")},
			"ucloud_lb_ssl":                 {Tok: makeResource(ucloudULB, "LBSsl")},
			"ucloud_lb_ssl_attachment":      {Tok: makeResource(ucloudULB, "LBSslAttachment")},
			"ucloud_disk":                   {Tok: makeResource(ucloudUDisk, "Disk")},
			"ucloud_disk_attachment":        {Tok: makeResource(ucloudUDisk, "DiskAttachment")},
			"ucloud_security_group":         {Tok: makeResource(ucloudUNet, "SecurityGroup")},
			"ucloud_isolation_group":        {Tok: makeResource(ucloudUHost, "IsolationGroup")},
			"ucloud_db_instance":            {Tok: makeResource(ucloudUDB, "DBInstance")},
			"ucloud_redis_instance":         {Tok: makeResource(ucloudUMem, "RedisInstance")},
			"ucloud_memcache_instance":      {Tok: makeResource(ucloudUMem, "MemcachedInstance")},
			"ucloud_nat_gateway":            {Tok: makeResource(ucloudVPC, "NATGateway")},
			"ucloud_nat_gateway_rule":       {Tok: makeResource(ucloudVPC, "NATGatewayRule")},
			"ucloud_vip":                    {Tok: makeResource(ucloudVPC, "VIP")},
			"ucloud_vpn_gateway":            {Tok: makeResource(ucloudIPSecVPN, "VPNGateway")},
			"ucloud_vpn_customer_gateway":   {Tok: makeResource(ucloudIPSecVPN, "VPNCustomerGateway")},
			"ucloud_vpn_connection":         {Tok: makeResource(ucloudIPSecVPN, "VPNConnection")},
			"ucloud_ufs_volume":             {Tok: makeResource(ucloudUFS, "UFSVolume")},
			"ucloud_us3_bucket":             {Tok: makeResource(ucloudUS3, "US3Bucket")},
			"ucloud_cube_pod":               {Tok: makeResource(ucloudCube, "CubePod")},
			"ucloud_uk8s_cluster":           {Tok: makeResource(ucloudUK8S, "UK8SCluster")},
			"ucloud_uk8s_node":              {Tok: makeResource(ucloudUK8S, "UK8SNode")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"ucloud_projects":              {Tok: makeDataSource(ucloudUAccount, "getProject")},
			"ucloud_zones":                 {Tok: makeDataSource(ucloudUAccount, "getZone")},
			"ucloud_images":                {Tok: makeDataSource(ucloudUHost, "getImage")},
			"ucloud_instances":             {Tok: makeDataSource(ucloudUHost, "getInstance")},
			"ucloud_eips":                  {Tok: makeDataSource(ucloudUNet, "getEIP")},
			"ucloud_security_groups":       {Tok: makeDataSource(ucloudUNet, "getSecurityGroup")},
			"ucloud_lbs":                   {Tok: makeDataSource(ucloudULB, "getLB")},
			"ucloud_lb_listeners":          {Tok: makeDataSource(ucloudULB, "getLBListeners")},
			"ucloud_lb_rules":              {Tok: makeDataSource(ucloudULB, "getLBRules")},
			"ucloud_lb_attachments":        {Tok: makeDataSource(ucloudULB, "getLBAttachment")},
			"ucloud_lb_ssls":               {Tok: makeDataSource(ucloudULB, "getLBSsl")},
			"ucloud_disks":                 {Tok: makeDataSource(ucloudUDisk, "getDisk")},
			"ucloud_db_instances":          {Tok: makeDataSource(ucloudUDB, "getDBInstance")},
			"ucloud_subnets":               {Tok: makeDataSource(ucloudVPC, "getSubnet")},
			"ucloud_vpcs":                  {Tok: makeDataSource(ucloudVPC, "getVPC")},
			"ucloud_nat_gateways":          {Tok: makeDataSource(ucloudVPC, "getNATGateway")},
			"ucloud_vpn_gateways":          {Tok: makeDataSource(ucloudIPSecVPN, "getVPNGateway")},
			"ucloud_vpn_customer_gateways": {Tok: makeDataSource(ucloudIPSecVPN, "getVPNCustomerGateway")},
			"ucloud_vpn_connections":       {Tok: makeDataSource(ucloudIPSecVPN, "getVPNConnection")},
			"ucloud_ufs_volumes":           {Tok: makeDataSource(ucloudUFS, "getUFSVolume")},
			"ucloud_us3_buckets":           {Tok: makeDataSource(ucloudUS3, "getUS3Bucket")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
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
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "3.*",
				"System.Collections.Immutable": "1.6.0",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
