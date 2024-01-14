package template

/**
There are n cities. Some of them are connected, while some are not. If city a is connected directly with city b, and city b is connected directly with city c, then city a is connected indirectly with city c.
A province is a group of directly or indirectly connected cities and no other cities outside of the group.
You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected, and isConnected[i][j] = 0 otherwise.
Return the total number of provinces.
@link https://leetcode.com/problems/number-of-provinces/
*/

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected)) // keep track of visited cities
	count := 0                                // number of provinces
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] { // if city is not visited yet
			dfs(isConnected, visited, i)
			count++
		}
	}
	return count
}

func dfs(isConnected [][]int, visited []bool, i int) {
	visited[i] = true // mark city as visited
	for j := 0; j < len(isConnected); j++ {
		if isConnected[i][j] == 1 && !visited[j] { // if city is connected and not visited yet
			dfs(isConnected, visited, j)
		}
	}
}

var Solution = findCircleNum
