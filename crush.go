package gocrush

import (
//	"log"
)

func Select(parent Node, input int64, count int, nodeType int, c Comparitor) []Node {
	var results []Node
	//if len(parent.Children) < count {
	//	panic("Asked for more node than are available")
	//}
	var rPrime = int64(0)
	for r := 1; r <= count; r++ {
		var failure = 0
		var loopbacks = 0
		var escape = false
		var retryOrigin = false
		var out Node
		for {
			retryOrigin = false
			var in = parent
			var skip = make(map[Node]bool)
			var retryNode = false
			for {
				retryNode = false
				rPrime = int64(r + failure)
				out = in.Select(input, rPrime)
				if out.GetType() != nodeType {
					in = out
					retryNode = true
				} else {
					if contains(results, out) {
						if !nodesAvailable(in, results, skip) {
							if loopbacks == 150 {
								escape = true
								break
							}
							loopbacks += 1
							retryOrigin = true
						} else {
							retryNode = true
						}
						failure += 1

					} else if c != nil && !c(out) {
						skip[out] = true
						if !nodesAvailable(in, results, skip) {
							if loopbacks == 150 {
								escape = true
								break
							}
							loopbacks += 1
							retryOrigin = true
						} else {
							retryNode = true
						}
						failure += 1
					} else if isDefunct(out) {
						failure++
						if loopbacks == 150 {
							escape = true
							break
						}
						loopbacks += 1
						retryOrigin = true
					} else {
						break
					}
				}
				if !retryNode {
					break
				}
			}
			if !retryOrigin {
				break
			}
		}
		if escape {
			continue
		}
		results = append(results, out)
	}
	return results
}

func nodesAvailable(parent Node, selected []Node, rejected map[Node]bool) bool {
	var children = parent.GetChildren()
	for _, child := range children {
		if !isDefunct(child) {
			if ok := contains(selected, child); !ok {
				if _, ok := rejected[child]; !ok {
					return true
				}
			}
		}
	}
	return false
}

func contains(s []Node, n Node) bool {
	for _, a := range s {
		if a == n {
			return true
		}
	}
	return false
}

func isDefunct(n Node) bool {
	if n.IsLeaf() && n.IsFailed() {
		return true
	}
	return false
}
