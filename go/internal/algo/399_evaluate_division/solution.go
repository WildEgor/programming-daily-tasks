package evaluate_division

/**
You are given an array of variable pairs equations and an array of real numbers values, where equations[i] = [Ai, Bi] and values[i] represent the equation Ai / Bi = values[i]. Each Ai or Bi is a string that represents a single variable.
You are also given some queries, where queries[j] = [Cj, Dj] represents the jth query where you must find the answer for Cj / Dj = ?.
Return the answers to all queries. If a single answer cannot be determined, return -1.0.
Note: The input is always valid. You may assume that evaluating the queries will not result in division by zero and that there is no contradiction.
Note: The variables that do not occur in the list of equations are undefined, so the answer cannot be determined for them.
@link https://leetcode.com/problems/evaluate-division
*/

// TODO
type key struct {
	a, b string
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	hash := map[key]float64{}
	for i, equation := range equations {
		a, b := equation[0], equation[1]
		hash[key{a, b}] = values[i]
		hash[key{b, a}] = 1 / values[i]
		hash[key{a, a}] = 1.0
		hash[key{b, b}] = 1.0
	}
	for k, v := range hash {
		dfs(k.a, k.b, v, hash)
	}
	var result []float64
	for _, query := range queries {
		v, ok := hash[key{query[0], query[1]}]
		if ok {
			result = append(result, v)
		} else {
			result = append(result, -1.0)
		}
	}
	return result
}

func dfs(a, b string, v float64, hash map[key]float64) {
	for k, vk := range hash {
		// a/b * b(k.a)/c(k.b)
		if b == k.a && a != k.b {
			newKey := key{a, k.b}
			_, ok := hash[newKey]
			if !ok {
				hash[newKey] = v * vk
				dfs(a, k.b, v*vk, hash)
			}
		}
	}
}

var Solution = calcEquation
