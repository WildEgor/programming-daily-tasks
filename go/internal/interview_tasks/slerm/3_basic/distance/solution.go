package __3_4

import "math"

/**
Написать функцию, работающую с float32. Функция будет вычислять  расстояние между двумя точками на поверхности Земли (координаты latitude/longitude).
Формулу вычисления найдите самостоятельно и после реализации сравните полученные значения с тем же самым расстоянием, полученным по картам Яндекс/Google.
*/

func distance(lat1, lon1 float32, lat2, lon2 float32) float32 {
	const (
		earthRadius = 6371 // km
	)
	lat1Rad := float64(lat1) * math.Pi / 180
	lon1Rad := float64(lon1) * math.Pi / 180
	lat2Rad := float64(lat2) * math.Pi / 180
	lon2Rad := float64(lon2) * math.Pi / 180
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return float32(earthRadius * c)
}

/**
Написать функцию, работающую с float32. На входе две координаты (latitude/longitude).
Определите, что они указывают на одну и ту же локацию.
Подсказка: сравнение вещественных числе на строгое равенство - некорректный подход.
*/

func distanceCompare(lat1, lon1 float32, lat2, lon2 float32) bool {
	const (
		float64EqualityThreshold = 1e-9
	)

	return math.Abs(float64(lat1-lat2)) <= float64EqualityThreshold && math.Abs(float64(lon1-lon2)) <= float64EqualityThreshold
}
