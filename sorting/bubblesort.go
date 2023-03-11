package sorting

type BubbleSort struct{} // classe

func (bs *BubbleSort) Sort(v []int) []int { // metodo
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

func (bs *BubbleSort) SortFloat(v []float64) []float64 { // metodo
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
