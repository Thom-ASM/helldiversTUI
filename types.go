package main

type Planet struct {
	Name          string
	Hash          int
	Index         int
	Initial_owner string
	Max_health    int
}

type SinglePlanet struct {
	Health           int
	Liberation       float32
	Players          int
	Regen_per_second int
}

type State struct {
	AllPlanets       []Planet
	FilteredPlanets  []Planet
	SelectedIdx      int
	FactionFilterIdx int
	ActivePlanet     SinglePlanet
	PaginationIdx    int
}

type AllPlanetsMsg []Planet

type SinglePlanetMsg SinglePlanet

func initialModel() State {
	return State{
		AllPlanets:       make([]Planet, 0),
		SelectedIdx:      260,
		FactionFilterIdx: 0,
		ActivePlanet:     SinglePlanet{},
		PaginationIdx:    0,
	}
}
