// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Load Balancer resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ucloud from "@pulumi/ucloud";
 *
 * const web = new ucloud.ulb.LB("web", {
 *     tag: "tf-example",
 * });
 * ```
 *
 * ## Import
 *
 * LB can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import ucloud:ulb/lB:LB example ulb-abc123456
 * ```
 */
export class LB extends pulumi.CustomResource {
    /**
     * Get an existing LB resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LBState, opts?: pulumi.CustomResourceOptions): LB {
        return new LB(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ucloud:ulb/lB:LB';

    /**
     * Returns true if the given object is an instance of LB.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LB {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LB.__pulumiType;
    }

    /**
     * , argument `chargeType` is deprecated for optimizing parameters.
     *
     * @deprecated attribute `charge_type` is deprecated for optimizing parameters
     */
    public readonly chargeType!: pulumi.Output<string | undefined>;
    /**
     * The time of creation for load balancer, formatted in RFC3339 time string.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * **Deprecated** attribute `expireTime` is deprecated for optimizing outputs.
     *
     * @deprecated attribute `expire_time` is deprecated for optimizing outputs
     */
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    /**
     * Indicate whether the load balancer is intranet mode.(Default: `"false"`)
     */
    public readonly internal!: pulumi.Output<boolean>;
    /**
     * It is a nested type which documented below.
     */
    public /*out*/ readonly ipSets!: pulumi.Output<outputs.ulb.LBIpSet[]>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The IP address of intranet IP. It is `""` if `internal` is `false`.
     */
    public /*out*/ readonly privateIp!: pulumi.Output<string>;
    /**
     * The remarks of the load balancer. (Default: `""`).
     */
    public readonly remark!: pulumi.Output<string>;
    /**
     * The ID of the associated security group. The securityGroup only takes effect for ULB instances of requestProxy mode and extranet mode at present.
     */
    public readonly securityGroup!: pulumi.Output<string>;
    /**
     * The ID of subnet that intranet load balancer belongs to. This argument is not required if default subnet.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * A tag assigned to load balancer, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
     */
    public readonly tag!: pulumi.Output<string | undefined>;
    /**
     * The ID of the VPC linked to the Load balancer, This argument is not required if default VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a LB resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LBArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LBArgs | LBState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LBState | undefined;
            inputs["chargeType"] = state ? state.chargeType : undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["expireTime"] = state ? state.expireTime : undefined;
            inputs["internal"] = state ? state.internal : undefined;
            inputs["ipSets"] = state ? state.ipSets : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["privateIp"] = state ? state.privateIp : undefined;
            inputs["remark"] = state ? state.remark : undefined;
            inputs["securityGroup"] = state ? state.securityGroup : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["tag"] = state ? state.tag : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as LBArgs | undefined;
            inputs["chargeType"] = args ? args.chargeType : undefined;
            inputs["internal"] = args ? args.internal : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["remark"] = args ? args.remark : undefined;
            inputs["securityGroup"] = args ? args.securityGroup : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tag"] = args ? args.tag : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["expireTime"] = undefined /*out*/;
            inputs["ipSets"] = undefined /*out*/;
            inputs["privateIp"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LB.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LB resources.
 */
export interface LBState {
    /**
     * , argument `chargeType` is deprecated for optimizing parameters.
     *
     * @deprecated attribute `charge_type` is deprecated for optimizing parameters
     */
    chargeType?: pulumi.Input<string>;
    /**
     * The time of creation for load balancer, formatted in RFC3339 time string.
     */
    createTime?: pulumi.Input<string>;
    /**
     * **Deprecated** attribute `expireTime` is deprecated for optimizing outputs.
     *
     * @deprecated attribute `expire_time` is deprecated for optimizing outputs
     */
    expireTime?: pulumi.Input<string>;
    /**
     * Indicate whether the load balancer is intranet mode.(Default: `"false"`)
     */
    internal?: pulumi.Input<boolean>;
    /**
     * It is a nested type which documented below.
     */
    ipSets?: pulumi.Input<pulumi.Input<inputs.ulb.LBIpSet>[]>;
    name?: pulumi.Input<string>;
    /**
     * The IP address of intranet IP. It is `""` if `internal` is `false`.
     */
    privateIp?: pulumi.Input<string>;
    /**
     * The remarks of the load balancer. (Default: `""`).
     */
    remark?: pulumi.Input<string>;
    /**
     * The ID of the associated security group. The securityGroup only takes effect for ULB instances of requestProxy mode and extranet mode at present.
     */
    securityGroup?: pulumi.Input<string>;
    /**
     * The ID of subnet that intranet load balancer belongs to. This argument is not required if default subnet.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * A tag assigned to load balancer, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
     */
    tag?: pulumi.Input<string>;
    /**
     * The ID of the VPC linked to the Load balancer, This argument is not required if default VPC.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LB resource.
 */
export interface LBArgs {
    /**
     * , argument `chargeType` is deprecated for optimizing parameters.
     *
     * @deprecated attribute `charge_type` is deprecated for optimizing parameters
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Indicate whether the load balancer is intranet mode.(Default: `"false"`)
     */
    internal?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    /**
     * The remarks of the load balancer. (Default: `""`).
     */
    remark?: pulumi.Input<string>;
    /**
     * The ID of the associated security group. The securityGroup only takes effect for ULB instances of requestProxy mode and extranet mode at present.
     */
    securityGroup?: pulumi.Input<string>;
    /**
     * The ID of subnet that intranet load balancer belongs to. This argument is not required if default subnet.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * A tag assigned to load balancer, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
     */
    tag?: pulumi.Input<string>;
    /**
     * The ID of the VPC linked to the Load balancer, This argument is not required if default VPC.
     */
    vpcId?: pulumi.Input<string>;
}
