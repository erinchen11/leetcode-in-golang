package happynumber

func isHappy(n int) bool {
	record := map[int]bool{}
	newNum := n

	for {
		if newNum == 1 {
			return true
		}
		// 如果 newNum不等於1,則去計算新的數
		newNum = getNew(newNum)
		// 檢查新數字是否存在record
		// 存在則傳回false
		if _, ok := record[newNum]; ok {
			return false
		} else {
			record[newNum] = true
		}
	}

}

func getNew(num int) (newNum int) {
	for num > 0 {
		//計算個位數
		newNum += (num % 10) * (num % 10)
		//換下一位數
		num /= 10
	}
	return newNum
}
