//package string_sum
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(StringSum(" 3+5 "))
}

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = fmt.Errorf("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = fmt.Errorf("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function

func StringSum(input string) (output string, err error) {
	input = strings.TrimSpace(input)

	if len(input) == 0 {
		return "", errorEmptyInput
	}

	operand := []string{}
	s := []int64{}

	for _, k := range input {
		if k == '+' || k == '-' {
			operand = append(operand, string(k))
		} else {
			d, err := strconv.ParseInt(string(k), 10, 64)
			if err != nil {
				return "", errorNotTwoOperands
			} else {
				s = append(s, d)
			}
		}
	}

	cntr := 0
	for _, t := range operand {
		if t == "-" {
			cntr++
		}
	}

	var sum int64

	if cntr%2 == 0 {
		for _, i := range s {
			sum += i
		}
	} else {
		for _, i := range s {
			sum -= i
		}
	}

	return string(sum), nil
}
