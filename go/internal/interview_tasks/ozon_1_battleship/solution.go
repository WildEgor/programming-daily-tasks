package main

import "fmt"

/**
Вы участвуете в разработке подсистемы проверки поля для игры <<Морской бой>>. Вам требуется написать проверку корректности количества кораблей на поле, учитывая их размеры. Напомним, что на поле должны быть:
∙четыре однопалубных корабля,
∙три двухпалубных корабля,
∙два трёхпалубных корабля,
∙один четырёхпалубный корабль.

Вам заданы 10 целых чисел от 1 до 4 Проверьте, что заданные размеры соответствуют требованиям выше.

Входные данные
5
2 1 3 1 2 3 1 1 4 2
1 1 1 2 2 2 3 3 3 4
1 1 1 1 2 2 2 3 3 4
4 3 3 2 2 2 1 1 1 1
4 4 4 4 4 4 4 4 4 4

Выходные данные
YES
NO
YES
YES
NO
*/

func isValidShipsConfiguration(sizes []int) bool {
	counts := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
	}

	for _, size := range sizes {
		counts[size]++
	}

	return counts[1] == 4 && counts[2] == 3 && counts[3] == 2 && counts[4] == 1
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		sizes := make([]int, 10)
		for j := 0; j < 10; j++ {
			fmt.Scan(&sizes[j])
		}

		if isValidShipsConfiguration(sizes) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
