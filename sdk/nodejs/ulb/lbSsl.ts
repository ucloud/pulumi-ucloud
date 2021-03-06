// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class LbSsl extends pulumi.CustomResource {
    /**
     * Get an existing LbSsl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LbSslState, opts?: pulumi.CustomResourceOptions): LbSsl {
        return new LbSsl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ucloud:ulb/lbSsl:LbSsl';

    /**
     * Returns true if the given object is an instance of LbSsl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LbSsl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LbSsl.__pulumiType;
    }

    public readonly caCert!: pulumi.Output<string | undefined>;
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly privateKey!: pulumi.Output<string>;
    public readonly userCert!: pulumi.Output<string>;

    /**
     * Create a LbSsl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LbSslArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LbSslArgs | LbSslState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LbSslState | undefined;
            inputs["caCert"] = state ? state.caCert : undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["privateKey"] = state ? state.privateKey : undefined;
            inputs["userCert"] = state ? state.userCert : undefined;
        } else {
            const args = argsOrState as LbSslArgs | undefined;
            if (!args || args.privateKey === undefined) {
                throw new Error("Missing required property 'privateKey'");
            }
            if (!args || args.userCert === undefined) {
                throw new Error("Missing required property 'userCert'");
            }
            inputs["caCert"] = args ? args.caCert : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["privateKey"] = args ? args.privateKey : undefined;
            inputs["userCert"] = args ? args.userCert : undefined;
            inputs["createTime"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LbSsl.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LbSsl resources.
 */
export interface LbSslState {
    readonly caCert?: pulumi.Input<string>;
    readonly createTime?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly privateKey?: pulumi.Input<string>;
    readonly userCert?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LbSsl resource.
 */
export interface LbSslArgs {
    readonly caCert?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly privateKey: pulumi.Input<string>;
    readonly userCert: pulumi.Input<string>;
}
