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
	return t.rootNode.search(key)
}

type node struct {
	elements []Element
	children []*node
	parent   *node

	tree *BTree
}

func (n node) search(key int) string {
	for i, element := range n.elements {
		if key == element.key {
			return element.value
		}
		if key < element.key {
			return n.children[i].search(key)
		}
	}
	if key > n.elements[len(n.elements)-1].key {
		return n.children[len(n.children)-1].search(key)
	}

	panic("seach algorithm error")
}

func (n *node) addItem(item Element) {
	if len(n.elements) == 0 {
		//only starting node can have 0 items so just add it
		n.elements = append(n.elements, item)
		return
	}

	//assuming leaf node
	for i, k := range n.elements {
		if k == item {
			panic("Haven't implemented behaviour for adding the same element twice into a tree")
		} else if k.key < item.key {
			continue
		} else if k.key > item.key {
			//shift everything after i along by 1 and replace index i with item
			n.elements = append(n.elements, Element{})
			copy(n.elements[i+1:], n.elements[i:])
			n.elements[i] = item

			goto postFor
		} else {
			panic("Impossible to reach here")
		}
	}
	//this item goes at the end of the list
	n.elements = append(n.elements, item)
postFor:
	if len(n.elements) <= 4 {
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
				tree:     n.tree,
				elements: n.elements[2:3],
			}
			otherChild := node{
				tree:     n.tree,
				elements: n.elements[:2],
			}
			n.elements = n.elements[3:]
			newRoot.children = append(newRoot.children, &otherChild, n)
			n.tree.rootNode = &newRoot
		}
	}
}
