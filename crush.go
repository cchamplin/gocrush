package gocrush

import "fmt"

const (
	defaultMaxRetries = uint8(3)
	//Root represents default root tree.
	Root = uint8(0)
	//DataCenter represents default first hierarchy element
	DataCenter = uint8(1)
	//Rack represents default second hierarchy element
	Rack = uint8(2)
	//Node represents default third hierarchy element
	Node = uint8(3)
	//Disk represents default fourth hierarchy element
	Disk = uint8(4)
)

// NodeTypes represents mapping of
// various tree levels
type NodeTypes map[string]uint8

var defaultNodeTypes = NodeTypes{
	"root":       Root,
	"dataCenter": DataCenter,
	"rack":       Rack,
	"node":       Node,
	"disk":       Disk,
}

// Client is an instance of gocrush.
type Client struct {
	nodeTypes  NodeTypes
	maxRetries uint8
}

// Config holds Client configuration.
type Config struct {
	NodeTypes  NodeTypes
	MaxRetries uint8
}

var conf *Config

// New instantinates gocrush library.
// It will expect zero or exactly one Config struct,
// returns Client which holds all methods and internal config.
func New(options ...*Config) (*Client, error) {
	conf = &Config{
		NodeTypes:  defaultNodeTypes,
		MaxRetries: defaultMaxRetries,
	}
	c := Client{}
	if len(options) == 1 {
		opts := options[0]
		if opts.NodeTypes != nil {
			if err := c.SetNodeTypes(opts); err != nil {
				return nil, fmt.Errorf("could not set node types: %v", err)
			}
		} else {
			c.nodeTypes = conf.NodeTypes
		}
		if opts.MaxRetries != 0 {
			if err := c.SetMaxRetries(opts); err != nil {
				return nil, fmt.Errorf("could not set max retries: %v", err)
			}
		} else {
			c.maxRetries = conf.MaxRetries
		}
	}
	return &c, nil
}

// SetNodeTypes sets node groups.
// This setting may change hashing so it is not adviced to be used
// after first use of hashing.
func (c *Client) SetNodeTypes(opts *Config) error {
	if len(opts.NodeTypes) != 0 {
		c.nodeTypes = opts.NodeTypes
		return nil
	}
	return fmt.Errorf("wrong options provided to SetNodeTypes")
}

// SetMaxRetries sets how many times we will retry
// to get node we are looking for before giving up
func (c *Client) SetMaxRetries(opts *Config) error {
	if opts.MaxRetries > 0 {
		c.maxRetries = opts.MaxRetries
		return nil
	}
	return fmt.Errorf("MaxRetries should be uint8 type")
}

// Select returns a slice of nodes which are it's childs and
// been selected by given algorithm.
func (c *Client) Select(parent CNode, input int64, requestedNodesCount uint8, nodeGroup uint8) []CNode {
	var results []CNode
	for replica := uint8(0); replica < requestedNodesCount; replica++ {
		var totalFailures uint8
		var result CNode
		for retryDescent := true; retryDescent; {
			retryDescent = false
			var replicaFailures uint8
			bucket := parent
			for retryBucket := true; retryBucket; {
				retryBucket = false
				if replica == 0 {
					result = bucket.Select(input, int64(replica+totalFailures))
				} else {
					result = bucket.Select(input, int64(replica+replicaFailures))
				}
				switch {
				case result.GetGroup() != nodeGroup:
					bucket = result
					retryBucket = true
				case contains(results, result):
					totalFailures++
					replicaFailures++
					if replicaFailures >= c.maxRetries {
						retryDescent = true
						break
					}
					retryBucket = true
				case isDefunct(result):
					totalFailures++
					replicaFailures++
					retryDescent = true
				}
			}
		}
		results = append(results, result)
	}
	return results
}

func contains(nodes []CNode, n CNode) bool {
	i := n.GetID()
	for _, node := range nodes {
		if node.GetID() == i {
			return true
		}
	}
	return false
}

func isDefunct(n CNode) bool {
	if n.IsLeaf() {
		if n.IsFailed() || n.IsOverloaded() {
			return true
		}
	}
	return false
}
