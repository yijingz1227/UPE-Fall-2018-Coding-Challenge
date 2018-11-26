package main

import "fmt"

const (
	UP            = "UP"
	DOWN          = "DOWN"
	LEFT          = "LEFT"
	RIGHT         = "RIGHT"
	OUT_OF_BOUNDS = "OUT_OF_BOUNDS"
	WALL          = "WALL"
	END           = "END"
	FINISHED      = "FINISHED"
	REACHED       = 1
	TIMEOUT       = 3
)

/*
	An infinite loop that will automatically restarts the maze-solving process if trySolveMaze() returns false
	The loop will stop if the mission is successful (maze is solve)
*/
func solveMaze() {
	for {
		if trySolveMaze() == true {
			fmt.Println("Mission Successful.")
			break
		}
		updateToken()
	}
}

/*
	A loop that will try to solve the maze
	it returns true if it completes all 12 levels
	it returns false if the maze gets expired in the process of solving the maze
*/
func trySolveMaze() bool {

	for getGameState().Status != FINISHED {
		//the y coordinate signifies rows in the maze matrix
		startingRow := getGameState().CurrentLocation[1]
		//the x coordinate signifies cols in the maze matrix
		startingCol := getGameState().CurrentLocation[0]
		//a boolean 2D slice to mark visited entries in the maze matrix
		visited := make([][]bool, getGameState().MazeSize[1])
		for row := range visited {
			visited[row] = make([]bool, getGameState().MazeSize[0])
		}
		//the starting position is always visited by default
		visited[startingRow][startingCol] = true
		//dfs into the maze
		if dfs(startingRow, startingCol, visited) == TIMEOUT {
			return false
		}
	}

	return true
}

/*
	A dfs algorithm that does the following:
		* At the given (row,col) position
			*Try to Go UP if it is not where you come from
            	*Case 1: you reach the exit, return REACHED to celebrate your success
				*Case 2: you realize the maze has expired, return TIMEOUT to restart the game
				*Case 3: the move is successful, but you are not quite there yet
					* mark your current spot to be visited, and try along this path
						* if this path leads to an dead end, maybe it's best to try another direction
				*Case 4: you bumped into a wall or walked out of bounds, hmm, maybe it's best to
						 choose another direction to goto
			*Try to Go DOWN, LEFT, RIGHT with similar logic
			*If none of the options are valid, you are in a dead end, return 0 to backtrack
*/

func dfs(row, col int, visited [][]bool) int {
	if inBound(row-1, col, visited) && !visited[row-1][col] {
		res := move(UP)

		if res == END {
			return REACHED
		} else if res == EXPIRED {
			return TIMEOUT
		}

		if validMove(res) {
			visited[row-1][col] = true
			echo := dfs(row-1, col, visited)
			if echo == REACHED || echo == TIMEOUT {
				return echo
			}
			move(DOWN)
		}

	}

	if inBound(row+1, col, visited) && !visited[row+1][col] {
		res := move(DOWN)

		if res == END {
			return REACHED
		} else if res == EXPIRED {
			return TIMEOUT
		}

		if validMove(res) {
			visited[row+1][col] = true
			echo := dfs(row+1, col, visited)
			if echo == REACHED || echo == TIMEOUT {
				return echo
			}
			move(UP)
		}

	}

	if inBound(row, col-1, visited) && !visited[row][col-1] {
		res := move(LEFT)

		if res == END {
			return REACHED
		} else if res == EXPIRED {
			return TIMEOUT
		}

		if validMove(res) {
			visited[row][col-1] = true
			echo := dfs(row, col-1, visited)
			if echo == REACHED || echo == TIMEOUT {
				return echo
			}
			move(RIGHT)
		}

	}

	if inBound(row, col+1, visited) && !visited[row][col+1] {
		res := move(RIGHT)

		if res == END {
			return REACHED
		} else if res == EXPIRED {
			return TIMEOUT
		}

		if validMove(res) {
			visited[row][col+1] = true
			echo := dfs(row, col+1, visited)
			if echo == REACHED || echo == TIMEOUT {
				return echo
			}
			move(LEFT)
		}
	}

	return 0
}

/*
	A function that returns a bool that takes a string and determines if a move is valid
*/
func validMove(res string) bool {
	if res == WALL || res == OUT_OF_BOUNDS {
		return false
	}
	return true
}

/*
	A function that returns a bool that determines if the slice access is inBound
*/

func inBound(row, col int, visited [][]bool) bool {
	if 0 <= row && row < len(visited) && 0 <= col && col < len(visited[0]) {
		return true
	}
	return false
}
