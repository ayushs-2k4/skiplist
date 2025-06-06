package main

import "fmt"

func main() {
	head := GetListNodeFromStartToFinish(
		1,
		func(cur, end int) bool { return cur <= end },
		func(i int) int { return i + 1 },
		100)

	//head := GetListNodeFromStartToFinish(
	//	'a',
	//	func(cur, end rune) bool { return cur <= end },
	//	func(r rune) rune { return r + 1 },
	//	'z',
	//)

	skipListHead := CreateExpressLine(head, 3)
	secondExpressLine := CreateExpressLineFromAnotherExpressLineWithConnection(skipListHead, 2)

	printLinkedList(head)
	fmt.Println()
	PrintSkipList(skipListHead)
	fmt.Println()
	PrintSkipList(secondExpressLine)
	fmt.Println()
}
