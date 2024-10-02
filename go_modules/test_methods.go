package main

import (
	"fmt"
	"math"
)

/*
Issue
	Considering the actual flow, there would be two differents test methods
 */
func getTestPackage(req *requestPackage, choice int) {
	req.MaxIterations = 100
	req.Tolerance = math.Pow10(-7)
	req.TypeOfError = 1

	switch choice {
	case 1:
		fmt.Println("You chose Búsquedas (Searches).")
		req.TypeOfFunction = "BUS"
		req.FunctionData = getTestSearchData()
	case 2:
		fmt.Println("You chose Bisección (Bisection).")
		req.TypeOfFunction = "BIS"
		req.FunctionData = getTestBisectionData()
	case 3:
		fmt.Println("You chose Regla falsa (False Position).")
		req.TypeOfFunction = "FAL"
		req.FunctionData = getTestFalsePositionData()
	case 4:
		fmt.Println("You chose Punto fijo (Fixed Point).")
		req.TypeOfFunction = "PUN"
		req.FunctionData = getTestFixedPointData()
	case 5:
		fmt.Println("You chose Newton (Newton's Method).")
		req.TypeOfFunction = "NEW"
		req.FunctionData = getTestNewtonData()
	case 6:
		fmt.Println("You chose Secante (Secant Method).")
		req.TypeOfFunction = "SEC"
		req.FunctionData = getTestSecantData()
	case 7:
		fmt.Println("You chose Raíces múltiples (Multiple Roots).")
		req.TypeOfFunction = "MUL"
		req.FunctionData = getTestMultipleRootsData()
	default:
		fmt.Println("Invalid choice, please select a number between 1 and 7.")
		return
	}
	return
}

func getTestResult(req requestPackage) {

}

// VALIDADO
func getTestNewtonData() NewtonData {
	return NewtonData{
			Function:     "ln(sin(x)^2 + 1) - 1/2",  // f(x)
		Derivative:   "2*(sin(x)^2 + 1)^(-1)*sin(x)*cos(x)",  // f'(x)
		InitialValue: 0.5,  // x0
	}
}


// VALIDADO
func getTestMultipleRootsData() MultipleRootsData {
	return MultipleRootsData{
		Function:         "exp(x) - x - 1",  // h(x)
		FirstDerivative:  "exp(x) - 1",      // h'(x)
		SecondDerivative: "exp(x)",          // h''(x)
		InitialValue:     1.0,               // x0
	}
}


// VALIDADO
func getTestBisectionData() BisectionData {
	return BisectionData{
		Function: "ln(sin(x)^2 + 1) - 1/2",  // f(x)
		Interval: [2]float64{0, 1},          // Interval [a, b]
	}
}

// VALIDADO
func getTestFixedPointData() FixedPointData {
	return FixedPointData{
		Function: "ln(sin(x)^2 + 1) - 1/2 - x",  // f1(x)
		RearrangedFunctions: []string{
			"ln(sin(x)^2 + 1) - 1/2",  // g(x)
		},
		InitialValue: -0.5,  // x0
	}
}

// VALIDADA
func getTestSecantData() SecantData {
	return SecantData{
		Function:      "ln(sin(x)^2+1)-1/2",  // f(x)
		InitialValue1: 0.5,               // x0
		InitialValue2: 1.0,               // x1
	}
}

// Validado
func getTestSearchData() SearchData {
	return SearchData{
		InitialValue: -3,
		Function: "ln(sin(x)^2 + 1) - 1/2",
		Delta: 0.5,
	}
}


func getTestFalsePositionData() FalsePositionData {
	return FalsePositionData{
		Function: "ln(sin(x)^2 + 1) - 1/2",  // f(x)
		Interval: [2]float64{0, 1},          // Interval [a, b]
	}
}
