package utils

func Intersect(a, b []uint) []uint {
	counter := make(map[uint]int)
	rlt := make([]uint, 0)
	for _, a := range a {
		counter[a]++
	}
	for _, b := range b {
		sz, _ := counter[b]
		if sz == 1 {
			rlt = append(rlt, b)
		}
	}
	return rlt
}
