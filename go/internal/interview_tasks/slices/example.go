package main

import "fmt"

func main() {
	var ms []int

	for i := 0; i < 10; i++ {
		ms = append(ms, i)
		fmt.Println(ms[i])
		fmt.Printf("len: %d, cap: %d\n", len(ms), cap(ms))

		test(ms)
	}

	for i, v := range ms {
		fmt.Println(i, v)
	}
}

func test(arr []int) {
	arr = append(arr, 1)
}
