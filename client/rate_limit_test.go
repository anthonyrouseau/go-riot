package client

import (
	"context"
	"net/http"
	"testing"

	"golang.org/x/time/rate"
)

func TestRateLimiting(t *testing.T) {
	testRoutes := []routeKey{
		routeLoLAllLeagueEntries,
		routeLobbyEvents,
		routeLolActiveGame,
		routeLolAccountMatches,
	}
	ctx := context.Background()
	t.Run("Dev Client", func(t *testing.T) {
		expectedPrimary := rate.NewLimiter(20, 20)
		expectedSecondary := rate.NewLimiter(5/6, 100)
		c := &client{
			variant:  DevClient,
			apiKey:   testAPIKey,
			client:   &http.Client{},
			region:   "na1",
			limiters: make(map[routeKey][]*rate.Limiter),
			host:     "api.riotgames.com",
		}
		for _, lim := range testRoutes {
			err := c.rateLimit(ctx, lim)
			if err != nil {
				t.Error(err)
			}
			primaryLim := c.limiters[lim][0]
			secondaryLim := c.limiters[lim][1]
			if primaryLim.Limit() != expectedPrimary.Limit() {
				t.Errorf("Primary limit expected to be %f but got %f", expectedPrimary.Limit(), primaryLim.Limit())
			}
			if secondaryLim.Limit() != 5/6 {
				t.Errorf("Secondary limit expected to be %f but got %f", expectedSecondary.Limit(), secondaryLim.Limit())
			}
			if primaryLim.Burst() != 20 {
				t.Errorf("Primary burst expected to be %d but got %d", expectedPrimary.Burst(), primaryLim.Burst())
			}
			if secondaryLim.Burst() != 100 {
				t.Errorf("Secondary burst expected to be %d but got %d", expectedSecondary.Burst(), secondaryLim.Burst())
			}
		}
	})
}
