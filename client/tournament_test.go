package client

import (
	"context"
	"testing"

	"github.com/anthonyrouseau/go-riot/tournament"
)

func TestTournamentMethods(t *testing.T) {
	const (
		testTournamentCode = "NA1908-TOURNAMENTCODE0001"
		testMatchID        = 0
		testRegion         = tournament.NA
		testURL            = "http://localhost/cb:80"
		testProviderID     = 168
		testTournamentID   = 1908
		testSummonerID     = "c0n56ouT0eGJLaVy8Sbfe628zfBkRbaKZZwByHVDQik"
	)
	ctx := context.Background()
	client, err := NewClient(testAPIKey, SetVariant(devClient))
	if err != nil {
		t.Error(err)
	}
	t.Run("TournamentMatchIds", func(t *testing.T) {
		t.Skip("skipping test until tournament route access")
		_, err := client.TournamentMatchIDs(ctx, testTournamentCode)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("TournamentMatch", func(t *testing.T) {
		t.Skip("skipping test until tournament route access")
		match, err := client.TournamentMatch(ctx, testMatchID, testTournamentCode)
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
	t.Run("TournamentCodes", func(t *testing.T) {
		testTournamentCodeRequestBody := &tournament.CodeRequestBody{
			SpectatorType:      tournament.None,
			TeamSize:           5,
			PickType:           tournament.TournamentDraft,
			AllowedSummonerIDs: nil,
			MapType:            tournament.SummonersRift,
			Metadata:           "metadata",
		}
		_, err := client.TournamentCodes(ctx, testTournamentID, testTournamentCodeRequestBody)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("TournamentCodeInfo", func(t *testing.T) {
		t.Skip("skipping test until tournament route access")
		codeInfo, err := client.TournamentCodeInfo(ctx, testTournamentCode)
		if err != nil {
			t.Error(err)
		}
		if codeInfo == nil {
			t.Error("Code Info was nil value")
		}
	})
	t.Run("UpdateTournamentCode", func(t *testing.T) {
		t.Skip("skipping test until tournament route access")
		err := client.UpdateTournamentCode(ctx, testTournamentCode)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("LobbyEvents", func(t *testing.T) {
		events, err := client.LobbyEvents(ctx, testTournamentCode)
		if err != nil {
			t.Error(err)
		}
		if events == nil {
			t.Error("Lobby Events was nil value")
		}
	})
	t.Run("TournamentProvider", func(t *testing.T) {
		_, err := client.TournamentProvider(ctx, testRegion, testURL)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("Tournament", func(t *testing.T) {
		_, err := client.Tournament(ctx, testProviderID)
		if err != nil {
			t.Error(err)
		}
	})
}
