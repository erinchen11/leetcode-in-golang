package betterrepeateddeletionalgo

func RemoveDuplicate(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	// 兩個pointer技巧
	// readPointer, writePointer
	// readPointer read all the elements to identify the duplicates.
	// writePointer keep track of the next position in the front
	// to write the next unique element we've found.

	// user the two pointer to remove the duplicate in-place
	//the first element shouldn't be touched, it's already in its correct place.
	writePointer := 1
	// go through each element in the Array.
	for readPointer := 1; readPointer < len(arr); readPointer++ {
		// If the current element we're reading is *difference*
		// to the previous element
		if arr[readPointer] != arr[readPointer-1] {
			// Copy it into the next position at the front,
			// tracked by writePointer
			arr[writePointer] = arr[readPointer]
			// And we need to now increment writePointer,
			// because the next element should be written one spcace over.
			writePointer++
		}

	}
	//this turns out to ve the correct length value.
	return writePointer

}
