package client

import (
	"github.com/anthonyrouseau/go-riot/account"
	"github.com/anthonyrouseau/go-riot/lol"
	"github.com/anthonyrouseau/go-riot/queue"
	"github.com/anthonyrouseau/go-riot/summoner"
)

func (c *client) LOLChallenger(queue.Name) (lol.LeagueInfo, error) {

}

func (c *client) LOLSummonerLeagueEntries(summoner.ID) ([]*lol.LeagueEntry, error) {

}

func (c *client) LOLAllLeagueEntries(queue.Name, lol.Tier, lol.Division) ([]*lol.LeagueEntry, error) {

}

func (c *client) LOLGrandmaster(queue.Name) (lol.LeagueInfo, error) {

}

func (c *client) LOLLeague(lol.LeagueID) (lol.LeagueInfo, error) {

}

func (c *client) LOLMaster(queue.Name) (lol.LeagueInfo, error) {

}

func (c *client) LOLMatch(lol.MatchID) (*lol.Match, error) {

}

func (c *client) LOLAccountMatches(account.ID) (*lol.MatchList, error) {

}

func (c *client) LOLMatchTimeline(lol.MatchID) (*lol.MatchTimeline, error) {

}
