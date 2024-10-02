package main

import (
	"bufio"
	"fmt"
	"strconv"
)

func getIntInput(reader *bufio.Reader, prompt string) int {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	value, _ := strconv.Atoi(input[:len(input)-2])
	println(value)
	return value
}

func getFloatInput(reader *bufio.Reader, prompt string) (value float64) {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	value, _ = strconv.ParseFloat(input[:len(input)-2], 64)
	println(value)
	return
}

func getStringInput(reader *bufio.Reader, prompt string) (input string) {
	fmt.Print(prompt)
	input, _ = reader.ReadString('\n')
	println(input)
	return
}

func calculateError(value float64, previousValue float64, typeOfError int) float64{
	//Absolute error
	if typeOfError == 1{
		return abs(value - previousValue)
	}else{
		return abs((value - previousValue)/value)*100
	}
}
