package main

import (
	"JorgPrueba/postFix"
	"fmt"
)

func main() {
	operators := postFix.Operators{
		Types: map[string]postFix.Operation{
			"+": &postFix.Sum{},
			"-": &postFix.Subtraction{},
			"*": &postFix.Multiply{},
			"/": &postFix.Division{},
		},
	}

	var infix string
	fmt.Print("Introduce una expresión infija: ")
	_, err := fmt.Scanln(&infix)
	if err != nil {
		fmt.Println("Error leyendo la expresión:", err)
		return
	}

	elements, err := postFix.SplitInfix(infix)
	if err != nil {
		return
	}

	expPostFix, err := postFix.InfixToPostFix(elements, operators)
	if err != nil {
		return
	}

	result, err := postFix.EvaluatePostFix(expPostFix)
	if err != nil {
		return
	}

	fmt.Println("")
	fmt.Printf("Infija: %s\n", infix)
	fmt.Printf("Posfija: %s\n", expPostFix)
	fmt.Printf("resultado: %f", result)
}
