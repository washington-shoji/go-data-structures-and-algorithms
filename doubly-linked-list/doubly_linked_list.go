package doublylinkedlist

import (
	"fmt"
)

type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) NodeBetweenValues(firstProperty int, secondProperty int) *Node {
	node := &Node{}
	nodeWith := &Node{}

	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty {
				nodeWith = node
				break
			}
		}
	}

	return nodeWith
}

func (linkedlist *LinkedList) AddToHead(property int) {
	node := &Node{}
	node.property = property
	node.nextNode = nil

	if linkedlist.headNode != nil {
		node.nextNode = linkedlist.headNode
		linkedlist.headNode.previousNode = node
	}
	linkedlist.headNode = node
}

func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	node := &Node{}
	nodeWith := &Node{}
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

func (linkedlist *LinkedList) AddAfter(nodeProperty int, property int) {
	node := &Node{}
	node.property = property
	node.nextNode = nil

	nodeWith := &Node{}
	nodeWith = linkedlist.NodeWithValue(nodeProperty)

	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}
}

func (linkedList *LinkedList) LastNode() *Node {
	node := &Node{}
	lastNode := &Node{}
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (linkedlist *LinkedList) AddToEnd(property int) {
	node := &Node{}
	node.property = property
	node.nextNode = nil

	lastNode := &Node{}
	lastNode = linkedlist.LastNode()

	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

func DoublyLinkedListExec() {
	linkedlist := &LinkedList{}
	linkedlist.AddToHead(1)
	linkedlist.AddToHead(3)
	linkedlist.AddToEnd(5)
	linkedlist.AddAfter(1, 7)

	fmt.Println(linkedlist.headNode.property)

	node := &Node{}
	node = linkedlist.NodeBetweenValues(1, 5)

	fmt.Println(node.property)
}
