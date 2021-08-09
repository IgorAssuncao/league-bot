package helpers

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetApiKeyFromEnv() string {
	riotApiKey := os.Getenv("RIOT_API_KEY")
	return riotApiKey
}

func MountEndpoint(endpointKeyword string, encryptedSummonerId string, championName string) string {
	riotBREndpoint := "https://br1.api.riotgames.com"

	switch endpointKeyword {
		case "summoner":
			endpoint := riotBREndpoint + "/lol/summoner/v4/summoners/by-name/Stealthy Carry"
			return endpoint

		case "championMastery":
			championId, err := GetChampionId(championName)
			if err != nil {
				log.Fatal(err)
			}
			endpoint := riotBREndpoint + "/lol/champion-mastery/v4/champion-masteries/by-summoner/" + encryptedSummonerId + "/by-champion/" + string(championId)
			return endpoint

		default:
			return riotBREndpoint
	}
}

func GetChampionId(name string) (string, error) {
	championsList := MountChampionsList()

	for _, champion := range championsList {
		if strings.ToLower(champion.Name) == strings.ToLower(name) {
			return strconv.FormatInt(int64(champion.Id), 10), nil
		}
	}

	return "", errors.New("id not found")
}