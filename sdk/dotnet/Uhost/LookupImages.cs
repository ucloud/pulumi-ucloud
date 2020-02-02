// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ucloud.uhost
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source providers a list of available image resources according to their availability zone, image ID and other fields.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/d/images.html.markdown.
        /// </summary>
        public static Task<LookupImagesResult> LookupImages(LookupImagesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<LookupImagesResult>("ucloud:uhost/lookupImages:lookupImages", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class LookupImagesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Availability zone where images are located. such as: `cn-bj2-02`. You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist).
        /// </summary>
        [Input("availabilityZone")]
        public string? AvailabilityZone { get; set; }

        /// <summary>
        /// The ID of image.
        /// </summary>
        [Input("imageId")]
        public string? ImageId { get; set; }

        /// <summary>
        /// The type of image. Possible values are: `base` as standard image, `business` as owned by market place, and `custom` as custom-image, all the image types will be retrieved by default.
        /// </summary>
        [Input("imageType")]
        public string? ImageType { get; set; }

        /// <summary>
        /// If more than one result is returned, use the most recent image.
        /// </summary>
        [Input("mostRecent")]
        public bool? MostRecent { get; set; }

        /// <summary>
        /// A regex string to filter resulting images by name. (Such as: `^CentOS 7.[1-2] 64` means CentOS 7.1 of 64-bit operating system or CentOS 7.2 of 64-bit operating system, "^Ubuntu 16.04 64" means Ubuntu 16.04 of 64-bit operating system).
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// The type of OS. Possible values are: `linux` and `windows`, all the OS types will be retrieved by default.
        /// </summary>
        [Input("osType")]
        public string? OsType { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public LookupImagesArgs()
        {
        }
    }

    [OutputType]
    public sealed class LookupImagesResult
    {
        /// <summary>
        /// Availability zone where image is located.
        /// </summary>
        public readonly string? AvailabilityZone;
        public readonly string? ImageId;
        public readonly string? ImageType;
        /// <summary>
        /// It is a nested type which documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.LookupImagesImagesResult> Images;
        public readonly bool? MostRecent;
        public readonly string? NameRegex;
        /// <summary>
        /// The type of OS.
        /// </summary>
        public readonly string? OsType;
        public readonly string? OutputFile;
        /// <summary>
        /// Total number of images that satisfy the condition.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private LookupImagesResult(
            string? availabilityZone,
            string? imageId,
            string? imageType,
            ImmutableArray<Outputs.LookupImagesImagesResult> images,
            bool? mostRecent,
            string? nameRegex,
            string? osType,
            string? outputFile,
            int totalCount,
            string id)
        {
            AvailabilityZone = availabilityZone;
            ImageId = imageId;
            ImageType = imageType;
            Images = images;
            MostRecent = mostRecent;
            NameRegex = nameRegex;
            OsType = osType;
            OutputFile = outputFile;
            TotalCount = totalCount;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class LookupImagesImagesResult
    {
        /// <summary>
        /// Availability zone where images are located. such as: `cn-bj2-02`. You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist).
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// The time of creation for image, formatted in RFC3339 time string.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of image if any.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// To identify if any particular feature belongs to the instance, the value is `NetEnhnced` as I/O enhanced instance for now.
        /// </summary>
        public readonly ImmutableArray<string> Features;
        /// <summary>
        /// The ID of image.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of image.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of OS.
        /// </summary>
        public readonly string OsName;
        /// <summary>
        /// The type of OS. Possible values are: `linux` and `windows`, all the OS types will be retrieved by default.
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// The size of image.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// The status of image. Possible values are `Available`, `Making` and `Unavailable`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The type of image.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private LookupImagesImagesResult(
            string availabilityZone,
            string createTime,
            string description,
            ImmutableArray<string> features,
            string id,
            string name,
            string osName,
            string osType,
            int size,
            string status,
            string type)
        {
            AvailabilityZone = availabilityZone;
            CreateTime = createTime;
            Description = description;
            Features = features;
            Id = id;
            Name = name;
            OsName = osName;
            OsType = osType;
            Size = size;
            Status = status;
            Type = type;
        }
    }
    }
}
