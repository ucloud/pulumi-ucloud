// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ucloud

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source providers a list of projects owned by user according to finance permission and name.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/d/projects.html.markdown.
func LookupProjects(ctx *pulumi.Context, args *LookupProjectsArgs, opts ...pulumi.InvokeOption) (*LookupProjectsResult, error) {
	var rv LookupProjectsResult
	err := ctx.Invoke("ucloud:/lookupProjects:lookupProjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking lookupProjects.
type LookupProjectsArgs struct {
	// To identify if the current account is granted with financial permission.
	IsFinance *bool `pulumi:"isFinance"`
	// A regex string to filter resulting projects by name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by lookupProjects.
type LookupProjectsResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	IsFinance  *bool   `pulumi:"isFinance"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// It is a nested type which documented below.
	Projects []LookupProjectsProject `pulumi:"projects"`
	// Total number of projects that satisfy the condition.
	TotalCount int `pulumi:"totalCount"`
}
