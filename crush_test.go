package gocrush

import (
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrushStraw(t *testing.T) {
	tree := makeStrawTree()
	g, err := New()
	if err != nil {
		panic(err)
	}
	nodes1 := g.Select(tree, 15, 3, Node)

	nodes2 := g.Select(tree, 4564564564, 3, Node)

	nodes3 := g.Select(tree, 8789342322, 3, Node)
	for _, node := range nodes1 {
		log.Printf("[STRAW] For key %d got node : %s", 15, node.GetID())
	}
	for _, node := range nodes2 {
		log.Printf("[STRAW] For key %d got node : %s", 4564564564, node.GetID())
	}
	for _, node := range nodes3 {
		log.Printf("[STRAW] For key %d got node : %s", 8789342322, node.GetID())
	}
	assert.Equal(t, len(tree.GetChildrens()), 4, "")
	//assert.Equal(t, len(tree.GetChildrens()), 5, "")
}

func TestCrushStrawTreeChange(t *testing.T) {
	tree := makeStrawTree()
	var key int64 = 64646436
	g, err := New()
	if err != nil {
		panic(err)
	}
	nodes := g.Select(tree, key, 3, Node)

	subTree, _ := tree.childrens[2].(*CrushNode)
	subSubTree, _ := subTree.childrens[2].(*CrushNode)

	subSubTree.childrens = append(subSubTree.childrens[:1], subSubTree.childrens[2:]...)
	subSubTree.selector = NewStrawSelector(subSubTree)
	for idx, node := range subSubTree.GetChildrens() {
		log.Printf("[STRAW] Node: (%d idx) %s", idx, node.GetID())
	}
	nodes2 := g.Select(tree, key, 3, Node)

	for _, node := range nodes {
		log.Printf("[STRAW] For key %d got node : %s", key, node.GetID())
	}

	for _, node := range nodes2 {
		log.Printf("[STRAW] For key %d got node : %s", key, node.GetID())
	}
	assert.Equal(t, len(tree.GetChildrens()), 4, "")
	//assert.Equal(t, len(tree.GetChildrens()), 5, "")
}

func TestCrushTree(t *testing.T) {
	tree := makeTreeTree()
	g, err := New()
	if err != nil {
		panic(err)
	}
	nodes1 := g.Select(tree, 15, 3, Node)

	nodes2 := g.Select(tree, 4564564564, 3, Node)

	nodes3 := g.Select(tree, 8789342322, 3, Node)
	for _, node := range nodes1 {
		log.Printf("[TREE] For key %d got node : %s", 15, node.GetID())
	}
	for _, node := range nodes2 {
		log.Printf("[TREE] For key %d got node : %s", 4564564564, node.GetID())
	}
	for _, node := range nodes3 {
		log.Printf("[TREE] For key %d got node : %s", 8789342322, node.GetID())
	}
	assert.Equal(t, len(tree.GetChildrens()), 4, "")
	//assert.Equal(t, len(tree.GetChildrens()), 5, "")
}

func TestCrushTreeTreeChange(t *testing.T) {
	tree := makeTreeTree()
	var key int64 = 64646436
	g, err := New()
	if err != nil {
		panic(err)
	}
	nodes := g.Select(tree, key, 3, Node)

	subTree, _ := tree.childrens[3].(*CrushNode)
	subSubTree, _ := subTree.childrens[0].(*CrushNode)

	subSubTree.childrens = append(subSubTree.childrens[:1], subSubTree.childrens[2:]...)
	subSubTree.selector = NewTreeSelector(subSubTree)
	for idx, node := range subSubTree.GetChildrens() {
		log.Printf("[TREE] Node: (%d idx) %s", idx, node.GetID())
	}
	nodes2 := g.Select(tree, key, 3, Node)

	for _, node := range nodes {
		log.Printf("[TREE] For key %d got node : %s", key, node.GetID())
	}

	for _, node := range nodes2 {
		log.Printf("[TREE] For key %d got node : %s", key, node.GetID())
	}
	assert.Equal(t, len(tree.GetChildrens()), 4, "")
	//assert.Equal(t, len(tree.GetChildrens()), 5, "")
}

func makeStrawTree() *CrushNode {
	var parent = new(CrushNode)
	parent.id = "root"
	parent.group = Root
	parent.weight = 0
	parent.childrens = make([]CNode, 4)
	for dc := 0; dc < 4; dc++ {
		var dcNode = new(CrushNode)
		dcNode.parent = parent
		dcNode.weight = 1
		dcNode.group = DataCenter
		dcNode.id = parent.id + ":DataCenter" + strconv.Itoa(dc)
		dcNode.childrens = make([]CNode, 4)

		parent.childrens[dc] = dcNode

		for rk := 0; rk < 4; rk++ {
			var rkNode = new(CrushNode)
			rkNode.parent = dcNode
			rkNode.group = Rack
			rkNode.weight = 1
			rkNode.id = dcNode.id + ":rack" + strconv.Itoa(rk)
			rkNode.childrens = make([]CNode, 4)

			dcNode.childrens[rk] = rkNode
			for nd := 0; nd < 4; nd++ {
				var ndNode = new(CrushNode)
				ndNode.parent = rkNode
				ndNode.group = Node
				ndNode.weight = 1
				ndNode.id = rkNode.id + ":Node" + strconv.Itoa(nd)
				ndNode.childrens = make([]CNode, 4)

				rkNode.childrens[nd] = ndNode
				for dsk := 0; dsk < 4; dsk++ {
					var dskNode = new(CrushNode)
					dskNode.parent = ndNode
					dskNode.group = Disk
					dskNode.weight = 1
					dskNode.id = ndNode.id + ":Disk" + strconv.Itoa(dsk)
					dskNode.selector = NewStrawSelector(dskNode)
					ndNode.childrens[dsk] = dskNode
				}
				ndNode.selector = NewStrawSelector(ndNode)
			}
			rkNode.selector = NewStrawSelector(rkNode)
		}
		dcNode.selector = NewStrawSelector(dcNode)
	}
	parent.selector = NewStrawSelector(parent)
	return parent
}

func makeSimpleStrawTree() *CrushNode {
	var parent = new(CrushNode)
	parent.id = "root"
	parent.group = Root
	parent.weight = 0
	parent.childrens = make([]CNode, 1)
	for dc := 0; dc < 1; dc++ {
		var dcNode = new(CrushNode)
		dcNode.parent = parent
		dcNode.weight = 1
		dcNode.group = DataCenter
		dcNode.id = parent.id + ":DataCenter" + strconv.Itoa(dc)
		dcNode.childrens = make([]CNode, 1)

		parent.childrens[dc] = dcNode

		for rk := 0; rk < 1; rk++ {
			var rkNode = new(CrushNode)
			rkNode.parent = dcNode
			rkNode.group = Rack
			rkNode.weight = 1
			rkNode.id = dcNode.id + ":rack" + strconv.Itoa(rk)
			rkNode.childrens = make([]CNode, 3)

			dcNode.childrens[rk] = rkNode
			for nd := 0; nd < 3; nd++ {
				var ndNode = new(CrushNode)
				ndNode.parent = rkNode
				ndNode.group = Node
				ndNode.weight = 1
				ndNode.id = rkNode.id + ":Node" + strconv.Itoa(nd)
				ndNode.childrens = make([]CNode, 1)

				rkNode.childrens[nd] = ndNode
				for dsk := 0; dsk < 1; dsk++ {
					var dskNode = new(CrushNode)
					dskNode.parent = ndNode
					dskNode.group = Disk
					dskNode.weight = 1
					dskNode.id = ndNode.id + ":Disk" + strconv.Itoa(dsk)
					dskNode.selector = NewStrawSelector(dskNode)
					ndNode.childrens[dsk] = dskNode
				}
				ndNode.selector = NewStrawSelector(ndNode)
			}
			rkNode.selector = NewStrawSelector(rkNode)
		}
		dcNode.selector = NewStrawSelector(dcNode)
	}
	parent.selector = NewStrawSelector(parent)
	return parent
}

func makeTreeTree() *CrushNode {
	var parent = new(CrushNode)
	parent.id = "root"
	parent.group = Root
	parent.weight = 0
	parent.childrens = make([]CNode, 4)
	for dc := 0; dc < 4; dc++ {
		var dcNode = new(CrushNode)
		dcNode.parent = parent
		dcNode.weight = 1
		dcNode.group = DataCenter
		dcNode.id = parent.id + ":DataCenter" + strconv.Itoa(dc)
		dcNode.childrens = make([]CNode, 4)

		parent.childrens[dc] = dcNode

		for rk := 0; rk < 4; rk++ {
			var rkNode = new(CrushNode)
			rkNode.parent = dcNode
			rkNode.group = Rack
			rkNode.weight = 1
			rkNode.id = dcNode.id + ":rack" + strconv.Itoa(rk)
			rkNode.childrens = make([]CNode, 4)

			dcNode.childrens[rk] = rkNode
			for nd := 0; nd < 4; nd++ {
				var ndNode = new(CrushNode)
				ndNode.parent = rkNode
				ndNode.group = Node
				ndNode.weight = 1
				ndNode.id = rkNode.id + ":Node" + strconv.Itoa(nd)
				ndNode.childrens = make([]CNode, 4)

				rkNode.childrens[nd] = ndNode
				for dsk := 0; dsk < 4; dsk++ {
					var dskNode = new(CrushNode)
					dskNode.parent = ndNode
					dskNode.group = Disk
					dskNode.weight = 1
					dskNode.id = ndNode.id + ":Disk" + strconv.Itoa(dsk)
					dskNode.selector = NewTreeSelector(dskNode)
					ndNode.childrens[dsk] = dskNode
				}
				ndNode.selector = NewTreeSelector(ndNode)
			}
			rkNode.selector = NewTreeSelector(rkNode)
		}
		dcNode.selector = NewTreeSelector(dcNode)
	}
	parent.selector = NewTreeSelector(parent)
	return parent
}

func makeSimpleTreeTree() *CrushNode {
	var parent = new(CrushNode)
	parent.id = "root"
	parent.group = Root
	parent.weight = 0
	parent.childrens = make([]CNode, 1)
	for dc := 0; dc < 1; dc++ {
		var dcNode = new(CrushNode)
		dcNode.parent = parent
		dcNode.weight = 1
		dcNode.group = DataCenter
		dcNode.id = parent.id + ":DataCenter" + strconv.Itoa(dc)
		dcNode.childrens = make([]CNode, 1)

		parent.childrens[dc] = dcNode

		for rk := 0; rk < 1; rk++ {
			var rkNode = new(CrushNode)
			rkNode.parent = dcNode
			rkNode.group = Rack
			rkNode.weight = 1
			rkNode.id = dcNode.id + ":rack" + strconv.Itoa(rk)
			rkNode.childrens = make([]CNode, 3)

			dcNode.childrens[rk] = rkNode
			for nd := 0; nd < 3; nd++ {
				var ndNode = new(CrushNode)
				ndNode.parent = rkNode
				ndNode.group = Node
				ndNode.weight = 1
				ndNode.id = rkNode.id + ":Node" + strconv.Itoa(nd)
				ndNode.childrens = make([]CNode, 1)

				rkNode.childrens[nd] = ndNode
				for dsk := 0; dsk < 1; dsk++ {
					var dskNode = new(CrushNode)
					dskNode.parent = ndNode
					dskNode.group = Disk
					dskNode.weight = 1
					dskNode.id = ndNode.id + ":Disk" + strconv.Itoa(dsk)
					dskNode.selector = NewTreeSelector(dskNode)
					ndNode.childrens[dsk] = dskNode
				}
				ndNode.selector = NewTreeSelector(ndNode)
			}
			rkNode.selector = NewTreeSelector(rkNode)
		}
		dcNode.selector = NewTreeSelector(dcNode)
	}
	parent.selector = NewTreeSelector(parent)
	return parent
}
