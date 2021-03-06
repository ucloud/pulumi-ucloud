// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ulb

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides a list of Load Balancer Attachment resources according to their Load Balancer Attachment ID.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/d/lb_attachments.html.markdown.
func LookupLbAttachments(ctx *pulumi.Context, args *LookupLbAttachmentsArgs, opts ...pulumi.InvokeOption) (*LookupLbAttachmentsResult, error) {
	var rv LookupLbAttachmentsResult
	err := ctx.Invoke("ucloud:ulb/lookupLbAttachments:lookupLbAttachments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking lookupLbAttachments.
type LookupLbAttachmentsArgs struct {
	// A list of LB Attachment IDs, all the LB Attachments belong to the Load Balancer listener will be retrieved if the ID is `""`.
	Ids []string `pulumi:"ids"`
	// The ID of a listener server.
	ListenerId string `pulumi:"listenerId"`
	// The ID of a load balancer.
	LoadBalancerId string  `pulumi:"loadBalancerId"`
	OutputFile     *string `pulumi:"outputFile"`
}

// A collection of values returned by lookupLbAttachments.
type LookupLbAttachmentsResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// It is a nested type which documented below.
	LbAttachments  []LookupLbAttachmentsLbAttachment `pulumi:"lbAttachments"`
	ListenerId     string                            `pulumi:"listenerId"`
	LoadBalancerId string                            `pulumi:"loadBalancerId"`
	OutputFile     *string                           `pulumi:"outputFile"`
	// Total number of LB Attachments that satisfy the condition.
	TotalCount int `pulumi:"totalCount"`
}
