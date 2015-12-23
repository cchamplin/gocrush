package gocrush

import ()

type TreeSelector struct {
	Node        Node
	weights     []int64
	totalWeight int64
}

func NewTreeSelector(n Node) *TreeSelector {
	var t = new(TreeSelector)
	if !n.IsLeaf() {
		t.Node = n
		var depth = calc_depth(len(n.GetChildren()))
		t.weights = make([]int64, 1<<uint(depth))

		//log.Printf("Tree with depth of %d for %d items and %d nodes", depth, len(n.Children), len(t.weights))
		for idx, child := range n.GetChildren() {
			if child == nil {
				panic("Null child")
			}
			node := ((idx + 1) << 1) - 1
			t.weights[node] = child.GetWeight()
			t.totalWeight += child.GetWeight()
			//log.Printf("Tree got node: %d for item %d", node, idx)
			for j := 1; j < depth; j++ {
				node = parent(node)
				t.weights[node] += child.GetWeight()
			}
		}
	}

	return t
}

/*func (t *TreeSelector) AddItem(n Node) {
	var newSize int = len(t.Node.GetChildren()) + 1
	var depth = calc_depth(newSize)
	node := (((newSize - 1) + 1) << 1) - 1
	var newSlice = make([]int64, 1<<uint(depth))
	copy(newSlice, t.weights)
	t.weights = newSlice
	t.weights[node] = n.GetWeight()
	var root int = len(t.weights) / 2
	if depth >= 2 && (node-1) == root {
		t.weights[root] = t.weights[root/2]
	}
	for j := 1; j < depth; j++ {
		node = parent(node)
		t.weights[node] += n.GetWeight()
	}
	t.Node.Children = append(t.Node.Children, n)
	t.totalWeight += n.Weight

}*/

func height(n int) int {
	var h int = 0
	for (n & 1) == 0 {
		h++
		n = n >> 1
	}
	return h
}
func on_right(n, h int) int {
	return n & (1 << uint(h+1))
}
func parent(n int) int {
	var h int = height(n)
	if on_right(n, h) > 0 {
		return n - (1 << uint(h))
	} else {
		return n + (1 << uint(h))
	}
}

func calc_depth(size int) int {
	if size == 0 {
		return 0
	}

	var depth int = 1
	var t int = size - 1
	for t > 0 {
		t = t >> 1
		depth++
	}
	return depth
}

func left(x int) int {
	var h = height(x)
	return x - (1 << uint(h-1))
}

func right(x int) int {
	var h = height(x)
	return x + (1 << uint(h-1))
}

func (s *TreeSelector) Select(input int64, round int64) Node {
	n := len(s.weights) >> 1

	for (n & 1) < 1 {

		var l int
		w := s.weights[n]
		hash := uint64(hash4(input, int64(n), round, Btoi(digestString(s.Node.GetId())))) * uint64(w)

		hash = hash >> 32

		l = left(n)

		if hash < uint64(s.weights[l]) {
			n = l
		} else {
			n = right(n)

		}
	}
	var result Node

	result = s.Node.GetChildren()[n>>1]
	return result
}
