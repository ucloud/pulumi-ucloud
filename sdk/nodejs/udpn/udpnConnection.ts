// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class UdpnConnection extends pulumi.CustomResource {
    /**
     * Get an existing UdpnConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UdpnConnectionState, opts?: pulumi.CustomResourceOptions): UdpnConnection {
        return new UdpnConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ucloud:udpn/udpnConnection:UdpnConnection';

    /**
     * Returns true if the given object is an instance of UdpnConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UdpnConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UdpnConnection.__pulumiType;
    }

    public readonly bandwidth!: pulumi.Output<number | undefined>;
    public readonly chargeType!: pulumi.Output<string | undefined>;
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    public readonly duration!: pulumi.Output<number | undefined>;
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    public readonly peerRegion!: pulumi.Output<string>;

    /**
     * Create a UdpnConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UdpnConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UdpnConnectionArgs | UdpnConnectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UdpnConnectionState | undefined;
            inputs["bandwidth"] = state ? state.bandwidth : undefined;
            inputs["chargeType"] = state ? state.chargeType : undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["duration"] = state ? state.duration : undefined;
            inputs["expireTime"] = state ? state.expireTime : undefined;
            inputs["peerRegion"] = state ? state.peerRegion : undefined;
        } else {
            const args = argsOrState as UdpnConnectionArgs | undefined;
            if (!args || args.peerRegion === undefined) {
                throw new Error("Missing required property 'peerRegion'");
            }
            inputs["bandwidth"] = args ? args.bandwidth : undefined;
            inputs["chargeType"] = args ? args.chargeType : undefined;
            inputs["duration"] = args ? args.duration : undefined;
            inputs["peerRegion"] = args ? args.peerRegion : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["expireTime"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(UdpnConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UdpnConnection resources.
 */
export interface UdpnConnectionState {
    readonly bandwidth?: pulumi.Input<number>;
    readonly chargeType?: pulumi.Input<string>;
    readonly createTime?: pulumi.Input<string>;
    readonly duration?: pulumi.Input<number>;
    readonly expireTime?: pulumi.Input<string>;
    readonly peerRegion?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UdpnConnection resource.
 */
export interface UdpnConnectionArgs {
    readonly bandwidth?: pulumi.Input<number>;
    readonly chargeType?: pulumi.Input<string>;
    readonly duration?: pulumi.Input<number>;
    readonly peerRegion: pulumi.Input<string>;
}
