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
		{IntArray: []int{}},
		{IntArray: []int{0}},
		{IntArray: []int{1, -1}},
		{IntArray: []int{0, 0, 0}},
		{IntArray: []int{1, 0, 1}},
		{IntArray: []int{0, 0, 1, -1, 1, -1}},
	}
	var expectedOutputs []Output = []Output{
		{Triplets: [][3]int{}},
		{Triplets: [][3]int{}},
		{Triplets: [][3]int{}},
		{Triplets: [][3]int{{0, 0, 0}}},
		{Triplets: [][3]int{}},
		{Triplets: [][3]int{{-1, 0, 1}}},
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
	IntArray []int
}

type Output struct {
	Triplets [][3]int
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
		if len(output.Output.Triplets) != len(expectedOutputs[i].Triplets) {
			testedCases[i] = fmt.Errorf("expected output length(%d) != output length (%d)", len(output.Output.Triplets), len(expectedOutputs[i].Triplets))
			continue
		}
		switch len(output.Output.Triplets) {
		case 0:
			workingCases++
			continue
		case 1:
			if output.Output.Triplets[0] != expectedOutputs[i].Triplets[0] {
				testedCases[i] = fmt.Errorf("expected output != output")
				continue
			}
			workingCases++
			continue
		default:
			outputMap := toSortedMap(output.Output.Triplets)
			if len(output.Output.Triplets) != len(outputMap) {
				testedCases[i] = fmt.Errorf("output contains a duplicate triplet")
				continue
			}
			expectedOutputMap := toSortedMap(expectedOutputs[i].Triplets)
			for key := range expectedOutputMap {
				_, found := outputMap[key]
				if !found {
					testedCases[i] = fmt.Errorf("expected output ({%d,%d,%d}) not found in the actual output", key[0], key[1], key[2])
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
			log.Printf("%s●%s Test %d passed!", ColorGreen, ColorReset, i+1)
		} else {
			log.Printf("%s●%s Test %d did not pass: %30s", ColorRed, ColorReset, i+1, testedCase.Error())
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
