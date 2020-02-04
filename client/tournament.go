package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/anthonyrouseau/go-riot/lol"
	"github.com/anthonyrouseau/go-riot/tournament"
)

//Can't test route without access to tournament api
func (c *client) TournamentMatchIDs(ctx context.Context, tourCode tournament.Code) ([]lol.MatchID, error) {
	/*url := fmt.Sprintf("https://%s.%s/lol/match/v4/matches/by-tournament-code/%s/ids", c.region, c.host, tourCode)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	var rec []lol.MatchID
	matchIDs, err := c.getValue(ctx, req, routeTournamentMatchIDs, &rec)
	if err != nil {
		return nil, err
	}
	return *matchIDs.(*[]lol.MatchID), nil*/
	return nil, errUnimplemented
}

//Can't test route without access to tournament api
func (c *client) TournamentMatch(ctx context.Context, matchID lol.MatchID, tourCode tournament.Code) (*lol.Match, error) {
	/*url := fmt.Sprintf("https://%s.%s/lol/match/v4/matches/%d/by-tournament-code/%s", c.region, c.host, matchID, tourCode)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	match, err := c.getValue(ctx, req, routeTournamentMatch, &lol.Match{})
	if err != nil {
		return nil, err
	}
	return match.(*lol.Match), nil*/
	return nil, errUnimplemented
}

func (c *client) TournamentCodes(ctx context.Context, tournamentID tournament.ID, body *tournament.CodeRequestBody, options ...tournament.CodeRequestOption) ([]tournament.Code, error) {
	if body == nil {
		return nil, errors.New("CodeRequestBody can not be nil pointer")
	}
	codeRequest := &tournament.CodeRequest{Count: 1}
	for _, opt := range options {
		opt(codeRequest)
	}
	//url := fmt.Sprintf("https://%s.%s/lol/tournament/v4/codes?count=%d&tournamentId=%d", regionalRouting(c.region), c.host, codeRequest.Count, tournamentID)
	//testing url
	url := fmt.Sprintf("https://%s.%s/lol/tournament-stub/v4/codes?count=%d&tournamentId=%d", regionalRouting(c.region), c.host, codeRequest.Count, tournamentID)
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyBuffer := bytes.NewBuffer(bodyBytes)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyBuffer)
	if err != nil {
		return nil, err
	}
	var rec []tournament.Code
	codes, err := c.getValue(ctx, req, routeTournamentCodes, &rec)
	if err != nil {
		return nil, err
	}
	return *codes.(*[]tournament.Code), nil
}

//Can't test route without access to tournament api
func (c *client) TournamentCodeInfo(ctx context.Context, tourCode tournament.Code) (*tournament.CodeInfo, error) {
	/*url := fmt.Sprintf("https://%s.%s/lol/tournament/v4/codes/%s", regionalRouting(c.region), c.host, tourCode)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	codeInfo, err := c.getValue(ctx, req, routeTournamentCodeInfo, &tournament.CodeInfo{})
	if err != nil {
		return nil, err
	}
	return codeInfo.(*tournament.CodeInfo), nil*/
	return nil, errUnimplemented
}

//Unclear how this route should work and can't test without access to tournament api
func (c *client) UpdateTournamentCode(ctx context.Context, tourCode tournament.Code) error {
	return errUnimplemented
}

func (c *client) LobbyEvents(ctx context.Context, tourCode tournament.Code) (*tournament.LobbyEvents, error) {
	//url := fmt.Sprintf("https://%s.%s/lol/tournament/v4/lobby-events/by-code/%s", regionalRouting(c.region), c.host, tourCode)
	//testing url
	url := fmt.Sprintf("https://%s.%s/lol/tournament-stub/v4/lobby-events/by-code/%s", regionalRouting(c.region), c.host, tourCode)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	lobbyEvents, err := c.getValue(ctx, req, routeLobbyEvents, &tournament.LobbyEvents{})
	if err != nil {
		return nil, err
	}
	return lobbyEvents.(*tournament.LobbyEvents), nil
}

func (c *client) TournamentProvider(ctx context.Context, region tournament.Region, cb string) (int32, error) {
	//url := fmt.Sprintf("https://%s.%s/lol/tournament/v4/providers", regionalRouting(c.region), c.host)
	//testing url
	url := fmt.Sprintf("https://%s.%s/lol/tournament-stub/v4/providers", regionalRouting(c.region), c.host)
	bodyBytes, err := json.Marshal(&tournament.ProviderRegistration{
		Region: region,
		URL:    cb,
	})
	if err != nil {
		return 0, err
	}
	bodyBuffer := bytes.NewBuffer(bodyBytes)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyBuffer)
	if err != nil {
		return 0, err
	}
	var rec int32
	providerID, err := c.getValue(ctx, req, routeTournamentProvider, &rec)
	if err != nil {
		return 0, err
	}
	return *providerID.(*int32), nil
}

func (c *client) Tournament(ctx context.Context, provID int32, options ...tournament.RegistrationOption) (tournament.ID, error) {
	//url := fmt.Sprintf("https://%s.%s/lol/tournament/v4/tournaments", regionalRouting(c.region), c.host)
	//testing url
	url := fmt.Sprintf("https://%s.%s/lol/tournament-stub/v4/tournaments", regionalRouting(c.region), c.host)
	registration := &tournament.Registration{
		Name:       nil,
		ProviderID: provID,
	}
	for _, opt := range options {
		opt(registration)
	}
	bodyBytes, err := json.Marshal(registration)
	if err != nil {
		return 0, err
	}
	bodyBuffer := bytes.NewBuffer(bodyBytes)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyBuffer)
	if err != nil {
		return 0, err
	}
	var rec tournament.ID
	tourID, err := c.getValue(ctx, req, routeTournament, &rec)
	if err != nil {
		return 0, err
	}
	return *tourID.(*tournament.ID), nil
}
