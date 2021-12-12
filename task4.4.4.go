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

	sum := 0 //Количество чисел больше 5

	if one >= 5 {
		sum++
	}

	if two >= 5 {
		sum++
	}

	if three >= 5 {
		sum++
	}

	if sum == 0 {
		fmt.Println("Среди введённых чисел нет больше или равных 5.")
	}

	if sum != 0 {
		fmt.Println("Среди введённых чисел", sum, "больше или равных 5.")
	}

}
