package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var arabic = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

var roman = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter an expression: ")
	expression, _ := reader.ReadString('\n')

	result, flag := calculate(expression)
	if flag == 1 {
		resul := intToRoman(int(result))
		fmt.Printf("Result: %s\n", resul)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
}

func calculate(expression string) (int64, int64) {
	expression = strings.TrimSuffix(expression, "\n")

	// разбиение  на операнды и оператора
	tokens := strings.Split(expression, " ")
	if len(tokens) != 3 {
		fmt.Println("Выдача паники, так как формат математической  операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		os.Exit(1)
	}

	var result int64
	var flag int64
	// Определение системы счисления
	ArabicNumbers := "1234567890"
	containsA := strings.ContainsAny(tokens[0], ArabicNumbers)
	containsB := strings.ContainsAny(tokens[2], ArabicNumbers)
	if containsA == true && containsB == true {

		operand1, err := strconv.ParseInt(tokens[0], 10, 64)
		if err != nil {
			fmt.Println("Invalid operand 1:", err)
			os.Exit(1)
		}

		operand2, err := strconv.ParseInt(tokens[2], 10, 64)
		if err != nil {
			fmt.Println("Invalid operand 2:", err)
			os.Exit(1)
		}

		operator := tokens[1]

		switch operator {
		case "+":
			result = operand1 + operand2
		case "-":
			result = operand1 - operand2
		case "*":
			result = operand1 * operand2
		case "/":
			if operand2 == 0 {
				fmt.Println("Division by zero")
				os.Exit(1)
			}
			result = operand1 / operand2
		default:
			fmt.Println("Invalid operator:", operator)
			os.Exit(1)
		}

	} else if containsA == false && containsB == false {
		flag = 1
		operand1 := romanToArabic(tokens[0])
		operand2 := romanToArabic(tokens[2])
		operator := tokens[1]

		switch operator {
		case "+":
			result = operand1 + operand2
		case "-":
			result = operand1 - operand2
		case "*":
			result = operand1 * operand2
		case "/":
			if operand2 == 0 {
				fmt.Println("Division by zero")
				os.Exit(1)
			}
			result = operand1 / operand2
		default:
			fmt.Println("Invalid operator:", operator)
			os.Exit(1)
			// RESULT 10_>X
		}
		if result < 0 {
			fmt.Println("Выдача паники, так как в римской системе нет отрицательных чисел.")
			os.Exit(1)
		}
	} else {
		fmt.Println("Выдача паники, так как используются одновременно разные системы счисления.")
	}
	return result, flag
}

func romanToArabic(roman string) int64 {
	romanMap := map[string]int64{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	var arabic int64 = 0
	i := 0
	for i < len(roman) {
		if i+1 < len(roman) && romanMap[roman[i:i+2]] > 0 {
			arabic += romanMap[roman[i:i+2]]
			i += 2
		} else {
			arabic += romanMap[string(roman[i])]
			i++
		}
	}
	return arabic
}
func intToRoman(number int) (result string) {

	if number >= 4000 || number <= 0 {
		return result
	}
	result = ""
	for i := len(arabic) - 1; i >= 0; i-- {
		for number >= arabic[i] {
			number -= arabic[i]
			result += roman[i]
		}
	}
	return result
}
