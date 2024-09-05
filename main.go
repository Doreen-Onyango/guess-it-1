package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"guessit/calculate"
)

func main() {
	var numbers []float64
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}
		numbers = append(numbers, number)

		lowLimit, upLimit := calculate.CalculateRange(numbers)
		fmt.Printf("%.0f %.0f\n", lowLimit, upLimit)
	}
}
