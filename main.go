package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Стрим начался!")
	defer func() {
		fmt.Println("Стрим закончился")
	}()

	for i := 0; i < 5; i++ {
		donate := float64(rand.Intn(10_000) + 1)
		currency := randomCurrency()

		fmt.Printf("Получен донат: %.2f %s\n", donate, currency)

		rubResult, err := convertDonateToRUB(donate, currency)

		if err != nil {
			fmt.Println("Ошибка конвертации:", err)
			fmt.Println()
			continue
		}

		fmt.Printf("Успешно конвертировано: %.2f RUB\n\n", rubResult)
		time.Sleep(1 * time.Second)
	}

}

func randomCurrency() string {
	currencies := []string{"USD", "EUR", "RUB", "CNY", "GBP"}

	randomIndex := rand.Intn(len(currencies))
	selectedCurrency := currencies[randomIndex]

	return selectedCurrency
}

func convertDonateToRUB(money float64, targetCurrency string) (float64, error) {
	const courseUSD = 85.5
	const courseEUR = 92.0
	var result float64

	if money <= 0 {
		return 0.0, errors.New("недостаточно средств")
	}
	if money >= 5000 {
		return 0.0, errors.New("превышен максимальный лимит пожертвования")
	}
	if targetCurrency != "USD" && targetCurrency != "EUR" {
		return 0.0, errors.New("поддерживается только USD и EUR")
	}

	switch targetCurrency {
	case "USD":
		result = money * courseUSD
	case "EUR":
		result = money * courseEUR
	}

	return result, nil
}
