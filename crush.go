package gocrush

const maxRetries uint8 = 150

// Select return a node. Main algorithm implementation
func Select(parent Node, input int64, requestedNodesCount uint8, nodeType int, comp Comparator) []Node {
	var results []Node
	for r := uint8(1); r <= requestedNodesCount; r++ {
		var retries, failure uint8
		var escape bool
		var child Node
		var skip = make(map[Node]bool)
		retry := true
		for retry {
			child = parent.Select(input, int64(r+failure))
			if child.GetType() != nodeType {
				parent = child
			}
			switch {
			case contains(results, child):
				if !nodesAvailable(parent, results, skip) {
					if retries >= maxRetries {
						escape = true
						break
					}
					retries++
				}
				failure++
			case comp != nil && !comp(child):
				skip[child] = true
				if !nodesAvailable(parent, results, skip) {
					if retries >= maxRetries {
						escape = true
						break
					}
					retries++
				}
				failure++
			case isDefunct(child):
				failure++
				if retries >= maxRetries {
					escape = true
					break
				}
				retries++
			default:
				retry = false
			}
		}
		if !escape {
			results = append(results, child)
		}
	}
	return results
}

func nodesAvailable(parent Node, selected []Node, rejected map[Node]bool) bool {
	for _, child := range parent.GetChildren() {
		if !isDefunct(child) && !contains(selected, child) && !rejected[child] {
			return true
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
