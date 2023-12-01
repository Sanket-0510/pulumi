// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class ExampleServer extends pulumi.CustomResource {
    /**
     * Get an existing ExampleServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ExampleServer {
        return new ExampleServer(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'example:index:ExampleServer';

    /**
     * Returns true if the given object is an instance of ExampleServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExampleServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExampleServer.__pulumiType;
    }

    public /*out*/ readonly name!: pulumi.Output<string | undefined>;

    /**
     * Create a ExampleServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ExampleServerArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["propertiesCollection"] = args ? args.propertiesCollection : undefined;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ExampleServer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ExampleServer resource.
 */
export interface ExampleServerArgs {
    propertiesCollection?: pulumi.Input<pulumi.Input<inputs.ServerPropertiesForReplicaArgs | inputs.ServerPropertiesForRestoreArgs>[]>;
}