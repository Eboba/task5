package main

import (
	"fmt"
)

func main() {
	fmt.Println("Банкомат.")

	fmt.Println("Введите сумму снятия со счёта:")
	var sum int
	fmt.Scan(&sum)

	if sum > 0 {
		if sum <= 100000 {
			if sum%100 == 0 {
				fmt.Println("Операция успешно выполнена.")
				fmt.Println("Вы сняли", sum, "рублей.")

			} else {
				fmt.Println("Операция не выполнена.")
				fmt.Println("Запрашиваемая сумма не кратна купюрам в банкомате.")
			}

		} else {
			fmt.Println("Операция не выполнена.")
			fmt.Println("Запрашиваемая сумма превышает лимит.")
		}
	} else {
		fmt.Println("Операция не выполнена.")
		fmt.Println("Запрашиваемая сумма должна быть больше 0.")
	}
}
