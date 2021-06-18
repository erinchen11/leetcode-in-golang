
package leetcode

func DuplicateZero(arr []int) {

	for i := 0; i < len(arr); i++{
		//如果是0且後面有地方可以複寫
		//就是index沒有超出範圍
		if arr[i] == 0 && i+1 < len(arr){
			arr = append(arr[:i+1],arr[i: len(arr)-1]... )
			i++
		}
	}


}