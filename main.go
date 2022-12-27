package main

import (
	"bytes"
	"fmt"
)

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

func (l linkedList) toString() string {
	buf := bytes.NewBufferString("")
	for ;l.length != 0; l.head = l.head.next {
		fmt.Fprintf(buf, "%d ", l.head.data)
		l.length-- 
	}
	return buf.String()
}

func (l *linkedList) deleteWithValue(value int) {
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	for h, n := l.head, l.length; n != 0; n, h = n-1, h.next {
		if h.next == nil {
			return
		}
		if h.next.data == value{
			h.next = h.next.next
			l.length--
			return
		}
	}
}

func main() {
	ll := linkedList{}
	n1 := &node{data: 18}
	n2 := &node{data: 32}
	n3 := &node{data: 999}
	n4 := &node{data: 542}
	n5 := &node{data: 435}

	ll.prepend(n1)
	ll.prepend(n2)
	ll.prepend(n3)
	ll.prepend(n4)
	ll.prepend(n5)

	fmt.Println(ll.toString())
	ll.deleteWithValue(8923)
	fmt.Println(ll.toString())
	ll.toString()
}