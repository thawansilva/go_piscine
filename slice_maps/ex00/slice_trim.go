package slice

func TrimEdges(s []int, n int) []int {
	slice_len := len(s)

	if n > slice_len || slice_len - n < n {
		return []int {};
	}

	return s[n:slice_len - n]
}
