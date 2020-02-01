package client

import (
	"github.com/anthonyrouseau/go-riot/lol"
	"github.com/anthonyrouseau/go-riot/tournament"
)

func (c *client) TournamentMatchIDs(tournament.Code) ([]lol.MatchID, error) {
	return nil, errUnimplemented
}

func (c *client) TournamentMatch(lol.MatchID, tournament.Code) (*lol.Match, error) {
	return nil, errUnimplemented
}

func (c *client) TournamentCodes() ([]tournament.Code, error) {
	return nil, errUnimplemented
}

func (c *client) TournamanetCodeInfo(tournament.Code) (*tournament.CodeInfo, error) {
	return nil, errUnimplemented
}

func (c *client) UpdateTournamentCode(tournament.Code) error {
	return errUnimplemented
}

func (c *client) LobbyEvents(tournament.Code) ([]*tournament.LobbyEvents, error) {
	return nil, errUnimplemented
}

func (c *client) TournamentProvider() (int32, error) {
	return -1, errUnimplemented
}

func (c *client) Tournament() (tournament.ID, error) {
	return -1, errUnimplemented
}
