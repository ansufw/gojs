package gojs

// added data with shift command

func UnshiftInt(sliceArray []int, data int) {
	slicePointer := &sliceArray
	temp := append([]int{data}, sliceArray...)
	*slicePointer = temp
}
