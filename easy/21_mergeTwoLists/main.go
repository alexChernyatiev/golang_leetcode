package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var res *ListNode
	if list1.Val < list2.Val {
		res = list1
		list1 = list1.Next
	} else {
		res = list2
		list2 = list2.Next
	}

	curr := res
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next
	}

	if list1 == nil {
		curr.Next = list2
	}

	if list2 == nil {
		curr.Next = list1
	}

	return res
}

func printList(list *ListNode) {
	curr := list
	for curr != nil {
		print(curr.Val, "->")
		curr = curr.Next
	}
	print("\n")
}

func main() {

	printList(
		mergeTwoLists(
			&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
			&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		),
	)
}
