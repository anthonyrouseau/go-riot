package client

import (
	"context"
	"testing"
)

func TestGeneralMethods(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient(testAPIKey, SetVariant(devClient))
	if err != nil {
		t.Error(err)
	}
	t.Run("ThirdPartyCode", func(t *testing.T) {
		_, err := client.ThirdPartyCode(ctx, testSummonerID)
		if err != nil {
			if err.Error() != errNotFound.Error() {
				t.Error(err)
			}
		}
	})
	t.Run("Status", func(t *testing.T) {
		status, err := client.Status(ctx)
		if err != nil {
			t.Error(err)
		}
		if status == nil {
			t.Error("Status was nil value")
		}
	})
}
