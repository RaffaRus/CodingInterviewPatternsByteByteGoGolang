package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func YourSolution(input Input) Output {
	return Output{}
}

func main() {
	var inputs []Input = []Input{
		{IntAscendingSortedArray: []int{}, Target: 0},
		{IntAscendingSortedArray: []int{1}, Target: 0},
		{IntAscendingSortedArray: []int{2, 3}, Target: 0},
		{IntAscendingSortedArray: []int{2, 4}, Target: 0},
		{IntAscendingSortedArray: []int{2, 3, 4}, Target: 0},
		{IntAscendingSortedArray: []int{-1, 2, 3}, Target: 0},
		{IntAscendingSortedArray: []int{-3, -2, -1}, Target: 0},
	}
	var expectedOutputs []Output = []Output{
		{Indexes: [][2]int{}},
		{Indexes: [][2]int{}},
		{Indexes: [][2]int{{0, 1}}},
		{Indexes: [][2]int{}},
		{Indexes: [][2]int{{1, 2}, {0, 2}}},
		{Indexes: [][2]int{{0, 2}}},
		{Indexes: [][2]int{{0, 1}}},
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

	fmt.Println("Apparently your solution worked, well done!")
}

type Input struct {
	IntAscendingSortedArray []int
	Target                  int
}

type Output struct {
	Indexes [][2]int
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
			for key, _ := range expectedOutputMap {
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
		log.Infof("All of %d test cases passed!")
		return
	}

	log.Warnf("Some of the %d test did not pass:")
	for i, testedCase := range testedCases {
		if testedCase == nil {
			log.Infof("Test %d passed!", i)
		} else {
			log.Infof("Test %d did not pass: %30w", i, testedCase.Error())
		}

	}

}

// TODO: make this take as input an Hashable element to create the key of the map
func toSortedMap(ins [][2]int) map[[2]int]struct{} {
	out := make(map[[2]int]struct{}, len(ins))
	for _, in := range ins {
		hash := [2]int{in[1], in[0]}
		if in[1] > in[0] {
			hash = [2]int{in[0], in[1]}
		}
		out[hash] = struct{}{}
	}
	return out
}
