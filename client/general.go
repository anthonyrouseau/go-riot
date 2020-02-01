package client

import (
	"github.com/anthonyrouseau/go-riot/shard"
	"github.com/anthonyrouseau/go-riot/summoner"
)

func (c *client) ThirdPartyCode(summoner.ID) (string, error) {
	return "", errUnimplemented
}

func (c *client) Status() (*shard.Status, error) {
	return nil, errUnimplemented
}
