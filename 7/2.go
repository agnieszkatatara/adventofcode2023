package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var data []string
	var no_map_prefix_data []string
	var next_map_check []string
	var current_number int
	var next_number int
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
	for i := 3; i < len(data); i++ {
		if !strings.Contains(data[i], "map") {
			no_map_prefix_data = append(no_map_prefix_data, data[i])
		}
	}
	var min int = 100000000000000
	//inputs := []int{2880930400, 17599561, 549922357, 200746426, 1378552684, 43534336, 155057073, 56546377, 824205101, 378503603, 1678376802, 130912435, 2685513694, 137778160, 2492361384, 188575752, 3139914842, 1092214826, 2989476473, 58874625}
	//inputs := []int{2880930400, 17599561}
	//inputs := []int{79, 14, 55, 13}
	//inputs := []int{2880930400, 17599561, 549922357}
	for liczby := 2880930400; liczby < 2880930400+175913; liczby++ {
		current_number = liczby
		next_map_check = next_map_check[:0]
		for i := 0; i < len(no_map_prefix_data); i++ {
			if no_map_prefix_data[i] == "" {
				//fmt.Println(next_map_check)
				sortBasedOnThirdElement(next_map_check)
				//fmt.Println("sorted")
				//fmt.Println(next_map_check)
				next_number = find_if_number_is_weird(current_number, next_map_check)
				current_number = next_number
				next_map_check = next_map_check[:0]
				//fmt.Println("should be empty")
				//fmt.Println(next_map_check)
				//fmt.Println("switch")
				//fmt.Println(current_number)
			}
			if no_map_prefix_data[i] != "" {
				//fmt.Println("dodaje")
				//fmt.Println(no_map_prefix_data[i])
				next_map_check = append(next_map_check, no_map_prefix_data[i])
				//fmt.Println("should be inc")
				//fmt.Println(next_map_check)
			}
		}
		//fmt.Println("last czeck")
		current_number = find_if_number_is_weird(current_number, next_map_check)
		//fmt.Println("all")
		//fmt.Println(current_number)
		//fmt.Println(inputs[liczby])
		//fmt.Println(current_number)
		if current_number < min {
			min = current_number
			fmt.Println(min)
		}
	}

}

func sortBasedOnThirdElement(list []string) {
	sort.Slice(list, func(i, j int) bool {
		// Split the strings by space
		splitI := strings.Split(list[i], " ")
		splitJ := strings.Split(list[j], " ")

		// Convert the third element to integer for comparison
		valI, errI := strconv.Atoi(splitI[1])
		valJ, errJ := strconv.Atoi(splitJ[1])

		// Handle potential conversion errors
		if errI != nil {
			panic(errI)
		}
		if errJ != nil {
			panic(errJ)
		}

		// Compare the integers
		return valI < valJ
	})
}

func find_if_number_is_weird(source int, list_of_sources_destinations []string) int {
	var in_string string

	left, right := 0, len(list_of_sources_destinations)-1
	var mid int
	result := source // Default to -1 if no larger number is found

	for left <= right {
		mid = left + (right-left)/2
		in_string = strings.Split(list_of_sources_destinations[mid], " ")[1]
		in_int, error := strconv.Atoi(in_string)
		if error != nil {
			fmt.Println("error while converting to int")
		}
		if in_int <= source {
			left = mid + 1
		} else {
			result = in_int
			right = mid - 1
		}
	}
	return result
}
