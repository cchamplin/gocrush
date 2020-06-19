package gocrush

import (
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	root       = 0
	dataCenter = 1
	rack       = 2
	node       = 3
	disk       = 4
)

func TestCrushStraw(t *testing.T) {
	tree := makeStrawTree()
	nodes1 := Select(tree, 15, 3, node)

	nodes2 := Select(tree, 4564564564, 3, node)

	nodes3 := Select(tree, 8789342322, 3, node)
	for _, node := range nodes1 {
		log.Printf("[STRAW] For key %d got node : %s", 15, node.GetID())
	}
	for _, node := range nodes2 {
		log.Printf("[STRAW] For key %d got node : %s", 4564564564, node.GetID())
	}
	for _, node := range nodes3 {
		log.Printf("[STRAW] For key %d got node : %s", 8789342322, node.GetID())
	}
	assert.Equal(t, len(tree.GetChildren()), 4, "")
	//assert.Equal(t, len(tree.GetChildren()), 5, "")
}

func TestCrushStrawTreeChange(t *testing.T) {
	tree := makeStrawTree()
	var key int64 = 64646436
	nodes := Select(tree, key, 3, node)

	subTree, _ := tree.Children[2].(*TestingNode)
	subSubTree, _ := subTree.Children[2].(*TestingNode)

	subSubTree.Children = append(subSubTree.Children[:1], subSubTree.Children[2:]...)
	subSubTree.Selector = NewStrawSelector(subSubTree)
	for idx, node := range subSubTree.GetChildren() {
		log.Printf("[STRAW] Node: (%d idx) %s", idx, node.GetID())
	}
	nodes2 := Select(tree, key, 3, node)

	for _, node := range nodes {
		log.Printf("[STRAW] For key %d got node : %s", key, node.GetID())
	}

	for _, node := range nodes2 {
		log.Printf("[STRAW] For key %d got node : %s", key, node.GetID())
	}
	assert.Equal(t, len(tree.GetChildren()), 4, "")
	//assert.Equal(t, len(tree.GetChildren()), 5, "")
}

func TestCrushTree(t *testing.T) {
	tree := makeTreeTree()
	nodes1 := Select(tree, 15, 3, node)

	nodes2 := Select(tree, 4564564564, 3, node)

	nodes3 := Select(tree, 8789342322, 3, node)
	for _, node := range nodes1 {
		log.Printf("[TREE] For key %d got node : %s", 15, node.GetID())
	}
	for _, node := range nodes2 {
		log.Printf("[TREE] For key %d got node : %s", 4564564564, node.GetID())
	}
	for _, node := range nodes3 {
		log.Printf("[TREE] For key %d got node : %s", 8789342322, node.GetID())
	}
	assert.Equal(t, len(tree.GetChildren()), 4, "")
	//assert.Equal(t, len(tree.GetChildren()), 5, "")
}

func TestCrushTreeTreeChange(t *testing.T) {
	tree := makeTreeTree()
	var key int64 = 64646436
	nodes := Select(tree, key, 3, node)

	subTree, _ := tree.Children[3].(*TestingNode)
	subSubTree, _ := subTree.Children[0].(*TestingNode)

	subSubTree.Children = append(subSubTree.Children[:1], subSubTree.Children[2:]...)
	subSubTree.Selector = NewTreeSelector(subSubTree)
	for idx, node := range subSubTree.GetChildren() {
		log.Printf("[TREE] Node: (%d idx) %s", idx, node.GetID())
	}
	nodes2 := Select(tree, key, 3, node)

	for _, node := range nodes {
		log.Printf("[TREE] For key %d got node : %s", key, node.GetID())
	}

	for _, node := range nodes2 {
		log.Printf("[TREE] For key %d got node : %s", key, node.GetID())
	}
	assert.Equal(t, len(tree.GetChildren()), 4, "")
	//assert.Equal(t, len(tree.GetChildren()), 5, "")
}

