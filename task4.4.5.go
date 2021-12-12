package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ресторан")

	fmt.Println("Введите день недели:")
	var day int
	fmt.Scan(&day)

	fmt.Println("Введите число гостей:")
	var people int
	fmt.Scan(&people)

	fmt.Println("Введите сумму чека:")
	var sum int
	fmt.Scan(&sum)

	var discount int
	var allowance int

	if day == 1 {
		discount = sum * 10 / 100
		fmt.Println("Скидка по понедельникам:", discount)

	} else if day == 5 {
		discount = sum * 5 / 100
		fmt.Println("Скидка по пятницам:", discount)
	}
	if people > 5 {
		allowance = sum * 10 / 100
		fmt.Println("Надбавка на обслуживание:", allowance)
	}

	fmt.Println("Сумма к оплате:", sum+allowance-discount)

}
