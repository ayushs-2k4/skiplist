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
