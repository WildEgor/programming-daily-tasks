package template

func tribonacci(n int) int {
	result := make([]int, 38+n)
	result[0] = 0
	result[1] = 1
	result[2] = 1

	for i := 0; i < n; i++ {
		result[i+3] = result[i] + result[i+1] + result[i+2]
	}

	return result[n]
}

var Solution = tribonacci
