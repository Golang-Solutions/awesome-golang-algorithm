package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	solution func Info
type SolutionFuncType func(*TreeNode, *TreeNode, *TreeNode) *TreeNode

var SolutionFuncList = []SolutionFuncType{
	lowestCommonAncestor,
}

//	test info struct
type Case struct {
	name   string
	root   *TreeNode
	p      *TreeNode
	q      *TreeNode
	expect *TreeNode
}

// 	test case
var cases = []Case{
	{
		name: "TestCase 1",
		root: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 0},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 5},
				},
			},
			Right: &TreeNode{
				Val:   8,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 9},
			},
		},
		p:      &TreeNode{Val: 2},
		q:      &TreeNode{Val: 8},
		expect: &TreeNode{Val: 6},
	},
	{
		name: "TestCase 2",
		root: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 0},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 5},
				},
			},
			Right: &TreeNode{
				Val:   8,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 9},
			},
		},
		p:      &TreeNode{Val: 2},
		q:      &TreeNode{Val: 4},
		expect: &TreeNode{Val: 2},
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.root, c.p, c.q)
				ast.Equal(c.expect.Val, actual.Val,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
