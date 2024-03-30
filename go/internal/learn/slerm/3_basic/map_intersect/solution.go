package map_intersect

/**
Напишите функцию для поиска пересечения двух map. Функция должна вернуть новую map, содержащую ключи, присутствующие в обеих исходных map.
*/

func mapKeyIntersect(m1 map[int]struct{}, m2 map[int]struct{}) []int {
	keys := make([]int, 0)

	for key, _ := range m1 {
		if _, ok := m2[key]; ok {
			keys = append(keys, key)
		}
	}

	return keys
}
