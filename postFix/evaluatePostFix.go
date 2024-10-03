package postFix

import (
	"fmt"
	"strconv"
)

func EvaluatePostFix(postFix []string) (float64, error) {
	var stack Stack
	opFactory := OperatorFactory{}

	for _, element := range postFix {

		if num, err := strconv.ParseFloat(element, 64); err == nil {
			stack.Push(num)
			continue
		} else {

			if len(stack.items) < 2 {
				return 0, fmt.Errorf("expresión inválida, faltan operandos")
			}

			operand1 := stack.Top()
			stack.Pop()
			operand2 := stack.Top()
			stack.Pop()

			operation, errOperation := opFactory.FactoryOperator(element)
			if errOperation != nil {
				return 0, errOperation
			}

			result := operation.Calculate(operand1, operand2)
			stack.Push(result)
		}
	}

	return stack.Top(), nil
}

type Stack struct {
	items []float64
}

func (s *Stack) Push(data float64) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Top() float64 {
	if s.IsEmpty() {
		return 0
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

