package btree

type NodeInt int

type BTree struct {
	rootNode *Node
}

func NewBTree() *BTree {
	t := &BTree{}
	t.rootNode = &Node{
		tree: t,
	}
	return t
}

type Node struct {
	keys     []NodeInt
	children []*Node
	parent   *Node

	tree *BTree
}

func (n *Node) addItem(item NodeInt) {
	if len(n.keys) == 0 {
		//only starting node can have 0 items so just add it
		n.keys = append(n.keys, item)
		return
	}

	//assuming leaf node
	for i, k := range n.keys {
		if k == item {
			panic("Haven't implemented behaviour for adding the same element twice into a tree")
		} else if k < item {
			continue
		} else if k > item {
			//shift everything after i along by 1 and replace index i with item
			n.keys = append(n.keys, 0)
			copy(n.keys[i+1:], n.keys[i:])
			n.keys[i] = item

			goto postFor
		} else {
			panic("Impossible to reach here")
		}
	}
	//this item goes at the end of the list
	n.keys = append(n.keys, item)
postFor:
	if len(n.keys) <= 4 {
		return
	} else {
		//rebalance and split tree
		if n.parent != nil {
			panic("Not Implemented")
			// n.parent.addItem(n.keys[2])
			// n.keys = n.keys[:2]
			// n.parent.children = append(n.parent.children, Node{
			// 	parent: n.parent,
			// 	keys:   n.keys[3:],
			// })
		} else {
			newRoot := Node{
				tree: n.tree,
				keys: n.keys[2:3],
			}
			otherChild := Node{
				tree: n.tree,
				keys: n.keys[3:],
			}
			n.keys = n.keys[:2]
			newRoot.children = append(newRoot.children, &otherChild, n)
			n.tree.rootNode = &newRoot
		}
	}
}
