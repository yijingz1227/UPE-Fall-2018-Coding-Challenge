package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	initialization()
	solveMaze()
	end := time.Now()
	summary(start, end)
}

/*
	The function initialization() will print a line to to notify the user that maze solving has started
	It will call updateToken() to give the tokenCache an initial value
*/
func initialization() {
	fmt.Println("Mission starts.")
	updateToken()
}

/*
	The function will return a brief summary about how much time (real-time) is spent on the mission.
*/
func summary(start, end time.Time) {
	elapsed := end.Sub(start)
	remainingSeconds := math.Mod(elapsed.Seconds(), 60)
	fmt.Printf("The elapsed time is: %d minutes, %d seconds", int(elapsed.Minutes()), int(remainingSeconds))
}
