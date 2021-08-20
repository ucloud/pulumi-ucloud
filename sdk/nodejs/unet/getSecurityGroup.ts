// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides a list of Security Group resources according to their Security Group ID, name and resource id.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ucloud from "@pulumi/ucloud";
 *
 * const example = ucloud.unet.getSecurityGroup({});
 * export const first = example.then(example => example.securityGroups[0].id);
 * ```
 */
export function getSecurityGroup(args?: GetSecurityGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityGroupResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ucloud:unet/getSecurityGroup:getSecurityGroup", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecurityGroup.
 */
export interface GetSecurityGroupArgs {
    /**
     * A list of Security Group IDs, all the Security Group resources belong to this region will be retrieved if the ID is `[]`.
     */
    ids?: string[];
    /**
     * A regex string to filter resulting Security Group resources by name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The type of Security Group. Possible values are: `recommendWeb` as the default Web security group that UCloud recommend to users, default opened port include 80, 443, 22, 3389, `recommendNonWeb` as the default non Web security group that UCloud recommend to users, default opened port include 22, 3389, `userDefined` as the security groups defined by users. You may refer to [security group](https://docs.ucloud.cn/network/firewall/firewall.html).
     */
    type?: string;
}

/**
 * A collection of values returned by getSecurityGroup.
 */
export interface GetSecurityGroupResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * It is a nested type which documented below.
     */
    readonly securityGroups: outputs.unet.GetSecurityGroupSecurityGroup[];
    /**
     * Total number of Security Group resources that satisfy the condition.
     */
    readonly totalCount: number;
    /**
     * The type of Security Group.
     */
    readonly type?: string;
}
