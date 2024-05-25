package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

/**
Условие задачи
Задана дата в формате "день месяц год" в виде трёх целых чисел. Гарантируется, что:
∙ день— это целое число от 1 до 31;
· месяц— это целое число от 1 до 12;
· год — это целое число от 1950 до 2300.
Проверьте, что заданные три числа соответствуют корректной дате (в современном григорианском календаре).
Напоминаем, что в соответствии с современным календарём год считает високосным, если для этого года верно хотя бы одно из утверждений:
· делится на 4, но при этом не делится на 100;
· делится на 400.
Например, годы 2012 и 2000 являются високосными, но годы 1999, 2022 и 2100 - нет.
*/

func checkDate(data string) string {
	pd := strings.Split(data, " ")
	if _, err := time.Parse("2006 01 02", fmt.Sprintf("%04s %02s %02s", pd[2], pd[1], pd[0])); err != nil {
		return "no"
	}

	return "yes"
}

func main() {
	r := bufio.NewReader(os.Stdin)
	c := -1

	for {
		txt, _ := r.ReadString('\r')
		txt2, _ := r.ReadString('\n')
		txt += txt2

		if c == -1 {
			c, _ = strconv.Atoi(txt)
			continue
		}

		fmt.Fprintln(os.Stdout, checkDate(txt))

		if c--; c == 0 {
			os.Exit(0)
		}
	}
}
