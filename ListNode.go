package main

import "fmt"

type ListNode struct {
	next *ListNode
	val  int
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		val:  val,
		next: nil,
	}
}

func NewListNodeWithPrev(val int, prev *ListNode) *ListNode {
	newNode := NewListNode(val)

	prev.next = newNode

	return newNode
}

func printLinkedList(head *ListNode) {
	ptr := head

	for ptr != nil {
		fmt.Print(ptr.val)
		fmt.Print(" ")

		ptr = ptr.next
	}
}

func GetListNodeFromStartToFinish(start, finish int) *ListNode {
	head := NewListNode(start)
	tail := head

	for i := start + 1; i <= finish; i++ {
		newNode := NewListNodeWithPrev(i, tail)
		tail = newNode
	}

	return head
}
