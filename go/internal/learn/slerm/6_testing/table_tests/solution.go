package table_tests

// go test -v -coverprofile cover.out .
// go tool cover -html cover.out -o cover.html

func factorial(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}
