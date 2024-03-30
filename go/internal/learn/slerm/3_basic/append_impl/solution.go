package append_impl

func appendTo(arr []int, values ...int) []int {
	l := len(arr) + len(values)

	if l > cap(arr) {
		c := l
		if c < 2*len(arr) {
			c = 2 * len(arr)
		}
		ns := make([]int, l, c)

		copy(ns, arr)
		copy(ns[len(arr):], values)

		return ns
	}

	arr = arr[:l]
	copy(arr[len(arr)-len(values):], values)

	return arr
}
