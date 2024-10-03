package postFix

import (
	"errors"
	"strconv"
)

func InfixToPostFix(elements []string, operators Operators) ([]string, error) {
	var postFix, stack []string

	for _, element := range elements {

		if _, err := strconv.ParseFloat(element, 64); err == nil {
			postFix = append(postFix, element)
			continue
		}

		if operators.IsOperator(element) {
			//lastElement := stack[len(stack)-1]
			for len(stack) > 0 && operators.IsOperator(stack[len(stack)-1]) && operators.OperatorPrecedence(stack[len(stack)-1], element) {
				postFix = append(postFix, stack[len(stack)-1])

				newStack := stack[:len(stack)-1]
				stack = newStack
			}
			stack = append(stack, element)
			continue
		}

		return nil, errors.New("error invalid element")
	}

	for len(stack) > 0 {
		postFix = append(postFix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return postFix, nil
}