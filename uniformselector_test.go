package gocrush

import (
	"github.com/stretchr/testify/assert"
	//"log"
	"strconv"
	"testing"
)

func TestUniformSelector(t *testing.T) {
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
	node := new(TestingNode)
	node.Children = make([]Node, 8)
	counter := make(map[string]Node)
	for i := 0; i < 8; i++ {
		child := new(TestingNode)
		child.Weight = 1
		child.ID = "Child" + strconv.Itoa(i)
		node.Children[i] = child
	}
	selector := NewUniformSelector(node)
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
