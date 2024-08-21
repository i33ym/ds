package main

import (
	"fmt"

	linkedlist "github.com/i33ym/ds/linked-list"
	"github.com/i33ym/ds/stack"
)

func main() {
	linkedList := linkedlist.NewLinkedList[string]()

	linkedList.Prepend("Abraham")
	linkedList.Append("Moses")
	linkedList.Prepend("Noah")
	linkedList.Append("Jesus")

	fmt.Println(linkedList.Logs())
	linkedList.Display()
	fmt.Println()

	stack := stack.NewStack[int]()

	stack.Push(14)
	stack.Push(23)
	stack.Push(40)

	fmt.Println(stack.Logs())
	stack.Display()

	fmt.Println()

	number, ok := stack.Pop()
	if !ok {
		panic("something terrible happaned (1)")
	}

	fmt.Println("The first to be popped:", number)

	number, ok = stack.Pop()
	if !ok {
		panic("something terrible happened (2)")
	}

	fmt.Println("The second to be popped:", number)

	number, ok = stack.Pop()
	if !ok {
		panic("something terrible happened (3)")
	}

	fmt.Println("The third to be popped:", number)

	_, ok = stack.Pop()
	if ok {
		panic("something that should not happen happened")
	}

	fmt.Println()

	fmt.Println(stack.Logs())
	stack.Display()
}
