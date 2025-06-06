package main

import "fmt"

func main() {
	head := GetListNodeFromStartToFinish(
		1,
		func(cur, end int) bool { return cur <= end },
		func(i int) int { return i + 1 },
		1000)

	//head := GetListNodeFromStartToFinish(
	//	'a',
	//	func(cur, end rune) bool { return cur <= end },
	//	func(r rune) rune { return r + 1 },
	//	'z',
	//)

	expressLines := CreateExpressLines(head, 1)

	for i := len(expressLines) - 1; i >= 0; i-- {
		PrintSkipList(expressLines[i])
		fmt.Println()
	}

	printLinkedList(head)
}
