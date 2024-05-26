package btree

import (
	"fmt"
	"testing"
)

func slicesDifferent[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return true
	}
	for i, v := range a {
		if v != b[i] {
			return true
		}
	}
	return false
}

func TestOneItem(t *testing.T) {
	node := node{}

	node.addItem(Element{7, "foo"})

	if slicesDifferent(node.elements, []Element{{7, "foo"}}) {
		t.Errorf("Node.addItem(7); provides %v; want [7]", elementToKeys(node.elements))
	}
}

func TestTwoItemsRightOrder(t *testing.T) {
	node := node{}

	node.addItem(Element{2, "foo"})
	node.addItem(Element{3, "foo"})

	if slicesDifferent(node.elements, []Element{{2, "foo"}, {3, "foo"}}) {
		t.Errorf("Added 2 items correctly ordered. Got %v; wanted [2, 3]", elementToKeys(node.elements))
	}
}

func TestTwoItemsWrongOrder(t *testing.T) {
	node := node{}

	node.addItem(Element{6, "foo"})
	node.addItem(Element{4, "foo"})

	if slicesDifferent(node.elements, []Element{{4, "foo"}, {6, "foo"}}) {
		t.Errorf("Added 2 items incorrectly ordered. Got %v; wanted [4, 6]", elementToKeys(node.elements))
	}
}

func TestAddBetweenTwoItems(t *testing.T) {
	node := node{}
	node.addItem(Element{15, "foo"})
	node.addItem(Element{18, "foo"})

	node.addItem(Element{17, "foo"})

	if slicesDifferent(node.elements, []Element{{15, "foo"}, {17, "foo"}, {18, "foo"}}) {
		t.Errorf("Added item between 2 other items. Got %v; wanted [15, 17, 18]", elementToKeys(node.elements))
	}
}

func TestAddBeforeTwoItems(t *testing.T) {
	node := node{}
	node.addItem(Element{15, "foo"})
	node.addItem(Element{18, "foo"})

	node.addItem(Element{13, "foo"})

	if slicesDifferent(node.elements, []Element{{13, "foo"}, {15, "foo"}, {18, "foo"}}) {
		t.Errorf("Added item before 2 other items. Got %v; wanted [13, 15, 18]", elementToKeys(node.elements))
	}
}

func TestAddAfterTwoItems(t *testing.T) {
	node := node{}
	node.addItem(Element{15, "foo"})
	node.addItem(Element{18, "foo"})

	node.addItem(Element{22, "foo"})

	if slicesDifferent(node.elements, []Element{{15, "foo"}, {18, "foo"}, {22, "foo"}}) {
		t.Errorf("Added item after 2 other items. Got %v; wanted [15, 18, 22]", elementToKeys(node.elements))
	}
}

func checkNode(got []Element, want []Element, message string) *string {
	if slicesDifferent(got, want) {
		err := fmt.Sprintf("%s. Got `%v`; Wanted `%v`", message, got, want)
		return &err
	}
	return nil
}

func buildThreeLayerTree() *BTree {
	tree := NewBTree()
	tree.Insert(NewElement(1, "leaf1"))
	tree.Insert(NewElement(2, "leaf2"))

	tree.Insert(NewElement(3, "node3"))

	tree.Insert(NewElement(4, "leaf4"))
	tree.Insert(NewElement(5, "leaf5"))

	tree.Insert(NewElement(6, "node6"))

	tree.Insert(NewElement(7, "leaf7"))
	tree.Insert(NewElement(8, "leaf8"))

	tree.Insert(NewElement(9, "root9"))

	tree.Insert(NewElement(10, "leaf10"))
	tree.Insert(NewElement(11, "leaf11"))

	tree.Insert(NewElement(12, "node12"))

	tree.Insert(NewElement(13, "leaf13"))
	tree.Insert(NewElement(14, "leaf14"))

	tree.Insert(NewElement(15, "node15"))

	tree.Insert(NewElement(16, "leaf16"))
	tree.Insert(NewElement(17, "leaf17"))

	tree.Insert(NewElement(18, "root18"))

	tree.Insert(NewElement(19, "leaf19"))
	tree.Insert(NewElement(20, "leaf20"))

	tree.Insert(NewElement(21, "node21"))

	tree.Insert(NewElement(22, "leaf22"))
	tree.Insert(NewElement(23, "node23"))

	tree.Insert(NewElement(24, "node24"))

	tree.Insert(NewElement(25, "leaf25"))
	tree.Insert(NewElement(26, "node26"))

	return tree
}

func TestThreeLayerRoot(t *testing.T) {
	tree := buildThreeLayerTree()

	if err := checkNode(tree.rootNode.elements, []Element{{9, "root9"}, {18, "root18"}}, "searching 3-layer tree for roots"); err != nil {
		t.Errorf(*err)
	}
}

func TestThreeLayerLeftNode(t *testing.T) {
	tree := buildThreeLayerTree()

	if err := checkNode(tree.rootNode.children[0].elements, []Element{{3, "node3"}, {6, "node6"}}, "searching 3-layer tree for left 2nd layer"); err != nil {
		t.Errorf(*err)
	}
}

func TestThreeLayerCentreNode(t *testing.T) {
	tree := buildThreeLayerTree()

	if err := checkNode(tree.rootNode.children[1].elements, []Element{{12, "node12"}, {15, "node15"}}, "searching 3-layer tree for centre 2nd layer"); err != nil {
		t.Errorf(*err)
	}
}

func TestThreeLayerRightNode(t *testing.T) {
	tree := buildThreeLayerTree()

	if err := checkNode(tree.rootNode.children[2].elements, []Element{{21, "node21"}, {24, "node24"}}, "searching 3-layer tree for right 2nd layer"); err != nil {
		t.Errorf(*err)
	}
}

func TestThreeLayerLeftLeaf(t *testing.T) {
	tree := buildThreeLayerTree()

	if err := checkNode(tree.rootNode.children[1].children[0].elements, []Element{{10, "leaf10"}, {11, "leaf11"}}, "searching 3-layer tree for a left leaf"); err != nil {
		t.Errorf(*err)
	}
}

func TestThreeLayerCenterLeaf(t *testing.T) {
	tree := buildThreeLayerTree()

	if err := checkNode(tree.rootNode.children[2].children[1].elements, []Element{{22, "leaf22"}, {23, "leaf23"}}, "searching 3-layer tree for a center leaf"); err != nil {
		t.Errorf(*err)
	}
}

func TestThreeLayerRightLeaf(t *testing.T) {
	tree := buildThreeLayerTree()

	if err := checkNode(tree.rootNode.children[0].children[2].elements, []Element{{7, "leaf7"}, {8, "leaf8"}}, "searching 3-layer tree for a center leaf"); err != nil {
		t.Errorf(*err)
	}
}
