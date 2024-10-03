package postFix

type Operation interface {
	Calculate(a, b float64) float64
}

type Sum struct{}

func (s *Sum) Calculate(a, b float64) float64 {
	return a + b
}

type Subtraction struct{}

func (s *Subtraction) Calculate(a, b float64) float64 {
	return a - b
}

type Multiply struct{}

func (s *Multiply) Calculate(a, b float64) float64 {
	return a * b
}

type Division struct{}

func (s *Division) Calculate(a, b float64) float64 {
	return b / a
}
