package stringsg_sum

//package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Sprint(StringSum(" -24 - 55 "))
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

	var operand = ""
	var str = ""
	var s []int

	for _, k := range input {
		if string(k) == "-" || string(k) == "+" {
			if operand == "-" && string(k) == "-" {
				operand = "+"
			} else {
				operand = "-"
			}

			d, err := strconv.Atoi(str)
			if len(s) == 2 {
				return "", errorNotTwoOperands
			} else {
				if err != nil {
					return "", err
				} else {
					s = append(s, d)
				}
			}
			str = ""
		} else {
			fmt.Println(string(k))
			str += string(k)
		}
	}

	var sum int

	for _, i := range s {
		switch operand {
		case "-":
			sum -= i
		case "+":
			sum += i
		}
	}

	return strconv.Itoa(sum), nil
}
