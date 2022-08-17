package solutions

func JumpingOnClouds(c []int32) int32 {
	// Write your code here
	var jumpCounter, l, i int32
	l = int32(len(c))
	for i = 0; i < l-1; {
		if i+2 < l && c[i+2] == 0 {
			jumpCounter++
			i += 2
			continue
		}
		if i+1 < l && c[i] == 0 {
			jumpCounter++
			i += 1
		}

	}

	return jumpCounter
}
