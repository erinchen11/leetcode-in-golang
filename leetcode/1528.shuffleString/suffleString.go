package leetcode

func restoreString(s string, indices []int) string {
	result := make([]rune, len(s))
	for k, v := range s {
		result[indices[k]] = v
	}
	return string(result)
}