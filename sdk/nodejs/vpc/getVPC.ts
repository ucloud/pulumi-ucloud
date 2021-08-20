// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides a list of VPC resources according to their VPC ID, name.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ucloud from "@pulumi/ucloud";
 *
 * const example = ucloud.vpc.getVPC({});
 * export const first = example.then(example => example.vpcs[0].id);
 * ```
 */
export function getVPC(args?: GetVPCArgs, opts?: pulumi.InvokeOptions): Promise<GetVPCResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ucloud:vpc/getVPC:getVPC", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "tag": args.tag,
    }, opts);
}

/**
 * A collection of arguments for invoking getVPC.
 */
export interface GetVPCArgs {
    /**
     * A list of VPC IDs, all the VPC resources belong to this region will be retrieved if the ID is `[]`.
     */
    ids?: string[];
    /**
     * A regex string to filter resulting VPC resources by name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * A tag assigned to VPC.
     */
    tag?: string;
}

/**
 * A collection of values returned by getVPC.
 */
export interface GetVPCResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * A tag assigned to VPC.
     */
    readonly tag?: string;
    /**
     * Total number of VPC resources that satisfy the condition.
     */
    readonly totalCount: number;
    /**
     * It is a nested type which documented below.
     */
    readonly vpcs: outputs.vpc.GetVPCVpc[];
}
