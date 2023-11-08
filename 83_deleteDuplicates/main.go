package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Println(head.Val)
	if head.Next != nil {
		fmt.Printf("->")
		printList(head.Next)
	}
}

//for algorithm
func deleteDuplicates(head *ListNode) *ListNode {
	res := head

	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}

	return res
}

//Recursive algorithm
// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}

// 	if head.Next != nil && head.Val == head.Next.Val {
// 		if head.Next.Next != nil {
// 			head.Next = head.Next.Next
// 			deleteDuplicates(head)
// 		} else {
// 			head.Next = nil
// 		}
// 	} else {
// 		deleteDuplicates(head.Next)
// 	}

// 	return head
// }

func main() {
	p := ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}}
	//p := ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}}}
	printList(&p)
	deleteDuplicates(&p)
	fmt.Printf("After")
	printList(&p)
}