func makeStrawTree() *TestingNode {
	var parent = new(TestingNode)
	parent.ID = "root"
	parent.Type = root
	parent.Weight = 0
	parent.Children = make([]Node, 4)
	for dc := 0; dc < 4; dc++ {
		var dcNode = new(TestingNode)
		dcNode.Parent = parent
		dcNode.Weight = 1
		dcNode.Type = dataCenter
		dcNode.ID = parent.ID + ":DataCenter" + strconv.Itoa(dc)
		dcNode.Children = make([]Node, 4)

		parent.Children[dc] = dcNode

		for rk := 0; rk < 4; rk++ {
			var rkNode = new(TestingNode)
			rkNode.Parent = dcNode
			rkNode.Type = rack
			rkNode.Weight = 1
			rkNode.ID = dcNode.ID + ":rack" + strconv.Itoa(rk)
			rkNode.Children = make([]Node, 4)

			dcNode.Children[rk] = rkNode
			for nd := 0; nd < 4; nd++ {
				var ndNode = new(TestingNode)
				ndNode.Parent = rkNode
				ndNode.Type = node
				ndNode.Weight = 1
				ndNode.ID = rkNode.ID + ":Node" + strconv.Itoa(nd)
				ndNode.Children = make([]Node, 4)

				rkNode.Children[nd] = ndNode
				for dsk := 0; dsk < 4; dsk++ {
					var dskNode = new(TestingNode)
					dskNode.Parent = ndNode
					dskNode.Type = disk
					dskNode.Weight = 1
					dskNode.ID = ndNode.ID + ":Disk" + strconv.Itoa(dsk)
					dskNode.Selector = NewStrawSelector(dskNode)
					ndNode.Children[dsk] = dskNode
				}
				ndNode.Selector = NewStrawSelector(ndNode)
			}
			rkNode.Selector = NewStrawSelector(rkNode)
		}
		dcNode.Selector = NewStrawSelector(dcNode)
	}
	parent.Selector = NewStrawSelector(parent)
	return parent
}

func makeSimpleStrawTree() *TestingNode {
	var parent = new(TestingNode)
	parent.ID = "root"
	parent.Type = root
	parent.Weight = 0
	parent.Children = make([]Node, 1)
	for dc := 0; dc < 1; dc++ {
		var dcNode = new(TestingNode)
		dcNode.Parent = parent
		dcNode.Weight = 1
		dcNode.Type = dataCenter
		dcNode.ID = parent.ID + ":DataCenter" + strconv.Itoa(dc)
		dcNode.Children = make([]Node, 1)

		parent.Children[dc] = dcNode

		for rk := 0; rk < 1; rk++ {
			var rkNode = new(TestingNode)
			rkNode.Parent = dcNode
			rkNode.Type = rack
			rkNode.Weight = 1
			rkNode.ID = dcNode.ID + ":rack" + strconv.Itoa(rk)
			rkNode.Children = make([]Node, 3)

			dcNode.Children[rk] = rkNode
			for nd := 0; nd < 3; nd++ {
				var ndNode = new(TestingNode)
				ndNode.Parent = rkNode
				ndNode.Type = node
				ndNode.Weight = 1
				ndNode.ID = rkNode.ID + ":Node" + strconv.Itoa(nd)
				ndNode.Children = make([]Node, 1)

				rkNode.Children[nd] = ndNode
				for dsk := 0; dsk < 1; dsk++ {
					var dskNode = new(TestingNode)
					dskNode.Parent = ndNode
					dskNode.Type = disk
					dskNode.Weight = 1
					dskNode.ID = ndNode.ID + ":Disk" + strconv.Itoa(dsk)
					dskNode.Selector = NewStrawSelector(dskNode)
					ndNode.Children[dsk] = dskNode
				}
				ndNode.Selector = NewStrawSelector(ndNode)
			}
			rkNode.Selector = NewStrawSelector(rkNode)
		}
		dcNode.Selector = NewStrawSelector(dcNode)
	}
	parent.Selector = NewStrawSelector(parent)
	return parent
}

