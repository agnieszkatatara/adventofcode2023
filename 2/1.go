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

func main2() {
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
		fmt.Println("here")
		fmt.Println(line)
		no_game_no := strings.SplitAfter(line, ": ")[1]
		fmt.Println("here2")
		fmt.Println(no_game_no)
		fmt.Println("here3")
		standard_in := strings.Replace(no_game_no, ";", ",", -1)
		standard_in = strings.Replace(standard_in, ", ", ",", -1)
		standard_in = strings.Replace(standard_in, "\n", "", -1)
		fmt.Println("here4")
		color_with_number := strings.SplitAfter(standard_in, ",")
		fmt.Println("here5")
		amounts := make(string[]int){}
		for _, color := range color_with_number {
			fmt.Println(color)
			color = strings.Replace(color, ", ", "", -1)
			color = strings.Replace(color, ",", "", -1)
			fmt.Println(color)
			fmt.Println("color")
			colour := strings.SplitAfter(color, " ")[1]
			amount := strings.SplitAfter(color, " ")[0]
			fmt.Println("kolor ditst")
			fmt.Println(string(colour))
			fmt.Println(string(amount))
			colour = strings.Replace(colour, " ", "", -1)
			amount = strings.Replace(amount, " ", "", -1)
			fmt.Println(string(colour))
			fmt.Println(string(amount))
			fmt.Println("here7")


			if !check_if_round_is_possible(colour, amount) {
				fmt.Printf("this game is impossible")
				impossible_rounds += 1
				//fmt.Printf("this game is impossible")
				//fmt.Printf("dla niemoliwego %s, %s", color_list[2], string(color_list[1]))
				break
			}
		}
		//fmt.Printf("My line no %d %s \n", i, standard_in)
		//fmt.Printf("And it is doable if it %d \n", impossible_rounds)

		if impossible_rounds == 0 {
			sum_of_workable_rounds = sum_of_workable_rounds + 1 + i
		}

		//fmt.Println("działa: ")
		//fmt.Println(sum_of_workable_rounds)
		if err != nil {
			break
		}
	}
	fmt.Println("działa: ")
	fmt.Println(sum_of_workable_rounds)
}

func check_if_round_is_possible(color string, number string) bool {
	fmt.Println("checking")
	fmt.Println(color)
	fmt.Println(number)
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
