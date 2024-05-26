package btree

type Element struct {
	key   int
	value string
}

func NewElement(k int, v string) Element {
	return Element{k, v}
}

func elementToKeys(kvs []Element) []int {
	keys := make([]int, 0)
	for _, v := range kvs {
		keys = append(keys, v.key)
	}
	return keys
}

type BTree struct {
	rootNode *node
}

func NewBTree() *BTree {
	t := &BTree{}
	t.rootNode = &node{
		tree: t,
	}
	return t
}

func (t *BTree) Insert(e Element) {
	t.rootNode.addItem(e)
}

func (t BTree) Search(key int) string {
	return "NOT IMPLEMENTED"
}

type node struct {
	keys     []Element
	children []*node
	parent   *node

	tree *BTree
}

// func (n Node) search() {

// }

func (n *node) addItem(item Element) {
	if len(n.keys) == 0 {
		//only starting node can have 0 items so just add it
		n.keys = append(n.keys, item)
		return
	}

	//assuming leaf node
	for i, k := range n.keys {
		if k == item {
			panic("Haven't implemented behaviour for adding the same element twice into a tree")
		} else if k.key < item.key {
			continue
		} else if k.key > item.key {
			//shift everything after i along by 1 and replace index i with item
			n.keys = append(n.keys, Element{})
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
			newRoot := node{
				tree: n.tree,
				keys: n.keys[2:3],
			}
			otherChild := node{
				tree: n.tree,
				keys: n.keys[3:],
			}
			n.keys = n.keys[:2]
			newRoot.children = append(newRoot.children, &otherChild, n)
			n.tree.rootNode = &newRoot
		}
	}
}
