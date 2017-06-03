# postcodedaylight

Sample function used to evaluate [faas](https://github.com/alexellis/faas) from [get-faas.com](http://docs.get-faas.com/).  The function accepts a UK postcode and returns the daylight hours for that place using a combination of endpoints offered by [postcodes.io](https://postcodes.io) and [sunrise-sunset.org](https://sunrise-sunset.org/api)

You can execute the function like this:

`curl http://localhost:8080/function/func_postcodedaylight -d "SW1A 1AA"`

(or use the [FaaS UI](http://localhost:8080/ui/) to pass the postcode as text)

## Installation

You can either install `postcodeDaylight` via your FaaS compose file or you can add it via the UI.

### Compose file

Add this to your FaaS `docker-compose.yml` 

```
    # Returns the amount of daylight to expect for a given postcode        
    postcodedaylight:
        image: rgee0/postcodedaylight:latest
        labels:
            function: "true"
        depends_on:
            - gateway
        networks:
            - functions
        environment:
            fprocess: "/go/bin/postcodedaylight"
            no_proxy: "gateway"
            https_proxy: $https_proxy
```
and then redeploy the FaaS func stack
`docker stack deploy -c docker-compose.yml func`

### UI

Use the `CREATE NEW FUNCTION` link and add these details:

- Image: `rgee0/postcodedaylight:latest`
- Service name: `func_postcodedaylight`
- fProcess: `/go/bin/postcodedaylight`
- Network: `func_functions`
