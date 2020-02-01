package client

import (
	"github.com/anthonyrouseau/go-riot/lol"
	"github.com/anthonyrouseau/go-riot/tournament"
)

func (c *client) TournamentMatchIDs(tournament.Code) ([]lol.MatchID, error) {

}

func (c *client) TournamentMatch(lol.MatchID, tournament.Code) (*lol.Match, error) {

}

func (c *client) TournamentCodes() ([]tournament.Code, error) {

}

func (c *client) TournamanetCodeInfo(tournament.Code) (*tournament.CodeInfo, error) {

}

func (c *client) UpdateTournamentCode(tournament.Code) {

}

func (c *client) LobbyEvents(tournament.Code) ([]tournament.LobbyEvents, error) {

}

func (c *client) TournamentProvider() (int32, error) {

}

func (c *client) Tournament() (tournament.ID, error) {

}
