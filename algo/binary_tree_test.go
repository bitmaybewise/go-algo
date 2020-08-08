package algo

import (
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	root := NewBinaryTree(1)
	if root.item != 1 {
		t.Error("BinaryTree item not initialized")
	}
	nils := map[string]*BinaryTree{"parent": root.parent, "left": root.left, "right": root.right}
	for k, v := range nils {
		if v != nil {
			t.Errorf("BinaryTree root's %s should be nil", k)
		}
	}
}

func TestBinaryTreeInsert(t *testing.T) {
	assertLeftRight := func(l *BinaryTree, r *BinaryTree, left int, right int) {
		if l.item != left {
			t.Errorf("Wrong left value, should be %d", left)
		}
		if r.item != right {
			t.Errorf("Wrong right value, should be %d", right)
		}
	}
	root := NewBinaryTree(10)
	root.Insert(5)
	root.Insert(15)
	assertLeftRight(root.left, root.right, 5, 15)
	root.Insert(4)
	root.Insert(7)
	root.Insert(12)
	root.Insert(17)
	assertLeftRight(root.left.left, root.left.right, 4, 7)
	assertLeftRight(root.right.left, root.right.right, 12, 17)
}

func TestBinaryTreeFindMinimum(t *testing.T) {
	root := NewBinaryTree(10)
	root.Insert(5)
	root.Insert(15)
	root.Insert(12)
	root.Insert(17)
	root.Insert(1)
	root.Insert(7)

	if root.FindMinimum().item != 1 {
		t.Error("Wrong minimum")
	}
}

func TestBinrayTreeFindMaximum(t *testing.T) {
	root := NewBinaryTree(10)
	root.Insert(5)
	root.Insert(15)
	root.Insert(12)
	root.Insert(17)
	root.Insert(1)
	root.Insert(7)

	if root.FindMaximum().item != 17 {
		t.Error("Wrong maximum")
	}
}

func TestBinaryTreeTraverse(t *testing.T) {
	root := NewBinaryTree(10)
	root.Insert(5)
	root.Insert(15)
	root.Insert(12)
	root.Insert(17)
	root.Insert(1)
	root.Insert(7)

	result := []int{}
	fn := func(leaf *BinaryTree) {
		result = append(result, leaf.item)
	}

	root.Traverse(fn)
	expected := []int{1, 5, 7, 10, 12, 15, 17}
	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Errorf("Wrong order of traverse at index %d, value should be %d but got %d", i, expected[i], result[i])
		}
	}
}

func TestBinaryTreeCount(t *testing.T) {
	root := NewBinaryTree(10)
	root.Insert(5)
	root.Insert(15)
	root.Insert(12)
	root.Insert(17)
	root.Insert(1)
	root.Insert(7)

	if count := root.Count(); count != 7 {
		t.Errorf("BinaryTree has wrong size, should be %d but got %d", 7, count)
	}
}

func TestBinaryTreeDelete(t *testing.T) {
	t.Run("When item not found", func(t *testing.T) {
		root := NewBinaryTree(10)
		result := root.Delete(11)
		if result != root && root.Count() == 1 {
			t.Errorf("BinaryTree single leaf wrongly deleted = %v", result)
		}
	})
	t.Run("When single root node", func(t *testing.T) {
		root := NewBinaryTree(10)
		result := root.Delete(10)
		if result != nil {
			t.Error("BinaryTree single root node still alive")
		}
	})
	t.Run("When leaf has no child", func(t *testing.T) {
		root := NewBinaryTree(10)
		root.Insert(11)
		root.Insert(12)
		result := root.Delete(12)
		if count := root.Count(); count != 2 && result == root {
			t.Errorf("BinaryTree wrongly deleted leaf without children, contains %d leaves instead of %d", count, 2)
		}
	})
	t.Run("When leaf has a single child to the right", func(t *testing.T) {
		root := NewBinaryTree(10)
		root.Insert(11)
		root.Insert(12)
		result := root.Delete(11)
		if count := root.Count(); count != 2 && result == root {
			t.Errorf("BinaryTree wrongly deleted leaf with single children, contains %d leaves instead of %d", count, 2)
		}
	})
	t.Run("When leaf has a single child to the left", func(t *testing.T) {
		root := NewBinaryTree(10)
		root.Insert(5)
		root.Insert(2)
		result := root.Delete(5)
		if count := root.Count(); count != 2 && result == root {
			t.Errorf("BinaryTree wrongly deleted leaf with single children, contains %d leaves instead of %d", count, 2)
		}
	})
	t.Run("When leaf has two children", func(t *testing.T) {
		root := NewBinaryTree(10)
		root.Insert(5)
		root.Insert(2)
		root.Insert(6)
		result := root.Delete(5)
		if count := root.Count(); count != 3 && result == root {
			t.Errorf("BinaryTree wrongly deleted leaf with single children, contains %d leaves instead of %d", count, 3)
		}
	})
	t.Run("When leaf has two children deeply nested to the left", func(t *testing.T) {
		root := NewBinaryTree(10)
		root.Insert(5)
		root.Insert(2)
		root.Insert(7)
		root.Insert(6)
		root.Insert(9)
		result := root.Delete(5)
		if count := root.Count(); count != 5 && result == root {
			t.Errorf("BinaryTree wrongly deleted leaf with single children, contains %d leaves instead of %d", count, 5)
		}
	})
	t.Run("When leaf has two children deeply nested to the right", func(t *testing.T) {
		root := NewBinaryTree(10)
		root.Insert(25)
		root.Insert(12)
		root.Insert(17)
		root.Insert(16)
		root.Insert(19)
		result := root.Delete(25)
		if count := root.Count(); count != 5 && result == root {
			t.Errorf("BinaryTree wrongly deleted leaf with single children, contains %d leaves instead of %d", count, 5)
		}
	})
}
