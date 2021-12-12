package main

import (
	"fmt"
)

func main() {
	fmt.Println("Прогрессивный налог")

	fmt.Println("Введите доход:")
	var sum int
	fmt.Scan(&sum)

	var tax int

	if sum <= 0 {
		fmt.Println("Укажите корректный доход.")

	} else if sum < 10000 {
		tax = sum * 13 / 100
		fmt.Println("C дохода", sum, "рублей придётся заплатить:", tax)

	} else if sum < 50000 {
		tax = (sum-10000)*20/100 + 10000*13/100
		fmt.Println("C дохода", sum, "рублей придётся заплатить:", tax)

	} else {
		tax = (sum-50000)*30/100 + (50000-10000)*20/100 + 10000*13/100
		fmt.Println("C дохода", sum, "рублей придётся заплатить:", tax)
	}

}
