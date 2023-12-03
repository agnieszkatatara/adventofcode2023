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

func main() {
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
				//fmt.Println("my i")
				//fmt.Println(i)
				//fmt.Println("I found my number as string(line[i])")
				//fmt.Println(string(line[i]))
				if no_first_number == 0 && i == 0 {
					no_first_number = 1
					fmt.Printf("my number %c \n", rune(line[i]))
					numbers_in_line = append(numbers_in_line, string(line[i]))
					fmt.Printf("I found a number super early %s \n", numbers_in_line[0])
					fmt.Println("I work")
				}
				if no_first_number == 0 {
					fmt.Printf("I pass this string %s\n", string(line[0:i]))
					success, output := find_textual_number_earlier(string(line[0:i]))

					if success {
						numbers_in_line = append(numbers_in_line, output)
						fmt.Printf("I add this digit %c \n", output)
					} else {
						//ddebug this
						numbers_in_line = append(numbers_in_line, string(line[i]))
						fmt.Printf("I added this digit %s\n", string(line[i]))
					}
					no_first_number = 1
				}
				last_digit_to_pass_string = string(line[i])
				last_digit_to_pass_index = i

			}

			fmt.Println("My line and length")
			fmt.Println(i)
			fmt.Println(len(line) - 1)
			if i == len(line)-1 {
				if no_first_number == 0 {
					fmt.Printf("There is no first digit I pass this string %s\n", string(line))
					success, output := find_textual_number_earlier(string(line))

					if success {
						numbers_in_line = append(numbers_in_line, output)
						fmt.Printf("I add this digit %c \n", output)
					}
				}
				fmt.Printf("I am in my last loop for the word %s and the last digit is %s \n", line, last_digit_to_pass_string)
				fmt.Printf("I am passing %s", string(line[last_digit_to_pass_index:]))
				success, output := find_textual_number_later(string(line[last_digit_to_pass_index:]))

				if success {
					fmt.Println("my later output %c, %c \n", success, output)
					numbers_in_line = append(numbers_in_line, output)
					fmt.Printf("I add later this digit %c \n", output)
				} else {
					//ddebug this
					numbers_in_line = append(numbers_in_line, string(line[last_digit_to_pass_index]))
					fmt.Printf("I added with later this digit %s\n", string(line[last_digit_to_pass_index]))
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
	fmt.Println("Integers in the list:")
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
		//fmt.Printf("looking in %d %s \n", i, str)
		index := strings.Index(string_to_find_number, str)
		if index != -1 {
			fmt.Printf("I found a hidden number in %s, and it is in index %d \n", string_to_find_number, index)
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
		fmt.Printf("looking in %d %s \n", i, str)
		index := strings.LastIndex(string_to_find_number, str)
		if index != -1 {
			fmt.Printf("I found a hidden number in %s, and it is in index %d \n", string_to_find_number, index)
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
		fmt.Printf("Later adding %string \n", strconv.Itoa(latest_value))
		return true, strconv.Itoa(latest_value)
	} else {
		fmt.Printf("Later not adding %s \n", "0")
		return false, "0"
	}
}
