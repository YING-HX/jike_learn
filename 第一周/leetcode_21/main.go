package main

/*21.合并两个升序的链表*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	a := &ListNode{}
	head := a
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}
	if list1 == nil {
		head.Next = list2
	}
	if list2 == nil {
		head.Next = list1
	}
	return a.Next
}
