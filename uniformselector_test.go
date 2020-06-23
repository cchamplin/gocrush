package gocrush

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniformSelector(t *testing.T) {
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
	selector := NewUniformSelector(node)

	for i := int64(0); i < 10000; i++ {
		// Get replicants
		for r := int64(0); r < 3; r++ {
			nn := selector.Select(i, r)
			counter[nn.GetID()]++
		}
	}

	assert.Equal(t, counter["Child0"], 3795, "")
	assert.Equal(t, counter["Child1"], 3677, "")
	assert.Equal(t, counter["Child2"], 3781, "")
	assert.Equal(t, counter["Child3"], 3742, "")
	assert.Equal(t, counter["Child4"], 3760, "")
	assert.Equal(t, counter["Child5"], 3799, "")
	assert.Equal(t, counter["Child6"], 3748, "")
	assert.Equal(t, counter["Child7"], 3698, "")
}

func TestUniformSelectorAdd(t *testing.T) {
	node := new(CrushNode)
	node.childrens = make([]CNode, 8)
	counter := make(map[string]CNode)
	for i := 0; i < 8; i++ {
		child := new(CrushNode)
		child.weight = 1
		child.id = "Child" + strconv.Itoa(i)
		node.childrens[i] = child
	}
	selector := NewUniformSelector(node)
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
	selector = NewUniformSelector(node)
	for i := int64(0); i < 5; i++ {
		for r := int64(0); r < 3; r++ {
			nn := selector.Select(i, r)
			if (i == 1 && r == 2) || (i == 2 && r == 2) || (i == 4 && r == 1) || (i == 4 && r == 2) {
				assert.Equal(t, counter[strconv.Itoa(int(i))+":"+strconv.Itoa(int(r))], nn, "")
			} else {
				assert.NotEqual(t, counter[strconv.Itoa(int(i))+":"+strconv.Itoa(int(r))], nn, "")
			}

		}
	}
}
