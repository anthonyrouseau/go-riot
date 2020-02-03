package client

type routeKey int

//Routes constants are used as keys to map client methods to rate limiters
const (
	routeLolAllSummonerChampionMastery routeKey = iota
	routeLolSummonerChampionMastery
	routeLolSummonerChampionMasteryScore
	routeLolChampionRotations
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
	routeLolActiveGame
	routeLolFeaturedGames
	routeLolSummonerByAccount
	routeLolSummonerByName
	routeLolSummonerByPUUID
	routeLolSummonerBySummonerID
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
