package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	BASEAPI = "http://ec2-34-216-8-43.us-west-2.compute.amazonaws.com/"
	UID     = "404786693"
	POST    = "session"
	GAME    = "game?"
	TOKEN   = "token="
)

type GameState struct {
	MazeSize        []int  `json:"maze_size"`        //width, height
	CurrentLocation []int  `json:"current_location"` //x,y
	Status          string `json:"status"`
	LevelsCompleted int    `json:"levels_completed"`
	TotalLevels     int    `json:"total_levels"`
}

type Token struct {
	Token string `json:"token"`
}

type Result struct {
	Result string `json:"result"`
}

func getToken() string {

	body := url.Values{}
	body.Set("uid", UID)

	resp, err := http.PostForm(BASEAPI+POST, body)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var tokenString string

	if resp.StatusCode == http.StatusOK {
		token := Token{}
		err := json.NewDecoder(resp.Body).Decode(&token)
		if err != nil {
			panic(err)
		}
		tokenString = token.Token
	}
	return tokenString
}

func getGameState() GameState {

	resp, err := http.Get(BASEAPI + GAME + TOKEN + getToken())

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var state GameState

	if resp.StatusCode == http.StatusOK {
		err := json.NewDecoder(resp.Body).Decode(&state)
		if err != nil {
			panic(err)
		}
	}
	return state
}

func move(dir string) string {

	body := url.Values{}
	body.Set("action", dir)

	resp, err := http.PostForm(BASEAPI+GAME+TOKEN+getToken(), body)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var res string

	if resp.StatusCode == http.StatusOK {
		var result Result
		err := json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			panic(err)
		}
		res = result.Result
	}

	return res
}
