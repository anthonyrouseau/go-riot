package client

import (
	"context"

	"github.com/anthonyrouseau/go-riot/shard"
	"github.com/anthonyrouseau/go-riot/summoner"
)

func (c *client) ThirdPartyCode(ctx context.Context, sumID summoner.ID) (string, error) {
	return "", errUnimplemented
}

func (c *client) Status(ctx context.Context) (*shard.Status, error) {
	return nil, errUnimplemented
}