func makeTreeTree() *TestingNode {
	var parent = new(TestingNode)
	parent.ID = "root"
	parent.Type = root
	parent.Weight = 0
	parent.Children = make([]Node, 4)
	for dc := 0; dc < 4; dc++ {
		var dcNode = new(TestingNode)
		dcNode.Parent = parent
		dcNode.Weight = 1
		dcNode.Type = dataCenter
		dcNode.ID = parent.ID + ":DataCenter" + strconv.Itoa(dc)
		dcNode.Children = make([]Node, 4)

		parent.Children[dc] = dcNode

		for rk := 0; rk < 4; rk++ {
			var rkNode = new(TestingNode)
			rkNode.Parent = dcNode
			rkNode.Type = rack
			rkNode.Weight = 1
			rkNode.ID = dcNode.ID + ":rack" + strconv.Itoa(rk)
			rkNode.Children = make([]Node, 4)

			dcNode.Children[rk] = rkNode
			for nd := 0; nd < 4; nd++ {
				var ndNode = new(TestingNode)
				ndNode.Parent = rkNode
				ndNode.Type = node
				ndNode.Weight = 1
				ndNode.ID = rkNode.ID + ":Node" + strconv.Itoa(nd)
				ndNode.Children = make([]Node, 4)

				rkNode.Children[nd] = ndNode
				for dsk := 0; dsk < 4; dsk++ {
					var dskNode = new(TestingNode)
					dskNode.Parent = ndNode
					dskNode.Type = disk
					dskNode.Weight = 1
					dskNode.ID = ndNode.ID + ":Disk" + strconv.Itoa(dsk)
					dskNode.Selector = NewTreeSelector(dskNode)
					ndNode.Children[dsk] = dskNode
				}
				ndNode.Selector = NewTreeSelector(ndNode)
			}
			rkNode.Selector = NewTreeSelector(rkNode)
		}
		dcNode.Selector = NewTreeSelector(dcNode)
	}
	parent.Selector = NewTreeSelector(parent)
	return parent
}

func makeSimpleTreeTree() *TestingNode {
	var parent = new(TestingNode)
	parent.ID = "root"
	parent.Type = root
	parent.Weight = 0
	parent.Children = make([]Node, 1)
	for dc := 0; dc < 1; dc++ {
		var dcNode = new(TestingNode)
		dcNode.Parent = parent
		dcNode.Weight = 1
		dcNode.Type = dataCenter
		dcNode.ID = parent.ID + ":DataCenter" + strconv.Itoa(dc)
		dcNode.Children = make([]Node, 1)

		parent.Children[dc] = dcNode

		for rk := 0; rk < 1; rk++ {
			var rkNode = new(TestingNode)
			rkNode.Parent = dcNode
			rkNode.Type = rack
			rkNode.Weight = 1
			rkNode.ID = dcNode.ID + ":rack" + strconv.Itoa(rk)
			rkNode.Children = make([]Node, 3)

			dcNode.Children[rk] = rkNode
			for nd := 0; nd < 3; nd++ {
				var ndNode = new(TestingNode)
				ndNode.Parent = rkNode
				ndNode.Type = node
				ndNode.Weight = 1
				ndNode.ID = rkNode.ID + ":Node" + strconv.Itoa(nd)
				ndNode.Children = make([]Node, 1)

				rkNode.Children[nd] = ndNode
				for dsk := 0; dsk < 1; dsk++ {
					var dskNode = new(TestingNode)
					dskNode.Parent = ndNode
					dskNode.Type = disk
					dskNode.Weight = 1
					dskNode.ID = ndNode.ID + ":Disk" + strconv.Itoa(dsk)
					dskNode.Selector = NewTreeSelector(dskNode)
					ndNode.Children[dsk] = dskNode
				}
				ndNode.Selector = NewTreeSelector(ndNode)
			}
			rkNode.Selector = NewTreeSelector(rkNode)
		}
		dcNode.Selector = NewTreeSelector(dcNode)
	}
	parent.Selector = NewTreeSelector(parent)
	return parent
}
