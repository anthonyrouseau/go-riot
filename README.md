# Go-Riot
This is a wrapper for the [Riot Games API](https://developer.riotgames.com/apis) implemented in Go.
This repository contains the client package which interacts with the Riot API as well as supporting packages that define Riot Types.
These supporting packages can be used independently as well.
For more information about the API including types, possible values, etc., please see the Riot Games API documentation.

## Installation

To install the client library use the command

`go get github.com/anthonyrouseau/go-riot/client`

Other packages can be installed using a similar format.

## Example Use

````Go
package main

import (
	"context"
	"fmt"
	"github.com/anthonyrouseau/go-riot/client"
	"github.com/anthonyrouseau/go-riot/queue"
	"log"
)

func main() {
	apiKey := "YOUR_API_KEY"
	//Creating a client requires an API key from Riot
	//A client defaults to UnspecifiedClient which will not use rate limiting
	//Specificing a client variant will rate limit according to the API Key type
	client, err := client.NewClient(apiKey, client.SetVariant(client.DevClient))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The client API key is: %s\n", client.APIKey())
	ctx := context.Background()
	leagueInfo, err := client.LOLChallenger(ctx, queue.RankedSolo5x5)
	if err != nil {
		//error handling
	}
	fmt.Printf("The Challenger league name is: %s\n", leagueInfo.Name)
}
````

## Optional Parameters 

Some of the functions (e.g. NewClient in the client package) make use of Functional Options. You can learn more
about them [here](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis). Use these functions to 
pass optional values to the API calls.

## Rate Limiting

If a client is created with a variant other than UnspecifiedClient, then calls using the client will be rate limited according
to the values allowed by Riot. The client is unable to determine rate limiting form the API key alone so make sure to set
the client variant if rate limiting will not be handled elsewhere.

## Examples
Examples of use can be found in the [examples](https://github.com/anthonyrouseau/go-riot/tree/master/examples) directory.
