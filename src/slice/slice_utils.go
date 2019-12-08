package slice

// SwapLastAndIndex Swaps last and the element on the specified index
func SwapLastAndIndex(slice []int, index int) {
	if len(slice) > index {
		SwapOnIndexes(slice, index, len(slice)-1)
	}
}

func SwapOnIndexes(slice []int, index1 int, index2 int) {
	replaced := slice[index1]
	slice[index1] = slice[index2]
	slice[index2] = replaced
}
