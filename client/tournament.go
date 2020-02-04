package client

import (
	"context"

	"github.com/anthonyrouseau/go-riot/lol"
	"github.com/anthonyrouseau/go-riot/tournament"
)

func (c *client) TournamentMatchIDs(ctx context.Context, tourCode tournament.Code) ([]lol.MatchID, error) {
	return nil, errUnimplemented
}

func (c *client) TournamentMatch(ctx context.Context, matchID lol.MatchID, tourCode tournament.Code) (*lol.Match, error) {
	return nil, errUnimplemented
}

func (c *client) TournamentCodes(ctx context.Context, tournamentID tournament.ID, options ...tournament.CodeRequestOption) ([]tournament.Code, error) {
	return nil, errUnimplemented
}

func (c *client) TournamanetCodeInfo(ctx context.Context, tourCode tournament.Code) (*tournament.CodeInfo, error) {
	return nil, errUnimplemented
}

func (c *client) UpdateTournamentCode(ctx context.Context, tourCode tournament.Code) error {
	return errUnimplemented
}

func (c *client) LobbyEvents(ctx context.Context, tourCode tournament.Code) ([]*tournament.LobbyEvents, error) {
	return nil, errUnimplemented
}

func (c *client) TournamentProvider(ctx context.Context, region tournament.Region, url string) (int32, error) {
	return -1, errUnimplemented
}

func (c *client) Tournament(ctx context.Context, provID int32, options ...tournament.RegistrationOption) (tournament.ID, error) {
	return -1, errUnimplemented
}
