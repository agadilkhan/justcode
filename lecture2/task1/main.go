package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	ans := &ListNode{}
	cur := ans

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			cur.Next = list2
			list2 = list2.Next
		} else {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}

	for list1 != nil {
		cur.Next = list1
		cur = cur.Next
		list1 = list1.Next
	}

	for list2 != nil {
		cur.Next = list2
		cur = cur.Next
		list2 = list2.Next
	}

	return ans.Next
}
