package gocrush

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"sort"
)

// HashingSelector holds Hashing selector.
type HashingSelector struct {
	tokenList tokenList
	tokenMap  map[int64]CNode
}

// NewHashingSelector returns new instance of HashingSelector
func NewHashingSelector(n CNode) *HashingSelector {
	var h = new(HashingSelector)
	var nodes = n.GetChildrens()
	var maxWeight int64 = 0
	h.tokenMap = make(map[int64]CNode)
	for _, node := range nodes {
		if node.GetWeight() > maxWeight {
			maxWeight = node.GetWeight()
		}
		var hash []byte
		for i := int64(0); i < 500*node.GetWeight()/maxWeight; i++ {
			var input []byte
			if len(hash) == 0 {
				input = []byte(node.GetID())
			} else {
				input = hash
			}
			hash = digestBytes(input)
			token := btoi(hash)
			if _, ok := h.tokenMap[token]; !ok {
				h.tokenMap[token] = node
			}
		}
	}
	h.tokenList = make([]int64, 0, len(h.tokenMap))
	for k := range h.tokenMap {
		h.tokenList = append(h.tokenList, k)
	}
	sort.Sort(h.tokenList)
	return h
}

type tokenList []int64

func (t tokenList) Len() int {
	return len(t)
}

func (t tokenList) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t tokenList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func digestInt64(input int64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, input)
	bytes := buf.Bytes()
	result := sha1.Sum(bytes)
	var hash []byte
	hash = make([]byte, 20)
	copy(hash[:], result[:20])
	return hash
}

func digestBytes(input []byte) []byte {
	result := sha1.Sum(input)
	var hash []byte
	hash = make([]byte, 20)
	copy(hash[:], result[:20])
	return hash
}

func digestString(input string) []byte {
	result := sha1.Sum([]byte(input))
	var hash []byte
	hash = make([]byte, 20)
	copy(hash[:], result[:20])
	return hash
}

func btoi(b []byte) int64 {
	var result int64
	buf := bytes.NewReader(b)
	binary.Read(buf, binary.LittleEndian, &result)
	return result
}

// Select implements Select interface
func (s *HashingSelector) Select(input int64, round int64) CNode {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, input)
	binary.Write(buf, binary.LittleEndian, round)
	bytes := buf.Bytes()
	hash := digestBytes(bytes)
	token := btoi(hash)
	return s.tokenMap[s.findToken(token)]
}

func (s *HashingSelector) findToken(token int64) int64 {
	i := sort.Search(len(s.tokenList), func(i int) bool { return s.tokenList[i] >= token })
	if i == len(s.tokenList) {
		return s.tokenList[i-1]
	}
	return s.tokenList[i]
}
