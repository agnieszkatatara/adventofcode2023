package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, error := os.Open("input.txt")
	multiplication := 0
	sum := 0
	if error != nil {
		fmt.Println("Missing file")
		log.Fatal(error)
	}

	reader := bufio.NewReader(content)
	for i := 0; i <= reader.Size(); i++ {
		line, err := reader.ReadString('\n')
		no_game_no := strings.SplitAfter(line, ": ")[1]
		standard_in := strings.Replace(no_game_no, ";", ",", -1)
		standard_in = strings.Replace(standard_in, ", ", ",", -1)
		standard_in = strings.Replace(standard_in, "\n", "", -1)
		color_with_number := strings.SplitAfter(standard_in, ",")
		amounts := make(map[string][]int)
		for _, color := range color_with_number {
			color = strings.Replace(color, ", ", "", -1)
			color = strings.Replace(color, ",", "", -1)
			colour := strings.SplitAfter(color, " ")[1]
			amount := strings.SplitAfter(color, " ")[0]
			colour = strings.Replace(colour, " ", "", -1)
			amount = strings.Replace(amount, " ", "", -1)
			amount_as_int, err := strconv.Atoi(amount)
			if err != nil {
				fmt.Println("Error converting value:", err)
			}
			amounts[colour] = append(amounts[colour], amount_as_int)
		}

		multiplication = findLargestNumber(amounts["blue"]) * findLargestNumber(amounts["green"]) * findLargestNumber(amounts["red"])
		sum = sum + multiplication
		fmt.Println(sum)
		if err != nil {
			break
		}
	}
}

func findLargestNumber(numbers []int) int {
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}
