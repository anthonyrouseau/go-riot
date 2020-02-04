package client

import (
	"context"
	"github.com/anthonyrouseau/go-riot/tft"
	"testing"
)

func TestTFTMethods(t *testing.T) {
	const (
		testSummonerID    = "c0n56ouT0eGJLaVy8Sbfe628zfBkRbaKZZwByHVDQik"
		testLeagueID      = "2417f7dd-cefe-3ef6-a8f1-556c780f45b5"
		testSummonerPUUID = "Uc6OVbDglCiUrKxrRvcqRmo6RY_EKiFFy5CqNet57wEcFyzA1ElI5C8WUuPvp1eLFgYrmLajMGpsyA"
		testMatchID       = "NA1_3229913302"
		testSummonerName  = "C9 Sneaky"
		testAccountID     = "182BwdvZQMIpZwCk4vGZiHs71qSK18Bo9Ll6zBU2LQ"
	)
	ctx := context.Background()
	client, err := NewClient(testAPIKey, SetVariant(DevClient))
	if err != nil {
		t.Error(err)
	}
	t.Run("TFTChallenger", func(t *testing.T) {
		leagueList, err := client.TFTChallenger(ctx)
		if err != nil {
			t.Error(err)
		}
		if leagueList == nil {
			t.Error("League List was nil value")
		} else {
			if leagueList.Tier != "CHALLENGER" {
				t.Errorf("Expected tier to be %s but got %s", "CHALLENGER", leagueList.Tier)
			}
		}
	})
	t.Run("TFTSummonerLeagueEntries", func(t *testing.T) {
		leagueEntries, err := client.TFTSummonerLeagueEntries(ctx, testSummonerID)
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
	t.Run("TFTAllLeagueEntries", func(t *testing.T) {
		leagueEntries, err := client.TFTAllLeagueEntries(ctx, tft.Diamond, tft.I)
		if err != nil {
			t.Error(err)
		}
		if leagueEntries == nil {
			t.Error("League entries was nil value")
		} else {
			if len(leagueEntries) > 0 && leagueEntries[0].Tier != tft.Diamond {
				t.Errorf("Expected tier to be %s but got %s", tft.Diamond, leagueEntries[0].Tier)
			}
		}
	})
	t.Run("TFTGrandmaster", func(t *testing.T) {
		leagueInfo, err := client.TFTGrandmaster(ctx)
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
	t.Run("TFTLeague", func(t *testing.T) {
		leagueInfo, err := client.TFTLeague(ctx, testLeagueID)
		if err != nil {
			t.Error(err)
		}
		if leagueInfo == nil {
			t.Error("League Info was nil value")
		} else {
			if leagueInfo.LeagueID != testLeagueID {
				t.Errorf("Expected league id to be %s but got %s", testLeagueID, leagueInfo.LeagueID)
			}
		}
	})
	t.Run("TFTMaster", func(t *testing.T) {
		leagueInfo, err := client.TFTMaster(ctx)
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
	t.Run("TFTMatchIDs", func(t *testing.T) {
		matches, err := client.TFTMatchIDs(ctx, testSummonerPUUID)
		if err != nil {
			t.Error(err)
		}
		if matches == nil {
			t.Error("Matches was nil value")
		}
	})
	t.Run("TFTMatch", func(t *testing.T) {
		match, err := client.TFTMatch(ctx, testMatchID)
		if err != nil {
			t.Error(err)
		}
		if match == nil {
			t.Error("Match was nil value")
		} else {
			if match.Metadata.MatchID != testMatchID {
				t.Errorf("Expected match id to be %s but got %s", testMatchID, match.Metadata.MatchID)
			}
		}
	})
	t.Run("TFTSummonerByAccount", func(t *testing.T) {
		summoner, err := client.TFTSummonerByAccount(ctx, testAccountID)
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
	t.Run("TFTSummonerByName", func(t *testing.T) {
		summoner, err := client.TFTSummonerByName(ctx, testSummonerName)
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
	t.Run("TFTSummonerByPUUID", func(t *testing.T) {
		summoner, err := client.TFTSummonerByPUUID(ctx, testSummonerPUUID)
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
	t.Run("TFTSummonerBySummonerID", func(t *testing.T) {
		summoner, err := client.TFTSummonerBySummonerID(ctx, testSummonerID)
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
}
