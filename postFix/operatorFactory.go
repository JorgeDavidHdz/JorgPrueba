package postFix

import "fmt"

type OperatorFactory struct{}

func (of *OperatorFactory) FactoryOperator(op string) (Operation, error) {
	switch op {
	case "+":
		return &Sum{}, nil
	case "-":
		return &Subtraction{}, nil
	case "*":
		return &Multiply{}, nil
	case "/":
		return &Division{}, nil
	default:
		return nil, fmt.Errorf("operador no soportado: %s", op)
	}
}
