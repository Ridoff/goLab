package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strconv"
	"go/course/factorialset/pkg/calculations"
)

func main() {
	// Example usage with command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: main <number>")
		os.Exit(1)
	}

	// Parse input from command line
	num, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatalf("Error parsing input: %v", err)
	}

	// Enable log if the second command-line argument is "true"
	logFlag := false
	if len(os.Args) > 2 && os.Args[2] == "true" {
		logFlag = true
	}

	// Initialize logrus
	logrus.SetFormatter(&logrus.TextFormatter{})

	// Perform calculation
	result, err := calculations.Calculate(num, logFlag)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	// Print the result
	fmt.Printf("Factorial of %d is: %d\n", num, result)
}
