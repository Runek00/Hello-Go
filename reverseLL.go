package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) toStr() string {
	var builder strings.Builder
	for node != nil {
		builder.WriteString(fmt.Sprint((*node).Val))
		if (*node).Next == nil {
			break
		}
		node = node.Next
	}
	return builder.String()

}

func example(n int) *ListNode {
	head := new(ListNode)
	head.Val = -2
	temp := new(ListNode)
	temp.Val = -1
	head.Next = temp
	for i := 0; i < n; i++ {
		temp.Next = new(ListNode)
		temp.Next.Val = i
		temp = temp.Next
	}
	return head
}

func ReverseLinkedList(head *ListNode) *ListNode {
	first := head
	last := head
	for last.Next != nil {
		temp := last.Next
		last.Next = temp.Next
		temp.Next = first
		first = temp
	}
	return first
}
