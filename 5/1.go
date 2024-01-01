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
	inputs := []int{2880930400, 17599561, 549922357, 200746426, 1378552684, 43534336, 155057073, 56546377, 824205101, 378503603, 1678376802, 130912435, 2685513694, 137778160, 2492361384, 188575752, 3139914842, 1092214826, 2989476473, 58874625}
	//inputs := []int{2880930400, 17599561}
	//inputs := []int{79, 14, 55, 13}
	//inputs := []int{2880930400, 17599561, 549922357}
	for liczby := 0; liczby < len(inputs); liczby++ {
		current_number = inputs[liczby]
		next_map_check = next_map_check[:0]
		for i := 0; i < len(no_map_prefix_data); i++ {
			if no_map_prefix_data[i] == "" {

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

func find_if_number_is_weird(source int, list_of_sources_destinations []string) int {
	var row []string

	for _, value := range list_of_sources_destinations {
		row = strings.Split(value, " ")
		int_value_source, error := strconv.Atoi(row[1])
		if error != nil {
			fmt.Println("error while converting to int")
		}
		incrementor, error := strconv.Atoi(row[2])
		if error != nil {
			fmt.Println("error while converting to int")
		}

		if incrementor+int_value_source-2 > source && int_value_source <= source {
			// fmt.Println("duzo print")
			// fmt.Println(int_value_source)
			// fmt.Println(source)
			// fmt.Println(incrementor)
			int_value_destination, error := strconv.Atoi(row[0])
			if error != nil {
				fmt.Println("error while converting to int")
			}
			difference := source - int_value_source
			//fmt.Printf("I am about to finish and return %d \n", int_value_destination+difference)
			return int_value_destination + difference
		}
	}
	//fmt.Printf("I am about to finish and I found nothing %d \n", source)
	return source
}
