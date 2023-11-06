package calculator

import "errors"

func Calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 != 0 {
			return num1 / num2, nil
		} else {
			return 0, errors.New("division by zero")
		}
	default:
		return 0, errors.New("invalid operator")
	}
}
