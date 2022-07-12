package string_sum

//package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(StringSum("24+55"))
	fmt.Println(StringSum("-24+55"))
	fmt.Println(StringSum("24-55"))
	fmt.Println(StringSum("-24-55"))
	fmt.Println(StringSum(" -24 - 55 "))
	fmt.Println(StringSum(""))
	fmt.Println(StringSum("11+23+43"))
	fmt.Println(StringSum("42"))
	fmt.Println(StringSum("24c+55"))
	fmt.Println(StringSum("24+55f"))
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
		if string(k) == " " {
			continue
		}

		if string(k) == "-" || string(k) == "+" {
			if operand == "-" && string(k) == "-" {
				operand = "+"
			} else {
				operand = "-"
			}

			if len(str) > 0 {
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
			}
		} else {
			str += string(k)
		}
	}

	if len(s) == 1 {
		return "", errorNotTwoOperands
	}

	if len(str) != 0 {
		d, err := strconv.Atoi(str)
		if err != nil {
			return "", err
		} else {
			s = append(s, d)
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
