package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	solution func Info
type SolutionFuncType func([]int) int

var SolutionFuncList = []SolutionFuncType{
	maxSubArray,
}

//	test info struct
type Case struct {
	name   string
	inputs []int
	expect int
}

// 	test case
var cases = []Case{
	{
		name:   "TestCase 1",
		inputs: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		expect: 6,
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.inputs)
				ast.Equal(c.expect, actual,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
