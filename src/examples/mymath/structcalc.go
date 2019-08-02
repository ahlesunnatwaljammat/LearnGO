package mymath

type calculator struct {
}

func New() calculator {
	return calculator{}
}

func (calculator) Sum(i, j int) int {
	return i + j
}

func (c calculator) Subtract(i int, j int) interface{} {
	return i - j
}

func (calculator) Multiply(i int, j int) interface{} {
	return i - j
}

func (calculator) Divide(i, j float64) float64 {
	return i / j
}
