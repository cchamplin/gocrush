package gocrush

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashingSelector(t *testing.T) {
	node := new(CrushNode)
	node.childrens = make([]CNode, 8)
	counter := make(map[string]int)
	for i := 0; i < 8; i++ {
		child := new(CrushNode)
		child.weight = 1
		child.id = "Child" + strconv.Itoa(i)
		node.childrens[i] = child
		counter[child.id] = 0
	}
	selector := NewHashingSelector(node)

	for i := int64(0); i < 10000; i++ {
		// Get replicants
		for r := int64(0); r < 3; r++ {
			nn := selector.Select(i, r)
			counter[nn.GetID()]++
		}
	}

	assert.Equal(t, counter["Child0"], 3673, "")
	assert.Equal(t, counter["Child1"], 3768, "")
	assert.Equal(t, counter["Child2"], 3646, "")
	assert.Equal(t, counter["Child3"], 3790, "")
	assert.Equal(t, counter["Child4"], 3525, "")
	assert.Equal(t, counter["Child5"], 4098, "")
	assert.Equal(t, counter["Child6"], 3710, "")
	assert.Equal(t, counter["Child7"], 3790, "")
}

func TestHashingSelectorAdd(t *testing.T) {
	node := new(CrushNode)
	node.childrens = make([]CNode, 8)
	counter := make(map[string]CNode)
	for i := 0; i < 8; i++ {
		child := new(CrushNode)
		child.weight = 1
		child.id = "Child" + strconv.Itoa(i)
		node.childrens[i] = child
	}
	selector := NewHashingSelector(node)
	for i := int64(0); i < 5; i++ {
		for r := int64(0); r < 3; r++ {
			nn := selector.Select(i, r)
			counter[strconv.Itoa(int(i))+":"+strconv.Itoa(int(r))] = nn
		}

	}
	child := new(CrushNode)
	child.weight = 1
	child.id = "Child9"
	node.childrens = append(node.childrens, child)
	selector = NewHashingSelector(node)
	for i := int64(0); i < 5; i++ {
		for r := int64(0); r < 3; r++ {
			nn := selector.Select(i, r)
			if i == 4 && r == 0 {
				assert.Equal(t, child, nn, "")
			} else {
				assert.Equal(t, counter[strconv.Itoa(int(i))+":"+strconv.Itoa(int(r))], nn, "")
			}
		}
	}
}
