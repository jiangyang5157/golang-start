package redblacktree

import (
	"fmt"
	"testing"
)

type key int

func (n key) LessThan(b interface{}) bool {
	value, _ := b.(key)
	return n < value
}

func Test_Print(t *testing.T) {
	tree := NewTree()

	tree.Insert(key(1), "123")
	tree.Insert(key(3), "234")
	tree.Insert(key(4), "dfa3")
	tree.Insert(key(6), "sd4")
	tree.Insert(key(5), "jcd4")
	tree.Insert(key(2), "bcd4")
	if tree.Size() != 6 {
		t.Error("Error size")
		return
	}
	tree.Print()
}

func Test_Find(t *testing.T) {
	tree := NewTree()

	tree.Insert(key(1), "123")
	tree.Insert(key(3), "234")
	tree.Insert(key(4), "dfa3")
	tree.Insert(key(6), "sd4")
	tree.Insert(key(5), "jcd4")
	tree.Insert(key(2), "bcd4")

	n := tree.Find(key(4))
	if n.Value != "dfa3" {
		t.Error("Error value")
		return
	}
	n.Value = "bdsf"
	if n.Value != "bdsf" {
		t.Error("Error value modify")
		return
	}
	value := tree.FindValue(key(5)).(string)
	if value != "jcd4" {
		t.Error("Error value after modifyed other node")
		return
	}
}

func Test_Iterator(t *testing.T) {
	tree := NewTree()

	tree.Insert(key(1), "123")
	tree.Insert(key(3), "234")
	tree.Insert(key(4), "dfa3")
	tree.Insert(key(6), "sd4")
	tree.Insert(key(5), "jcd4")
	tree.Insert(key(2), "bcd4")

	it := tree.Iterator()
	for it != nil {
		fmt.Printf("[%v]{value=%v}\n", it.Key, it.Value)
		it = it.Next()
	}
}

func Test_Delete(t *testing.T) {
	tree := NewTree()

	tree.Insert(key(1), "123")
	tree.Insert(key(3), "234")
	tree.Insert(key(4), "dfa3")
	tree.Insert(key(6), "sd4")
	tree.Insert(key(5), "jcd4")
	tree.Insert(key(2), "bcd4")

	for i := 1; i <= 6; i++ {
		tree.Delete(key(i))
		if tree.Size() != 6-i {
			t.Error("Delete Error")
		}
	}
	tree.Insert(key(1), "bcd4")
	tree.Clear()
	tree.Print()
	if tree.FindValue(key(1)) != nil {
		t.Error("Can't clear")
		return
	}
}
