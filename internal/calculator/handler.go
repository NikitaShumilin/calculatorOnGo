package calculator

import (
	"errors"
)

type Calculator struct {
	FirstNumber  int
	SecondNumber int
	Operation    string
}

func (c *Calculator) Calculate() (int, error) {
	err := c.validate()
	if err != nil {
		return 0, err
	}
	switch c.Operation {
	case "+":
		return c.FirstNumber + c.SecondNumber, nil
	case "-":
		return c.FirstNumber - c.SecondNumber, nil
	case "*":
		return c.FirstNumber * c.SecondNumber, nil
	case "/":
		return c.FirstNumber / c.SecondNumber, nil
	}
	return 0, errors.New("операция не была проведена, неизвестная ошибка")
}

func (c *Calculator) validate() error {
	if c.FirstNumber <= 0 || c.FirstNumber > 10 {
		return errors.New("некореектное первое число")
	}

	if c.SecondNumber <= 0 || c.SecondNumber > 10 {
		return errors.New("некореектное второе число")
	}

	_, ok := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}[c.Operation]

	if !ok {
		return errors.New("неверный оператор")
	}

	return nil
}
