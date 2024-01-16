package reorder_routes_to_make_all_paths_lead_to_the_city_zero

/**
There are n cities numbered from 0 to n - 1 and n - 1 roads such that there is only one way to travel between two different cities (this network form a tree). Last year, The ministry of transport decided to orient the roads in one direction because they are too narrow.
Roads are represented by connections where connections[i] = [ai, bi] represents a road from city ai to city bi.
This year, there will be a big event in the capital (city 0), and many people want to travel to this city.
Your task consists of reorienting some roads such that each city can visit the city 0. Return the minimum number of edges changed.
It's guaranteed that each city can reach city 0 after reorder.
@link https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero
*/

// TODO
func minReorder(n int, connections [][]int) int {
	visited := make([]bool, n)
	toC := make([][]int, n)
	fromC := make([][]int, n)

	for _, pair := range connections {
		start := pair[0]
		end := pair[1]

		toC[start] = append(toC[start], end)
		fromC[end] = append(fromC[end], start)
	}

	res := 0
	mArr := []int{0}
	visited[0] = true

	for len(mArr) > 0 {
		temp := mArr[0]

		for _, city := range fromC[temp] {
			if !visited[city] {
				mArr = append(mArr, city)
				visited[city] = true
			}
		}

		for _, city := range toC[temp] {
			if !visited[city] {
				res++
				mArr = append(mArr, city)
				visited[city] = true
			}
		}
		mArr = mArr[1:]
	}
	return res
}

var Solution = minReorder
