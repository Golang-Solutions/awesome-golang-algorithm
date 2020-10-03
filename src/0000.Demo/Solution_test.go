package Solution

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)
	//	test cases
	cases := []struct {
		name   string
		inputs bool
		expect bool
	}{
		{"TestCase", true, true},
		{"TestCase", true, true},
		{"TestCase", false, false},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := Solution(c.inputs)
			ast.Equal(actual, c.expect, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.inputs)
		})
	}
}

//	Benchmark Test
func BenchmarkSolution(b *testing.B) {

}

//	Example Test
func ExampleSolution() {

}
