package main

import (
	"context"
	"fmt"
	"github.com/anthonyrouseau/go-riot/client"
	"github.com/anthonyrouseau/go-riot/queue"
	"log"
)

func main() {
	apiKey := "RGAPI-16d30c4a-0678-42b4-b1dc-c1e8ed1b1a90"
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
