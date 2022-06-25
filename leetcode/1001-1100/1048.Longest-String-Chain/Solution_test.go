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
		inputs []string
		expect int
	}{
		{"TestCase1", []string{"a", "b", "ba", "bca", "bda", "bdca"}, 4},
		{"TestCase2", []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}, 5},
		{"TestCase3", []string{"abcd", "dbqca"}, 1},
		{"TestCase4", []string{"a"}, 1},
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
