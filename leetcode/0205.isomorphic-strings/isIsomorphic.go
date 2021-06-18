package isomorphicstrings

func isIsomorphic(s string, t string) bool {
	// if len(s) == 0 {
	// 	return true
	// }

	// //建立兩個map
	// s2t := make(map[byte]byte)
	// t2s := make(map[byte]byte)
	// //建立雙映射關係
	// for i,_ := range s {
	// 	x, y := s[i], t[i]
	// 	if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
	// 		return false
	// 	}
	// 	s2t[x] = y
	// 	t2s[y] = x
	// }
	// return true
	sByte := []byte(s)
	tByte := []byte(t)
	if (s == "" && t != "") || (len(sByte) != len(tByte)) {
		return false
	}

	s2t := map[byte]byte{}
	t2s := map[byte]byte{}
	// 遍歷 t 找出s的映射

	for index, b := range tByte {
		//如果t2s map中沒有tByte的該字元, 則存入該字元以及對應的s字元
		if _, ok := t2s[b]; !ok {
			// 如果 s2t map 中沒有 sByte的該字元, 則存入該字元以及對應的t字元
			if _, ok := s2t[sByte[index]]; !ok {
				t2s[b] = sByte[index]
				//同時 s2t map也存入 s的該字元以及對應的t字元
				s2t[sByte[index]] = b
			} else {
				//如果在s2t map找不到對應的t字元則回傳false
				if s2t[sByte[index]] != b {
					return false
				}
			}
		} else {
			//如果在 t2s map 找不到對應的s字元則回傳false
			if t2s[b] != sByte[index] {
				return false
			}
		}

	}
	return true

}
