package main

import (
	"strings"
	"unicode"
)

// Función principal que preprocesa la expresión
func preprocessExpression(expression string) string {
	// Eliminar espacios en blanco innecesarios
	expression = strings.ReplaceAll(expression, " ", "")
	// Iniciar el proceso de reemplazo
	result, _ := replacePowers(expression)
	return result
}

// Función recursiva que reemplaza los operadores ^ por pow()
func replacePowers(expression string) (string, int) {
	var result strings.Builder
	i := 0
	for i < len(expression) {
		ch := expression[i]
		if ch == '(' {
			// Extraer la subexpresión completa dentro de los paréntesis
			subExpr, length := matchParentheses(expression[i:])
			// Procesar la expresión interna recursivamente
			subExprProcessed, _ := replacePowers(subExpr[1 : len(subExpr)-1]) // Quitamos los paréntesis externos
			result.WriteString("(" + subExprProcessed + ")")
			i += length - 1 // Ajustar el índice 'i' correctamente
	} else if ch == '^' {
			// Encontrar la base y el exponente
			baseExpr, baseLength := extractBase(result.String())
			exponentExpr, exponentLength := extractExponent(expression[i+1:])
			// Reemplazar por pow(base, exponente)
			powExpr := "pow(" + baseExpr + "," + exponentExpr + ")"
			// Ajustar el resultado y los índices
			resultStr := result.String()
			result.Reset()
			result.WriteString(resultStr[:len(resultStr)-baseLength])
			result.WriteString(powExpr)
			i += exponentLength // Saltar el exponente procesado
		} else if isFunctionStart(expression[i:]) {
			// Procesar funciones matemáticas
			funcExpr, length := extractFunction(expression[i:])
			result.WriteString(funcExpr)
			i += length - 1 // Ajustar índice
		} else {
			result.WriteByte(ch)
		}
		i++
	}
	return result.String(), i
}

// Función para extraer la base antes del operador ^
func extractBase(expression string) (string, int) {
	i := len(expression) - 1
	parenthesisCounter := 0
	for i >= 0 {
		ch := expression[i]
		if ch == ')' {
			parenthesisCounter++
		} else if ch == '(' {
			parenthesisCounter--
		}
		if parenthesisCounter == 0 && (ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '^') {
			break
		}
		i--
	}
	return expression[i+1:], len(expression) - (i + 1)
}


// Función para extraer el exponente después del operador ^
func extractExponent(expression string) (string, int) {
	if len(expression) == 0 {
		return "", 0
	}
	if expression[0] == '(' {
		// Exponente es una expresión entre paréntesis
		subExpr, length := matchParentheses(expression)
		return subExpr, length
	} else if isLetter(expression[0]) {
		// Exponente es una variable o función
		name, length := extractName(expression)
		if len(expression) > length && expression[length] == '(' {
			// Es una función
			funcExpr, funcLength := extractFunction(expression)
			return funcExpr, funcLength
		}
		return name, length
	} else {
		// Exponente es un número
		number, length := extractNumber(expression)
		return number, length
	}
}

// Función para extraer una función matemática
func extractFunction(expression string) (string, int) {
	name, nameLength := extractName(expression)
	if len(expression) <= nameLength || expression[nameLength] != '(' {
		return name, nameLength
	}
	// Encontrar la expresión dentro de los paréntesis
	subExpr, length := matchParentheses(expression[nameLength:])
	totalLength := nameLength + length
	// Procesar la expresión interna recursivamente
	subExprProcessed, _ := replacePowers(subExpr[1 : len(subExpr)-1]) // Quitamos los paréntesis externos
	return name + "(" + subExprProcessed + ")", totalLength
}


// Función para extraer un nombre (variable o función)
func extractName(expression string) (string, int) {
	i := 0
	for i < len(expression) && isLetter(expression[i]) {
		i++
	}
	return expression[:i], i
}

// Función para extraer un número
func extractNumber(expression string) (string, int) {
	i := 0
	for i < len(expression) && (unicode.IsDigit(rune(expression[i])) || expression[i] == '.') {
		i++
	}
	return expression[:i], i
}

// Función para verificar si es el inicio de una función matemática
func isFunctionStart(expression string) bool {
	functions := []string{"sin", "cos", "tan", "ln", "log", "sqrt"}
	for _, f := range functions {
		if strings.HasPrefix(expression, f+"(") {
			return true
		}
	}
	return false
}

// Función para verificar si un carácter es una letra
func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch))
}

// Función para hacer coincidir los paréntesis y extraer la expresión interna
func matchParentheses(expression string) (string, int) {
	if expression[0] != '(' {
		return "", 0
	}
	parenthesisCounter := 0
	i := 0
	for i < len(expression) {
		ch := expression[i]
		if ch == '(' {
			parenthesisCounter++
		} else if ch == ')' {
			parenthesisCounter--
			if parenthesisCounter == 0 {
				return expression[:i+1], i + 1
			}
		}
		i++
	}
	// Si llegamos aquí, los paréntesis están desbalanceados
	return expression, i
}



