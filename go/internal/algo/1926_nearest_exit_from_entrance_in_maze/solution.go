package nearest_exit_from_entrance_in_maze

/**
You are given an m x n matrix maze (0-indexed) with empty cells (represented as '.') and walls (represented as '+'). You are also given the entrance of the maze, where entrance = [entrancerow, entrancecol] denotes the row and column of the cell you are initially standing at.
In one step, you can move one cell up, down, left, or right. You cannot step into a cell with a wall, and you cannot step outside the maze. Your goal is to find the nearest exit from the entrance. An exit is defined as an empty cell that is at the border of the maze. The entrance does not count as an exit.
Return the number of steps in the shortest path from the entrance to the nearest exit, or -1 if no such path exists.
@link https://leetcode.com/problems/nearest-exit-from-entrance-in-maze
*/

// TODO
func nearestExit(maze [][]byte, entrance []int) int {
	rowQ := []int{entrance[0]}
	colQ := []int{entrance[1]}
	steps := 0
	layers := 1
	nextLayer := 0
	maze[entrance[0]][entrance[1]] = '+'

	explore := func(row, col int) {
		// perform check on row/col
		if row < 0 || col < 0 {
			return
		}
		if row >= len(maze) || col >= len(maze[0]) {
			return
		}
		if maze[row][col] == '+' {
			return
		}
		maze[row][col] = '+'
		rowQ = append(rowQ, row)
		colQ = append(colQ, col)
		nextLayer = nextLayer + 1
	}

	for len(rowQ) > 0 {
		row, col := rowQ[0], colQ[0]
		rowQ, colQ = rowQ[1:], colQ[1:]
		if steps > 0 &&
			(row == 0 || row == len(maze)-1 || col == 0 || col == len(maze[0])-1) {
			return steps
		}
		explore(row+1, col)
		explore(row-1, col)
		explore(row, col+1)
		explore(row, col-1)
		// perform the layers check
		layers = layers - 1
		if layers == 0 {
			layers, nextLayer = nextLayer, 0
			steps = steps + 1
		}
	}

	return -1
}

var Solution = nearestExit
