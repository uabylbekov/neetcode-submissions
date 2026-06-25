/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}

	current := head

	for current != nil && current.Next != nil {
		gdc := getGcd(current.Val, current.Next.Val)
		next := &ListNode{Val: gdc, Next: current.Next}
		current.Next = next

		current = next.Next
	}

	return head
}

func getGcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}
