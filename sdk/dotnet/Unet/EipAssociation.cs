// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ucloud.unet
{
    /// <summary>
    /// Provides an EIP Association resource for associating Elastic IP to UHost Instance, Load Balancer, etc.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/r/eip_association.html.markdown.
    /// </summary>
    public partial class EipAssociation : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of EIP.
        /// </summary>
        [Output("eipId")]
        public Output<string> EipId { get; private set; } = null!;

        /// <summary>
        /// The ID of resource with EIP attached.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// **Deprecated**, attribute `resource_type` is deprecated for optimizing parameters.
        /// </summary>
        [Output("resourceType")]
        public Output<string> ResourceType { get; private set; } = null!;


        /// <summary>
        /// Create a EipAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EipAssociation(string name, EipAssociationArgs args, CustomResourceOptions? options = null)
            : base("ucloud:unet/eipAssociation:EipAssociation", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private EipAssociation(string name, Input<string> id, EipAssociationState? state = null, CustomResourceOptions? options = null)
            : base("ucloud:unet/eipAssociation:EipAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EipAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EipAssociation Get(string name, Input<string> id, EipAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new EipAssociation(name, id, state, options);
        }
    }

    public sealed class EipAssociationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of EIP.
        /// </summary>
        [Input("eipId", required: true)]
        public Input<string> EipId { get; set; } = null!;

        /// <summary>
        /// The ID of resource with EIP attached.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// **Deprecated**, attribute `resource_type` is deprecated for optimizing parameters.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        public EipAssociationArgs()
        {
        }
    }

    public sealed class EipAssociationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of EIP.
        /// </summary>
        [Input("eipId")]
        public Input<string>? EipId { get; set; }

        /// <summary>
        /// The ID of resource with EIP attached.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// **Deprecated**, attribute `resource_type` is deprecated for optimizing parameters.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        public EipAssociationState()
        {
        }
    }
}
