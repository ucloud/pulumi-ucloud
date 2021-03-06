// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ucloud.vpc
{
    /// <summary>
    /// Provides a VIP resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/r/vip.html.markdown.
    /// </summary>
    public partial class Vip : Pulumi.CustomResource
    {
        /// <summary>
        /// The time of creation for VIP, formatted in RFC3339 time string.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The ip address of the VIP.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The remarks of the VIP. (Default: `""`).
        /// </summary>
        [Output("remark")]
        public Output<string> Remark { get; private set; } = null!;

        /// <summary>
        /// The ID of subnet. If defined `vpc_id`, the `subnet_id` is Required. 
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// A tag assigned to VIP, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        /// </summary>
        [Output("tag")]
        public Output<string?> Tag { get; private set; } = null!;

        /// <summary>
        /// The ID of VPC linked to the VIP. 
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Vip resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vip(string name, VipArgs args, CustomResourceOptions? options = null)
            : base("ucloud:vpc/vip:Vip", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Vip(string name, Input<string> id, VipState? state = null, CustomResourceOptions? options = null)
            : base("ucloud:vpc/vip:Vip", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Vip resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vip Get(string name, Input<string> id, VipState? state = null, CustomResourceOptions? options = null)
        {
            return new Vip(name, id, state, options);
        }
    }

    public sealed class VipArgs : Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The remarks of the VIP. (Default: `""`).
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// The ID of subnet. If defined `vpc_id`, the `subnet_id` is Required. 
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        /// <summary>
        /// A tag assigned to VIP, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// The ID of VPC linked to the VIP. 
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public VipArgs()
        {
        }
    }

    public sealed class VipState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time of creation for VIP, formatted in RFC3339 time string.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The ip address of the VIP.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The remarks of the VIP. (Default: `""`).
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// The ID of subnet. If defined `vpc_id`, the `subnet_id` is Required. 
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// A tag assigned to VIP, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// The ID of VPC linked to the VIP. 
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public VipState()
        {
        }
    }
}
