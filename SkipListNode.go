package main

import (
	"fmt"
)

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

func CreateExpressLineFromAnotherExpressLineWithConnection[T any](head *SkipListNode[T], skipAmount int) (expressLine *SkipListNode[T], count int) {
	skipAmount++

	totalSize := 1

	currentNode := head
	expressLineHead := NewSkipListNode(head.val)
	expressLineHead.below = currentNode
	expressLineTail := expressLineHead

	for currentNode != nil {
		for i := 0; i < skipAmount && currentNode != nil; i++ {
			currentNode = currentNode.next
		}

		if currentNode != nil {
			thisSkipListNode := NewSkipListNodeWithPrev(currentNode.val, expressLineTail)
			thisSkipListNode.below = currentNode
			expressLineTail = thisSkipListNode

			totalSize++
		}
	}

	return expressLineHead, totalSize
}

func CreateExpressLines[T any](head *ListNode[T], skipAmount int) []*SkipListNode[T] {
	if skipAmount < 1 {
		return nil
	}

	var ans []*SkipListNode[T]

	firstExpressLine := CreateExpressLine(head, skipAmount)
	ans = append(ans, firstExpressLine)
	lastCount := 100

	for lastCount > 2 {
		lastExpressLine := ans[len(ans)-1]
		newExpressLine, count := CreateExpressLineFromAnotherExpressLineWithConnection(lastExpressLine, skipAmount)

		ans = append(ans, newExpressLine)

		lastCount = count
	}

	return ans
}

func isValueInBetweenTwoValues[T comparable](a, b, val T, isALessThanB func(a, b T) bool) bool {
	return (a == val || isALessThanB(a, val)) && (isALessThanB(val, b) || b == val)
}

func Search[T comparable](head *SkipListNode[T], val T, isALessThanB func(a, b T) bool) *ListNode[T] {
	if head == nil {
		return nil
	}

	if isALessThanB(val, head.val.val) { //val < head.val.val
		return nil
	}

	if head.below == nil {
		curr := head

		for curr != nil && curr.next != nil {
			if isValueInBetweenTwoValues(curr.val.val, curr.next.val.val, val, isALessThanB) {
				return SearchInListNode(curr.val, curr.next.val, val)
			}

			curr = curr.next
		}

		if curr != nil {
			return SearchInListNode(curr.val, nil, val)
		}

		return nil
	}

	if head.next == nil {
		return Search(head.below, val, isALessThanB)
	}

	curr := head

	for curr != nil && curr.next != nil {
		if isValueInBetweenTwoValues(curr.val.val, curr.next.val.val, val, isALessThanB) {
			return Search(curr.below, val, isALessThanB)
		}

		curr = curr.next
	}

	if curr != nil {
		return Search(curr.below, val, isALessThanB)
	}

	return nil
}

func PrintSkipList[T any](head *SkipListNode[T]) {
	ptr := head

	for ptr != nil {
		fmt.Print(ptr.val.val)
		fmt.Print(" ")

		ptr = ptr.next
	}
}
