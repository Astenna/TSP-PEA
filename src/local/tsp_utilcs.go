package local

func CalculateCost(nodes []int, adjacencyMatrix [][]int) int {
	var result int
	last := nodes[0]
	for _, node := range nodes[1:] {
		result = result + adjacencyMatrix[last][node]
		last = node
	}
	result = result + adjacencyMatrix[nodes[len(nodes)-1]][nodes[0]]
	return result
}
