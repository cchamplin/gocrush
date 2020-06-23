package gocrush

// TreeSelector implements Select interface
type TreeSelector struct {
	Node        CNode
	weights     []int64
	totalWeight int64
}

// NewTreeSelector returns new TreeSelector
func NewTreeSelector(n CNode) *TreeSelector {
	var t = new(TreeSelector)
	if !n.IsLeaf() {
		t.Node = n
		var depth = depth(len(n.GetChildrens()))
		t.weights = make([]int64, 1<<uint(depth))

		//log.Printf("Tree with depth of %d for %d items and %d nodes", depth, len(n.childrens), len(t.weights))
		for idx, child := range n.GetChildrens() {
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
	var newSize int = len(t.Node.GetChildrens()) + 1
	var depth = depth(newSize)
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
	t.Node.childrens = append(t.Node.childrens, n)
	t.totalWeight += n.weight

}*/

func height(n int) int {
	var h int = 0
	for (n & 1) == 0 {
		h++
		n = n >> 1
	}
	return h
}

func parent(n int) int {
	var h int = height(n)
	if n&(1<<uint(h+1)) > 0 {
		return n - (1 << uint(h))
	}
	return n + (1 << uint(h))
}

func depth(size int) int {
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

// Select returns a node
func (s *TreeSelector) Select(input int64, round int64) CNode {
	n := len(s.weights) >> 1
	for (n & 1) < 1 {
		var l int
		w := s.weights[n]
		hash := uint64(hash4(input, int64(n), round, btoi(digestString(s.Node.GetID())))) * uint64(w)
		hash = hash >> 32
		l = left(n)
		if hash < uint64(s.weights[l]) {
			n = l
		} else {
			n = right(n)
		}
	}
	var result CNode
	result = s.Node.GetChildrens()[n>>1]
	return result
}
