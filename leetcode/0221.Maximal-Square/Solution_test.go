package Solution

import (
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]byte
		expect int
	}{
		{"TestCase", [][]byte{
			{1, 0, 1, 0, 0},
			{1, 0, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 0, 0, 1, 0},
		}, 4},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			//got := maximalSquare(c.inputs)
			//if !reflect.DeepEqual(got, c.expect) {
			//	t.Fatalf("expected: %v, but got: %v, with inputs: %v",
			//		c.expect, got, c.inputs)
			//}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {

}

//	使用案列
func ExampleSolution() {

}
