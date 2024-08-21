package main

import (
	"fmt"

	linkedlist "github.com/i33ym/ds/linked-list"
)

func main() {
	items := linkedlist.NewLinkedList[string]()

	items.Prepend("Abraham")
	items.Append("Moses")
	items.Prepend("Noah")
	items.Append("Jesus")

	fmt.Println(items.Logs())
	fmt.Println()

	items.Display()
}
