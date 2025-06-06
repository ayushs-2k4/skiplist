package main

import "fmt"

type ListNode[T any] struct {
	next *ListNode[T]
	val  T
}

func NewListNode[T any](val T) *ListNode[T] {
	return &ListNode[T]{
		val:  val,
		next: nil,
	}
}

func NewListNodeWithPrev[T any](val T, prev *ListNode[T]) *ListNode[T] {
	newNode := NewListNode(val)

	prev.next = newNode

	return newNode
}

func printLinkedList[T any](head *ListNode[T]) {
	ptr := head

	for ptr != nil {
		fmt.Print(ptr.val)
		fmt.Print(" ")

		ptr = ptr.next
	}
}

func SearchInListNode[T comparable](startNode *ListNode[T], endNode *ListNode[T], val T) *ListNode[T] {
	currentNode := startNode

	for currentNode != nil && currentNode != endNode {
		if currentNode.val == val {
			return currentNode
		}

		currentNode = currentNode.next
	}

	if endNode == nil || currentNode == nil {
		return nil
	}

	if endNode.val == val {
		return endNode
	}

	return nil
}

func GetListNodeFromArray[T any](arr []T) *ListNode[T] {
	if len(arr) == 0 {
		return nil
	}

	head := NewListNode(arr[0])
	tail := head

	for i := 1; i < len(arr); i++ {
		newNode := NewListNode(arr[i])
		tail.next = newNode
		tail = tail.next
	}

	return head
}

func GetListNodeFromStartToFinish[T any](
	start T,
	isEnd func(current, finish T) bool,
	step func(T) T,
	finish T,
) *ListNode[T] {
	head := NewListNode(start)
	tail := head
	for val := step(start); isEnd(val, finish); val = step(val) {
		tail = NewListNodeWithPrev(val, tail)
	}
	return head
}
