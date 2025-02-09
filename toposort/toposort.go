package toposort

import "fmt"

func SortDigraph(digraph map[string][]string) ([]string, error) {

	// count how many indegrees every node has
	indegree := map[string]int{}
	for fromNode, toNodes := range digraph {
		indegree[fromNode] += 0
		for _, toNode := range toNodes {
			indegree[toNode] += 1
		}
	}

	// start with empty results
	sorted := make([]string, 0, len(indegree))
	nodesWithoutInDegree := findOneIndegreeZero(indegree)

	// repeat while we can find a node with indegree 0
	for len(nodesWithoutInDegree) > 0 {
		node := nodesWithoutInDegree[len(nodesWithoutInDegree)-1]
		nodesWithoutInDegree = nodesWithoutInDegree[:len(nodesWithoutInDegree)-1]
		sorted = append(sorted, node)
		delete(indegree, node)
		for _, toNode := range digraph[node] {
			indegree[toNode] -= 1
			if indegree[toNode] == 0 {
				nodesWithoutInDegree = append(nodesWithoutInDegree, toNode)
			}
		}
	}

	if len(indegree) == 0 {
		return sorted, nil
	}

	return nil, fmt.Errorf("digraph has a cycle. no topological sort exists.")
}

func findOneIndegreeZero(indegree map[string]int) []string {
	for node, count := range indegree {
		if count == 0 {
			return []string{node}
		}
	}
	return nil
}
