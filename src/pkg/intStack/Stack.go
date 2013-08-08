package intStack

type Stack []int

func (s *Stack) Push(v int) {
    *s = append(*s, v)
}

func (s *Stack) Pop() int {
    ret := (*s)[len(*s)-1]
    *s = (*s)[0:len(*s)-1]
    return ret
}

func NewStack() *Stack {
	var s Stack
	return &s
}