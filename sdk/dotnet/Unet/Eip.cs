// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ucloud.unet
{
    /// <summary>
    /// Provides an Elastic IP resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/r/eip.html.markdown.
    /// </summary>
    public partial class Eip : Pulumi.CustomResource
    {
        /// <summary>
        /// Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). The ranges for bandwidth are: 1-200 for pay by traffic, 1-800 for pay by bandwidth. (Default: `1`).
        /// </summary>
        [Output("bandwidth")]
        public Output<int> Bandwidth { get; private set; } = null!;

        [Output("chargeMode")]
        public Output<string> ChargeMode { get; private set; } = null!;

        /// <summary>
        /// Elastic IP charge type. Possible values are: `year` as pay by year, `month` as pay by month, `dynamic` as pay by hour (specific permission required). (Default: `month`).
        /// </summary>
        [Output("chargeType")]
        public Output<string> ChargeType { get; private set; } = null!;

        /// <summary>
        /// The time of creation for EIP, formatted in RFC3339 time string.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The duration that you will buy the resource. (Default: `1`). It is not required when `dynamic` (pay by hour), the value is `0` when `month`(pay by month) and the instance will be valid till the last day of that month.
        /// * `charge_mode` -(Optional) Elastic IP charge mode. Possible values are: `traffic` as pay by traffic, `bandwidth` as pay by bandwidth, `share_bandwidth` as share bandwidth mode. (Default: `bandwidth`for the Elastic IP, `share_bandwidth` for the Elastic IP with share bandwidth mode).
        /// </summary>
        [Output("duration")]
        public Output<int?> Duration { get; private set; } = null!;

        /// <summary>
        /// The expiration time for EIP, formatted in RFC3339 time string.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// Type of Elastic IP routes. Possible values are: `international` as international BGP IP and `bgp` as china mainland BGP IP.
        /// </summary>
        [Output("internetType")]
        public Output<string> InternetType { get; private set; } = null!;

        /// <summary>
        /// It is a nested type which documented below.
        /// </summary>
        [Output("ipSets")]
        public Output<ImmutableArray<Outputs.EipIpSets>> IpSets { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Public IP address of Elastic IP.
        /// </summary>
        [Output("publicIp")]
        public Output<string> PublicIp { get; private set; } = null!;

        /// <summary>
        /// The remarks of the EIP. (Default: `""`).
        /// </summary>
        [Output("remark")]
        public Output<string> Remark { get; private set; } = null!;

        /// <summary>
        /// It is a nested type which documented below.
        /// </summary>
        [Output("resource")]
        public Output<Outputs.EipResource> Resource { get; private set; } = null!;

        /// <summary>
        /// EIP status. Possible values are: `used` as in use, `free` as available and `freeze` as associating.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A tag assigned to Elastic IP, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        /// </summary>
        [Output("tag")]
        public Output<string?> Tag { get; private set; } = null!;


        /// <summary>
        /// Create a Eip resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Eip(string name, EipArgs args, CustomResourceOptions? options = null)
            : base("ucloud:unet/eip:Eip", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Eip(string name, Input<string> id, EipState? state = null, CustomResourceOptions? options = null)
            : base("ucloud:unet/eip:Eip", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Eip resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Eip Get(string name, Input<string> id, EipState? state = null, CustomResourceOptions? options = null)
        {
            return new Eip(name, id, state, options);
        }
    }

    public sealed class EipArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). The ranges for bandwidth are: 1-200 for pay by traffic, 1-800 for pay by bandwidth. (Default: `1`).
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        [Input("chargeMode")]
        public Input<string>? ChargeMode { get; set; }

        /// <summary>
        /// Elastic IP charge type. Possible values are: `year` as pay by year, `month` as pay by month, `dynamic` as pay by hour (specific permission required). (Default: `month`).
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// The duration that you will buy the resource. (Default: `1`). It is not required when `dynamic` (pay by hour), the value is `0` when `month`(pay by month) and the instance will be valid till the last day of that month.
        /// * `charge_mode` -(Optional) Elastic IP charge mode. Possible values are: `traffic` as pay by traffic, `bandwidth` as pay by bandwidth, `share_bandwidth` as share bandwidth mode. (Default: `bandwidth`for the Elastic IP, `share_bandwidth` for the Elastic IP with share bandwidth mode).
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// Type of Elastic IP routes. Possible values are: `international` as international BGP IP and `bgp` as china mainland BGP IP.
        /// </summary>
        [Input("internetType", required: true)]
        public Input<string> InternetType { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The remarks of the EIP. (Default: `""`).
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// A tag assigned to Elastic IP, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        public EipArgs()
        {
        }
    }

    public sealed class EipState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). The ranges for bandwidth are: 1-200 for pay by traffic, 1-800 for pay by bandwidth. (Default: `1`).
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        [Input("chargeMode")]
        public Input<string>? ChargeMode { get; set; }

        /// <summary>
        /// Elastic IP charge type. Possible values are: `year` as pay by year, `month` as pay by month, `dynamic` as pay by hour (specific permission required). (Default: `month`).
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// The time of creation for EIP, formatted in RFC3339 time string.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The duration that you will buy the resource. (Default: `1`). It is not required when `dynamic` (pay by hour), the value is `0` when `month`(pay by month) and the instance will be valid till the last day of that month.
        /// * `charge_mode` -(Optional) Elastic IP charge mode. Possible values are: `traffic` as pay by traffic, `bandwidth` as pay by bandwidth, `share_bandwidth` as share bandwidth mode. (Default: `bandwidth`for the Elastic IP, `share_bandwidth` for the Elastic IP with share bandwidth mode).
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// The expiration time for EIP, formatted in RFC3339 time string.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        /// <summary>
        /// Type of Elastic IP routes. Possible values are: `international` as international BGP IP and `bgp` as china mainland BGP IP.
        /// </summary>
        [Input("internetType")]
        public Input<string>? InternetType { get; set; }

        [Input("ipSets")]
        private InputList<Inputs.EipIpSetsGetArgs>? _ipSets;

        /// <summary>
        /// It is a nested type which documented below.
        /// </summary>
        public InputList<Inputs.EipIpSetsGetArgs> IpSets
        {
            get => _ipSets ?? (_ipSets = new InputList<Inputs.EipIpSetsGetArgs>());
            set => _ipSets = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Public IP address of Elastic IP.
        /// </summary>
        [Input("publicIp")]
        public Input<string>? PublicIp { get; set; }

        /// <summary>
        /// The remarks of the EIP. (Default: `""`).
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// It is a nested type which documented below.
        /// </summary>
        [Input("resource")]
        public Input<Inputs.EipResourceGetArgs>? Resource { get; set; }

        /// <summary>
        /// EIP status. Possible values are: `used` as in use, `free` as available and `freeze` as associating.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// A tag assigned to Elastic IP, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        public EipState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class EipIpSetsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of Elastic IP routes. Possible values are: `international` as international BGP IP and `bgp` as china mainland BGP IP.
        /// </summary>
        [Input("internetType")]
        public Input<string>? InternetType { get; set; }

        [Input("ip")]
        public Input<string>? Ip { get; set; }

        public EipIpSetsGetArgs()
        {
        }
    }

    public sealed class EipResourceGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the resource with EIP attached.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The type of resource with EIP attached. Possible values are `instance` as instance, `lb` as load balancer.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public EipResourceGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class EipIpSets
    {
        /// <summary>
        /// Type of Elastic IP routes. Possible values are: `international` as international BGP IP and `bgp` as china mainland BGP IP.
        /// </summary>
        public readonly string InternetType;
        public readonly string Ip;

        [OutputConstructor]
        private EipIpSets(
            string internetType,
            string ip)
        {
            InternetType = internetType;
            Ip = ip;
        }
    }

    [OutputType]
    public sealed class EipResource
    {
        /// <summary>
        /// The ID of the resource with EIP attached.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The type of resource with EIP attached. Possible values are `instance` as instance, `lb` as load balancer.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private EipResource(
            string id,
            string type)
        {
            Id = id;
            Type = type;
        }
    }
    }
}
