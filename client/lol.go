package client

import (
	"github.com/anthonyrouseau/go-riot/account"
	"github.com/anthonyrouseau/go-riot/lol"
	"github.com/anthonyrouseau/go-riot/queue"
	"github.com/anthonyrouseau/go-riot/summoner"
)

func (c *client) LOLChallenger(queue.Name) (*lol.LeagueInfo, error) {
	return nil, errUnimplemented
}

func (c *client) LOLSummonerLeagueEntries(summoner.ID) ([]*lol.LeagueEntry, error) {
	return nil, errUnimplemented
}

func (c *client) LOLAllLeagueEntries(queue.Name, lol.Tier, lol.Division) ([]*lol.LeagueEntry, error) {
	return nil, errUnimplemented
}

func (c *client) LOLGrandmaster(queue.Name) (*lol.LeagueInfo, error) {
	return nil, errUnimplemented
}

func (c *client) LOLLeague(lol.LeagueID) (*lol.LeagueInfo, error) {
	return nil, errUnimplemented
}

func (c *client) LOLMaster(queue.Name) (*lol.LeagueInfo, error) {
	return nil, errUnimplemented
}

func (c *client) LOLMatch(lol.MatchID) (*lol.Match, error) {
	return nil, errUnimplemented
}

func (c *client) LOLAccountMatches(account.ID) (*lol.MatchList, error) {
	return nil, errUnimplemented
}

func (c *client) LOLMatchTimeline(lol.MatchID) (*lol.MatchTimeline, error) {
	return nil, errUnimplemented
}
