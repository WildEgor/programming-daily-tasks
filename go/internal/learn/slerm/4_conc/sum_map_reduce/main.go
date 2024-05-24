package main

import (
	"fmt"
	"sync"
)

/**
Напишите программу, которая суммирует числа с использованием горутин и выдает общую сумму (простой map/reduce).
На входе основной функции - слайс каналов, из которых читаются данные (для примера: это данные из разных источников).
Все горутины читают исходные каналы и складывают частичные суммы к себе, пока в исходных каналах есть данные и пока они не закрыты. (операция map)
Создайте выходной канал, в который после дочитывания данных суммирующие горутины запишут свои результаты. Просуммируйте результаты из выходного канала. (операция reduce)
Итоговая схема: каналы с исходными данными ⇒ набор суммирующих горутин ⇒ сбор финальной суммы.
Цель этой задачи: освоиться с sync.WaitGroup, с чтением/записью каналов, с блокировками выполнения горутин.
*/

func main() {
	var chs []chan int64
	for range 10 {
		var ch = make(chan int64, 10)
		for num := range 10 {
			ch <- int64(num) // 0 + 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 = 45 x 10 = 450
		}
		close(ch)
		chs = append(chs, ch)
	}

	result := sumChannels(chs)

	fmt.Println(result)
}

// sumChannels
// in - слайс входных каналов, в которые приходят числа
// Признак окончания данных в канале - канал закрыт
func sumChannels(inputs []chan int64) int64 {
	rc := make(chan int64)
	var result int64

	var wg sync.WaitGroup
	wg.Add(len(inputs))

	for _, ch := range inputs {
		go func(ch chan int64) {
			defer wg.Done()

			var sum int64
			for num := range ch {
				sum += num
			}

			rc <- sum
		}(ch)
	}

	go func() {
		wg.Wait()
		close(rc)
	}()

	for sum := range rc {
		result += sum
	}

	return result
}
