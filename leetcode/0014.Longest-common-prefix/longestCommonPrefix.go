package leetcode

//先找出最短的字串當prefix去表叫

func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	res := ""

	for i, v := range strs[0] {
		c := byte(v)
		for _, s := range strs {
			if len(s) < i+1 || s[i] != c {
				return res
			}

		}
		res += string(c)
	}
	return res

}
