package main

func SwapLastAndIndex(slice []int, index int) {
	if len(slice) > index {
		replaced := slice[len(slice)-1]
		slice[len(slice)-1] = slice[index]
		slice[index] = replaced
	}
}
