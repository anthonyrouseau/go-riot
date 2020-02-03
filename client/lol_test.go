package client

import (
	"context"
	"github.com/anthonyrouseau/go-riot/queue"
	"testing"
)

const (
	testAPIKey     = "RGAPI-a348803e-d2f9-42c4-a442-c8a5b71dd12d"
	testSummonerID = "c0n56ouT0eGJLaVy8Sbfe628zfBkRbaKZZwByHVDQik"
)

func TestLOLChallenger(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient(testAPIKey, SetVariant(devClient))
	if err != nil {
		t.Error(err)
	}
	leagueInfo, err := client.LOLChallenger(ctx, queue.RankedSolo5x5)
	if err != nil {
		t.Error(err)
	}
	if leagueInfo == nil {
		t.Error("League Info was nil value")
	}
	if leagueInfo.Tier != "CHALLENGER" {
		t.Errorf("Expected tier to be %s but got %s", "CHALLENGER", leagueInfo.Tier)
	}
}

func TestLOLSummonerLeagueEntries(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient(testAPIKey, SetVariant(devClient))
	if err != nil {
		t.Error(err)
	}
	leagueEntries, err := client.LOLSummonerLeagueEntries(ctx, testSummonerID)
	if err != nil {
		t.Error(err)
	}
	if leagueEntries == nil {
		t.Error("League entries was nil value")
	}
	if leagueEntries[0].SummonerID != testSummonerID {
		t.Errorf("Expected summoner id to be %s but got %s", testSummonerID, leagueEntries[0].SummonerID)
	}
}
