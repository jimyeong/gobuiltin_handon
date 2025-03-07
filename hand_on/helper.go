package hand_on

func reverse(s []int) []int {
	// swap in place, call by val
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func reverseWithArrayPointer[T any](arr *[]T) {
	length := len(*arr)
	for i := 0; i < length/2; i++ {
		(*arr)[i], (*arr)[length-1-i] = (*arr)[length-1-i], (*arr)[i]
	}
}

func rotate(slice []int) []int {
	size := len(slice)
	if size < 2 {
		return slice
	}
	slice = reverse(slice[:])
	slice = reverse(slice[1:])
	return slice
}

func rorateNth(slice []int, nth int) []int {
	for i := 0; i == nth; i++ {
		slice = rotate(slice[:])
	}
	return slice
}

func main() {
	data := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reverseWithArrayPointer(&data)

	// fmt.Println("%q\n", nonempty(data))
	// fmt.Println("%q\n", nonempty2(data))
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// In place techniques
func nonempty(strings []string) []string {
	var i int = 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	// not i+1 => the code ends with i++
	return strings[:i]
}

func removeIdx(slice []int, idx int) []int {
	copy(slice[idx:], slice[idx+1:])
	return slice[:len(slice)-1]
}

// remove item without preserving the order
func removeIdxWithoutOrder(slice []int, idx int) []int {
	slice[idx] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func delete_empty_strings(slc []string) []string {
	var index int = 0
	var last_char string = ""
	for i := 0; i < len(slc); i++ {
		if slc[i] != last_char {
			last_char = slc[i]
			slc[index] = slc[i]
			index++
		}
	}
	return slc[:index]
}
