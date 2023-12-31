package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	var data []string
	var removed_card_prefix string = ""
	var numbers []string
	var total_sum int = 0
	content, error := os.Open("input.txt")

	if error != nil {
		fmt.Println("Finished reading a file")
		os.Exit(2)
	}
	reader := bufio.NewReader(content)

	// Reading board data
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		data = append(data, line)
	}

	// Parsing winning numbers

	for _, value := range data {
		removed_card_prefix = strings.Split(value, ":")[1]
		removed_card_prefix = strings.ReplaceAll(removed_card_prefix, "  ", " ")
		numbers = strings.Split(removed_card_prefix, " | ")
		total_sum = total_sum + find_number_of_winning_numbers(numbers)
	}
	fmt.Println(total_sum)
}

func find_number_of_winning_numbers(numbers []string) int {
	var winning_numbers []string
	var guessed_numbers []string
	winning_numbers = strings.Split(numbers[0], " ")
	guessed_numbers = strings.Split(numbers[1], " ")
	winning_numbers = winning_numbers[1:]
	var sum int = 0
	for _, number := range guessed_numbers {
		if slices.Contains(winning_numbers, number) {
			sum = sum + 1
		}
	}

	switch {
	case sum == 1:
		return 1
	case sum == 2:
		return 2
	case sum > 2:
		return int(math.Pow(2, float64(sum-1)))
	default:
		return sum
	}
}
