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

var fiveOfAKind []string
var fourOfAKind []string
var fullHouse []string
var threeOfAKind []string
var twoPairs []string
var onePair []string
var highCard []string

func main() {
	content, error := os.Open("input.txt")
	var data [][]string
	var mapOfHands map[string]int
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

		data = append(data, readTwoConsecutiveStrings(line))
	}
	//fmt.Println(data)
	for i := 0; i < len(data); i++ {
		typesOfHands(data[i][0])
	}
	// fmt.Println("all")
	mapOfHands = turnListIntoMap(data)
	//fmt.Println(mapOfHands)
	//fmt.Println(sortList(sortList(sortList(sortList(sortList(sortList(sortList(sortList(fourOfAKind)))))))))

	for i := 0; i < 20001; i++ {
		fiveOfAKind = sortList(fiveOfAKind)
		fourOfAKind = sortList(fourOfAKind)
		fullHouse = sortList(fullHouse)
		threeOfAKind = sortList(threeOfAKind)
		twoPairs = sortList(twoPairs)
		onePair = sortList(onePair)
		highCard = sortList(highCard)
	}
	// fmt.Println("fiveOfAKind")
	// fmt.Println(fiveOfAKind)
	// fmt.Println("fourOfAKind")
	// fmt.Println(fourOfAKind)
	// fmt.Println("fullHouse")
	// fmt.Println(fullHouse)
	// fmt.Println("threeOfAKind")
	// fmt.Println(threeOfAKind)
	// fmt.Println("twoPairs")
	// fmt.Println(twoPairs)
	// fmt.Println("onePair")
	// fmt.Println(onePair)
	// fmt.Println("highCard")
	// fmt.Println(highCard)
	// fmt.Println("checking stronger functiom")
	var result int = 0
	for i := 0; i < len(highCard); i++ {
		result = result + (i+1)*mapOfHands[highCard[i]]
	}
	for i := 0; i < len(onePair); i++ {
		result = result + (i+1+len(highCard))*mapOfHands[onePair[i]]
	}
	for i := 0; i < len(twoPairs); i++ {
		result = result + (i+1+len(highCard)+len(onePair))*mapOfHands[twoPairs[i]]
	}
	for i := 0; i < len(threeOfAKind); i++ {
		result = result + (i+1+len(highCard)+len(onePair)+len(twoPairs))*mapOfHands[threeOfAKind[i]]
	}
	for i := 0; i < len(fullHouse); i++ {
		result = result + (i+1+len(highCard)+len(onePair)+len(twoPairs)+len(threeOfAKind))*mapOfHands[fullHouse[i]]
	}
	for i := 0; i < len(fourOfAKind); i++ {
		result = result + (i+1+len(highCard)+len(onePair)+len(twoPairs)+len(threeOfAKind)+len(fullHouse))*mapOfHands[fourOfAKind[i]]
	}
	for i := 0; i < len(fiveOfAKind); i++ {
		result = result + (i+1+len(highCard)+len(onePair)+len(twoPairs)+len(threeOfAKind)+len(fullHouse)+len(fourOfAKind))*mapOfHands[fiveOfAKind[i]]
	}

	fmt.Println(result)
}

func sortList(list []string) []string {
	if len(list) == 0 {
		return list
	}
	for i, _ := range list[:len(list)-1] {
		if list[i][0] != list[i+1][0] {
			list[i], list[i+1] = checkWhichIsWeaker(list[i], list[i+1], 0)
		} else if list[i][1] != list[i+1][1] {
			list[i], list[i+1] = checkWhichIsWeaker(list[i], list[i+1], 1)
		} else if list[i][2] != list[i+1][2] {
			list[i], list[i+1] = checkWhichIsWeaker(list[i], list[i+1], 2)
		} else if list[i][3] != list[i+1][3] {
			list[i], list[i+1] = checkWhichIsWeaker(list[i], list[i+1], 3)
		} else {
			list[i], list[i+1] = checkWhichIsWeaker(list[i], list[i+1], 4)
		}
	}
	return list
}

func readTwoConsecutiveStrings(s string) []string {
	fields := strings.Fields(s)
	numbers := make([]string, 0, 1)

	for _, number := range fields {
		numbers = append(numbers, number)
	}

	return numbers
}

func turnListIntoMap(list [][]string) map[string]int {
	set := make(map[string]int)
	for _, pair := range list {
		if len(pair) == 2 {
			number, err := strconv.Atoi(pair[1])
			if err != nil {
				return nil
			}
			set[pair[0]] = number
		}
	}
	return set
}

func checkWhichIsWeaker(one string, two string, value int) (string, string) {
	var values map[string]int
	values = map[string]int{"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2}
	if values[string(one[value])] >= values[string(two[value])] {
		return two, one
	} else {
		return one, two
	}
}

func typesOfHands(one string) {
	occurrences := make(map[string]int)
	for _, char := range one {
		occurrences[string(char)]++
	}
	assignCardsToType(occurrences, one)
}

func assignCardsToType(occurences map[string]int, hand string) {

	for _, value := range occurences {

		if value == 5 {
			fiveOfAKind = append(fiveOfAKind, hand)
			return
		}

		if (value == 4 || value == 1) && len(occurences) == 2 {
			fourOfAKind = append(fourOfAKind, hand)
			return
		}

		if (value == 3 || value == 2) && len(occurences) == 2 {
			fullHouse = append(fullHouse, hand)
			return
		}

		if len(occurences) == 3 && value == 1 {
			continue
		}

		if len(occurences) == 3 && value == 2 {
			twoPairs = append(twoPairs, hand)
			return
		}

		if len(occurences) == 3 && value == 3 {
			threeOfAKind = append(threeOfAKind, hand)
			return
		}

		if len(occurences) == 4 {
			onePair = append(onePair, hand)
			return
		}

		if len(occurences) == 5 {
			highCard = append(highCard, hand)
			return
		}
	}
	fmt.Println("this shouldnt happen")
	return

}
