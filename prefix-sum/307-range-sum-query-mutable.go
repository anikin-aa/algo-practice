package main

// 1st non-optimal approach
// recalculate on each update
type NumArray307 struct {
	holder []int
	sum    []int
}

func Constructor307(nums []int) NumArray307 {
	n := len(nums)
	sum := make([]int, n+1)

	for i := 1; i < len(sum); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	return NumArray307{holder: nums, sum: sum}
}

func (n *NumArray307) Update(index int, val int) {
	n.holder[index] = val
	for i := index + 1; i < len(n.sum); i++ {
		n.sum[i] = n.holder[i-1] + n.sum[i-1]
	}
}

func (n *NumArray307) SumRange(l int, r int) int {
	return n.sum[r+1] - n.sum[l]
}

// 2nd approach with segment tree
// https://cp-algorithms.com/data_structures/segment_tree.html#implementation
type SegmentTree struct {
	values []int
	tree   []int
}

func NewSegmentTree(from []int) *SegmentTree {
	tree := SegmentTree{
		values: from,
		tree:   make([]int, 4*len(from)),
	}

	tree.Build(1, 0, len(from)-1)

	return &tree
}

func (t *SegmentTree) Build(cur, left, right int) {
	if left == right {
		t.tree[cur] = t.values[left]
	} else {
		mid := (left + right) / 2

		t.Build(2*cur, left, mid)
		t.Build(2*cur+1, mid+1, right)

		t.tree[cur] = t.tree[2*cur] + t.tree[2*cur+1]
	}
}

func (t *SegmentTree) Query(cur, tLeft, tRight, left, right int) int {
	if left > right {
		return 0
	}
	if tLeft == left && tRight == right {
		return t.tree[cur]
	}
	mid := (tLeft + tRight) / 2
	return t.Query(cur*2, tLeft, mid, left, min(right, mid)) + t.Query(cur*2+1, mid+1, tRight, max(left, mid+1), right)
}

func (t *SegmentTree) Update(cur, tLeft, tRight, idxAt, val int) {
	if tLeft == tRight {
		t.tree[cur] = val
		return
	}
	mid := (tLeft + tRight) / 2
	if idxAt <= mid {
		t.Update(cur*2, tLeft, mid, idxAt, val)
	} else {
		t.Update(cur*2+1, mid+1, tRight, idxAt, val)
	}

	t.tree[cur] = t.tree[2*cur] + t.tree[2*cur+1]
}

// func main() {
// 	vals := []int{1, 2, 3, 4, 5}

// 	tree := NewSegmentTree(vals)

// 	fmt.Println(tree.Query(1, 0, len(vals)-1, 0, 4))

// 	tree.Update(1, 0, len(vals)-1, 0, 5)

// 	fmt.Println(tree.Query(1, 0, len(vals)-1, 0, 4))

// }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
