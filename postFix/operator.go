package postFix

import "errors"

type Operators struct {
	Types map[string]Operation
}

func (o *Operators) IsOperator(element string) bool {
	return o.Types[element] != nil
}

func (o *Operators) Get(element string) (Operation, error) {
	operation := o.Types[element]
	if operation == nil {
		return nil, errors.New("operation invalid")
	}
	return operation, nil
}

func (o *Operators) OperatorPrecedence(a, b string) bool {
	return precedence[a] >= precedence[b]
}

var precedence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}
