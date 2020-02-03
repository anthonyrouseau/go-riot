package client

type routeKey int

//Routes constants are used as keys to map client methods to rate limiters
const (
	routeAllSummonerChampionMastery routeKey = iota
	routeSummonerChampionMastery
	routeSummonerChampionMasteryScore
	routeChampionRotations
	routeLolChallenger
	routeLolSummonerLeagueEntries
	routeLoLAllLeagueEntries
	routeLolGrandmaster
	routeLolLeague
	routeLolMaster
	routeLolMatch
	routeStatus
	routeLolAccountMatches
	routeLolMatchTimeline
	routeActiveGame
	routeFeaturedGames
	routeSummonerByAccount
	routeSummonerByName
	routeSummonerByPUUID
	routeSummonerBySummonerID
	routeTftChallenger
	routeTftSummonerLeagueEntries
	routeTftGrandmaster
	routeTftLeague
	routeTftMaster
	routeTftMatchIDs
	routeTftMatch
	routeTftSummonerByAccount
	routeTftSummonerByName
	routeTftSummonerByPUUID
	routeTftSummonerBySummonerID
	routeThirdPartyCode
	routeTournamentMatchIDs
	routeTournamentMatch
	routeTournamentCodes
	routeTournamentCodeInfo
	routeUpdateTournamentCode
	routeLobbyEvents
	routeTournamentProvider
	routeTournament
)
