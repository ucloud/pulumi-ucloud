// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Eip extends pulumi.CustomResource {
    /**
     * Get an existing Eip resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EipState, opts?: pulumi.CustomResourceOptions): Eip {
        return new Eip(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ucloud:unet/eip:Eip';

    /**
     * Returns true if the given object is an instance of Eip.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Eip {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Eip.__pulumiType;
    }

    public readonly bandwidth!: pulumi.Output<number>;
    public readonly chargeMode!: pulumi.Output<string>;
    public readonly chargeType!: pulumi.Output<string>;
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    public readonly duration!: pulumi.Output<number | undefined>;
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    public readonly internetType!: pulumi.Output<string>;
    public /*out*/ readonly ipSets!: pulumi.Output<outputs.unet.EipIpSet[]>;
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly publicIp!: pulumi.Output<string>;
    public readonly remark!: pulumi.Output<string>;
    public /*out*/ readonly resource!: pulumi.Output<outputs.unet.EipResource>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly tag!: pulumi.Output<string | undefined>;

    /**
     * Create a Eip resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EipArgs | EipState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EipState | undefined;
            inputs["bandwidth"] = state ? state.bandwidth : undefined;
            inputs["chargeMode"] = state ? state.chargeMode : undefined;
            inputs["chargeType"] = state ? state.chargeType : undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["duration"] = state ? state.duration : undefined;
            inputs["expireTime"] = state ? state.expireTime : undefined;
            inputs["internetType"] = state ? state.internetType : undefined;
            inputs["ipSets"] = state ? state.ipSets : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["publicIp"] = state ? state.publicIp : undefined;
            inputs["remark"] = state ? state.remark : undefined;
            inputs["resource"] = state ? state.resource : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tag"] = state ? state.tag : undefined;
        } else {
            const args = argsOrState as EipArgs | undefined;
            if (!args || args.internetType === undefined) {
                throw new Error("Missing required property 'internetType'");
            }
            inputs["bandwidth"] = args ? args.bandwidth : undefined;
            inputs["chargeMode"] = args ? args.chargeMode : undefined;
            inputs["chargeType"] = args ? args.chargeType : undefined;
            inputs["duration"] = args ? args.duration : undefined;
            inputs["internetType"] = args ? args.internetType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["remark"] = args ? args.remark : undefined;
            inputs["tag"] = args ? args.tag : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["expireTime"] = undefined /*out*/;
            inputs["ipSets"] = undefined /*out*/;
            inputs["publicIp"] = undefined /*out*/;
            inputs["resource"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Eip.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Eip resources.
 */
export interface EipState {
    readonly bandwidth?: pulumi.Input<number>;
    readonly chargeMode?: pulumi.Input<string>;
    readonly chargeType?: pulumi.Input<string>;
    readonly createTime?: pulumi.Input<string>;
    readonly duration?: pulumi.Input<number>;
    readonly expireTime?: pulumi.Input<string>;
    readonly internetType?: pulumi.Input<string>;
    readonly ipSets?: pulumi.Input<pulumi.Input<inputs.unet.EipIpSet>[]>;
    readonly name?: pulumi.Input<string>;
    readonly publicIp?: pulumi.Input<string>;
    readonly remark?: pulumi.Input<string>;
    readonly resource?: pulumi.Input<inputs.unet.EipResource>;
    readonly status?: pulumi.Input<string>;
    readonly tag?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Eip resource.
 */
export interface EipArgs {
    readonly bandwidth?: pulumi.Input<number>;
    readonly chargeMode?: pulumi.Input<string>;
    readonly chargeType?: pulumi.Input<string>;
    readonly duration?: pulumi.Input<number>;
    readonly internetType: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly remark?: pulumi.Input<string>;
    readonly tag?: pulumi.Input<string>;
}
