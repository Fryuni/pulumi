// *** WARNING: this file was generated by the Lumi IDL Compiler (CLIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";

import {InternetGateway} from "./internetGateway";
import {VPC} from "./vpc";

export class VPCGatewayAttachment extends lumi.Resource implements VPCGatewayAttachmentArgs {
    public readonly name: string;
    public readonly vpc: VPC;
    public readonly internetGateway: InternetGateway;

    constructor(name: string, args: VPCGatewayAttachmentArgs) {
        super();
        if (name === undefined) {
            throw new Error("Missing required resource name");
        }
        this.name = name;
        if (args.vpc === undefined) {
            throw new Error("Missing required argument 'vpc'");
        }
        this.vpc = args.vpc;
        if (args.internetGateway === undefined) {
            throw new Error("Missing required argument 'internetGateway'");
        }
        this.internetGateway = args.internetGateway;
    }
}

export interface VPCGatewayAttachmentArgs {
    readonly vpc: VPC;
    readonly internetGateway: InternetGateway;
}


