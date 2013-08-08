package scc

import(
	"../graph"
	"../intStack"
)

func FindSccs(graph graphs.Graph) SccList {
	iterator := doFirstPass(graph)
	return doSecondPass(graph, iterator)
}

func doFirstPass(graph graphs.Graph) graphs.GraphIterator {
	graphRev := graph.Reverse()
	
	stack := intStack.NewStack()
	iterator := graphs.NewNaiveIterator(graphRev.NumVertices())
	dfs(graphRev, iterator, func(dfsRoot int, v *graphs.Vertex) {
		stack.Push(v.Label)
	})
	
	return graphs.NewStackIterator(*stack)
}

func doSecondPass(graph graphs.Graph, iterator graphs.GraphIterator) SccList {
	sccs := map[int]int{}
	dfs(graph, iterator, func(dfsRoot int, v *graphs.Vertex) {
		count, found := sccs[dfsRoot]
		if found {
			sccs[dfsRoot] = count + 1
		} else {
			sccs[dfsRoot] = 1
		}
	})
	
	return NewSccList(sccs)
}

func dfs(graph graphs.Graph, graphIterator graphs.GraphIterator, visitor func(int, *graphs.Vertex)) {
	var nodeDfs func(v *graphs.Vertex, dfsRoot int)
	nodeDfs = func(v *graphs.Vertex, dfsRoot int) {
		v.Explored = true;
		
		for _, neighbour := range graph.GetNeighbours(v) {
			if !neighbour.Explored {
				nodeDfs(neighbour, dfsRoot)
			}
		}
		
		visitor(dfsRoot, v)
	}

	for graphIterator.HasNext(graph) {
		vertex := graphIterator.Next(graph)
		if !vertex.Explored {
			nodeDfs(vertex, vertex.Label)
		}
	}
}