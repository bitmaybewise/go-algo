package algo

type BinaryTree struct {
	item   int
	parent *BinaryTree
	left   *BinaryTree
	right  *BinaryTree
}

func NewBinaryTree(item int) *BinaryTree {
	return &BinaryTree{item: item}
}

func (t *BinaryTree) Insert(item int) {
	if item < t.item {
		if t.left == nil {
			t.left = &BinaryTree{item: item, parent: t}
			return
		}
		t.left.Insert(item)
	}
	if item > t.item {
		if t.right == nil {
			t.right = &BinaryTree{item: item, parent: t}
			return
		}
		t.right.Insert(item)
	}
}

func (t *BinaryTree) FindMinimum() *BinaryTree {
	if t.left != nil {
		return t.left.FindMinimum()
	}
	return t
}

func (t *BinaryTree) FindMaximum() *BinaryTree {
	if t.right != nil {
		return t.right.FindMaximum()
	}
	return t
}

func (t *BinaryTree) Traverse(fn func(leaf *BinaryTree)) {
	if t.left != nil {
		t.left.Traverse(fn)
	}
	fn(t)
	if t.right != nil {
		t.right.Traverse(fn)
	}
}

func (t *BinaryTree) Count() (count int) {
	t.Traverse(func(_ *BinaryTree) { count++ })
	return
}

func (t *BinaryTree) Delete(item int) *BinaryTree {
	if t.left != nil && t.item > item {
		t.left = t.left.Delete(item)
	}
	if t.right != nil && t.item < item {
		t.right = t.right.Delete(item)
	}
	if t.item != item { // when item not found neither on the right nor left
		return t
	}
	if t.left == nil {
		return t.right
	}
	if t.right == nil {
		return t.left
	}
	min := t.right.FindMinimum().item
	t.right = t.right.Delete(min)
	t.item = min
	return t
}
