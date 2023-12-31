package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	var data []string
	var removed_card_prefix string = ""
	var numbers []string
	var total_sum int = 0
	var matched_numbers int = 0
	content, error := os.Open("input.txt")

	if error != nil {
		fmt.Println("Finished reading a file")
		os.Exit(2)
	}
	reader := bufio.NewReader(content)

	cards := make(map[int]int)

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

	for i := 1; i <= len(data); i++ {
		cards[i] = 1
	}
	// Parsing winning numbers

	for i, value := range data {
		var y int = 0
		removed_card_prefix = strings.Split(value, ":")[1]
		removed_card_prefix = strings.ReplaceAll(removed_card_prefix, "  ", " ")
		numbers = strings.Split(removed_card_prefix, " | ")
		matched_numbers = find_number_of_winning_numbers(numbers)
		fmt.Println("I matched_numbers")
		fmt.Println(matched_numbers)

		if matched_numbers > 0 {
			for y = 1; y <= matched_numbers; y++ {
				fmt.Println("My y is ")
				fmt.Println(y)
				fmt.Println("I do number one")
				fmt.Println(i + 1)
				if i+1+y <= len(data) {
					fmt.Println("I am adding one to %d", i+1+y)
					cards[i+1+y] = cards[i+1+y] + cards[i+1]
				}
			}
		}
	}
	fmt.Println(cards)
	fmt.Println("sum of cards")
	for _, value := range cards {
		total_sum = total_sum + value
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
	return sum
}
