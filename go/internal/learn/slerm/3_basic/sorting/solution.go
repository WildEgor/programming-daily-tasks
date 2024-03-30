package sorting

/**
Задание со звездочкой*
Напишите функцию для сортировки слайса любым способом. Поэксперементируйте с разными вариантами реализации:
изменение in place без возврата значения;
возврат из функции нового значения.
Сравните время выполнения для очень больших размеров слайсов. Подумайте, почему один вариант работает быстрее.
Для проверки времени выполнения изучите, как применяется пакет time.
*/

func bubbleSort(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		cV := arr[i]
		pI := i - 1
		for pI >= 0 && cV < arr[pI] {
			arr[pI+1] = arr[pI]
			pI--
		}
		arr[pI+1] = cV
	}

	return arr
}

func insertSortAlt(arr []int) {
	for i := 1; i < len(arr); i++ {
		cV := arr[i]
		pI := i - 1
		for pI >= 0 && cV < arr[pI] {
			arr[pI+1] = arr[pI]
			pI--
		}
		arr[pI+1] = cV
	}
}
