package helpers

import (
	"encoding/json"
	"log"
	"net/http"

	"TwitchLeagueBot/src/models"
)

func ParseResponseBodyToLeagueAccount(response *http.Response) models.LeagueAccount {
	var leagueAccount models.LeagueAccount

	if err := json.NewDecoder(response.Body).Decode(&leagueAccount); err != nil {
		log.Fatal(err)
	}

	return leagueAccount
}