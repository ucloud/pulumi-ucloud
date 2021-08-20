// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ucloud.Unet
{
    public static class GetEIP
    {
        /// <summary>
        /// This data source provides a list of EIP resources (Elastic IP address) according to their EIP ID.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Ucloud = Pulumi.Ucloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Ucloud.Unet.GetEIP.InvokeAsync());
        ///         this.First = example.Apply(example =&gt; example.Eips[0].IpSets[0].Ip);
        ///     }
        /// 
        ///     [Output("first")]
        ///     public Output&lt;string&gt; First { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEIPResult> InvokeAsync(GetEIPArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEIPResult>("ucloud:unet/getEIP:getEIP", args ?? new GetEIPArgs(), options.WithVersion());
    }


    public sealed class GetEIPArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Elastic IP IDs, all the EIPs belong to this region will be retrieved if the ID is `[]`.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter resulting eips by name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetEIPArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEIPResult
    {
        /// <summary>
        /// It is a nested type which documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetEIPEipResult> Eips;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// Total number of Elastic IPs that satisfy the condition.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetEIPResult(
            ImmutableArray<Outputs.GetEIPEipResult> eips,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            string? outputFile,

            int totalCount)
        {
            Eips = eips;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
