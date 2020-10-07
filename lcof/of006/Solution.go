package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

//	递归写法
func reversePrint(head *ListNode) []int {
	ans := make([]int, 0)
	if head == nil {
		return ans
	}
	ans = reversePrint(head.Next)
	ans = append(ans, head.Val)
	return ans
}

//	反转琏表
func reversePrint2(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	pre, cur, next, ans := &ListNode{}, head, head.Next, []int{}
	for cur != nil {
		next = cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}
	for pre.Next != nil {
		ans = append(ans, pre.Val)
		pre = pre.Next
	}
	return ans
}
