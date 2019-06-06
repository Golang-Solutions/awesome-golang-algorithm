package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution1(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect int
	}{
		{"TestCase", []int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
		{"TestCase", []int{1, 2, 3, 4, 5}, 4},
		{"TestCase", []int{7, 6, 4, 3, 1}, 0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+strconv.Itoa(i), func(t *testing.T) {
			got := maxProfit(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect int
	}{
		{"TestCase", []int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
		{"TestCase", []int{1, 2, 3, 4, 5}, 4},
		{"TestCase", []int{7, 6, 4, 3, 1}, 0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+strconv.Itoa(i), func(t *testing.T) {
			got := maxProfit2(c.inputs)
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
