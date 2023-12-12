package main

import (
	calculators "calculator_test_work/internal/calculator"
	"fmt"
	"strconv"

	"calculator_test_work/internal/helper"
)

func main() {
	var err error

	fmt.Println("Введите выражение в формате [число] [операция] [число]:")

	firstNumber := ""
	secondNumber := ""
	operationType := ""
	_, err = fmt.Scanln(&firstNumber, &operationType, &secondNumber)
	if err != nil {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}

	calculator := calculators.Calculator{
		Operation: operationType,
	}

	isRomeFirstNumber := false
	if number, err := strconv.Atoi(firstNumber); err != nil {
		isRomeFirstNumber = true
		calculator.FirstNumber, err = helper.RomanToNumber(firstNumber)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		calculator.FirstNumber = number
	}

	isRomeSecondNumber := false
	if number, err := strconv.Atoi(secondNumber); err != nil {
		isRomeSecondNumber = true
		calculator.SecondNumber, err = helper.RomanToNumber(secondNumber)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		calculator.SecondNumber = number
	}

	if isRomeSecondNumber != isRomeFirstNumber {
		fmt.Println("оба аргумента должны быть одного типа")
	}

	result, err := calculator.Calculate()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if result < 1 && isRomeFirstNumber {
		fmt.Println("результат работы с римскими цифрами должен быть положительным.")
		return
	}

	if isRomeFirstNumber {
		romeResult, err := helper.IntToRoman(result)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(romeResult)
	} else {
		fmt.Println(result)
	}
}
