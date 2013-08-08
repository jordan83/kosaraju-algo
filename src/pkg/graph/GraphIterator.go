package graphs

import (
	"../intStack"
)

type GraphIterator interface {
	HasNext(graph Graph) bool
	Next(graph Graph) *Vertex
}

//-----------------------------------------------------------------------------

type NaiveIterator struct {
	curIndex int
}

func NewNaiveIterator(start int) GraphIterator {
	return &NaiveIterator {
		curIndex: start,
	}
}

func (i *NaiveIterator) HasNext(graph Graph) bool {
	return i.curIndex - 1 >= 0
}

func (i *NaiveIterator) Next(graph Graph) *Vertex {
	i.curIndex = i.curIndex - 1
	if i.curIndex >= 0  {
		return graph.GetVertex(i.curIndex)
	}
	return nil
}

//-----------------------------------------------------------------------------

type StackIterator struct {
	stack *intStack.Stack
}

func NewStackIterator(s intStack.Stack) StackIterator {
	return StackIterator {
		stack: &s,
	}
}

func (s StackIterator) HasNext(graph Graph) bool {
	return len(*(s.stack)) > 0
}

func (s StackIterator) Next(graph Graph) *Vertex {
	return graph.GetVertex(s.stack.Pop() -1)
}