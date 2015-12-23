package gocrush

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"sort"
)

type HashingSelector struct {
	tokenList tokenList
	tokenMap  map[int64]Node
}

func NewHashingSelector(n Node) *HashingSelector {
	var h = new(HashingSelector)
	var nodes = n.GetChildren()
	var maxWeight int64 = 0
	for _, node := range nodes {
		maxWeight = Max64(maxWeight, node.GetWeight())
	}
	h.tokenMap = make(map[int64]Node)
	for _, node := range nodes {
		var count int64 = 500 * node.GetWeight() / maxWeight
		var hash []byte
		for i := int64(0); i < count; i++ {
			var input []byte
			if len(hash) == 0 {
				input = []byte(node.GetId())
			} else {
				input = hash
			}
			hash = digestBytes(input)
			token := Btoi(hash)
			if _, ok := h.tokenMap[token]; !ok {
				h.tokenMap[token] = node
			}
		}
	}
	h.tokenList = make([]int64, 0, len(h.tokenMap))
	for k, _ := range h.tokenMap {
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

func Max64(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
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

func Btoi(b []byte) int64 {
	var result int64
	buf := bytes.NewReader(b)
	binary.Read(buf, binary.LittleEndian, &result)
	return result
}

func (s *HashingSelector) Select(input int64, round int64) Node {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, input)
	binary.Write(buf, binary.LittleEndian, round)
	bytes := buf.Bytes()
	hash := digestBytes(bytes)
	token := Btoi(hash)
	return s.tokenMap[s.findToken(token)]
}

func (s *HashingSelector) findToken(token int64) int64 {
	i := sort.Search(len(s.tokenList), func(i int) bool { return s.tokenList[i] >= token })
	if i == len(s.tokenList) {
		return s.tokenList[i-1]
	}
	return s.tokenList[i]
}
