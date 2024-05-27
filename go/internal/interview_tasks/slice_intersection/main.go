package main

import "fmt"

func intersection(po, pt []int) []int {
	r := make([]int, 0)

	sm := make(map[int]int)
	for _, el := range po {
		if _, ok := sm[el]; !ok {
			sm[el] += 1
		}
	}

	for _, el := range pt {
		if counter, ok := sm[el]; ok && counter > 0 {
			counter -= 1
			r = append(r, el)
		}
	}

	return r
}

func main() {
	po := []int{23, 3, 1, 2}
	pt := []int{6, 2, 4, 23}

	r := intersection(po, pt)

	fmt.Println(r)
}
