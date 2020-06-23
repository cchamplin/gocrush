package gocrush

import "fmt"

// CNode represents CRUSH node
type CNode interface {
	GetChildrens() []CNode
	GetGroup() uint8
	GetWeight() int64
	GetID() string
	IsFailed() bool
	IsOverloaded() bool
	GetSelector() Selector
	SetSelector(Selector)
	GetParent() CNode
	IsLeaf() bool
	Select(input int64, round int64) CNode
}

// Selector interface is implemented by
// hashing algorithms.
type Selector interface {
	Select(input int64, round int64) CNode
}

// Overload may be used to skip given node manually.
// If function returns TRUE, than node will be disabled.
type Overload func(CNode) bool

// CrushNode represents single crush node and
// implements CNode interface
type CrushNode struct {
	id         string
	group      uint8
	parent     CNode
	childrens  []CNode
	selector   Selector
	weight     int64
	failed     bool
	overloaded bool
}

// CNodeConfig defines CrushNode configuration.
type CNodeConfig struct {
	ID         string // required for new node
	Group      uint8  // required for new node
	Parent     CNode
	Childrens  []CNode
	Selector   Selector
	Weight     int64
	Failed     bool
	Overloaded bool
}

var cNodeConf *CNodeConfig

// NewNode creates node entity.
// It will expect zero or exactly one Config struct,
// returns CNode which holds all methods and internal config.
func NewNode(options ...*CNodeConfig) (*CrushNode, error) {
	cNodeConf = &CNodeConfig{}
	c := CrushNode{}
	if len(options) == 1 {
		opts := options[0]
		if err := c.SetID(opts.ID); err != nil {
			return nil, fmt.Errorf("could not set ID for node: %v", err)
		}
		if err := c.SetGroup(opts.Group); err != nil {
			return nil, fmt.Errorf("could not set group: %v", err)
		}
		if len(opts.Childrens) != 0 {
			if err := c.AddChildrens(opts.Childrens); err != nil {
				return nil, fmt.Errorf("could not add childs: %v", err)
			}
		}
		if opts.Weight != 0 {
			if err := c.SetWeight(opts.Weight); err != nil {
				return nil, fmt.Errorf("could node set weight: %v", err)
			}
		}
		if opts.Parent != nil {
			if err := c.SetParent(opts.Parent); err != nil {
				return nil, fmt.Errorf("could node set parent: %v", err)
			}
		}

	}
	return &c, nil
}

// GetID returns node ID.
func (c *CrushNode) GetID() string {
	return c.id
}

// SetID sets node ID. This is required field for each node.
func (c *CrushNode) SetID(ID string) error {
	if ID == "" {
		return fmt.Errorf("ID cannot be empty")
	}
	c.id = ID
	return nil
}

// GetGroup returns horizontal group of the node,
// eg. DataCenter, Disk if default settings are used.
func (c *CrushNode) GetGroup() uint8 {
	return c.group
}

// SetGroup sets horizontal group of the node,
// eg. DataCenter, Disk if default settings are used.
func (c *CrushNode) SetGroup(opt uint8) error {
	c.group = opt
	return nil
}

// GetParent returns parent object of curent node.
func (c *CrushNode) GetParent() CNode {
	return c.parent
}

// SetParent sets parent node.
func (c *CrushNode) SetParent(opt CNode) error {
	if c.GetGroup() == 0 {
		return fmt.Errorf("root node cannot have parent")
	}
	if opt.GetID() == c.GetID() {
		return fmt.Errorf("cannot add myself as a parent")
	}
	c.parent = opt
	return nil
}

// GetChildrens returns slice containing childrens of
// selected node
func (c *CrushNode) GetChildrens() []CNode {
	return c.childrens
}

// AddChildren sets node groups.
// This setting may change hashing so it is not adviced to be used
// after first use of hashing.
func (c *CrushNode) AddChildren(opt CNode) error {
	if opt.GetID() == c.GetID() {
		return fmt.Errorf("cannot add myself as child")
	}
	if c.parent != nil {
		if opt.GetID() == c.parent.GetID() {
			return fmt.Errorf("cannot add parent as a child")
		}
	}
	for _, childs := range c.childrens {
		if opt.GetID() == childs.GetID() {
			return fmt.Errorf("this child node already exist")
		}
	}
	c.childrens = append(c.childrens, opt)
	return nil
}

// AddChildrens adds multiple childrens.
func (c *CrushNode) AddChildrens(opt []CNode) error {
	for _, child := range opt {
		if err := c.AddChildren(child); err != nil {
			return fmt.Errorf("could not add child: %v", err)
		}
	}
	return nil
}

// GetSelector returns selector for given node
func (c *CrushNode) GetSelector() Selector {
	return c.selector
}

// SetSelector changes selector for given CrushNode
func (c *CrushNode) SetSelector(s Selector) {
	c.selector = s
}

// Select implements selector interface and
// returns result from selected algorithm for
// selected criteria.
func (c *CrushNode) Select(input int64, round int64) CNode {
	return c.GetSelector().Select(input, round)
}

// GetWeight returns weight of the node.
func (c *CrushNode) GetWeight() int64 {
	return c.weight
}

// SetWeight sets weight for a node.
func (c *CrushNode) SetWeight(opt int64) error {
	c.weight = opt
	return nil
}

// IsFailed returns true if node is marked as failed.
func (c *CrushNode) IsFailed() bool {
	return c.failed
}

// SetFailed marks node as failed.
func (c *CrushNode) SetFailed() error {
	c.failed = true
	return nil
}

// SetHealthy marks node as healthy - failed false.
func (c *CrushNode) SetHealthy() error {
	c.failed = false
	return nil
}

// IsOverloaded returns true if node is marked as overloaded.
func (c *CrushNode) IsOverloaded() bool {
	return c.overloaded
}

// SetOverloaded marks node as overloaded.
func (c *CrushNode) SetOverloaded() error {
	c.overloaded = true
	return nil
}

// UnsetOverloaded marks node as not overloaded.
func (c *CrushNode) UnsetOverloaded() error {
	c.overloaded = false
	return nil
}

// IsLeaf returns true if node is a leaf.
// (doesn't have any children)
func (c *CrushNode) IsLeaf() bool {
	return len(c.childrens) == 0
}

// IsRoot returns true if node is a Root.
// (doesn't have any parents)
func (c *CrushNode) IsRoot() bool {
	if c.parent != nil {
		return false
	}
	return true
}
