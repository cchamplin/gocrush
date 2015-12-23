package gocrush

import (
	"fmt"
	"sort"
)

type UnweightedHashSelector struct {
	tokenList utokenList
	tokenMap  map[uint64]Node
}
type utokenList []uint64

func (t utokenList) Len() int {
	return len(t)
}
func (t utokenList) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t utokenList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func hashVal(bKey []byte) uint64 {
	return ((uint64(bKey[3]) << 24) |
		(uint64(bKey[2]) << 16) |
		(uint64(bKey[1]) << 8) |
		(uint64(bKey[0])))
}

func NewUnweightedHashSelector(n Node) *UnweightedHashSelector {
	var s = new(UnweightedHashSelector)
	if !n.IsLeaf() {
		nodes := n.GetChildren()
		s.tokenMap = make(map[uint64]Node)
		var factor int = 60 * len(nodes) * len(nodes)
		idx := 0
		s.tokenList = make([]uint64, len(nodes)*factor*3)
		for _, node := range nodes {
			var bKey []byte
			for c := 0; c < factor; c++ {
				bKey = digestString(fmt.Sprintf("%s-%s", node.GetId(), c))
				for i := 0; i < 3; i++ {
					key := hashVal(bKey[i*4 : i*4+4])
					s.tokenMap[key] = node
					s.tokenList[idx] = key
					idx += 1
				}
			}
		}
	}
	sort.Sort(s.tokenList)
	return s
}
func (s *UnweightedHashSelector) Select(input int64, round int64) Node {
	var hash = hash2(input, round)
	token := uint64(hash)
	return s.tokenMap[s.findToken(token)]
}

func (s *UnweightedHashSelector) findToken(token uint64) uint64 {
	i := sort.Search(len(s.tokenList), func(i int) bool { return s.tokenList[i] > token })
	if i >= len(s.tokenList) {
		return s.tokenList[0]
	}
	return s.tokenList[uint64(i)]
}
