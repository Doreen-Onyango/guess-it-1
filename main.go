package main

import (
	"bufio"
	"fmt"
	"math"
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
		if len(numbers) > 1 {
			lowLimit, upLimit := calculate.CalculateRange(numbers)
			fmt.Printf("%v %.v\n", int(math.Round(lowLimit)), int(math.Round(upLimit)))
		}

	}
}
