package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	var interval float64 = 1.28
	var validInt []int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		if len(validInt) == 0 {
			validInt = append(validInt, n)
		}

		average := calculateMean(validInt)
		deviation := calculateStandardDeviation(validInt, average)

		if !(n > int(average*2) || n < int(average/2)) {
			validInt = append(validInt, n)
		}

		low := average - (deviation * interval)
		high := average + (deviation * interval)

		fmt.Printf("%d %d\n", int(math.Round(low)), int(math.Round(high)))
	}
}

func calculateMean(numbers []int) float64 {
	length := len(numbers)
	var total float64
	for _, n := range numbers {
		total += float64(n)
	}
	mean := total / float64(length)
	return mean
}

func calculateStandardDeviation(numbers []int, average float64) float64 {
	var deviation float64
	for _, n := range numbers {
		deviation += (float64(n) - average) * (float64(n) - average)
	}
	deviation = float64(deviation) / float64(len(numbers))
	return math.Sqrt(deviation)
}
