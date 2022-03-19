package calculator_test

import (
	calculator "github.com/faisal-ibrahim/testing-in-golang/getting-started-with-unittest-in-go"
	"testing"
)

type TestCase struct {
	value    int
	expected bool
	actual   bool
}

func TestCalculateIsArmstrong(t *testing.T) {
	t.Run("Should return true for 371", func(t *testing.T) {
		testCase := TestCase{
			value:    371,
			expected: true,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

	t.Run("Should return true for 370", func(t *testing.T) {
		testCase := TestCase{
			value:    370,
			expected: true,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})
}

func TestNegativeCalculateIsArmstrong(t *testing.T) {
	t.Run("Should return false for 350", func(t *testing.T) {
		testCase := TestCase{
			value:    350,
			expected: false,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

	t.Run("Should return false for 300", func(t *testing.T) {
		testCase := TestCase{
			value:    300,
			expected: false,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})
}
