package linkedlist

import (
	"fmt"
)

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int) {
	node := &Node{}
	node.property = property
	node.nextNode = nil

	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
	}

	linkedList.headNode = node

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

func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	node := &Node{}
	node.property = property
	node.nextNode = nil

	nodeWith := &Node{}

	nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
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

func (linkedList *LinkedList) AddToEnd(property int) {
	node := &Node{}
	node.property = property
	node.nextNode = nil

	lastNode := &Node{}

	lastNode = linkedList.LastNode()

	if lastNode != nil {
		lastNode.nextNode = node
	} else {
		linkedList.headNode = node
	}

}

func (linkedList *LinkedList) IterateList() {

	node := &Node{}
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)

	}
}

func LinkedListExec() {

	linkedList := &LinkedList{}

	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)

	linkedList.IterateList()
}
