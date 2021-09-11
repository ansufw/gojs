package gojs

// added data with shift command

/*
func UnshiftInt(sliceArray []int, data int) {
	temp = append([]int{data}, *sliceArray...)
	sliceArray = temp
}
*/

func UnshiftInt(sliceArray *[]int, data int) {
	temp := *sliceArray
	*sliceArray = append([]int{data}, temp...)
}
