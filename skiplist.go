package main

import "fmt"

func printIsValuePresent[T int](head *SkipListNode[T], val T) {
	node := Search(head, val, func(a, b T) bool {
		return a < b
	})

	if node != nil {
		fmt.Println("Yes, present")
	} else {
		fmt.Println("No, absent")
	}
}

func main() {
	//head := GetListNodeFromStartToFinish(1, func(cur, end int) bool { return cur <= end }, func(i int) int { return i + 1 }, 10)
	head := GetListNodeFromArray([]int{1, 2, 3, 4, 5})

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
	fmt.Println()

	printIsValuePresent(expressLines[len(expressLines)-1], 4)
}
