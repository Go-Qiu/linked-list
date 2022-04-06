package main

import (
	"fmt"
)

type node struct {
	item string
	next *node
}

/*
	declare the linked list variable
	(and thus Go will allocate a memory address to it)
*/
type linkedList struct {
	head *node
	size int
}

/*
	function add a node to the linked list
*/
func (p *linkedList) addNode(name string) error {
	newNode := &node{item: name, next: nil}

	if p.head == nil {
		// empty
		p.head = newNode
	} else {
		// not empty.
		currentNode := p.head

		// travesal across the nodes in the linked list until the last node
		// that has its next value is nil.

		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	p.size++
	return nil
}

/*
	function to print all the nodes in the linked list
*/
func (p *linkedList) printAllNodes() error {
	if p.head == nil {
		fmt.Println("There are no nodes in the list.")
		return nil
	}
	currentNode := p.head
	fmt.Printf("%v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%v\n", currentNode.item)
	}
	return nil
}

/*
	function to remove a node from the linked list
*/
func (p *linkedList) removeNode(name string) error {
	if p.head == nil {
		// empty. cannot remove any node.
		fmt.Println("There are no nodes in the list.")
		return nil
	}

	// linked list is not empty.
	// ok to proceed further.
	currentNode := p.head
	previousNode := p.head
	nextNode := currentNode.next
	hob := 0

	// loop until item matches name
	for currentNode.item != name {

		if currentNode.next != nil {
			// node in the middle of the list
			previousNode = currentNode
			currentNode = currentNode.next
			nextNode = currentNode.next
		} else {
			// node at the end of the list
			nextNode = nil
		}
		fmt.Printf("hob count -> %d\n", hob)

	}

	// 'remove' the link to the item of the node that has a match.
	if hob == 0 {
		p.head = p.head.next
	} else {
		previousNode.next = nextNode
		p.size--
		fmt.Printf("\"%s\" has been removed from the linked list.\n", name)
	}

	return nil
}

// func (p *linkedList) removeNodeByIndex(index int) (string, error) {

// 	if p.head == nil {
// 		return "", errors.New("Empty linked list.")
// 	}

// 	if index < 1 || index > p.size {
// 		return "", errors.New("Invalid index position.")
// 	}

// 	var item string
// 	if index == 1 {
// 		// remove at 1st node
// 		item = p.head.item
// 		p.head = p.head.next
// 	} else {

// 		currentNode := p.head
// 		prevNode := p.head // currentNode
// 		for i := 1; i <= index-1; i++ {
// 			prevNode = currentNode
// 			currentNode = currentNode.next
// 		}
// 	}

// 	if index == p.size {

// 	}
// }

func main() {

	// init the linked list variable
	myList := &linkedList{nil, 0}

	// add the nodes
	fmt.Print("Adding nodes to linked list ...\n")
	myList.addNode("Mary")
	myList.addNode("Tom")
	myList.addNode("Dan")

	// print out all the nodes in the linked list
	fmt.Println("Printing all nodes of the list ...")
	myList.printAllNodes()
	fmt.Println("")

	fmt.Println("Removing a node ...")
	// remove a specific node in the linked list
	myList.removeNode("Mary")
	myList.printAllNodes()
}
