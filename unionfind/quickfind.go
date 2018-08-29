package unionfind

type QuickFind struct {
	id []int
}

func (qf *QuickFind) New(n int) []int {
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}
	qf.id = id
	return id
}

func (qf *QuickFind) Connected(p int, q int) bool {
	return qf.id[p] == qf.id[q]
}

func (qf *QuickFind) Union(p int, q int) {
	pid := qf.id[p]
	qid := qf.id[q]
	for i := range qf.id {
		if qf.id[i] == pid {
			qf.id[i] = qid
		}
	}
}
