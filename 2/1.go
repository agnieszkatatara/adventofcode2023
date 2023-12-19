package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var available_cubes = map[string]int{"red": 12, "green": 13, "blue": 14}

func main() {
	content, error := os.Open("test_input.txt")
	if error != nil {
		fmt.Println("Missing file")
		log.Fatal(error)
	}

	sum_of_workable_rounds := 0
	reader := bufio.NewReader(content)
	for i := 0; i <= reader.Size(); i++ {
		impossible_rounds := 0
		line, err := reader.ReadString('\n')
		no_game_no := strings.SplitAfter(line, ": ")[1]
		standard_in := strings.Replace(no_game_no, ";", ",", -1)
		standard_in = strings.Replace(standard_in, ", ", ",", -1)
		standard_in = strings.Replace(standard_in, "\n", "", -1)
		color_with_number := strings.SplitAfter(standard_in, ",")
		amounts := make(string[]int){}
		for _, color := range color_with_number {
			color = strings.Replace(color, ", ", "", -1)
			color = strings.Replace(color, ",", "", -1)
			colour := strings.SplitAfter(color, " ")[1]
			amount := strings.SplitAfter(color, " ")[0]
			colour = strings.Replace(colour, " ", "", -1)
			amount = strings.Replace(amount, " ", "", -1)

			if !check_if_round_is_possible(colour, amount) {
				impossible_rounds += 1
				break
			}
		}
		if impossible_rounds == 0 {
			sum_of_workable_rounds = sum_of_workable_rounds + 1 + i
		}

		if err != nil {
			break
		}
	}
	fmt.Println("Score: ")
	fmt.Println(sum_of_workable_rounds)
}

func check_if_round_is_possible(color string, number string) bool {
	value, error := strconv.Atoi(number)
	if error != nil {
		fmt.Println("Error converting value:", error)
	}

	if available_cubes[color] >= value {
		return true
	} else {
		return false
	}
}
