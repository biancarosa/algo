package sorting

type number interface {
	float64 | int | float32 | uint
}
type BubbleSort[T number] struct{} // classe

func (bs *BubbleSort[T]) Sort(v []T) []T { // metodo
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
