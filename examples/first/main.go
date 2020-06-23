package main

import (
	"fmt"
	"strconv"

	"github.com/gocrush"
)

func main() {
	tree, err := makeNewStrawTree()
	if err != nil {
		panic(err)
	}
	g, err := gocrush.New(&gocrush.Config{
		MaxRetries: 10,
	})
	if err != nil {
		panic(err)
	}
	selectedNodes := g.Select(tree, 12312313131, 3, gocrush.Node)
	for _, node := range selectedNodes {
		fmt.Printf("Node ID: %v, isLeaf: %v, Type: %v, Parent: %v, IsFailed: %v\n", node.GetID(), node.IsLeaf(), node.GetGroup(), node.GetParent().GetID(), node.IsFailed())

	}
}

func makeNewStrawTree() (*gocrush.CrushNode, error) {
	parent, err := gocrush.NewNode(&gocrush.CNodeConfig{
		ID:     "root",
		Group:  gocrush.Root,
		Weight: 0,
	})
	if err != nil {
		return nil, fmt.Errorf("Could not initialize parent node: %v", err)
	}
	for dc := 0; dc < 4; dc++ {
		dataCenter, err := gocrush.NewNode(&gocrush.CNodeConfig{
			ID:     parent.GetID() + ":dataCenter" + strconv.Itoa(dc),
			Group:  gocrush.DataCenter,
			Weight: 1,
			Parent: parent,
		})
		if err != nil {
			return nil, fmt.Errorf("Could not initialize dataCenter node %v: %v", dc, err)
		}
		parent.AddChildren(dataCenter)
		for rk := 0; rk < 4; rk++ {
			rack, err := gocrush.NewNode(&gocrush.CNodeConfig{
				ID:     parent.GetID() + ":rack" + strconv.Itoa(rk),
				Group:  gocrush.DataCenter,
				Weight: 1,
				Parent: dataCenter,
			})
			if err != nil {
				return nil, fmt.Errorf("Could not initialize rack node %v: %v", rk, err)
			}
			dataCenter.AddChildren(rack)
			for nd := 0; nd < 4; nd++ {
				node, err := gocrush.NewNode(&gocrush.CNodeConfig{
					ID:     parent.GetID() + ":node" + strconv.Itoa(nd),
					Group:  gocrush.DataCenter,
					Weight: 1,
					Parent: rack,
				})
				if err != nil {
					return nil, fmt.Errorf("Could not initialize node %v: %v", nd, err)
				}
				rack.AddChildren(node)
				for dsk := 0; nd < 4; nd++ {
					disk, err := gocrush.NewNode(&gocrush.CNodeConfig{
						ID:     parent.GetID() + ":disk" + strconv.Itoa(dsk),
						Group:  gocrush.DataCenter,
						Weight: 1,
						Parent: node,
					})
					if err != nil {
						return nil, fmt.Errorf("Could not initialize disk node %v: %v", dsk, err)
					}
					disk.SetSelector(gocrush.NewStrawSelector(disk))
					node.AddChildren(disk)
				}
				node.SetSelector(gocrush.NewStrawSelector(node))
			}
			rack.SetSelector(gocrush.NewStrawSelector(rack))
		}
		dataCenter.SetSelector(gocrush.NewStrawSelector(dataCenter))
	}
	parent.SetSelector(gocrush.NewStrawSelector(parent))
	return parent, nil
}
