package main

import (
	"fmt"
	"log"
	"slices"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
)

func main() {

	log.SetFlags(0)
	var inputs [][]int = [][]int{
		{},
		{0},
		{1},
		{0, 1},
		{1, 2},
		{0, 1, 2, 3, 0, 4, 5, 0},
		{1, 2, 3, 4, 5},
		{0, 1, 2, 3, 4, 5, 0},
	}
	var expectedOutputs [][]int = [][]int{
		{},
		{0},
		{1},
		{1, 0},
		{1, 2},
		{1, 2, 3, 4, 5, 0, 0, 0},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 0, 0},
	}
	if len(inputs) != len(expectedOutputs) {
		log.Fatalf("error with the test setup: len(input) != len(expectedOutput)")
	}

	var outputs []OutputErr = make([]OutputErr, len(expectedOutputs))
	for i, input := range inputs {
		output, err := wrapper(YourSolution, input)
		outputs[i] = OutputErr{
			Output: &output,
			Error:  err,
		}
	}
	verifySolution(outputs, expectedOutputs)
}

type OutputErr struct {
	Output *[]int
	Error  error
}

func wrapper(f func([]int) []int, input []int) (output []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("solution paniced: %v", r)
		}
	}()
	return f(input), nil
}

func verifySolution(outputs []OutputErr, expectedOutputs [][]int) {

	var numberOfCases = len(outputs)
	if len(outputs) != len(expectedOutputs) {
		log.Fatalf("lenght mismatch between output(%d) and expected output(%d)", len(outputs), len(expectedOutputs))
	}

	workingCases := 0
	testedCases := make([]error, numberOfCases)
	for i, output := range outputs {
		if len(*output.Output) != len(expectedOutputs[i]) {
			testedCases[i] = fmt.Errorf("expected output length != output length")
			continue
		}
		switch len(*output.Output) {
		case 0:
			workingCases++
			continue
		default:
			if slices.Equal(*output.Output, expectedOutputs[i]) {
				workingCases++
				continue
			} else {
				testedCases[i] = fmt.Errorf("expected output != output ")
				continue
			}
		}
	}

	if workingCases == numberOfCases {
		log.Printf("%s●%s All of %d test cases passed!", ColorGreen, ColorReset, numberOfCases)
		return
	}

	log.Printf("%s●%s Some of the %d test did not pass:", ColorYellow, ColorReset, numberOfCases)
	for i, testedCase := range testedCases {
		if testedCase == nil {
			log.Printf("  %s●%s Test %d passed!", ColorGreen, ColorReset, i)
		} else {
			log.Printf("  %s●%s Test %d did not pass: %30s", ColorRed, ColorReset, i, testedCase.Error())
		}

	}

}
