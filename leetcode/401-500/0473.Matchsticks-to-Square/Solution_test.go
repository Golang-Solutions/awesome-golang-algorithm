package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect bool
	}{
		{"TestCase1", []int{1, 1, 2, 2, 2}, true},
		{"TestCase2", []int{3, 3, 3, 3, 4}, false},
		{"TestCase3", []int{6, 6, 6, 6, 4, 2, 2}, false},
		{"TestCase4", []int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}, true},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
