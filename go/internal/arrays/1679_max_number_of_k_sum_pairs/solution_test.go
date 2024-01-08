package arrays_1679_max_number_of_k_sum_pairs

import "testing"

func Test_maxOperations_1(t *testing.T) {
	result := Solution([]int{1, 2, 3, 4}, 5)
	expect := 2

	if result != expect {
		t.Errorf("Expected %d, got %d", expect, result)
	}
}

func Test_maxOperations_2(t *testing.T) {
	result := Solution([]int{3, 1, 3, 4, 3}, 6)
	expect := 1

	if result != expect {
		t.Errorf("Expected %d, got %d", expect, result)
	}
}
