package stack

import (
	"bytes"
	"fmt"
	"sync"
)

type Stack[T any] struct {
	items []T
	size  uint
	logs  bytes.Buffer
	mu    sync.Mutex
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: nil,
		size:  0,
		logs:  bytes.Buffer{},
	}
}

func (stack *Stack[T]) Push(item T) {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	stack.items = append(stack.items, item)
	stack.size++

	stack.logs.WriteString(fmt.Sprintf("A new item %+v is pushed to the stack. Current number of items in the stack: %d\n", item, stack.size))
}

func (stack *Stack[T]) Pop() (T, bool) {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	var zero T

	if stack.size == 0 {
		return zero, false
	}

	item := stack.items[0]
	stack.items = stack.items[1:]
	stack.size--

	stack.logs.WriteString(fmt.Sprintf("An item %+v is popped from the stack. Current number of items in the stack: %d\n", item, stack.size))
	return item, true
}

func (stack *Stack[T]) Logs() string {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	return stack.logs.String()
}

func (stack *Stack[T]) Display() {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	if stack.size == 0 {
		fmt.Println("Empty stack")
		return
	}

	for index := 0; index < int(stack.size)-1; index++ {
		fmt.Printf("%+v -> ", stack.items[index])
	}

	fmt.Printf("%+v\n", stack.items[stack.size-1])
}
