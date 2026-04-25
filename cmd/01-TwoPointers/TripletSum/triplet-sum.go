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
	var inputs []Input = []Input{
		{IntAscendingSortedArray: []int{}},
		{IntAscendingSortedArray: []int{0}},
		{IntAscendingSortedArray: []int{1, -1}},
		{IntAscendingSortedArray: []int{0, 0, 0}},
		{IntAscendingSortedArray: []int{1, 0, 1}},
		{IntAscendingSortedArray: []int{0, 0, 1, -1, 1, -1}},
	}
	var expectedOutputs []Output = []Output{
		{Indexes: [][3]int{}},
		{Indexes: [][3]int{}},
		{Indexes: [][3]int{}},
		{Indexes: [][3]int{{0, 0, 0}}},
		{Indexes: [][3]int{}},
		{Indexes: [][3]int{{-1, 0, 1}}},
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

type Input struct {
	IntAscendingSortedArray []int
}

type Output struct {
	Indexes [][3]int
}

type OutputErr struct {
	Output *Output
	Error  error
}

func wrapper(f func(Input) Output, input Input) (output Output, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("solution paniced: %v", r)
		}
	}()
	return f(input), nil
}

func verifySolution(outputs []OutputErr, expectedOutputs []Output) {

	var numberOfCases = len(outputs)
	if len(outputs) != len(expectedOutputs) {
		log.Fatalf("lenght mismatch between output(%d) and expected output(%d)", len(outputs), len(expectedOutputs))
	}

	workingCases := 0
	testedCases := make([]error, numberOfCases)
	for i, output := range outputs {
		if len(output.Output.Indexes) != len(expectedOutputs[i].Indexes) {
			testedCases[i] = fmt.Errorf("expected output length != output length")
			continue
		}
		switch len(output.Output.Indexes) {
		case 0:
			workingCases++
			continue
		case 1:
			if output.Output.Indexes[0] != expectedOutputs[i].Indexes[0] {
				testedCases[i] = fmt.Errorf("expected output != output")
				continue
			}
			workingCases++
			continue
		default:
			outputMap := toSortedMap(output.Output.Indexes)
			expectedOutputMap := toSortedMap(expectedOutputs[i].Indexes)
			for key := range expectedOutputMap {
				_, found := outputMap[key]
				if !found {
					testedCases[i] = fmt.Errorf("expected output ({%d,%d}) not found in the actual output", key[0], key[1])
					continue
				}
			}
			workingCases++
		}
	}

	if workingCases == numberOfCases {
		log.Printf("%s●%s All of %d test cases passed!", ColorGreen, ColorReset, numberOfCases)
		return
	}

	log.Printf("%s●%s Some of the %d test did not pass:", ColorYellow, ColorReset, numberOfCases)
	for i, testedCase := range testedCases {
		if testedCase == nil {
			log.Printf("%s●%s Test %d passed!", ColorGreen, ColorReset, i)
		} else {
			log.Printf("%s●%s Test %d did not pass: %30s", ColorRed, ColorReset, i, testedCase.Error())
		}

	}

}

func toSortedMap(ins [][3]int) map[[3]int]struct{} {
	out := make(map[[3]int]struct{}, len(ins))
	for _, in := range ins {
		slices.Sort(in[:])
		hash := in
		out[hash] = struct{}{}
	}
	return out
}
