// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ucloud.Ipsecvpn.Inputs
{

    public sealed class VPNConnectionIkeConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authentication algorithm of IPSec negotiation. Possible values: `sha1`, `md5`. (Default: `sha1`)
        /// </summary>
        [Input("authenticationAlgorithm")]
        public Input<string>? AuthenticationAlgorithm { get; set; }

        /// <summary>
        /// The Diffie-Hellman group used by IKE negotiation. Possible values: `1`, `2`, `5`, `14`, `15`, `16`. (Default:`15`)
        /// </summary>
        [Input("dhGroup")]
        public Input<string>? DhGroup { get; set; }

        /// <summary>
        /// The encryption algorithm of IPSec negotiation. Possible values: `aes128`, `aes192`, `aes256`, `aes512`, `3des`. (Default: `aes128`).
        /// </summary>
        [Input("encryptionAlgorithm")]
        public Input<string>? EncryptionAlgorithm { get; set; }

        /// <summary>
        /// The negotiation exchange mode of IKE V1 of VPN gateway. Possible values: `main` (main mode), `aggressive` (aggressive mode). (Default: `main`)
        /// </summary>
        [Input("exchangeMode")]
        public Input<string>? ExchangeMode { get; set; }

        /// <summary>
        /// The version of the IKE protocol which only be supported IKE V1 protocol at present. Possible values: ikev1. (Default: ikev1)
        /// </summary>
        [Input("ikeVersion")]
        public Input<string>? IkeVersion { get; set; }

        /// <summary>
        /// The identification of the VPN gateway.
        /// </summary>
        [Input("localId")]
        public Input<string>? LocalId { get; set; }

        /// <summary>
        /// The key used for authentication between the VPN gateway and the Customer gateway which contains 1-128 characters and only support English, numbers and special characters: `!@#$%^&amp;*()_+-=[]:,./'~`.
        /// </summary>
        [Input("preSharedKey", required: true)]
        public Input<string> PreSharedKey { get; set; } = null!;

        /// <summary>
        /// The identification of the Customer gateway.
        /// </summary>
        [Input("remoteId")]
        public Input<string>? RemoteId { get; set; }

        /// <summary>
        /// The Security Association lifecycle as the result of IPSec negotiation. Unit: second. Range: 1200-604800. (Default: `3600`)
        /// </summary>
        [Input("saLifeTime")]
        public Input<int>? SaLifeTime { get; set; }

        public VPNConnectionIkeConfigArgs()
        {
        }
    }
}
