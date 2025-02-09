package toposort

import (
	"fmt"
)

func SortDigraph(digraph map[string][]string) ([]string, error) {

	// figure out how many indegrees every node has
	indegree := map[string]int{}
	for fromNode, toNodes := range digraph {
		indegree[fromNode] += 0
		for _, toNode := range toNodes {
			indegree[toNode] += 1
		}
	}

	// start with empty results
	sorted := []string{}

	// repeat until there are no more nodes
Next:
	for len(indegree) > 0 {
		for node, count := range indegree {
			if count == 0 {
				sorted = append(sorted, node)
				delete(indegree, node)
				for _, toNode := range digraph[node] {
					indegree[toNode] -= 1
				}
				continue Next
			}
		}
		// there are no nodes with indegree 0, therefore there is a cycle
		return nil, fmt.Errorf("digraph has a cycle. no topological sort exists.")
	}

	return sorted, nil
}
