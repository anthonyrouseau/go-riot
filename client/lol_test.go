package client

import (
	"context"
	"github.com/anthonyrouseau/go-riot/queue"
	"testing"
)

const (
	testAPIKey = "RGAPI-a348803e-d2f9-42c4-a442-c8a5b71dd12d"
)

func TestLOLChallenger(t *testing.T) {
	ctx := context.Background()
	client, err := NewDevClient(testAPIKey)
	if err != nil {
		t.Error(err)
	}
	leagueInfo, err := client.LOLChallenger(ctx, queue.RankedSolo5x5)
	if err != nil {
		t.Error(err)
	}
	if leagueInfo == nil {
		t.Error("Returned value is nil")
	}
}
