package main

import (
	"fmt"
)

func main() {
	fmt.Println("Баллы ЕГЭ.")

	fmt.Println("Введите результат первого экзамена:")
	var scorOne int
	fmt.Scan(&scorOne)

	fmt.Println("Введите результат второго экзамена:")
	var scorTwo int
	fmt.Scan(&scorTwo)

	fmt.Println("Введите результат третьего экзамена:")
	var scorThree int
	fmt.Scan(&scorThree)

	if scorOne < 0 {
		fmt.Println("Неверно указаны баллы первого экзамена")
	} else if scorOne > 100 {
		fmt.Println("Неверно указаны баллы первого экзамена")
	} else if scorTwo < 0 {
		fmt.Println("Неверно указаны баллы второго экзамена")
	} else if scorTwo > 100 {
		fmt.Println("Неверно указаны баллы второго экзамена")
	} else if scorThree > 100 {
		fmt.Println("Неверно указаны баллы третьего экзамена")
	} else if scorThree < 0 {
		fmt.Println("Неверно указаны баллы третьего экзамена")
	} else {

		scorSum := scorOne + scorTwo + scorThree
		fmt.Println("Сумма проходных баллов: 275")
		fmt.Println("Количество набранных баллов:", scorSum)
		if scorSum < 275 {
			fmt.Println("Вы не поступили.")
		} else {
			fmt.Println("Вы поступили.")
		}
	}
}
