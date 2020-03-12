package math

func Calc(operation string) func(a, b float64) float64 {
	switch operation {
	case "+":
		return func(a, b float64) float64 {
			return a + b
		}
	case "-":
		return func(a, b float64) float64 {
			return a - b
		}
	case "*":
		return func(a, b float64) float64 {
			return a * b
		}
	case "/":
		return func(a, b float64) float64 {
			return a / b
		}
	default:
		return func(a, b float64) float64 {
			return a + b
		}
	}
}
