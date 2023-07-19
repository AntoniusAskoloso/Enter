package main

import (
	"fmt"
	"strings"
)

type Calculator struct {
}

///////////////////////////////// Математические операции\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\

func (Calculator) plus(a, b int) int {

	return a + b
}
func (Calculator) minus(a, b int) int {

	return a - b
}

func (Calculator) mno(a, b int) int {

	return a * b
}
func (Calculator) del(a, b int) int {

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
	}

	if convertRomanToArabic(aStr) < 1 && convertRomanToArabic(bStr) > 0 || convertRomanToArabic(aStr) > 0 && convertRomanToArabic(bStr) < 1 {
		panic("Вывод ошибки, так как используются одновременно разные системы счисления.")
	}

	if a < 0 || a > 10 {
		panic("Вывод ошибки, на вход принимаются числа от 1 до 10 включительно.")
		return
	}
	if b < 0 || b > 10 {
		panic("Вывод ошибки, на вход принимаются числа от 1 до 10 включительно.")
		return
	}
	if len(input[0:operatorIndex]) > 3 || len(input[operatorIndex+1:]) > 3 {
		panic("Вывод ошибки, некорректный ввод.")
		return
	}

	var result int
	switch input[operatorIndex] {
	case '+':
		result = operator.plus(a, b)
	case '-':
		result = operator.minus(a, b)
	case '*':
		result = operator.mno(a, b)
	case '/':
		result = operator.del(a, b)
	default:
	}
	if operator.del(a, b) > 10 {
		panic("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
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
	if result < 0 || result > 10 {
		panic("Вывод ошибки, на вход принимаются числа от 1 до 10 включительно.")

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
