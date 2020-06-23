package gocrush

import (
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkCrushStraw(b *testing.B) {
	b.StopTimer()
	tree := makeBenchStrawTree()
	r := rand.New(rand.NewSource(544564))
	b.StartTimer()
	g, err := New()
	if err != nil {
		panic(err)
	}
	for x := 0; x < b.N; x++ {
		g.Select(tree, r.Int63(), 3, Node)
	}
	b.StopTimer()
}

func BenchmarkCrushTree(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(544564))
	tree := makeBenchTreeTree()
	b.StartTimer()
	g, err := New()
	if err != nil {
		panic(err)
	}
	for x := 0; x < b.N; x++ {
		g.Select(tree, r.Int63(), 3, Node)

	}
	b.StopTimer()
}

func makeBenchStrawTree() *CrushNode {
	var parent = new(CrushNode)
	parent.id = "root"
	parent.group = Root
	parent.weight = 0
	parent.childrens = make([]CNode, 5)
	for dc := 0; dc < 5; dc++ {
		var dcNode = new(CrushNode)
		dcNode.parent = parent
		dcNode.weight = 1
		dcNode.group = DataCenter
		dcNode.id = parent.id + ":DataCenter" + strconv.Itoa(dc)
		dcNode.childrens = make([]CNode, 50)

		parent.childrens[dc] = dcNode

		for rk := 0; rk < 50; rk++ {
			var rkNode = new(CrushNode)
			rkNode.parent = dcNode
			rkNode.group = Rack
			rkNode.weight = 1
			rkNode.id = dcNode.id + ":rack" + strconv.Itoa(rk)
			rkNode.childrens = make([]CNode, 5000)

			dcNode.childrens[rk] = rkNode
			for nd := 0; nd < 5000; nd++ {
				var ndNode = new(CrushNode)
				ndNode.parent = rkNode
				ndNode.group = Node
				ndNode.weight = 1
				ndNode.id = rkNode.id + ":Node" + strconv.Itoa(nd)
				ndNode.childrens = make([]CNode, 2)

				rkNode.childrens[nd] = ndNode
				for dsk := 0; dsk < 2; dsk++ {
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

func makeBenchSimpleStrawTree() *CrushNode {
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

func makeBenchTreeTree() *CrushNode {
	var parent = new(CrushNode)
	parent.id = "root"
	parent.group = Root
	parent.weight = 0
	parent.childrens = make([]CNode, 5)
	for dc := 0; dc < 5; dc++ {
		var dcNode = new(CrushNode)
		dcNode.parent = parent
		dcNode.weight = 1
		dcNode.group = DataCenter
		dcNode.id = parent.id + ":DataCenter" + strconv.Itoa(dc)
		dcNode.childrens = make([]CNode, 50)

		parent.childrens[dc] = dcNode

		for rk := 0; rk < 50; rk++ {
			var rkNode = new(CrushNode)
			rkNode.parent = dcNode
			rkNode.group = Rack
			rkNode.weight = 1
			rkNode.id = dcNode.id + ":rack" + strconv.Itoa(rk)
			rkNode.childrens = make([]CNode, 5000)

			dcNode.childrens[rk] = rkNode
			for nd := 0; nd < 5000; nd++ {
				var ndNode = new(CrushNode)
				ndNode.parent = rkNode
				ndNode.group = Node
				ndNode.weight = 1
				ndNode.id = rkNode.id + ":Node" + strconv.Itoa(nd)
				ndNode.childrens = make([]CNode, 2)

				rkNode.childrens[nd] = ndNode
				for dsk := 0; dsk < 2; dsk++ {
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

func makeBenchSimpleTreeTree() *CrushNode {
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
