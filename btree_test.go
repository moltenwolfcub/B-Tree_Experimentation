package btree

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	tree := NewBTree()
	tree.rootNode.addItem(5)
	fmt.Println(tree.rootNode.keys)
	tree.rootNode.addItem(7)
	tree.rootNode.addItem(2)
	fmt.Println(tree.rootNode.keys)
	tree.rootNode.addItem(3)
	fmt.Println(tree.rootNode.keys)
	tree.rootNode.addItem(6)
	fmt.Println(tree.rootNode.keys)
}
