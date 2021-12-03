package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	body, err := ioutil.ReadFile("input_case.txt")

	if err != nil {
		log.Fatalf("Unable To Open/Read File: %v", err)
	}

	numSlice := strings.Split(strings.TrimSpace(string(body)), "\n")
	count := 0
	previous, _ := strconv.Atoi(numSlice[0])
	for i := 1; i < len(numSlice); i++ {
		num, _ := strconv.Atoi(numSlice[i])
		if num > previous {
			count++
		}
		previous = num
	}

	fmt.Println(count)

}
