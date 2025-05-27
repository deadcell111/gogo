package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var a, b, op string
	fmt.Print("Input 2 variables in binary and operation: +,-,*,/: ")
	fmt.Scan(&a, &b, &op)
	fmt.Print("Result in Binary: ")
	result, err := BinaryCalculator(a, b, op)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
func BinaryCalculator(a, b, op string) (string, error) {
	var res int
	if CheckBinary(a) && CheckBinary(b) {
		num1 := BinaryToNum(a)
		num2 := BinaryToNum(b)
		switch op {
		case "+":
			res = num1 + num2
		case "-":
			res = num1 - num2
		case "*":
			res = num1 * num2
		case "/":
			res = num1 / num2
		default:
			return "", fmt.Errorf("unknown operation: %s")
		}
	} else {
		return "", fmt.Errorf("variables are not binary: %s")
	}
	return NumToBinary(res), nil
}
func CheckBinary(a string) bool {
	for _, ch := range a {
		if ch != '1' && ch != '0' {
			return false
		}
	}
	return true
}
func NumToBinary(x int) string {
	if x == 0 {
		return "0"
	}
	res := ""
	for x > 0 {
		digit := x % 2
		res = strconv.Itoa(digit) + res
		x /= 2
	}
	return res
}
func BinaryToNum(a string) int {
	res := 0
	for i := 0; i < len(a); i++ {
		if a[i] == '1' {
			res += 1 << (len(a) - i - 1)
		}
	}
	return res
}
