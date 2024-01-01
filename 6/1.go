package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, error := os.Open("input.txt")
	var data [][]int
	if error != nil {
		fmt.Println("Finished reading a file")
		os.Exit(2)
	}
	reader := bufio.NewReader(content)

	// Reading data
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		data = append(data, readThreeConsecutiveNumbers(line))
	}
	fmt.Println(data)
	var score int = 1
	var waysToWin int
	for i := 0; i < len(data[0]); i++ {
		fmt.Println("checking for")
		fmt.Println(data[0][i])
		waysToWin = checkWaysToWin(data[0][i], data[1][i])
		score = score * waysToWin
	}
	fmt.Println(score)

}

func checkWaysToWin(seconds int, distance int) int {
	var seconds_holding int
	var seconds_swimming int
	var distance_swam int
	var ways_to_win int = 0

	for seconds_holding = 0; seconds_holding < seconds; seconds_holding++ {
		seconds_swimming = seconds - seconds_holding
		distance_swam = seconds_holding * seconds_swimming
		// fmt.Println("seconds holding")
		// fmt.Println(seconds_holding)
		// fmt.Println("seconds swimming")
		// fmt.Println(seconds_swimming)
		// fmt.Println("distance_swam")
		// fmt.Println(distance_swam)
		// fmt.Println("next round")
		// fmt.Println("")
		if distance < distance_swam {
			ways_to_win++
		}
	}
	return ways_to_win
}

func readThreeConsecutiveNumbers(s string) []int {
	fields := strings.Fields(s)
	numbers := make([]int, 0, 1)

	for _, field := range fields[1:] {
		number, err := strconv.Atoi(field)
		if err != nil {
			return nil
		}
		numbers = append(numbers, number)
	}

	return numbers
}
