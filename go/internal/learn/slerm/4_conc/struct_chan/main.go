package main

import (
	"fmt"
	"time"
)

/**
Напишите программу, которая в качестве типа данных канала использует структуру с полем указателем.
Передайте значение в канал, измените значение по указателю внутри структуры в читающей горутине, а так же другие поля, которые не являются указателями.
Проверьте значения полей структуры в исходной горутине после изменений полученной копии структуры в читающей горутине.
*/

type User struct {
	Name string
	Age  int
}

type Account struct {
	user   *User
	Amount int
}

func main() {
	ch := make(chan *Account, 1)

	s := Account{
		user: &User{
			Name: "Test",
			Age:  10,
		},
		Amount: 100,
	}

	ch <- &s

	go func() {
		a := <-ch
		a.user.Name = "World"
		a.user.Age = 20
		a.Amount = 200
	}()

	time.Sleep(time.Second)

	fmt.Println(s.user.Name, s.user.Age)
}
