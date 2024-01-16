package nearest_exit_from_entrance_in_maze

/**
You are given an m x n grid where each cell can have one of three values:
0 representing an empty cell,
1 representing a fresh orange, or
2 representing a rotten orange.
Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.
Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.
@link https://leetcode.com/problems/rotting-oranges
*/

// TODO
func orangesRotting(grid [][]int) int {
	q := [][2]int{}
	fo := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				q = append(q, [2]int{i, j})
			}
			if grid[i][j] == 1 {
				fo += 1
			}

		}
	}
	if fo == 0 {
		return 0
	}
	c := bfs(grid, q)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return -1
			}

		}
	}
	return c - 1
}

func bfs(grid [][]int, q [][2]int) int {
	count := 0
	row := len(grid)
	column := len(grid[0])
	for len(q) > 0 {
		for _, v := range q {

			i, j := v[0], v[1]
			grid[i][j] = 2

			if i-1 >= 0 && grid[i-1][j] == 1 {
				grid[i-1][j] = 2
				q = append(q, [2]int{i - 1, j})
			}
			if i+1 < row && grid[i+1][j] == 1 {
				grid[i+1][j] = 2
				q = append(q, [2]int{i + 1, j})

			}
			if j-1 >= 0 && grid[i][j-1] == 1 {
				grid[i][j-1] = 2
				q = append(q, [2]int{i, j - 1})

			}
			if j+1 < column && grid[i][j+1] == 1 {
				grid[i][j+1] = 2
				q = append(q, [2]int{i, j + 1})
			}
			q = q[1:]

		}
		count += 1

	}
	return count

}

var Solution = orangesRotting
