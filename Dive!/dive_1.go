package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	horizontal int = 0
	depth      int = 0
)

func CalculatePos(position string) {
	ws_stripped_position := strings.TrimSpace(position)

	position_params := strings.Split(ws_stripped_position, " ")

	if len(position_params) < 2 {
		return
	}

	num, _ := strconv.Atoi(position_params[1])

	switch position_params[0] {
	case "forward":
		horizontal += num
	case "down":
		depth += num
	case "up":
		depth -= num
	}
}

func main() {
	body, err := ioutil.ReadFile("input_case.txt")

	if err != nil {
		log.Fatalf("Unable To Open/Read File: %v", err)
	}

	positional_params_slice := strings.Split(string(body), "\n")

	for _, positional_param := range positional_params_slice {
		CalculatePos(positional_param)
	}

	fmt.Println(horizontal * depth)

}
