package Solution

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//	序列化
func dumpTreeToString(tree *TreeNode) string {
	var str string = ""
	dumpTreeToStringHelper(tree, &str)
	return str
}

func dumpTreeToStringHelper(tree *TreeNode, str *string) {
	if tree == nil {
		*str += "#"
		return
	}
	*str += strconv.Itoa(tree.Val)
	dumpTreeToStringHelper(tree.Left, str)
	dumpTreeToStringHelper(tree.Right, str)
}

//	反序列化
func conFromPreStr(str string) *TreeNode {
	var strIndex int = 0
	return conFromPreStrHelper(str, &strIndex)
}

func conFromPreStrHelper(str string, index *int) *TreeNode {
	if len(str) == *index {
		return nil
	}
	if (str)[*index] == '#' {
		(*index) = *index + 1
		return nil
	}
	v, _ := strconv.Atoi(string((str)[*index]))
	root := &TreeNode{Val: v}
	(*index) = *index + 1
	root.Left = conFromPreStrHelper(str, index)
	root.Right = conFromPreStrHelper(str, index)

	return root
}
