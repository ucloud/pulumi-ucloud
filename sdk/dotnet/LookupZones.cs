// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ucloud
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides a list of available zones in the current region.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/d/zones.html.markdown.
        /// </summary>
        public static Task<LookupZonesResult> LookupZones(LookupZonesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<LookupZonesResult>("ucloud:/lookupZones:lookupZones", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class LookupZonesArgs : Pulumi.InvokeArgs
    {
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public LookupZonesArgs()
        {
        }
    }

    [OutputType]
    public sealed class LookupZonesResult
    {
        public readonly string? OutputFile;
        /// <summary>
        /// Total number of zones that satisfy the condition.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// It is a nested type which documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.LookupZonesZonesResult> Zones;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private LookupZonesResult(
            string? outputFile,
            int totalCount,
            ImmutableArray<Outputs.LookupZonesZonesResult> zones,
            string id)
        {
            OutputFile = outputFile;
            TotalCount = totalCount;
            Zones = zones;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class LookupZonesZonesResult
    {
        /// <summary>
        /// The ID of availability zone.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private LookupZonesZonesResult(string id)
        {
            Id = id;
        }
    }
    }
}
