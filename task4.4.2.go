package main

import (
	"fmt"
)

func main() {
	fmt.Println("Три числа.")

	fmt.Println("Введите первое число:")
	var one int
	fmt.Scan(&one)

	fmt.Println("Введите второе число:")
	var two int
	fmt.Scan(&two)

	fmt.Println("Введите третье число:")
	var three int
	fmt.Scan(&three)

	if one <= 5 {

		if two <= 5 {

			if three <= 5 {
				fmt.Println("Среди введённых чисел нет числа больше 5.")

			} else {
				fmt.Println("Среди введённых чисел есть число больше 5.")
			}
		} else {
			fmt.Println("Среди введённых чисел есть число больше 5.")
		}
	} else {
		fmt.Println("Среди введённых чисел есть число больше 5.")
	}
}
