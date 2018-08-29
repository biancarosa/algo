package sorting

type BubbleSort struct{}

func (bs *BubbleSort) Sort(v []int) []int {
	for i, a := range v {
		for j, b := range v {
			if b > a {
				tmp := v[i]
				v[i] = v[j]
				v[j] = tmp
			}
		}
	}
	return v
}
