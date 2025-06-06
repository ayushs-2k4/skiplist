package main

import "fmt"

type SkipListNode struct {
	next *SkipListNode
	val  *ListNode
}

func NewSkipListNode(val *ListNode) *SkipListNode {
	return &SkipListNode{
		next: nil,
		val:  val,
	}
}

func NewSkipListNodeWithPrev(val *ListNode, prev *SkipListNode) *SkipListNode {
	newNode := NewSkipListNode(val)

	prev.next = newNode

	return newNode
}

func CreateExpressLine(head *ListNode, skipAmount int) *SkipListNode {
	currentNode := head
	expressLineHead := NewSkipListNode(head)
	expressLineTail := expressLineHead

	for currentNode != nil {
		for i := 0; i < skipAmount && currentNode != nil; i++ {
			currentNode = currentNode.next
		}

		if currentNode != nil {
			thisSkipListNode := NewSkipListNodeWithPrev(currentNode, expressLineTail)
			expressLineTail = thisSkipListNode
		}
	}

	return expressLineHead
}

func PrintSkipList(head *SkipListNode) {
	ptr := head

	for ptr != nil {
		fmt.Print(ptr.val.val)
		fmt.Print(" ")

		ptr = ptr.next
	}
}
