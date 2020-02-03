package client

import (
	"context"
	"github.com/anthonyrouseau/go-riot/lol"
	"github.com/anthonyrouseau/go-riot/queue"
	"testing"
)

const (
	testAPIKey        = "RGAPI-16d30c4a-0678-42b4-b1dc-c1e8ed1b1a90"
	testSummonerID    = "c0n56ouT0eGJLaVy8Sbfe628zfBkRbaKZZwByHVDQik"
	testLeagueID      = "c60807e8-6afb-38fd-ab9b-ae8588dc8b27"
	testMatchID       = 3285199726
	testAccountID     = "182BwdvZQMIpZwCk4vGZiHs71qSK18Bo9Ll6zBU2LQ"
	testSummonerName  = "C9 Sneaky"
	testSummonerPUUID = "Uc6OVbDglCiUrKxrRvcqRmo6RY_EKiFFy5CqNet57wEcFyzA1ElI5C8WUuPvp1eLFgYrmLajMGpsyA"
	testChampionID    = 202
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
	t.Run("LOLMaster", func(t *testing.T) {
		leagueInfo, err := client.LOLMaster(ctx, queue.RankedSolo5x5)
		if err != nil {
			t.Error(err)
		}
		if leagueInfo == nil {
			t.Error("League Info was nil value")
		} else {
			if leagueInfo.Tier != "MASTER" {
				t.Errorf("Expected tier to be %s but got %s", "MASTER", leagueInfo.Tier)
			}
		}
	})
	t.Run("LOLMatch", func(t *testing.T) {
		match, err := client.LOLMatch(ctx, testMatchID)
		if err != nil {
			t.Error(err)
		}
		if match == nil {
			t.Error("Match was nil value")
		} else {
			if match.GameID != testMatchID {
				t.Errorf("Expected match id to be %d but got %d", testMatchID, match.GameID)
			}
		}
	})
	t.Run("LOLAccountMatches", func(t *testing.T) {
		matchList, err := client.LOLAccountMatches(ctx, testAccountID)
		if err != nil {
			t.Error(err)
		}
		if matchList == nil {
			t.Error("MatchList was nil value")
		}
	})
	t.Run("LOLMatchTimeline", func(t *testing.T) {
		matchTimeline, err := client.LOLMatchTimeline(ctx, testMatchID)
		if err != nil {
			t.Error(err)
		}
		if matchTimeline == nil {
			t.Error("MatchTimeline was nil value")
		}
	})
	t.Run("LOLSummonerByAccount", func(t *testing.T) {
		summoner, err := client.LOLSummonerByAccount(ctx, testAccountID)
		if err != nil {
			t.Error(err)
		}
		if summoner == nil {
			t.Error("Summoner was nil value")
		} else {
			if summoner.AccountID != testAccountID {
				t.Errorf("Expected summoner account id to be %s but got %s", testAccountID, summoner.AccountID)
			}
		}
	})
	t.Run("LOLSummonerByName", func(t *testing.T) {
		summoner, err := client.LOLSummonerByName(ctx, testSummonerName)
		if err != nil {
			t.Error(err)
		}
		if summoner == nil {
			t.Error("Summoner was nil value")
		} else {
			if summoner.Name != testSummonerName {
				t.Errorf("Expected summoner name to be %s but got %s", testSummonerName, summoner.Name)
			}
		}
	})
	t.Run("LOLSummonerByPUUID", func(t *testing.T) {
		summoner, err := client.LOLSummonerByPUUID(ctx, testSummonerPUUID)
		if err != nil {
			t.Error(err)
		}
		if summoner == nil {
			t.Error("Summoner was nil value")
		} else {
			if summoner.PUUID != testSummonerPUUID {
				t.Errorf("Expected summoner PUUID to be %s but got %s", testSummonerPUUID, summoner.PUUID)
			}
		}
	})
	t.Run("LOLSummonerBySummonerID", func(t *testing.T) {
		summoner, err := client.LOLSummonerBySummonerID(ctx, testSummonerID)
		if err != nil {
			t.Error(err)
		}
		if summoner == nil {
			t.Error("Summoner was nil value")
		} else {
			if summoner.PUUID != testSummonerPUUID {
				t.Errorf("Expected summoner id to be %s but got %s", testSummonerID, summoner.ID)
			}
		}
	})
	t.Run("LOLAllSummonerChampionMastery", func(t *testing.T) {
		champMastery, err := client.LOLAllSummonerChampionMastery(ctx, testSummonerID)
		if err != nil {
			t.Error(err)
		}
		if champMastery == nil {
			t.Error("Summoner was nil value")
		}
	})
	t.Run("LOLSummonerChampionMastery", func(t *testing.T) {
		champMastery, err := client.LOLSummonerChampionMastery(ctx, testSummonerID, testChampionID)
		if err != nil {
			t.Error(err)
		}
		if champMastery == nil {
			t.Error("Summoner was nil value")
		}
	})
	t.Run("LOLSummonerChampionMasteryScore", func(t *testing.T) {
		_, err := client.LOLSummonerChampionMasteryScore(ctx, testSummonerID)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("LOLChampionRotations", func(t *testing.T) {
		rotation, err := client.LOLChampionRotations(ctx)
		if err != nil {
			t.Error(err)
		}
		if rotation == nil {
			t.Error("Rotation was nil value")
		}
	})
	t.Run("LOLFeaturedGames", func(t *testing.T) {
		featured, err := client.LOLFeaturedGames(ctx)
		if err != nil {
			t.Error(err)
		}
		if featured == nil {
			t.Error("Featured was nil value")
		}
	})
	t.Run("LOLActiveGame", func(t *testing.T) {
		active, err := client.LOLActiveGame(ctx, testSummonerID)
		if err != nil {
			if err.Error() != errNotFound.Error() {
				t.Error(err)
			}
		} else {
			if active == nil {
				t.Error("Active game was nil value")
			}
		}
	})
}
