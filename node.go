package gocrush

// Node represents CRUSH node.
type Node interface {
	GetChildren() []Node
	GetType() int
	GetWeight() int64
	GetID() string
	IsFailed() bool
	IsOverloaded() bool
	GetSelector() Selector
	SetSelector(Selector)
	GetParent() Node
	IsLeaf() bool
	Select(input int64, round int64) Node
}

// Overload may be used to skip given node manually.
// If function returns TRUE, than node will be disabled.
type Overload func(Node) bool

// CrushNode is a mock used by unit tests
type CrushNode struct {
	Selector Selector
}

// TestingNode is a mock used by unit tests
type TestingNode struct {
	Children []Node
	CrushNode
	Weight   int64
	Parent   Node
	Failed   bool
	Overloaded bool
	ID       string
	Type     int
}

// GetSelector is part of TestingNode
func (n CrushNode) GetSelector() Selector {
	return n.Selector
}

// SetSelector is part of TestingNode
func (n CrushNode) SetSelector(Selector Selector) {
	n.Selector = Selector
}

// Select is part of TestingNode
func (n CrushNode) Select(input int64, round int64) Node {
	return n.GetSelector().Select(input, round)
}

// IsFailed return true if node is marked as failed.
func (n TestingNode) IsFailed() bool {
	return n.Failed
}

// IsOverloaded returns true if node is marked as overloaded.
func (n TestingNode) IsOverloaded() bool {
	return n.Overloaded
}

// IsLeaf is part of TestingNode
func (n TestingNode) IsLeaf() bool {
	return len(n.Children) == 0
}

// GetParent is part of TestingNode
func (n TestingNode) GetParent() Node {
	return n.Parent
}

// GetID is part of TestingNode
func (n TestingNode) GetID() string {
	return n.ID
}

// GetWeight is part of TestingNode
func (n TestingNode) GetWeight() int64 {
	return n.Weight
}

// GetType is part of TestingNode
func (n TestingNode) GetType() int {
	return n.Type
}

// GetChildren is part of TestingNode
func (n TestingNode) GetChildren() []Node {
	return n.Children
}

// TestCompare is part of TestingNode
func TestCompare(n Node) bool {
	tNode, ok := n.(TestingNode)
	if ok {
		return len(tNode.Children) > 0
	}
	return false
}
