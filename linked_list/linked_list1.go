package linked_list

import "fmt"

// node element to keep element and next node together
type node struct {
	el   interface{}
	next *node
}

// LinkedList of nodes
type LinkedList struct {
	head   *node
	length int
}

func NewLinkedList(value int) *LinkedList {
	n := node{el: value}
	return &LinkedList{
		head:   &n,
		length: 1,
	}
}

func (l LinkedList) List() {
	cur := l.head
	for cur.next != nil {
		fmt.Println(cur.el)
		cur = cur.next
	}
	fmt.Println(cur.el)
}

func (l LinkedList) Get(index int) interface{} {
	cur := l.head

	if index > l.length {
		return nil
	}

	for pos := 0; pos < index; pos++ {
		if cur == nil {
			return nil
		}
		cur = cur.next
	}

	return cur.el
}

// Search searches the element e and returns its index
func (l LinkedList) Search(e interface{}) interface{} {
	pos := 0
	cur := l.head
	for cur.el != e {
		cur = cur.next
		pos++
	}
	return pos
}

// Add inserts the element e to the list at position i
func (l *LinkedList) Add(index int, e interface{}) {
	if index > l.length - 1 {
		panic("Can't add to linked list - index is bigger than length of linked list")
	}
	n := node{el: e}
	if index == 0 {
		n.next = l.head
		l.head = &n
		return
	}
	prev := l.head
	cur := l.head.next
	pos := 1
	for pos < index {
		prev = prev.next
		cur = cur.next
		pos++
	}
	prev.next = &n
	n.next = cur
	l.length += 1
}

// Delete removes the element at position i
func (l *LinkedList) Delete(index int) {
	if index > l.length - 1  {
		panic("Not valid index")
	}
	if index == 0 { // deleting the head
		l.head = l.head.next
		return
	}
	prev := l.head
	cur := l.head.next
	pos := 1
	for pos < index {
		prev = prev.next
		cur = cur.next
		pos++
	}
	prev.next = cur.next
	l.length -= 1
}
