package btree

import (
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

	if slicesDifferent(node.keys, []Element{{7, "foo"}}) {
		t.Errorf("Node.addItem(7); provides %v; want [7]", elementToKeys(node.keys))
	}
}

func TestTwoItemsRightOrder(t *testing.T) {
	node := node{}

	node.addItem(Element{2, "foo"})
	node.addItem(Element{3, "foo"})

	if slicesDifferent(node.keys, []Element{{2, "foo"}, {3, "foo"}}) {
		t.Errorf("Added 2 items correctly ordered. Got %v; wanted [2, 3]", elementToKeys(node.keys))
	}
}

func TestTwoItemsWrongOrder(t *testing.T) {
	node := node{}

	node.addItem(Element{6, "foo"})
	node.addItem(Element{4, "foo"})

	if slicesDifferent(node.keys, []Element{{4, "foo"}, {6, "foo"}}) {
		t.Errorf("Added 2 items incorrectly ordered. Got %v; wanted [4, 6]", elementToKeys(node.keys))
	}
}

func TestAddBetweenTwoItems(t *testing.T) {
	node := node{}
	node.addItem(Element{15, "foo"})
	node.addItem(Element{18, "foo"})

	node.addItem(Element{17, "foo"})

	if slicesDifferent(node.keys, []Element{{15, "foo"}, {17, "foo"}, {18, "foo"}}) {
		t.Errorf("Added item between 2 other items. Got %v; wanted [15, 17, 18]", elementToKeys(node.keys))
	}
}

func TestAddBeforeTwoItems(t *testing.T) {
	node := node{}
	node.addItem(Element{15, "foo"})
	node.addItem(Element{18, "foo"})

	node.addItem(Element{13, "foo"})

	if slicesDifferent(node.keys, []Element{{13, "foo"}, {15, "foo"}, {18, "foo"}}) {
		t.Errorf("Added item before 2 other items. Got %v; wanted [13, 15, 18]", elementToKeys(node.keys))
	}
}

func TestAddAfterTwoItems(t *testing.T) {
	node := node{}
	node.addItem(Element{15, "foo"})
	node.addItem(Element{18, "foo"})

	node.addItem(Element{22, "foo"})

	if slicesDifferent(node.keys, []Element{{15, "foo"}, {18, "foo"}, {22, "foo"}}) {
		t.Errorf("Added item after 2 other items. Got %v; wanted [15, 18, 22]", elementToKeys(node.keys))
	}
}
