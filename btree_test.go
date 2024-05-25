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
	node := Node{}

	node.addItem(7)

	if slicesDifferent(node.keys, []NodeInt{7}) {
		t.Errorf("Node.addItem(7); provides %v; want [7]", node.keys)
	}
}

func TestTwoItemsRightOrder(t *testing.T) {
	node := Node{}

	node.addItem(2)
	node.addItem(3)

	if slicesDifferent(node.keys, []NodeInt{2, 3}) {
		t.Errorf("Added 2 items correctly ordered. Got %v; wanted [2, 3]", node.keys)
	}
}

func TestTwoItemsWrongOrder(t *testing.T) {
	node := Node{}

	node.addItem(6)
	node.addItem(4)

	if slicesDifferent(node.keys, []NodeInt{4, 6}) {
		t.Errorf("Added 2 items incorrectly ordered. Got %v; wanted [4, 6]", node.keys)
	}
}

func TestAddBetweenTwoItems(t *testing.T) {
	node := Node{}
	node.addItem(15)
	node.addItem(18)

	node.addItem(17)

	if slicesDifferent(node.keys, []NodeInt{15, 17, 18}) {
		t.Errorf("Added item between 2 other items. Got %v; wanted [15, 17, 18]", node.keys)
	}
}

func TestAddBeforeTwoItems(t *testing.T) {
	node := Node{}
	node.addItem(15)
	node.addItem(18)

	node.addItem(13)

	if slicesDifferent(node.keys, []NodeInt{13, 15, 18}) {
		t.Errorf("Added item before 2 other items. Got %v; wanted [13, 15, 18]", node.keys)
	}
}

func TestAddAfterTwoItems(t *testing.T) {
	node := Node{}
	node.addItem(15)
	node.addItem(18)

	node.addItem(22)

	if slicesDifferent(node.keys, []NodeInt{15, 18, 22}) {
		t.Errorf("Added item after 2 other items. Got %v; wanted [15, 18, 22]", node.keys)
	}
}
