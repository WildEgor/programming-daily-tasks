package template

import "github.com/wildegor/daily_tasks/go/internal/data_structures/hash_set"

func singleNumber(nums []int) int {
	set := hash_set.NewHashSet()

	for _, num := range nums {
		if set.Has(num) {
			set.Remove(num)
		} else {
			set.Add(num)
		}
	}

	return set.First().(int)
}

func singleNumber2(nums []int) int {
	p := 0
	for _, num := range nums {
		/**
		The reason this works is because of the properties of the XOR operation.
		In particular, for any number x, x XOR x equals 0, and x XOR 0 equals x.
		Also, XOR is associative and commutative, which means the order
		of the operations does not matter.
		*/
		p ^= num
	}

	return p
}

var Solution = singleNumber2
