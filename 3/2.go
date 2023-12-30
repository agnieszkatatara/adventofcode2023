package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var board_data [][]byte
	content, error := os.Open("input.txt")
	var board_safety [][]int
	var board_data_string [][]string

	if error != nil {
		fmt.Println("Finished reading a file")
		os.Exit(2)
	}
	reader := bufio.NewReader(content)

	// Reading board data
	for i := 0; i <= reader.Size(); i++ {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(error)
		}
		line = bytes.TrimSuffix(line, []byte("\n"))
		board_data = append(board_data, line)
	}
	// Producing board of strings and empty board of safety
	for _, bd := range board_data {
		var row []string
		for _, b := range bd {
			row = append(row, string(b))
		}
		board_data_string = append(board_data_string, row)
	}
	// Printing string board of data
	fmt.Println("string board_data")
	fmt.Println(board_data_string)

	// Creating board of safety
	board_safety = make([][]int, len(board_data))
	for i := range board_data {
		board_safety[i] = make([]int, len(board_data[i]))
	}

	// Filling in board of safety
	gear_number := 1
	for i, bd := range board_data_string {
		for j, b := range bd {
			if strings.ContainsAny(b, "*") {
				mark_safe_spaces_based_on_symbol_location(i, j, board_safety, gear_number)
				gear_number = gear_number + 1
			}
		}
	}

	// Ready board of safety
	fmt.Println("board of safety")
	fmt.Println(board_safety)

	// Find consecutive digits on board of safety
	var digits string
	var safe_digit_flag bool
	var gear_number_saved int
	fmt.Println(safe_digit_flag)
	var gear_value = make(map[int][]int)
	for i, bd := range board_data_string {
		for j, b := range bd {
			if strings.ContainsAny(b, "0123456789") {
				digits = digits + b
				if board_safety[i][j] > 0 {
					safe_digit_flag = true
					gear_number_saved = board_safety[i][j]
				}
				if j == 139 {
					if safe_digit_flag {
						if digits != "" {
							value, error := strconv.Atoi(digits)
							if error != nil {
								fmt.Println("error while converting to int")
							}
							gear_value[gear_number_saved] = append(gear_value[gear_number_saved], value)
						}
					}
					safe_digit_flag = false
					digits = ""
				}
			} else {
				if safe_digit_flag {
					if digits != "" {
						value, error := strconv.Atoi(digits)
						if error != nil {
							fmt.Println("error while converting to int")
						}

						gear_value[gear_number_saved] = append(gear_value[gear_number_saved], value)
					}
				}
				safe_digit_flag = false
				digits = ""
			}
		}
	}
	fmt.Println(gear_value)
	sum := 0

	// Iterate through map, find if key has 2 values and multiply
	for _, values := range gear_value {
		if len(values) == 2 {
			sum = sum + (values[1] * values[0])
		}
	}
	fmt.Println(sum)

}

func mark_safe_spaces_based_on_symbol_location(x int, y int, board_safety [][]int, gear int) [][]int {
	if x != 0 && y != 0 {
		board_safety[x][y] = gear
		board_safety[x+1][y] = gear
		board_safety[x][y+1] = gear
		board_safety[x+1][y+1] = gear
		board_safety[x-1][y] = gear
		board_safety[x][y-1] = gear
		board_safety[x-1][y-1] = gear
		board_safety[x+1][y-1] = gear
		board_safety[x-1][y+1] = gear
	}
	if x == 0 && y != 0 {
		board_safety[x][y] = gear
		board_safety[x+1][y] = gear
		board_safety[x][y+1] = gear
		board_safety[x+1][y+1] = gear
		board_safety[x][y-1] = gear
		board_safety[x+1][y-1] = gear
	}

	if x != 0 && y == 0 {
		board_safety[x][y] = gear
		board_safety[x+1][y] = gear
		board_safety[x][y+1] = gear
		board_safety[x+1][y+1] = gear
		board_safety[x-1][y] = gear
		board_safety[x-1][y+1] = gear
	}

	return board_safety
}
