package main

import (
	"errors"
	"fmt"
)

/*
In a lot of problems, you may be asked to reverse the links between a set of nodes of a linked list.
Often, the constraint is that you need to do this in-place, i.e.,
using the existing node objects and without using extra memory.
This is where this pattern is useful.

How do I identify when to use this pattern:
If you’re asked to reverse a linked list without using extra memory
Problems featuring in-place reversal of linked list pattern:

1-Reverse a Sub-list (medium)
2-Reverse every K-element Sub-list (medium)
*/

// ListNode Definition for singly-linked list.
// Node represents a node of linked list

type Node struct {
	Val  int
	next *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	head *Node
	len  int
}

func main() {
	linkedList := LinkedList{}
	/*linkedList.Insert(1)
	linkedList.Insert(2)
	linkedList.Insert(4)*/
	linkedList.Insert(3)
	linkedList.Insert(5)
	//ReverseLinkedList(linkedList.head)
	//linkedList.head = reverseList(linkedList.head)
	linkedList.head = reverseBetween(linkedList.head, 1, 2)
	linkedList.Print()

}

// Insert inserts new node at the end of  from linked list
func (l *LinkedList) Insert(val int) {
	n := Node{}
	n.Val = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

// Print displays all the nodes from linked list
func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("Node: ", ptr.Val)
		ptr = ptr.next
	}
}

// Search returns node position with given value from linked list
func (l *LinkedList) Search(val int) int {
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.Val == val {
			return i
		}
		ptr = ptr.next
	}
	return -1
}

// InsertAt inserts new node at given position
func (l *LinkedList) InsertAt(pos int, value int) {
	// create a new node
	newNode := Node{}
	newNode.Val = value
	// validate the position
	if pos < 0 {
		return
	}
	if pos == 0 {
		l.head = &newNode
		l.len++
		return
	}
	if pos > l.len {
		return
	}
	n := l.GetAt(pos)
	newNode.next = n
	prevNode := l.GetAt(pos - 1)
	prevNode.next = &newNode
	l.len++
}

// GetAt returns node at given position from linked list
func (l *LinkedList) GetAt(pos int) *Node {
	ptr := l.head
	if pos < 0 {
		return ptr
	}
	if pos > (l.len - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}

// DeleteAt deletes node at given position from linked list
func (l *LinkedList) DeleteAt(pos int) error {
	// validate the position
	if pos < 0 {
		fmt.Println("position can not be negative")
		return errors.New("position can not be negative")
	}
	if l.len == 0 {
		fmt.Println("No nodes in list")
		return errors.New("no nodes in list")
	}
	prevNode := l.GetAt(pos - 1)
	if prevNode == nil {
		fmt.Println("Node not found")
		return errors.New("node not found")
	}
	prevNode.next = l.GetAt(pos).next
	l.len--
	return nil
}

// DeleteVal deletes node having given value from linked list
func (l *LinkedList) DeleteVal(val int) error {
	ptr := l.head
	if l.len == 0 {
		fmt.Println("List is empty")
		return errors.New("list is empty")
	}
	for i := 0; i < l.len; i++ {
		if ptr.Val == val {
			if i > 0 {
				prevNode := l.GetAt(i - 1)
				prevNode.next = l.GetAt(i).next
			} else {
				l.head = ptr.next
			}
			l.len--
			return nil
		}
		ptr = ptr.next
	}
	fmt.Println("Node not found")
	return errors.New("node not found")
}

/*
in this method we just travers the array
*/

func ReverseLinkedList(head *Node) {
	if head == nil {
		return
	}
	ReverseLinkedList(head.next)
	//here is the returned value
	fmt.Printf("%d ", head.Val)
}

/*
reverseLinkedList
use two pointers
*/
func reverseList(head *Node) *Node {
	if head.next == nil {
		return head
	}
	var prev *Node = nil
	var next *Node = nil
	for head != nil {
		next = head.next
		head.next = prev
		prev = head
		head = next
	}
	head = prev
	return head
}

/*
Given the head of a singly linked list and two integers left and right
where left <= right, reverse the nodes of the list from position left to position right, and return the reversed list.

*/
func reverseBetween(head *Node, left int, right int) *Node {
	if head == nil {
		return nil
	}
	if head.next == nil {
		return head
	}
	var leftNode *Node = head
	counter := 1
	for counter < left {
		leftNode = leftNode.next
		counter++
	}
	for left != right && right > left {
		counter := left
		var rightNode *Node = leftNode
		for counter != right {
			rightNode = rightNode.next
			counter++
		}
		tempVal := leftNode.Val
		leftNode.Val = rightNode.Val
		rightNode.Val = tempVal
		leftNode = leftNode.next
		left++
		right--
	}
	return head
}
