package main

import (
	"bufio"
	"fmt"
	"os"
)

type requestPackage struct {
	TypeOfFunction string
	FunctionData   interface{}
	MaxIterations  int
	Tolerance      float64
	TypeOfError    int
}

type SearchData struct {
	InitialValue float64
	Delta        float64
	Function     string
}

type NewtonData struct {
	Function     string
	Derivative   string
	InitialValue float64
}

type MultipleRootsData struct {
	Function         string
	FirstDerivative  string
	SecondDerivative string
	InitialValue     float64
}

type BisectionData struct {
	Function string
	Interval [2]float64
}

type FixedPointData struct {
	Function     string   // This is the main function f(x)
	RearrangedFunctions []string // Array of rearranged functions (can be one or more)
	InitialValue float64  // Starting value x0
}


type SecantData struct {
	Function      string
	InitialValue1 float64
	InitialValue2 float64}

type FalsePositionData struct {
	Function string
	Interval [2]float64
}

/*
Elements of the system:

	Main (CLI)
	Test results (For both testing the program and provide default values for the page)
	Numerical Methods (The set of functions used to iterate over a function)
	Auxiliar functions (Functions that serves a secondary goal, for example, DRY)
	Server interface.
*/
func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		var req requestPackage
		fmt.Println("Choose the task")
		fmt.Println("1. Test a method")
		fmt.Println("2. Evaluate a function")
		choiceString := "Enter the number of your choice: "
		taskChoice := getIntInput(reader, choiceString)

		// Display the menu to the user
		fmt.Println("Choose a mathematical method:")
		fmt.Println("1. Búsquedas (Searches)")
		fmt.Println("2. Bisección (Bisection)")
		fmt.Println("3. Regla falsa (False Position)")
		fmt.Println("4. Punto fijo (Fixed Point)")
		fmt.Println("5. Newton (Newton's Method)")
		fmt.Println("6. Secante (Secant Method)")
		fmt.Println("7. Raíces múltiples (Multiple Roots)")
		choice := getIntInput(reader, "Enter the number of your choice: ")
		switch taskChoice {
		case 1:
			fmt.Println("It entered case 1")
			getTestPackage(&req, choice)
			break
		case 2:
			fmt.Println("It entered case 2")
			buildPackage(reader, &req, choice)
			break
		default:
			fmt.Println("The number entered is not a valid choice.")
			continue
		}

		result, err := iterateNumericMethod(req)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Result: %.16f\n", result)
		}
	}
}



func buildPackage(reader *bufio.Reader, req *requestPackage, choice int) {
	req.MaxIterations = getIntInput(reader, "Enter the maximum number of iterations: ")
	req.Tolerance = getFloatInput(reader, "Enter the tolerance: ")
	req.TypeOfError = getIntInput(reader, "Enter the type of error (1 for absolute error, 2 for percentage error): ") // Switch statement to handle user's choice and function-specific data
	switch choice {
	case 1:
		fmt.Println("You chose Búsquedas (Searches).")
		req.TypeOfFunction = "BUS"
	case 2:
		fmt.Println("You chose Bisección (Bisection).")
		req.TypeOfFunction = "BIS"
		req.FunctionData = collectBisectionData(reader)
	case 3:
		fmt.Println("You chose Regla falsa (False Position).")
		req.TypeOfFunction = "FAL"
	case 4:
		fmt.Println("You chose Punto fijo (Fixed Point).")
		req.TypeOfFunction = "PUN"
		req.FunctionData = collectFixedPointData(reader)
	case 5:
		fmt.Println("You chose Newton (Newton's Method).")
		req.TypeOfFunction = "NEW"
		req.FunctionData = collectNewtonData(reader)
	case 6:
		fmt.Println("You chose Secante (Secant Method).")
		req.TypeOfFunction = "SEC"
		req.FunctionData = collectSecantData(reader)
	case 7:
		fmt.Println("You chose Raíces múltiples (Multiple Roots).")
		req.TypeOfFunction = "MUL"
		req.FunctionData = collectMultipleRootsData(reader)
	default:
		fmt.Println("Invalid choice, please select a number between 1 and 7.")
		return
	}

	// Call the central iteration function

}

// Helper function to collect Newton's method data
func collectNewtonData(reader *bufio.Reader) NewtonData {
	function := getStringInput(reader, "Enter the function (in terms of x): ")
	derivative := getStringInput(reader, "Enter the derivative of the function: ")
	initialValue := getFloatInput(reader, "Enter the initial value: ")
	return NewtonData{
		Function:     function,
		Derivative:   derivative,
		InitialValue: initialValue,
	}
}

// Helper function to collect data for Multiple Roots
func collectMultipleRootsData(reader *bufio.Reader) MultipleRootsData {
	function := getStringInput(reader, "Enter the function (in terms of x): ")
	firstDerivative := getStringInput(reader, "Enter the first derivative: ")
	secondDerivative := getStringInput(reader, "Enter the second derivative: ")
	initialValue := getFloatInput(reader, "Enter the initial value: ")
	return MultipleRootsData{
		Function:         function,
		FirstDerivative:  firstDerivative,
		SecondDerivative: secondDerivative,
		InitialValue:     initialValue,
	}
}

// Helper function to collect data for Bisection
func collectBisectionData(reader *bufio.Reader) BisectionData {
	function := getStringInput(reader, "Enter the function (in terms of x): ")
	a := getFloatInput(reader, "Enter the start of the interval: ")
	b := getFloatInput(reader, "Enter the end of the interval: ")
	return BisectionData{
		Function: function,
		Interval: [2]float64{a, b},
	}
}

// Helper function to collect data for Fixed Point
func collectFixedPointData(reader *bufio.Reader) FixedPointData {
	function := getStringInput(reader, "Enter the function (in terms of x): ")
	initialValue := getFloatInput(reader, "Enter the initial value: ")
	return FixedPointData{
		Function:     function,
		InitialValue: initialValue,
	}
}

// Helper function to collect data for Secant Method
func collectSecantData(reader *bufio.Reader) SecantData {
	function := getStringInput(reader, "Enter the function (in terms of x): ")
	initialValue1 := getFloatInput(reader, "Enter the first initial value: ")
	initialValue2 := getFloatInput(reader, "Enter the second initial value: ")
	return SecantData{
		Function:      function,
		InitialValue1: initialValue1,
		InitialValue2: initialValue2,
	}
}
