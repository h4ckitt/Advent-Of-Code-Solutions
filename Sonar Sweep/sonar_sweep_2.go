package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ConvertToInt(nums []string) []int {
	ints := make([]int, len(nums))

	for index, elem := range nums {
		ints[index], _ = strconv.Atoi(elem)
	}

	return ints
}

func Add(nums []int) int {
	sum := 0

	for _, elem := range nums {
		sum += elem
	}

	return sum
}

func main() {
	body, err := ioutil.ReadFile("input_case.txt")

	if err != nil {
		log.Fatal(err)
	}

	numSlice := strings.Split(strings.TrimSpace(string(body)), "\n")
	intSlice := ConvertToInt(numSlice)

	count := 0
	previous := Add(intSlice[0:3])

	for i := 1; i < len(intSlice)-2; i++ {
		sum := Add(intSlice[i : i+3])

		if sum > previous {
			count++
		}

		previous = sum
	}

	fmt.Println(count)
}
