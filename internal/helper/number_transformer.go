package helper

import (
	"errors"
)

var romanNumbersToEncode = []struct {
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

var romanNumbersToDecode = map[int]string{
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X",
}

func IntToRoman(num int) (string, error) {
	if num <= 0 || num > 100 {
		return "", errors.New("недопустимое число")
	}

	var result string
	for _, pair := range romanNumbersToEncode {
		for num >= pair.Value {
			result += pair.Symbol
			num -= pair.Value
		}
	}
	return result, nil
}

func RomanToNumber(number string) (int, error) {
	for key, val := range romanNumbersToDecode {
		if val == number {
			return key, nil
		}
	}
	return 0, errors.New("недопустимый римский символ")
}
