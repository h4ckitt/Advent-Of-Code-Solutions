package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	firstFiveDigits  [12]int
	secondFiveDigits [12]int
	gammaRate        []string
	epsilonRate      []string
)

func CalculatePoints(bin string) {
	ws_stripped_bin := strings.TrimSpace(bin)

	if ws_stripped_bin == "" {
		return
	}

	for index, digit := range []byte(ws_stripped_bin) {
		if digit == 49 {
			firstFiveDigits[index] += 1
		} else {
			secondFiveDigits[index] += 1
		}
	}
}

func AssignRates() {
	for i := 0; i < 12; i++ {
		if firstFiveDigits[i] > secondFiveDigits[i] {
			gammaRate[i] = "1"
			epsilonRate[i] = "0"
		} else {
			gammaRate[i] = "0"
			epsilonRate[i] = "1"
		}
	}
}

func ComputeResults() int64 {
	gammaBinValue := strings.Join(gammaRate, "")
	epsilonBinValue := strings.Join(epsilonRate, "")

	gammaIntValue, _ := strconv.ParseInt(gammaBinValue, 2, 16)
	epsilonIntValue, _ := strconv.ParseInt(epsilonBinValue, 2, 16)

	return gammaIntValue * epsilonIntValue
}

func main() {
	firstFiveDigits = [...]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	secondFiveDigits = [...]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	gammaRate = make([]string, 12)
	epsilonRate = make([]string, 12)

	body, err := ioutil.ReadFile("input_case.txt")

	if err != nil {
		log.Fatalf("Couldn't Open File For Reading: %v", err)
	}

	bodySlice := strings.Split(string(body), "\n")

	for _, elem := range bodySlice {
		CalculatePoints(elem)
	}

	AssignRates()

	fmt.Println(ComputeResults())

}
