package main

import "fmt"

func main() {

	in := []int{11, 2, 3, 4, 5, 32, 62, 1, 22, 34, 4, 112, 23}
	s := 1

	fp, lp := 0, len(in)-1

	for fp <= lp {
		midp := ((lp - fp) / 2) + fp

		if in[midp] == s {
			fmt.Print(midp)
			return
		} else if in[midp] > s {
			lp = midp - 1
		} else if in[midp] < s {
			fp = midp + 1
		}
	}
}
