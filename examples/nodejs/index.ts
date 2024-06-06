import * as pulumi from "@pulumi/pulumi";
import * as synadia-cloud from "@pulumi/synadia-cloud";

const myRandomResource = new synadia-cloud.Random("myRandomResource", {length: 24});
export const output = {
    value: myRandomResource.result,
};
