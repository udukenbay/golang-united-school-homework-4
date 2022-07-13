package string_sum

//package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(StringSum("24+55"))      //both_operands_positive
	fmt.Println(StringSum("-24+55"))     //first_operand_negative
	fmt.Println(StringSum("24-55"))      //"second_operand_negative"
	fmt.Println(StringSum("-24-55"))     //both_operands_negative
	fmt.Println(StringSum(" -24 - 55 ")) //with_whitespace
	fmt.Println(StringSum(""))           //empty_input
	fmt.Println(StringSum("11+23+43"))   //three_operands
	fmt.Println(StringSum("42"))         //one_operand
	fmt.Println(StringSum("78c+55"))     //letters_in_first_operand
	fmt.Println(StringSum("24+55f"))     //letters_in_second_operand
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
		return "", fmt.Errorf("input is empty")
	}

	stringRG, _ := regexp.Compile(`[a-z]`)
	stringsList := stringRG.FindAllString(input, -1)

	if len(stringsList) > 0 {
		i, err := strconv.Atoi(stringsList[0])
		if err != nil {
			return "", err
		}
	}

	numbersRG, _ := regexp.Compile(`[0-9]+`)
	numbersList := numbersRG.FindAllString(input, -1)

	if len(numbersList) == 1 || len(numbersList) == 3 {
		return "", fmt.Errorf("expecting two operands, but received more or less")
	}

	symbolsRG, _ := regexp.Compile(`[-+]`)
	symbolsList := symbolsRG.FindAllString(input, -1)

	var sum = 0

	a, err := strconv.Atoi(numbersList[0])
	b, err := strconv.Atoi(numbersList[1])
	if err != nil {
		return "", err
	}

	if len(symbolsList) == 1 {
		if symbolsList[0] == "-" {
			sum = a - b
		} else {
			sum = a + b
		}
	} else {
		if symbolsList[0] == "-" && symbolsList[1] == "-" {
			sum = -(a + b)
		} else if symbolsList[0] == "-" && symbolsList[1] == "+" {
			sum = b - a
		} else if symbolsList[0] == "+" && symbolsList[1] == "-" {
			sum = a - b
		} else {
			sum = a + b
		}
	}

	return strconv.Itoa(sum), nil
}
