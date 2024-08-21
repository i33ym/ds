package linkedlist

import (
	"bytes"
	"fmt"
	"sync"
)

type LinkedList[T any] struct {
	head *node[T]
	size uint
	logs bytes.Buffer
	mu   sync.Mutex
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		size: 0,
	}
}

func (linkedList *LinkedList[T]) Display() {
	linkedList.mu.Lock()
	defer linkedList.mu.Unlock()

	if linkedList.head == nil {
		fmt.Println("Empty list")
		return
	}

	for head := linkedList.head; head != nil; head = head.next {
		if head.next == nil {
			fmt.Printf("%+v\n", head.item)
			break
		}

		fmt.Printf("%+v -> ", head.item)
	}
}

func (linkedList *LinkedList[T]) Logs() string {
	return linkedList.logs.String()
}

type node[T any] struct {
	item T
	next *node[T]
}

func (linkedList *LinkedList[T]) Prepend(item T) {
	linkedList.mu.Lock()
	defer linkedList.mu.Unlock()

	linkedList.head = &node[T]{
		item: item,
		next: linkedList.head,
	}

	linkedList.size++
	linkedList.logs.WriteString(fmt.Sprintf("A new item %+v is prepended. Current number of items in the list: %d\n", item, linkedList.size))
}

func (linkedList *LinkedList[T]) Append(item T) {
	linkedList.mu.Lock()
	defer linkedList.mu.Unlock()

	if linkedList.head == nil {
		linkedList.head = &node[T]{
			item: item,
			next: linkedList.head,
		}

		linkedList.size++
		linkedList.logs.WriteString(fmt.Sprintf("A new item %+v is appended. Current number of items in the list: %d\n", item, linkedList.size))

		return
	}

	temp := &node[T]{}
	for head := linkedList.head; head != nil; head = head.next {
		if head.next == nil {
			temp = head
			break
		}
		
		temp = head
	}

	temp.next = &node[T]{
		item: item,
		next: temp.next,
	}

	linkedList.size++
	linkedList.logs.WriteString(fmt.Sprintf("A new item %+v is appended. Current number of items in the list: %d\n", item, linkedList.size))
}
