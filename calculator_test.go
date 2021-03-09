package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	/*numbers := []int{15, 4}
	sum := 0
	for _, number := range numbers {
		sum += number
	}*/
	var want float64 = 19
	got := calculator.Add(15, 4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = -2
	got := calculator.Subtract(2, 4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
		//t.Skip("Tests may be skipped")
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 0
	got := calculator.Multiply(5, 0)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	var want float64 = 5
	got := calculator.Divide(20, 4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.SquareRoot(16)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
