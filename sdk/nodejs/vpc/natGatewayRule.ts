// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Nat Gateway resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ucloud from "@pulumi/ucloud";
 *
 * const fooVPC = new ucloud.vpc.VPC("fooVPC", {
 *     tag: "tf-acc",
 *     cidrBlocks: ["192.168.0.0/16"],
 * });
 * const fooSubnet = new ucloud.vpc.Subnet("fooSubnet", {
 *     tag: "tf-acc",
 *     cidrBlock: "192.168.1.0/24",
 *     vpcId: fooVPC.id,
 * });
 * const fooEIP = new ucloud.unet.EIP("fooEIP", {
 *     bandwidth: 1,
 *     internetType: "bgp",
 *     chargeMode: "bandwidth",
 *     tag: "tf-acc",
 * });
 * const fooSecurityGroup = ucloud.unet.getSecurityGroup({
 *     type: "recommend_web",
 * });
 * const defaultZone = ucloud.uaccount.getZone({});
 * const defaultImage = defaultZone.then(defaultZone => ucloud.uhost.getImage({
 *     availabilityZone: defaultZone.zones[0].id,
 *     nameRegex: "^CentOS 7.[1-2] 64",
 *     imageType: "base",
 * }));
 * const fooInstance = new ucloud.uhost.Instance("fooInstance", {
 *     vpcId: fooVPC.id,
 *     subnetId: fooSubnet.id,
 *     availabilityZone: defaultZone.then(defaultZone => defaultZone.zones[0].id),
 *     imageId: defaultImage.then(defaultImage => defaultImage.images[0].id),
 *     instanceType: "n-basic-1",
 *     chargeType: "dynamic",
 *     tag: "tf-acc",
 * });
 * const fooNATGateway = new ucloud.vpc.NATGateway("fooNATGateway", {
 *     vpcId: fooVPC.id,
 *     subnetIds: [fooSubnet.id],
 *     eipId: fooEIP.id,
 *     tag: "tf-acc",
 *     enableWhiteList: false,
 *     securityGroup: fooSecurityGroup.then(fooSecurityGroup => fooSecurityGroup.securityGroups[0].id),
 * });
 * const fooNATGatewayRule = new ucloud.vpc.NATGatewayRule("fooNATGatewayRule", {
 *     natGatewayId: fooNATGateway.id,
 *     protocol: "tcp",
 *     srcEipId: fooEIP.id,
 *     srcPortRange: "88",
 *     dstIp: fooInstance.privateIp,
 *     dstPortRange: "80",
 * });
 * const bar = new ucloud.vpc.NATGatewayRule("bar", {
 *     natGatewayId: fooNATGateway.id,
 *     protocol: "tcp",
 *     srcEipId: fooEIP.id,
 *     srcPortRange: "90-100",
 *     dstIp: fooInstance.privateIp,
 *     dstPortRange: "90-100",
 * });
 * ```
 */
export class NATGatewayRule extends pulumi.CustomResource {
    /**
     * Get an existing NATGatewayRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NATGatewayRuleState, opts?: pulumi.CustomResourceOptions): NATGatewayRule {
        return new NATGatewayRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ucloud:vpc/nATGatewayRule:NATGatewayRule';

    /**
     * Returns true if the given object is an instance of NATGatewayRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NATGatewayRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NATGatewayRule.__pulumiType;
    }

    /**
     * The private ip of instance bound to the jNAT gateway.
     */
    public readonly dstIp!: pulumi.Output<string>;
    /**
     * The range of port numbers of the private ip, range: 1-65535. (eg: `port` or `port1-port2`).
     */
    public readonly dstPortRange!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the Nat Gateway.
     */
    public readonly natGatewayId!: pulumi.Output<string>;
    /**
     * The protocol of the Nat Gateway Rule. Possible values: `tcp`, `udp`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The ID of eip associate to the Nat Gateway.
     */
    public readonly srcEipId!: pulumi.Output<string>;
    /**
     * The range of port numbers of the eip, range: 1-65535. (eg: `port` or `port1-port2`).
     */
    public readonly srcPortRange!: pulumi.Output<string>;

    /**
     * Create a NATGatewayRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NATGatewayRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NATGatewayRuleArgs | NATGatewayRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NATGatewayRuleState | undefined;
            inputs["dstIp"] = state ? state.dstIp : undefined;
            inputs["dstPortRange"] = state ? state.dstPortRange : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["natGatewayId"] = state ? state.natGatewayId : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["srcEipId"] = state ? state.srcEipId : undefined;
            inputs["srcPortRange"] = state ? state.srcPortRange : undefined;
        } else {
            const args = argsOrState as NATGatewayRuleArgs | undefined;
            if ((!args || args.dstIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstIp'");
            }
            if ((!args || args.dstPortRange === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstPortRange'");
            }
            if ((!args || args.natGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'natGatewayId'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.srcEipId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcEipId'");
            }
            if ((!args || args.srcPortRange === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcPortRange'");
            }
            inputs["dstIp"] = args ? args.dstIp : undefined;
            inputs["dstPortRange"] = args ? args.dstPortRange : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["natGatewayId"] = args ? args.natGatewayId : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["srcEipId"] = args ? args.srcEipId : undefined;
            inputs["srcPortRange"] = args ? args.srcPortRange : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NATGatewayRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NATGatewayRule resources.
 */
export interface NATGatewayRuleState {
    /**
     * The private ip of instance bound to the jNAT gateway.
     */
    dstIp?: pulumi.Input<string>;
    /**
     * The range of port numbers of the private ip, range: 1-65535. (eg: `port` or `port1-port2`).
     */
    dstPortRange?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * The ID of the Nat Gateway.
     */
    natGatewayId?: pulumi.Input<string>;
    /**
     * The protocol of the Nat Gateway Rule. Possible values: `tcp`, `udp`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The ID of eip associate to the Nat Gateway.
     */
    srcEipId?: pulumi.Input<string>;
    /**
     * The range of port numbers of the eip, range: 1-65535. (eg: `port` or `port1-port2`).
     */
    srcPortRange?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NATGatewayRule resource.
 */
export interface NATGatewayRuleArgs {
    /**
     * The private ip of instance bound to the jNAT gateway.
     */
    dstIp: pulumi.Input<string>;
    /**
     * The range of port numbers of the private ip, range: 1-65535. (eg: `port` or `port1-port2`).
     */
    dstPortRange: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * The ID of the Nat Gateway.
     */
    natGatewayId: pulumi.Input<string>;
    /**
     * The protocol of the Nat Gateway Rule. Possible values: `tcp`, `udp`.
     */
    protocol: pulumi.Input<string>;
    /**
     * The ID of eip associate to the Nat Gateway.
     */
    srcEipId: pulumi.Input<string>;
    /**
     * The range of port numbers of the eip, range: 1-65535. (eg: `port` or `port1-port2`).
     */
    srcPortRange: pulumi.Input<string>;
}
