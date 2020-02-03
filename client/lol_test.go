package client

import (
	"context"
	"github.com/anthonyrouseau/go-riot/lol"
	"github.com/anthonyrouseau/go-riot/queue"
	"testing"
)

const (
	testAPIKey     = "RGAPI-a348803e-d2f9-42c4-a442-c8a5b71dd12d"
	testSummonerID = "c0n56ouT0eGJLaVy8Sbfe628zfBkRbaKZZwByHVDQik"
	testLeagueID   = "c60807e8-6afb-38fd-ab9b-ae8588dc8b27"
)

func TestLOLMethods(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient(testAPIKey, SetVariant(devClient))
	if err != nil {
		t.Error(err)
	}
	t.Run("LOLChallenger", func(t *testing.T) {
		leagueInfo, err := client.LOLChallenger(ctx, queue.RankedSolo5x5)
		if err != nil {
			t.Error(err)
		}
		if leagueInfo == nil {
			t.Error("League Info was nil value")
		} else {
			if leagueInfo.Tier != "CHALLENGER" {
				t.Errorf("Expected tier to be %s but got %s", "CHALLENGER", leagueInfo.Tier)
			}
		}
	})
	t.Run("LOLSummonerLeagueEntries", func(t *testing.T) {
		leagueEntries, err := client.LOLSummonerLeagueEntries(ctx, testSummonerID)
		if err != nil {
			t.Error(err)
		}
		if leagueEntries == nil {
			t.Error("League entries was nil value")
		} else {
			if leagueEntries[0].SummonerID != testSummonerID {
				t.Errorf("Expected summoner id to be %s but got %s", testSummonerID, leagueEntries[0].SummonerID)
			}
		}
	})
	t.Run("LOLAllLeagueEntries", func(t *testing.T) {
		leagueEntries, err := client.LOLAllLeagueEntries(ctx, queue.RankedSolo5x5, lol.Diamond, lol.I)
		if err != nil {
			t.Error(err)
		}
		if leagueEntries == nil {
			t.Error("League entries was nil value")
		} else {
			if len(leagueEntries) > 0 && leagueEntries[0].Tier != lol.Diamond {
				t.Errorf("Expected tier to be %s but got %s", lol.Diamond, leagueEntries[0].Tier)
			}
		}
	})
	t.Run("LOLGrandmaster", func(t *testing.T) {
		leagueInfo, err := client.LOLGrandmaster(ctx, queue.RankedSolo5x5)
		if err != nil {
			t.Error(err)
		}
		if leagueInfo == nil {
			t.Error("League Info was nil value")
		} else {
			if leagueInfo.Tier != "GRANDMASTER" {
				t.Errorf("Expected tier to be %s but got %s", "GRANDMASTER", leagueInfo.Tier)
			}
		}
	})
	t.Run("LOLLeague", func(t *testing.T) {
		leagueInfo, err := client.LOLLeague(ctx, testLeagueID)
		if err != nil {
			t.Error(err)
		}
		if leagueInfo == nil {
			t.Error("League Info was nil value")
		} else {
			if leagueInfo.ID != testLeagueID {
				t.Errorf("Expected league id to be %s but got %s", testLeagueID, leagueInfo.ID)
			}
		}
	})
}
