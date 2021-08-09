package main

import (
	"TwitchLeagueBot/src/helpers"
	"fmt"
	"io/ioutil"
	"log"

	"TwitchLeagueBot/src/http"
)

func main() {
	resp, err := http.MakeRequest("get", helpers.MountEndpoint("summoner", "", ""))
	if err != nil {
		log.Fatal(err)
	}

	leagueAcc := helpers.ParseResponseBodyToLeagueAccount(resp)
	championMastery, err := http.MakeRequest(
		"get",
		helpers.MountEndpoint(
			"championMastery",
			leagueAcc.Id,
			"lucian"),
		)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(championMastery.Body)
	sb := string(body)
	fmt.Print(sb)
}
