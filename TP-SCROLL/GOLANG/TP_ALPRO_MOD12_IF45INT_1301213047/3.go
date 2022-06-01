package main

func findSeq(tab []int, v int) int {
	for idx, val := range tab {
		if val == v {
			return idx
		}
	}

	// default return val
	return -1
}
