package scc

type SccList struct {
	Items []Scc
}

type Scc struct {
	Leader int
	Size int
}

func NewSccList(sccMap map[int]int) SccList {
	sccs := []Scc{}
	for k, v := range sccMap {
		sccs = append(sccs, Scc {k, v})
	}
	return SccList{sccs}
}

func (sccs SccList) Len() int {
	return len(sccs.Items)
}

func (sccs SccList) Less(i, j int) bool {
	return sccs.Items[i].Size > sccs.Items[j].Size
}

func (sccs SccList) Swap(i, j int) {
	tmp := sccs.Items[i]
	sccs.Items[i] = sccs.Items[j]
	sccs.Items[j] = tmp
}