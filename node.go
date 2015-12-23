package gocrush

type Node interface {
	GetChildren() []Node
	GetType() int
	GetWeight() int64
	GetId() string
	IsFailed() bool
	GetSelector() Selector
	SetSelector(Selector)
	GetParent() Node
	IsLeaf() bool
	Select(input int64, round int64) Node
}

type Comparitor func(Node) bool

type CrushNode struct {
	Selector Selector
}

type TestingNode struct {
	Children []Node
	CrushNode
	Weight int64
	Parent Node
	Failed bool
	Id     string
	Type   int
}

func (n CrushNode) GetSelector() Selector {
	return n.Selector
}

func (n CrushNode) SetSelector(Selector Selector) {
	n.Selector = Selector
}

func (n CrushNode) Select(input int64, round int64) Node {
	return n.GetSelector().Select(input, round)
}

func (n TestingNode) IsFailed() bool {
	return n.Failed
}

func (n TestingNode) IsLeaf() bool {
	return len(n.Children) == 0
}

func (n TestingNode) GetParent() Node {
	return n.Parent
}

func (n TestingNode) GetId() string {
	return n.Id
}

func (n TestingNode) GetWeight() int64 {
	return n.Weight
}

func (n TestingNode) GetType() int {
	return n.Type
}

func (n TestingNode) GetChildren() []Node {
	return n.Children
}

func TestCompare(n Node) bool {
	tNode, ok := n.(TestingNode)
	if ok {
		return len(tNode.Children) > 0
	}
	return false
}
