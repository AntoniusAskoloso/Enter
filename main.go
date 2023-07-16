package main

import (
	"fmt"
	"strings"
)

type Calculator struct {
}

///////////////////////////////// Математические операции\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\

func (Calculator) add(a, b int) int {
	return a + b
}
func (Calculator) subtract(a, b int) int {
	return a - b
}
func (Calculator) multiply(a, b int) int {
	return a * b
}
func (Calculator) divide(a, b int) int {
	return a / b

}

///////////////////////////////////// Проверка арабского числа\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
func (Calculator) isArabicNumber(num string) bool {
	for _, char := range num {
		if char < '0' || char > '9' {
			return false
		}

	}
	return true
}

////////////////////////////////////// Проверка римского числа\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
func (Calculator) isRomanNumber(num string) bool {
	romanNumerals := map[string]bool{

		"I":    true,
		"II":   true,
		"III":  true,
		"IV":   true,
		"V":    true,
		"VI":   true,
		"VII":  true,
		"VIII": true,
		"IX":   true,
		"X":    true,
		"L":    true,
	}

	return romanNumerals[num]
}

func main() {
	operator := Calculator{}

	fmt.Println("Введите выражение: ")
	input := ""
	fmt.Scanln(&input)

	input = strings.ReplaceAll(input, " ", "")
	operatorIndex := strings.IndexAny(input, "+-*/")

	if operatorIndex == -1 {
		panic("Вывод ошибки, так как строка не является математической операцией.")

	}

	aStr := input[0:operatorIndex]
	bStr := input[operatorIndex+1:]

	var a, b int
	var isArabic bool
	if operator.isArabicNumber(aStr) && operator.isArabicNumber(bStr) {
		fmt.Sscanf(aStr, "%d", &a)
		fmt.Sscanf(bStr, "%d", &b)
		isArabic = true
	} else if operator.isRomanNumber(aStr) && operator.isRomanNumber(bStr) {
		a = convertRomanToArabic(aStr)
		b = convertRomanToArabic(bStr)
		isArabic = false

	} else if convertRomanToArabic(aStr) < 1 && convertRomanToArabic(bStr) > 0 || convertRomanToArabic(aStr) > 0 && convertRomanToArabic(bStr) < 1 {
		panic("Вывод ошибки, так как используются одновременно разные системы счисления.")
	}
	var result int

	if len(input) > 7 {
		panic("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	switch input[operatorIndex] {
	case '+':
		result = operator.add(a, b)
	case '-':
		result = operator.subtract(a, b)
	case '*':
		result = operator.multiply(a, b)
	case '/':
		result = operator.divide(a, b)

	default:
	}

	if isArabic {
		fmt.Println("Результат:", result)
	} else {
		fmt.Println("Результат:", convertArabicToRoman(result))
	}

}

////////////////////////////////////// Функция для конвертации римского числа в арабское\\\\\\\\\\\\\\\\\\\\\\\\\\
func convertRomanToArabic(romanNum string) int {
	romanNumerals := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	result := 0
	prevValue := 0

	for i := len(romanNum) - 1; i >= 0; i-- {
		currValue := romanNumerals[string(romanNum[i])]

		if currValue >= prevValue {
			result += currValue
		} else {
			result -= currValue
		}

		prevValue = currValue
	}

	return result
}

////////////////////////////// Функция для конвертации арабского числа в римское\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
func convertArabicToRoman(arabicNum int) string {
	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""

	for _, numeral := range romanNumerals {
		for arabicNum >= numeral.Value {
			result += numeral.Symbol
			arabicNum -= numeral.Value
		}

	}
	if arabicNum < 0 {
		panic("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
	}

	return result
}
