package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main2() {
	content, error := os.Open("input.txt")
	numbers := []int{}
	sum := 0
	if error != nil {
		fmt.Println("Missing file")
		log.Fatal(error)
	}
	reader := bufio.NewReader(content)
	for y := 0; y <= reader.Size(); y++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
		var last_digit_to_pass_string string
		var last_digit_to_pass_index int
		numbers_in_line := []string{}
		no_first_number := 0
		for i := 0; i < len(line); i++ {
			if unicode.IsNumber(rune(line[i])) {
				if no_first_number == 0 && i == 0 {
					no_first_number = 1
					numbers_in_line = append(numbers_in_line, string(line[i]))
				}
				if no_first_number == 0 {
					success, output := find_textual_number_earlier(string(line[0:i]))
					if success {
						numbers_in_line = append(numbers_in_line, output)
					} else {
						numbers_in_line = append(numbers_in_line, string(line[i]))
					}
					no_first_number = 1
				}
				last_digit_to_pass_string = string(line[i])
				last_digit_to_pass_index = i

			}

			if i == len(line)-1 {
				if no_first_number == 0 {
					success, output := find_textual_number_earlier(string(line))

					if success {
						numbers_in_line = append(numbers_in_line, output)
					}
				}
				success, output := find_textual_number_later(string(line[last_digit_to_pass_index:]))

				if success {
					numbers_in_line = append(numbers_in_line, output)
				} else {
					numbers_in_line = append(numbers_in_line, string(line[last_digit_to_pass_index]))
				}
			}
		}

		two_digits_string := string(numbers_in_line[0]) + string(numbers_in_line[len(numbers_in_line)-1])
		num, err := strconv.Atoi(two_digits_string)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}
	for _, value := range numbers {
		fmt.Println(value)
	}

	for _, digit := range numbers {
		sum += digit
	}
	fmt.Printf("Sum: %d", sum)

}

func find_textual_number_earlier(string_to_find_number string) (bool, string) {
	numbers_in_text := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	already_existed := false
	lowest_index := 0
	earliest_value := 0
	for i, str := range numbers_in_text {
		index := strings.Index(string_to_find_number, str)
		if index != -1 {
			if already_existed == false {
				already_existed = true
				lowest_index = index
				earliest_value = i
			} else {
				if lowest_index > index {
					already_existed = true
					lowest_index = index
					earliest_value = i
				}
			}

		}
	}

	if already_existed {
		return true, strconv.Itoa(earliest_value)
	} else {
		return false, "0"
	}
}

func find_textual_number_later(string_to_find_number string) (bool, string) {
	numbers_in_text := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	already_existed := false
	highest_index := 0
	latest_value := 0
	for i, str := range numbers_in_text {
		index := strings.LastIndex(string_to_find_number, str)
		if index != -1 {
			if already_existed == false {
				already_existed = true
				highest_index = index
				latest_value = i
			} else {
				if highest_index < index {
					already_existed = true
					highest_index = index
					latest_value = i
				}
			}
		}
	}

	if already_existed {
		return true, strconv.Itoa(latest_value)
	} else {
		return false, "0"
	}
}
