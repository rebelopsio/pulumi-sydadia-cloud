using System.Collections.Generic;
using System.Linq;
using Pulumi;
using synadia-cloud = Pulumi.synadia-cloud;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new synadia-cloud.Random("myRandomResource", new()
    {
        Length = 24,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myRandomResource.Result },
        },
    };
});

