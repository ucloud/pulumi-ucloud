// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function lookupLbRules(args: LookupLbRulesArgs, opts?: pulumi.InvokeOptions): Promise<LookupLbRulesResult> & LookupLbRulesResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<LookupLbRulesResult> = pulumi.runtime.invoke("ucloud:ulb/lookupLbRules:lookupLbRules", {
        "ids": args.ids,
        "listenerId": args.listenerId,
        "loadBalancerId": args.loadBalancerId,
        "outputFile": args.outputFile,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking lookupLbRules.
 */
export interface LookupLbRulesArgs {
    readonly ids?: string[];
    readonly listenerId: string;
    readonly loadBalancerId: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by lookupLbRules.
 */
export interface LookupLbRulesResult {
    readonly ids?: string[];
    readonly lbRules: outputs.ulb.LookupLbRulesLbRule[];
    readonly listenerId: string;
    readonly loadBalancerId: string;
    readonly outputFile?: string;
    readonly totalCount: number;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
