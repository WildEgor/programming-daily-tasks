package main

import "fmt"

/**
Напишите простую функцию итерации по ключам map.
Запустите несколько раз и посмотрите на выдаваемый порядок ключей.
Объяснение результатов можно найти в документации https://go.dev/ref/spec#For_range.
Также напечатайте в консоль map целиком (просто передав в fmt.Println) и запустите несколько раз.
Обратите внимание на порядок ключей. Объяснение полученного результата https://tip.golang.org/doc/go1.12#fmt
(код https://cs.opensource.google/go/go/+/refs/tags/go1.18:src/fmt/print.go;l=769 )
*/

func main() {
	dict := make(map[string]bool)

	dict["red"] = true
	dict["yellow"] = false
	dict["green"] = false

	for key, _ := range dict {
		fmt.Println(key)
	}

	fmt.Println(dict)
}
