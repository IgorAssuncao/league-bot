package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"TwitchLeagueBot/src/models"
)

func MountChampionsList() []models.Champion {
	jsonFile, err := os.Open("./src/data/champions.json")
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var championsList []models.Champion
	if err := json.Unmarshal(data, &championsList); err != nil {
		log.Fatal(err)
	}

	return championsList
}

