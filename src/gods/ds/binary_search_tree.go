package ds

import (
	"fmt"
	"sync"
)

type BTreeNode struct {
	key       int
	value     string
	leftNode  *BTreeNode
	rightNode *BTreeNode
}

func (node *BTreeNode) GetKey() int {
	return node.key
}

func (node *BTreeNode) GetValue() string {
	return node.value
}

func NewBTreeNode(key int, value string, left, right *BTreeNode) *BTreeNode {
	return &BTreeNode{key, value, left, right}
}

type BTree struct {
	rootNode *BTreeNode
	lock     sync.RWMutex
}

func (bt *BTree) Root() *BTreeNode {
	return bt.rootNode
}

func (bt *BTree) isRoot(key int) bool {
	return bt.Root().GetKey() == key
}

func (bt *BTree) InsertElement(key int, value string) {
	bt.lock.Lock()
	defer bt.lock.Unlock()

	treeNode := NewBTreeNode(key, value, nil, nil)

	if bt.Root() == nil {
		bt.rootNode = treeNode
	} else {
		bt.insertTreeNode(bt.rootNode, treeNode)
	}
}

func (bt *BTree) insertTreeNode(rootNode, newNode *BTreeNode) {
	if rootNode.GetKey() > newNode.GetKey() {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newNode
		} else {
			bt.insertTreeNode(rootNode.leftNode, newNode)
		}
	} else {
		if rootNode.rightNode == nil {
			rootNode.rightNode = newNode
		} else {
			bt.insertTreeNode(rootNode.rightNode, newNode)
		}
	}
}

func (bt *BTree) InOrderTraversal() {
	if bt.Root() != nil {
		bt.inOrderTraversal(bt.Root())
	}
}

func (bt *BTree) inOrderTraversal(node *BTreeNode) {
	if node == nil {
		return
	}
	bt.inOrderTraversal(node.leftNode)
	fmt.Println(node.GetValue())
	bt.inOrderTraversal(node.rightNode)
}

func NewBinaryTree() *BTree {
	return &BTree{}
}

func ExecuteBinaryTree1() {
	btree := NewBinaryTree()
	btree.InsertElement(10, "10")
	btree.InsertElement(5, "5")
	btree.InsertElement(20, "20")

	btree.InOrderTraversal()
}
