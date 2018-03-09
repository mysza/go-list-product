package fun

func LargestProduct(list []int) int {
	l, s := 0, 0
	for _, v := range list {
		if v > l {
			s = l
			l = v
		} else if v > s {
			s = v
		}
	}
	return l * s
}
