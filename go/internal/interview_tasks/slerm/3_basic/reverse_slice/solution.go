package reverse_slice

/**
Напишите функцию для реверса слайса (вернуть слайс в обратном порядке)
*/

func reverse(slice []int) []int {
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
	return slice
}
