package to_freq_map

/**
Напишите функцию для преобразования слайса строк в map.
Ключами в map должны быть строки в слайсе, а значениями в map должно быть количество раз, которое каждая строка появляется в слайсе.
Например, для слайса []string{”a”, “b”, “a”} получим map, где для ключа “a” будет значение 2, а для ключа “b” будет значение 1.
*/

func toFrequencyMap(s []string) map[string]int {
	r := make(map[string]int)

	for _, value := range s {
		r[value] += 1
	}

	return r
}
