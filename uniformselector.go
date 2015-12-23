package gocrush

import ()

type UniformSelector struct {
	Node        Node
	totalWeight int64

	perm  int64
	perms []int64
	// This is not thread safe, which makes me unhappy....
	curInput int64
}

func NewUniformSelector(n Node) *UniformSelector {
	var u = new(UniformSelector)
	if !n.IsLeaf() {
		u.totalWeight = int64(len(n.GetChildren())) * n.GetWeight()
		u.Node = n
		u.perms = make([]int64, len(n.GetChildren()))
		u.perm = 0
		u.curInput = -1
	}

	return u
}

func (s *UniformSelector) Select(input int64, round int64) Node {
	var size = len(s.Node.GetChildren())
	var pr int64 = int64(round % int64(size))
	if s.curInput != input || s.perm == 0 {
		s.curInput = input
		if pr == 0 {
			hash := hash3(input, Btoi(digestString(s.Node.GetId())), 0) % int64(size)
			s.perms[0] = hash
			s.perm = 0xffff
			return s.Node.GetChildren()[hash]
		}
		for i := 0; i < size; i++ {
			s.perms[i] = int64(i)
		}
		s.perm = 0
	} else if s.perm == 0xffff {
		for i := 1; i < size; i++ {
			s.perms[i] = int64(i)
		}
		s.perms[s.perms[0]] = 0
		s.perm = 1
	}
	for s.perm <= pr {
		var p = s.perm
		if p < int64(size-1) {
			hash := hash3(input, Btoi(digestString(s.Node.GetId())), p) % (int64(size) - p)
			if hash > 0 {
				var t = s.perms[p+hash]
				s.perms[p+hash] = s.perms[p]
				s.perms[p] = t
			}
		}
		s.perm += 1
	}
	return s.Node.GetChildren()[s.perms[pr]]
}
