package http

import (
	"log"
	"net/http"
	"time"

	"TwitchLeagueBot/src/helpers"
)

func createHttpClient() *http.Client {
	client := http.Client {
		Timeout: time.Second * 10,
	}
	return &client
}

func createRequest(method string, url string) (req *http.Request, err error) {
	return http.NewRequest(method, url, nil)
}

func MakeRequest(method string, url string) (*http.Response, error) {
	httpClient := createHttpClient()
	req, err := createRequest(method, url)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Riot-Token", helpers.GetApiKeyFromEnv())
	resp, err := httpClient.Do(req)
	//defer resp.Body.Close()
	return resp, err
}
