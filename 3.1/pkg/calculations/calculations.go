package calculations

import (
	"errors"
	"github.com/sirupsen/logrus"
)

func Calculate(num int64, log_flag bool) (int64, error) {
	if (num < 0) {
		return 0, errors.New("given number cannot be lower than zero")
	}
	if (num == 0) {
		return 0, errors.New("if given number will be 0, that will lead to 0 result")
	}

	if log_flag {
		logrus.Info("Start calculations...")
		logrus.Infof("Calculate %d!", num)
		defer logrus.Info("Calculations complete!")
	}

	result := int64(1)
	for i := int64(1); i <= num; i++ {
		result *= i
	}

	return result, nil
}

func FactorialSolver(num int64) int64 {
	if num == 0 {
		return 1
	}
	return num * FactorialSolver(num - 1)
}