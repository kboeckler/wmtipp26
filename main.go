package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofor-little/env"
	"io"
	"net/http"
)

func main() {
	env := loadEnv()

	var fr FixturesResult
	request := "/v3/fixtures?date=2024-07-01&league=4&season=2024"

	res := makeRequest(env, request, &fr)

	responseHeader := res.Header

	fmt.Println(res.StatusCode)
	fmt.Println(responseHeader)
	fmt.Println(responseHeader["X-Ratelimit-Requests-Remaining"])
	fmt.Println(fr)

	for _, item := range fr.Response {
		fmt.Println(item.Teams.Home.Name + " vs " + item.Teams.Away.Name)
	}

}

func makeRequest(env Env, request string, resultObject any) *http.Response {
	url := fmt.Sprintf("https://%s%s", env.ApiHost, request)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", env.ApiKey)
	req.Header.Add("x-rapidapi-host", env.ApiHost)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	_ = json.Unmarshal(body, resultObject)

	return res
}

func loadEnv() Env {
	if err := env.Load(".env"); err != nil {
		panic(err)
	}

	apiKey, err := env.MustGet("API_KEY")
	if err != nil {
		panic(err)
	}
	apiHost, err := env.MustGet("API_HOST")
	if err != nil {
		panic(err)
	}

	return Env{apiKey, apiHost}
}
