package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
)

func makeHttpRequest[T any](httpClient *http.Client, url string) T {
	res, err := httpClient.Get(url)

	if err != nil {

		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	var AllPlanets T

	json.Unmarshal(body, &AllPlanets)

	return AllPlanets

}

func fetchAllPlanets_msg() tea.Msg {
	res := makeHttpRequest[[]Planet](client, "https://helldivers-2.fly.dev/api/801/planets")
	return AllPlanetsMsg(res)
}

func fetchSinglePlanet_msg(id int) tea.Msg {
	res := makeHttpRequest[SinglePlanet](client, fmt.Sprintf("https://helldivers-2.fly.dev/api/801/planets/%d/status", id))
	return SinglePlanetMsg(res)
}

func fetchPlanetInfo(id int) tea.Cmd {

	return func() tea.Msg {
		return fetchSinglePlanet_msg(id)
	}
}
