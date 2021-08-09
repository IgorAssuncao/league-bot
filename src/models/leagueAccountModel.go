package models

type LeagueAccount struct {
	Id 				string `json:"id"`
	AccountId 		string `json:"accountId"`
	Puuid 			string `json:"puuid"`
	Name 			string `json:"name"`
	ProfileIconId 	int `json:"profileIconId"`
	RevisionDate 	int `json:"revisionDate"`
	SummonerLevel 	int `json:"summonerLevel"`
}
