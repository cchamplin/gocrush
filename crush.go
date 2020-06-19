package gocrush

const maxRetries uint8 = 3

// Select returns a node.
func Select(parent Node, input int64, requestedNodesCount uint8, nodeType int) []Node {
	var results []Node
	for replica := uint8(0); replica < requestedNodesCount; replica++ {
		var totalFailures uint8
		var result Node
		for retryDescent := true; retryDescent; retryDescent = false {
			var replicaFailures uint8
			bucket := parent
			for retryBucket := true; retryBucket; retryBucket = false {
				var selector uint8
				if replica == 0 {
					selector = replica + totalFailures
				} else {
					selector = replica + replicaFailures
				}
				result = bucket.Select(input, int64(selector))
				switch {
				case result.GetType() != nodeType:
					bucket = result
					retryBucket = true
				case contains(results, result):
					totalFailures++
					replicaFailures++
					if replicaFailures >= maxRetries {
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

func contains(nodes []Node, n Node) bool {
	i := n.GetID()
	for _, node := range nodes {
		if node.GetID() == i {
			return true
		}
	}
	return false
}

func isDefunct(n Node) bool {
	if n.IsLeaf() {
		if n.IsFailed() || n.IsOverloaded() {
			return true
		}
	}
	return false
}
