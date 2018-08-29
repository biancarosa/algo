package sorting

type InsertionSort struct{}

func (bs *InsertionSort) Sort(v []int) []int {
	for i := range v {
		elem := v[i]
		j := i
		for j > 0 && elem < v[j-1] {
			v[j] = v[j-1]
			j = j - 1
		}
		v[j] = elem
	}
	return v
}
