package sliceExtensions

// SwapLastAndIndex Swaps last and the element on the specified index
func SwapLastAndIndex(slice []int, index int) {
	if len(slice) > index {
		SwapOnIndexes(slice, index, len(slice)-1)
	}
}

// SwapOnIndexes swaps elements between two specified indexes
func SwapOnIndexes(slice []int, index1 int, index2 int) {
	replaced := slice[index1]
	slice[index1] = slice[index2]
	slice[index2] = replaced
}

// CalculateCost calculated cost of path based on given adjacencyMatrix
func CalculateCost(adjacencyMatrix [][]int, path []int) int {
	var result int
	last := path[0]
	for _, node := range path[1:] {
		result = result + adjacencyMatrix[last][node]
		last = node
	}
	result = result + adjacencyMatrix[path[len(path)-1]][path[0]]
	return result
}
