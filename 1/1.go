package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	sum := 0
	content, error := os.Open("input.txt")
	if error != nil {
		fmt.Println("Missing file")
		log.Fatal(error)
	}

	reader := bufio.NewReader(content)
	numbers := []int{}
	for i := 0; i <= reader.Size(); i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
		numbers_in_line := []byte{}
		for i := 0; i < len(line); i++ {
			if unicode.IsNumber(rune(line[i])) {
				numbers_in_line = append(numbers_in_line, line[i])
			}

		}
		two_digits_string := string(numbers_in_line[0]) + string(numbers_in_line[len(numbers_in_line)-1])
		num, err := strconv.Atoi(two_digits_string)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}

	for _, digit := range numbers {
		sum += digit
	}
	fmt.Printf("Sum: %d", sum)
}
