package datastructures

import "fmt"

// single linked list
// double linked list

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func LinkedListMain() {
	myList := linkedList{}

	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 16}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)

	fmt.Println(myList)

}
