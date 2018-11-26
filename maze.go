package main

const (
	UP            = "UP"
	DOWN          = "DOWN"
	LEFT          = "LEFT"
	RIGHT         = "RIGHT"
	OUT_OF_BOUNDS = "OUT_OF_BOUNDS"
	WALL          = "WALL"
	SUCCESS       = "SUCCESS"
	END           = "END"
	FINISHED      = "FINISHED"
)

func solveMaze() {

	for getGameState().Status != FINISHED {
		startingRow := getGameState().CurrentLocation[1]
		startingCol := getGameState().CurrentLocation[0]
		visited := make([][]bool, getGameState().MazeSize[1])
		for row := range visited {
			visited[row] = make([]bool, getGameState().MazeSize[0])
		}
		visited[startingRow][startingCol] = true
		dfs(startingRow, startingCol, visited)
	}

}

func dfs(row, col int, visited [][]bool) int {
	if inBound(row-1, col, visited) && !visited[row-1][col] {
		res := move(UP)
		if res == END {
			return 1
		}
		if validMove(res) {
			visited[row-1][col] = true
			if dfs(row-1, col, visited) == 1 {
				return 1
			}
			move(DOWN)
		}

	}

	if inBound(row+1, col, visited) && !visited[row+1][col] {
		res := move(DOWN)
		if res == END {
			return 1
		}
		if validMove(res) {
			visited[row+1][col] = true
			if dfs(row+1, col, visited) == 1 {
				return 1
			}
			move(UP)
		}

	}

	if inBound(row, col-1, visited) && !visited[row][col-1] {
		res := move(LEFT)
		if res == END {
			return 1
		}
		if validMove(res) {
			visited[row][col-1] = true
			if dfs(row, col-1, visited) == 1 {
				return 1
			}
			move(RIGHT)
		}

	}

	if inBound(row, col+1, visited) && !visited[row][col+1] {
		res := move(RIGHT)
		if res == END {
			return 1
		}
		if validMove(res) {
			visited[row][col+1] = true
			if dfs(row, col+1, visited) == 1 {
				return 1
			}
			move(LEFT)
		}
	}

	return 0
}

func validMove(res string) bool {
	if res == WALL || res == OUT_OF_BOUNDS {
		return false
	}
	return true
}

func inBound(row, col int, visited [][]bool) bool {
	if 0 <= row && row < len(visited) && 0 <= col && col < len(visited[0]) {
		return true
	}
	return false
}
