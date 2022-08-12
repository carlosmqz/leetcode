package solutions

/*
John Watson knows of an operation called a right circular rotation on an array of integers. One rotation operation moves the last array element to the first position and shifts all remaining elements right one. To test Sherlock's abilities, Watson provides Sherlock with an array of integers. Sherlock is to perform the rotation operation a number of times then determine the value of the element at a given position.

For each array, perform a number of right circular rotations and return the values of the elements at the given indices.

Example


Here is the number of rotations on , and holds the list of indices to report. First we perform the two rotations: Now return the values from the zero-based indices and as indicated in the array.



*/
func CircularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	var result []int32
	n := int32(len(a)) //3
	for _, m := range queries {
		// (3 - (2%3)+ 1)%3)
		//	3 - (2) + 1
		result = append(result, a[(n-(k%n)+m)%n])
	}
	return result
}
