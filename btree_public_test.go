package btree_test

import (
	"testing"

	"github.com/moltenwolfcub/btree"
)

func TestSearchSmallRoot(t *testing.T) {
	tree := btree.NewBTree()
	tree.Insert(btree.NewElement(6, "information"))

	got := tree.Search(6)

	if got != "information" {
		t.Errorf("searching tree with only 1 item in root. Got '%v'; Wanted 'information'", got)
	}
}
func TestSearchFullRoot(t *testing.T) {
	tree := btree.NewBTree()
	tree.Insert(btree.NewElement(6, "information"))
	tree.Insert(btree.NewElement(8, "differentInfo"))
	tree.Insert(btree.NewElement(82, "lorem ipsum"))
	tree.Insert(btree.NewElement(1, "interestingData"))

	got1 := tree.Search(8)
	got2 := tree.Search(82)

	if got1 != "differentInfo" {
		t.Errorf("searching tree with only a full root. Got '%v'; Wanted 'differentInfo'", got1)
	}
	if got2 != "lorem ipsum" {
		t.Errorf("searching tree with only a full root. Got '%v'; Wanted 'lorem ipsum'", got2)
	}
}
func TestSearchChildNode(t *testing.T) {
	tree := btree.NewBTree()
	tree.Insert(btree.NewElement(10, "root"))
	tree.Insert(btree.NewElement(2, "lefty"))
	tree.Insert(btree.NewElement(5, "lester"))
	tree.Insert(btree.NewElement(13, "righty"))
	tree.Insert(btree.NewElement(3098, "roger"))

	roger := tree.Search(3098)
	lefty := tree.Search(2)

	if roger != "roger" {
		t.Errorf("searching tree's child nodes. Got '%v'; Wanted 'roger'", roger)
	}
	if lefty != "lefty" {
		t.Errorf("searching tree with only a full root. Got '%v'; Wanted 'lefty'", lefty)
	}
}
func TestSearchRootNodeWithChild(t *testing.T) {
	tree := btree.NewBTree()
	tree.Insert(btree.NewElement(10, "root"))
	tree.Insert(btree.NewElement(2, "lefty"))
	tree.Insert(btree.NewElement(5, "lester"))
	tree.Insert(btree.NewElement(13, "righty"))
	tree.Insert(btree.NewElement(3098, "roger"))

	got := tree.Search(10)

	if got != "root" {
		t.Errorf("searching tree for root with child nodes present. Got '%v'; Wanted 'root'", got)
	}
}
