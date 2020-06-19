package gocrush

import (
	"github.com/stretchr/testify/assert"
	//"log"
	"strconv"
	"testing"
)

func TestTreeSelector(t *testing.T) {
	node := new(TestingNode)
	node.Children = make([]Node, 8)
	counter := make(map[string]int)
	for i := 0; i < 8; i++ {
		child := new(TestingNode)
		child.Weight = 1
		child.ID = "Child" + strconv.Itoa(i)
		node.Children[i] = child
		counter[child.ID] = 0
	}
	selector := NewTreeSelector(node)

	for i := int64(0); i < 10000; i++ {
		// Get replicants
		for r := int64(0); r < 3; r++ {
			nn := selector.Select(i, r)
			counter[nn.GetID()]++
		}
	}

	assert.Equal(t, counter["Child0"], 3769, "")
	assert.Equal(t, counter["Child1"], 3693, "")
	assert.Equal(t, counter["Child2"], 3758, "")
	assert.Equal(t, counter["Child3"], 3665, "")
	assert.Equal(t, counter["Child4"], 3799, "")
	assert.Equal(t, counter["Child5"], 3653, "")
	assert.Equal(t, counter["Child6"], 3799, "")
	assert.Equal(t, counter["Child7"], 3864, "")
}

func TestTreeSelectorAdd(t *testing.T) {
	node := new(TestingNode)
	node.Children = make([]Node, 8)
	counter := make(map[string]Node)
	for i := 0; i < 8; i++ {
		child := new(TestingNode)
		child.Weight = 1
		child.ID = "Child" + strconv.Itoa(i)
		node.Children[i] = child
	}
	selector := NewTreeSelector(node)
	for i := int64(0); i < 5; i++ {
		for r := int64(0); r < 3; r++ {
			nn := selector.Select(i, r)
			counter[strconv.Itoa(int(i))+":"+strconv.Itoa(int(r))] = nn

		}

	}
	child := new(TestingNode)
	child.Weight = 1
	child.ID = "Child9"
	node.Children = append(node.Children, child)
	selector = NewTreeSelector(node)
	for i := int64(0); i < 5; i++ {
		for r := int64(0); r < 3; r++ {
			nn := selector.Select(i, r)
			if (i == 2 && r == 1) || (i == 4 && r == 2) {
				assert.Equal(t, child, nn, "")
			} else {
				assert.Equal(t, counter[strconv.Itoa(int(i))+":"+strconv.Itoa(int(r))], nn, "")
			}
		}
	}
}
