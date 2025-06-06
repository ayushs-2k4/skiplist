package main

import "fmt"

func main() {
	head := GetListNodeFromStartToFinish(1, 100)
	skipListHead := CreateExpressLine(head, 5)

	printLinkedList(head)
	fmt.Println()
	PrintSkipList(skipListHead)
	fmt.Println()
}
