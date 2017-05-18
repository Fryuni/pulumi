// *** WARNING: this file was generated by the Lumi IDL Compiler (CLIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";

import {Authorizer} from "./authorizer";
import {Model} from "./model";
import {Resource} from "./resource";
import {RestAPI} from "./restAPI";

export let AWSIAMAuthorization: AuthorizationType = "AWS_IAM";
export let AWSIntegration: IntegrationType = "AWS";
export let AWSProxyIntegration: IntegrationType = "AWS_PROXY";
export let CognitoAuthorization: AuthorizationType = "COGNITO_USER_POOLS";
export let CustomAuthorization: AuthorizationType = "CUSTOM";
export let HTTPIntegration: IntegrationType = "HTTP";
export let HTTPProxyIntegration: IntegrationType = "HTTP_PROXY";
export let LoggingErrorLevel: LoggingLevel = "ERROR";
export let LoggingInfoLevel: LoggingLevel = "INFO";
export let LoggingOff: LoggingLevel = "OFF";
export let MockIntegration: IntegrationType = "MOCK";
export let NoAuthorization: AuthorizationType = "NONE";
export let PassthroughNever: PassthroughBehavior = "NEVER";
export let PassthroughWhenNoMatch: PassthroughBehavior = "WHEN_NO_MATCH";
export let PassthroughWhenNoTemplates: PassthroughBehavior = "WHEN_NO_TEMPLATES";

export type AuthorizationType =
    "AWS_IAM" |
    "COGNITO_USER_POOLS" |
    "CUSTOM" |
    "NONE";

export interface Integration {
    type: IntegrationType;
    cacheKeyParameters?: string[];
    cacheNamespace?: string;
    credentials?: string;
    integrationHTTPMethod?: string;
    integrationResponse?: IntegrationResponse[];
    passthroughBehavior?: PassthroughBehavior;
    requestParameters?: {[key: string]: string};
    requestTemplates?: {[key: string]: string};
    uri?: string;
}

export interface IntegrationResponse {
    responseParameters?: {[key: string]: string};
    responseTemplates?: {[key: string]: string};
    selectionPattern?: string;
    statusCode?: string;
}

export type IntegrationType =
    "AWS" |
    "AWS_PROXY" |
    "HTTP" |
    "HTTP_PROXY" |
    "MOCK";

export type LoggingLevel =
    "ERROR" |
    "INFO" |
    "OFF";

export class Method extends lumi.Resource implements MethodArgs {
    public readonly name: string;
    public httpMethod: string;
    public apiResource: Resource;
    public restAPI: RestAPI;
    public apiKeyRequired?: boolean;
    public authorizationType?: AuthorizationType;
    public authorizer?: Authorizer;
    public integration?: Integration;
    public methodResponses?: MethodResponse[];
    public requestModels?: {[key: string]: Model};
    public requestParameters?: {[key: string]: boolean};

    constructor(name: string, args: MethodArgs) {
        super();
        if (name === undefined) {
            throw new Error("Missing required resource name");
        }
        this.name = name;
        if (args.httpMethod === undefined) {
            throw new Error("Missing required argument 'httpMethod'");
        }
        this.httpMethod = args.httpMethod;
        if (args.apiResource === undefined) {
            throw new Error("Missing required argument 'apiResource'");
        }
        this.apiResource = args.apiResource;
        if (args.restAPI === undefined) {
            throw new Error("Missing required argument 'restAPI'");
        }
        this.restAPI = args.restAPI;
        this.apiKeyRequired = args.apiKeyRequired;
        this.authorizationType = args.authorizationType;
        this.authorizer = args.authorizer;
        this.integration = args.integration;
        this.methodResponses = args.methodResponses;
        this.requestModels = args.requestModels;
        this.requestParameters = args.requestParameters;
    }
}

export interface MethodArgs {
    httpMethod: string;
    apiResource: Resource;
    restAPI: RestAPI;
    apiKeyRequired?: boolean;
    authorizationType?: AuthorizationType;
    authorizer?: Authorizer;
    integration?: Integration;
    methodResponses?: MethodResponse[];
    requestModels?: {[key: string]: Model};
    requestParameters?: {[key: string]: boolean};
}

export interface MethodResponse {
    statusCode: string;
    responseModels?: {[key: string]: Model};
    responseParameters?: {[key: string]: boolean};
}

export interface MethodSetting {
    cacheDataEncrypted?: boolean;
    cacheTTLInSeconds?: number;
    cachingEnabled?: boolean;
    dataTraceEnabled?: boolean;
    httpMethod?: string;
    loggingLevel?: LoggingLevel;
    metricsEnabled?: boolean;
    resourcePath?: string;
    throttlingBurstLimit?: number;
    throttlingRateLimit?: number;
}

export type PassthroughBehavior =
    "NEVER" |
    "WHEN_NO_MATCH" |
    "WHEN_NO_TEMPLATES";


