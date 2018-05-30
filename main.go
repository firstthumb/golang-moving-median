package main

import (
	"movingmedian/slidingwindow"
	"io/ioutil"
	"strings"
	"movingmedian/median"
	"strconv"
	"fmt"
)

type Input struct {
	filename 	string
	windowSize	int
}

var INPUTS = []Input{
	{"test.csv", 3},
	{"Round 1. Software engineering test cases - test2.csv", 100},
	{"Round 1. Software engineering test cases - test3.csv", 1000},
	{"Round 1. Software engineering test cases - test4.csv", 10000},
}

func main() {
	for _, input := range INPUTS {
		processFile(input.filename, input.windowSize)
	}

	fmt.Println("** COMPLETED **")
}

func processFile(filename string, windowSize int) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	inputs := strings.Split(string(dat), "\n")
	outputs := calculateMedians(windowSize, inputs)

	ioutil.WriteFile("output_" + filename, []byte(strings.Join(outputs[:], "\n")), 0644)

	fmt.Println("\n\nOutput file is created successfully => " + ("output_" + filename))

	return nil
}

func calculateMedians(windowSize int, inputs []string) []string {
	w := slidingwindow.New(windowSize)
	outputs := make([]string, len(inputs))
	for i, input := range inputs {
		value, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Printf("Cannot convert input to integer : %v", err)
			continue
		}
		w.Add(value)
		medianValue, err := median.FindMedian(w.Slice())
		if err != nil {
			outputs[i] = "-1"
		} else {
			outputs[i] = strconv.FormatFloat(medianValue, 'f', 1, 64)
		}
	}
	return outputs
}

