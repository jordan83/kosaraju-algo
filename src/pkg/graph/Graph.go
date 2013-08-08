package graphs

type Vertex struct {
	Label int
	Explored bool
	neighbours []*Vertex
}

type Graph interface {
	GetVertex(index int) *Vertex
	GetNeighbours(v *Vertex) []*Vertex
	AddNeighbour(v *Vertex, n *Vertex)
	Reverse() Graph
	NumVertices() int
}

type AdjacencyList struct {
	vertices *[]Vertex
}

func NewGraph(numVertices int) Graph {
	vertices := make([]Vertex, numVertices)
	for i := 0; i < len(vertices); i++ {
		vertex := &vertices[i]
		vertex.Explored = false
		vertex.Label = i + 1
	}
	return AdjacencyList {
		&vertices,
	}
}

func (adjacencyList AdjacencyList) GetVertex(index int) *Vertex {
	tmp := *adjacencyList.vertices
	return &tmp[index]
}

func (adjacencyList AdjacencyList) GetNeighbours(v *Vertex) []*Vertex {
	return v.neighbours
}

func(adjacencyList AdjacencyList) AddNeighbour(v *Vertex, n *Vertex) {
	v.neighbours = append(v.neighbours, n)
}

func (adjacencyList AdjacencyList) Reverse() Graph {
	vertices := *adjacencyList.vertices
	graph := NewGraph(len(vertices))
	
	for i := 0; i < len(vertices); i++ {
		vertex := &vertices[i]
		tail := graph.GetVertex(i)
		
		for _,neighbour := range vertex.neighbours { 
			head := graph.GetVertex(neighbour.Label -1)
			head.neighbours = append(head.neighbours, tail) 
		}
	}
	return graph
}

func (adjacencyList AdjacencyList) NumVertices() int {
	return len(*adjacencyList.vertices)
}