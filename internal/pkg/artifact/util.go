package artifact

func intersect(orgA, orgB []uint) []uint {
	counter := make(map[uint]int)
	rlt := make([]uint, 0)
	for _, a := range orgA {
		counter[a]++
	}
	for _, b := range orgB {
		sz, _ := counter[b]
		if sz == 1 {
			rlt = append(rlt, b)
		}
	}
	return rlt
}
