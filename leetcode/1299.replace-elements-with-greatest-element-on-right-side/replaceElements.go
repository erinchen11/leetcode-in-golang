package replaceelementswithgreatestelementonrightside

func ReplaceElements(arr []int) []int {
	if len(arr) == 0 {
		return []int{-1}
	}

	x, y := -1, 0
	//從尾部開始遍歷
	for i := len(arr) - 1; i >= 0; i--{
		//當前的element存給y,後面與x比較
		y = arr[i]
		//更新當前值
		arr[i] = x
		x = max(x, y)
	}
	return arr




}

func max(a, b int) int {
	if a > b{
		return a
	}
	return b
}