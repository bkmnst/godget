package godget

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func AddTransaction(transaction string) {
	f, err := os.OpenFile("transactions", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString(transaction + "\n"); err != nil {
		panic(err)
	}
}

func ListTransactions() string {
	f, err := os.ReadFile("transactions")
	if err != nil {
		log.Fatal(err)
	}

	str := string(f)
	return str
}

func CalculateBalance() float64 {
	balance := 0.00

	transactions := ListTransactions()
	lines := strings.Split(transactions, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		operation := line[:1]
		amount, err := strconv.ParseFloat(line[1:], 64)
		if err != nil {
			log.Fatal(err)
		}

		switch operation {
		case "+":
			balance += amount
		case "-":
			balance -= amount
		}
	}

	return balance
}
