// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ucloud.uhost
{
    /// <summary>
    /// Provides a Cloud Disk Attachment resource for attaching Cloud Disk to UHost Instance.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/r/disk_attachment.html.markdown.
    /// </summary>
    public partial class DiskAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The Zone to attach the disk in.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// The ID of disk that needs to be attached
        /// </summary>
        [Output("diskId")]
        public Output<string> DiskId { get; private set; } = null!;

        /// <summary>
        /// The ID of host instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a DiskAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DiskAttachment(string name, DiskAttachmentArgs args, CustomResourceOptions? options = null)
            : base("ucloud:uhost/diskAttachment:DiskAttachment", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private DiskAttachment(string name, Input<string> id, DiskAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("ucloud:uhost/diskAttachment:DiskAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DiskAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DiskAttachment Get(string name, Input<string> id, DiskAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new DiskAttachment(name, id, state, options);
        }
    }

    public sealed class DiskAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Zone to attach the disk in.
        /// </summary>
        [Input("availabilityZone", required: true)]
        public Input<string> AvailabilityZone { get; set; } = null!;

        /// <summary>
        /// The ID of disk that needs to be attached
        /// </summary>
        [Input("diskId", required: true)]
        public Input<string> DiskId { get; set; } = null!;

        /// <summary>
        /// The ID of host instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public DiskAttachmentArgs()
        {
        }
    }

    public sealed class DiskAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Zone to attach the disk in.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The ID of disk that needs to be attached
        /// </summary>
        [Input("diskId")]
        public Input<string>? DiskId { get; set; }

        /// <summary>
        /// The ID of host instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public DiskAttachmentState()
        {
        }
    }
}