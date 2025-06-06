package main

import "fmt"

type SkipListNode[T any] struct {
	next  *SkipListNode[T]
	below *SkipListNode[T] // address of SkipListNode of the below lane which holds the same value
	val   *ListNode[T]
}

func NewSkipListNode[T any](val *ListNode[T]) *SkipListNode[T] {
	return &SkipListNode[T]{
		next: nil,
		val:  val,
	}
}

func NewSkipListNodeWithPrev[T any](val *ListNode[T], prev *SkipListNode[T]) *SkipListNode[T] {
	newNode := NewSkipListNode(val)

	prev.next = newNode

	return newNode
}

func CreateExpressLine[T any](head *ListNode[T], skipAmount int) *SkipListNode[T] {
	skipAmount++

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

func CreateExpressLineFromAnotherExpressLineWithConnection[T any](head *SkipListNode[T], skipAmount int) *SkipListNode[T] {
	skipAmount++

	currentNode := head
	expressLineHead := NewSkipListNode(head.val)
	currentNode.below = expressLineHead
	expressLineTail := expressLineHead

	for currentNode != nil {
		for i := 0; i < skipAmount && currentNode != nil; i++ {
			currentNode = currentNode.next
		}

		if currentNode != nil {
			thisSkipListNode := NewSkipListNodeWithPrev(currentNode.val, expressLineTail)
			currentNode.below = thisSkipListNode
			expressLineTail = thisSkipListNode
		}
	}

	return expressLineHead
}

func PrintSkipList[T any](head *SkipListNode[T]) {
	ptr := head

	for ptr != nil {
		fmt.Print(ptr.val.val)
		fmt.Print(" ")

		ptr = ptr.next
	}
}
