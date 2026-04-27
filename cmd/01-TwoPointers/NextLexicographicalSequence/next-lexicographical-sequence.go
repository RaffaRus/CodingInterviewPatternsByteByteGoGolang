package main

import (
	"fmt"
	"log"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
)

func main() {

	log.SetFlags(0)
	var inputs []string = []string{
		"",
		"aba",
		"fedcba",
		"pqr",
		"aabb",
		"knl",
		"hefg",
	}
	var expectedOutputs []string = []string{
		"",
		"baa",
		"abcdef",
		"prq",
		"abab",
		"lkn",
		"hegf",
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
	Output *string
	Error  error
}

func wrapper(f func(string) string, input string) (output string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("solution paniced: %v", r)
		}
	}()
	return f(input), nil
}

func verifySolution(outputs []OutputErr, expectedOutputs []string) {

	var numberOfCases = len(outputs)
	if len(outputs) != len(expectedOutputs) {
		log.Fatalf("lenght mismatch between output(%d) and expected output(%d)", len(outputs), len(expectedOutputs))
	}

	workingCases := 0
	testedCases := make([]error, numberOfCases)
	for i, output := range outputs {
		if output.Error != nil {
			testedCases[i] = fmt.Errorf("error encountered: %w", output.Error)
			continue
		}
		if *output.Output != expectedOutputs[i] {
			testedCases[i] = fmt.Errorf("expected output != output")
			continue
		} else {
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
			log.Printf("  %s●%s Test %d passed!", ColorGreen, ColorReset, i)
		} else {
			log.Printf("  %s●%s Test %d did not pass: %30s", ColorRed, ColorReset, i, testedCase.Error())
		}

	}

}
