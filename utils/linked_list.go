package utils

import "fmt"

type IntegerNode struct {
	element int
	next    *IntegerNode
}

type IntegerLinkedList struct {
	length int
	head   *IntegerNode
	tail   *IntegerNode
}

type IndexOutOfBoundError struct {
	index  int
	length int
}

func (e IndexOutOfBoundError) Error() string {
	return fmt.Sprintf("Index %d is already out of bounds (%d)", e.index, e.length)
}

func NewIntegerLinkedList() *IntegerLinkedList {
	list := new(IntegerLinkedList)
	return list
}

func NewIntegerLinkedListRange(from int, to int) *IntegerLinkedList {
	list := new(IntegerLinkedList)
	for i := from; i <= to; i++ {
		list.Add(i)
	}
	return list
}

func (l *IntegerLinkedList) Add(element int) int {
	node := new(IntegerNode)
	node.element = element
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = l.tail.next
	}
	l.length += 1
	return l.length - 1
}

func (l *IntegerLinkedList) Get(index int) (int, error) {
	var node *IntegerNode = l.head
	counter := 0
	for {
		if counter == index {
			break
		}
		node = node.next
		if node == nil {
			return 0, &IndexOutOfBoundError{counter, l.length}
		}
		counter += 1
	}
	return node.element, nil
}

func (l *IntegerLinkedList) RemoveFirst(element int) bool {
	var node *IntegerNode = l.head
	if node.element == element {
		l.head = l.head.next
		l.length -= 1
		return true
	}
	var previous *IntegerNode = node
	node = node.next
	for {
		if node == nil {
			break
		}
		if node.element == element {
			previous.next = node.next
			l.length -= 1
			return true
		}
		previous = node
		node = node.next
	}
	return false
}

func (l *IntegerLinkedList) Length() int {
	return l.length
}

func (l *IntegerLinkedList) Clone() *IntegerLinkedList {
	list := NewIntegerLinkedList()
	length := l.Length()
	for i := 0; i < length; i++ {
		value, _ := l.Get(i)
		list.Add(value)
	}
	return list
}
