package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/anthonyrouseau/go-riot/shard"
	"github.com/anthonyrouseau/go-riot/summoner"
)

func (c *client) ThirdPartyCode(ctx context.Context, sumID summoner.ID) (string, error) {
	url := fmt.Sprintf("https://%s.%s/lol/platform/v4/third-party-code/by-summoner/%s", c.region, c.host, sumID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}
	var rec string
	code, err := c.getValue(ctx, req, routeThirdPartyCode, &rec)
	if err != nil {
		return "", err
	}
	return *code.(*string), nil
}

func (c *client) Status(ctx context.Context) (*shard.Status, error) {
	url := fmt.Sprintf("https://%s.%s/lol/status/v3/shard-data", c.region, c.host)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	status, err := c.getValue(ctx, req, routeStatus, &shard.Status{})
	if err != nil {
		return nil, err
	}
	return status.(*shard.Status), nil
}
