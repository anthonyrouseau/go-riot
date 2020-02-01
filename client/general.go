package client

import (
	"github.com/anthonyrouseau/go-riot/lol"
	"github.com/anthonyrouseau/go-riot/shard"
	"github.com/anthonyrouseau/go-riot/summoner"
)

func (c *client) ThirdPartyCode(summoner.ID) (string, error) {
	return "", errUnimplemented
}

func (c *client) ChampionRotations() (*lol.Rotation, error) {
	return nil, errUnimplemented
}

func (c *client) Status() (*shard.Status, error) {
	return nil, errUnimplemented
}

func (c *client) FeaturedGames() (*lol.FeaturedGames, error) {
	return nil, errUnimplemented
}
