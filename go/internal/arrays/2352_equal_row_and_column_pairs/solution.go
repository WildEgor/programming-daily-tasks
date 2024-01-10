package equal_row_and_column_pairs

func equalPairs(grid [][]int) int {
	gL := len(grid)
	d := make(map[[200]int]int) // allocate memory for 200 ints
	temp := [200]int{}

	// copy each row
	for i := 0; i < gL; i++ {
		copy(temp[:], grid[i])
		d[temp]++
	}

	result := 0
	// copy each column and compare with mapped rows
	for i := 0; i < gL; i++ {
		for j := 0; j < gL; j++ {
			temp[j] = grid[j][i]
		}
		if v, ok := d[temp]; ok {
			result += v
		}
	}

	return result
}

var Solution = equalPairs
