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
		inputs *ListNode
		expect *ListNode
	}{
		{"TestCase", MakeListNode([]int{1, 2, -3, 3, 1}), MakeListNode([]int{3, 1})},
		{"TestCase", MakeListNode([]int{1, 2, 3, -3, 4}), MakeListNode([]int{1, 2, 4})},
		{"TestCase", MakeListNode([]int{1, 2, 3, -3, -2}), MakeListNode([]int{1})},
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
