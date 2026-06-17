/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    third := &ListNode{}
	extra := 0
	current := third

	for l1 != nil || l2 != nil || extra != 0 {
		sum := extra
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		extra = sum/10

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return third.Next
}
