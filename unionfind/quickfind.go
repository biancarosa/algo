package unionfind

type QuickFind interface {
	New(n int) []int
	Connected(p int, q int) bool
	Union(p int, q int) []int
}

type QuickFindImpl struct {
	id []int
}

func (qf *QuickFindImpl) New(n int) []int {
	id := make([]int, n)
	for i, _ := range id {
		id[i] = i
	}
	qf.id = id
	return id
}

func (qf *QuickFindImpl) Connected(p int, q int) bool {
	return qf.id[p] == qf.id[q]
}

func (qf *QuickFindImpl) Union(p int, q int) {
	pid := qf.id[p]
	qid := qf.id[q]
	for i, _ := range qf.id {
		if qf.id[i] == pid {
			qf.id[i] = qid
		}
	}
}
