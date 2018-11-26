package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	/*URL related*/
	URL   = "http://ec2-34-216-8-43.us-west-2.compute.amazonaws.com/"
	UID   = "404786693"
	POST  = "session"
	GAME  = "game?"
	TOKEN = "token="
	/*Status strings to signal needs to update a token*/
	EXPIRED = "EXPIRED"
	NONE    = "NONE"
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

//global variable that will be updated when updateToken() is called
var tokenCache string

//A function that will update tokenCache (a global variable) when called
func updateToken() {

	body := url.Values{}
	body.Set("uid", UID)

	resp, err := http.PostForm(URL+POST, body)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		token := Token{}
		err := json.NewDecoder(resp.Body).Decode(&token)
		if err != nil {
			panic(err)
		}
		tokenCache = token.Token
	} else {
		errHttp := fmt.Errorf("%s", "Unsuccessful http request for token")
		panic(errHttp)
	}

}

/*
   A function that will return an GameState object when called.
   If the current cached token is invalid, it will update the token
   and then return a new GameState object with a follow-up GET request
*/
func getGameState() GameState {

	resp, err := http.Get(URL + GAME + TOKEN + tokenCache)

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
	} else {
		errHttp := fmt.Errorf("%s", "Unsuccessful http request for state")
		panic(errHttp)
	}

	return state
}

/*
   A function that will return a string that is either "WALL" "OUT_OF_BOUND" "SUCCESS" Or "END"
   Similarly to getGameState, if the current cached token is invalid, it will
   update the token and return a new string with a follow-up POST request
*/
func move(dir string) string {

	body := url.Values{}
	body.Set("action", dir)

	resp, err := http.PostForm(URL+GAME+TOKEN+tokenCache, body)

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
	} else {
		errHttp := fmt.Errorf("%s", "Unsuccessful http request to move")
		panic(errHttp)
	}

	return res
}
