// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./disk";
export * from "./diskAttachment";
export * from "./getDisk";

// Import resources to register:
import { Disk } from "./disk";
import { DiskAttachment } from "./diskAttachment";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "ucloud:udisk/disk:Disk":
                return new Disk(name, <any>undefined, { urn })
            case "ucloud:udisk/diskAttachment:DiskAttachment":
                return new DiskAttachment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("ucloud", "udisk/disk", _module)
pulumi.runtime.registerResourceModule("ucloud", "udisk/diskAttachment", _module)