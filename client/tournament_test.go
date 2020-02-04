package client

import (
	"context"
	"github.com/anthonyrouseau/go-riot/tournament"
	"testing"
)

func TestTournamentMethods(t *testing.T) {
	const (
		testTournamentCode = ""
		testMatchID        = 0
		testRegion         = tournament.NA
		testURL            = ""
		testProviderID     = 0
		testTournamentID   = 0
	)
	ctx := context.Background()
	client, err := NewClient(testAPIKey, SetVariant(devClient))
	if err != nil {
		t.Error(err)
	}
	t.Run("TournamentMatchIds", func(t *testing.T) {
		_, err := client.TournamentMatchIDs(ctx, testTournamentCode)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("TournamentMatch", func(t *testing.T) {
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
		_, err := client.TournamentCodes(ctx, testTournamentID)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("TournamentCodeInfo", func(t *testing.T) {
		codeInfo, err := client.TournamanetCodeInfo(ctx, testTournamentCode)
		if err != nil {
			t.Error(err)
		}
		if codeInfo == nil {
			t.Error("Code Info was nil value")
		}
	})
	t.Run("UpdateTournamentCode", func(t *testing.T) {
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
