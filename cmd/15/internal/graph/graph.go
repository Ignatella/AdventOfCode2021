package graph

type Node struct {
	I, J     int
	Risk     int
	Previous *Node
}

type Graph [][]Node

func NewNode(i, j, risk int) Node {
	return Node{I: i, J: j, Risk: risk}
}

func (g Graph) ToNodeSet() NodeSet {
	var v = make(NodeSet, 0)

	for _, row := range g {
		for _, ver := range row {
			v = append(v, ver)
		}
	}

	return v
}

func (g Graph) GetNeighbors(n Node) NodeSet {
	neighbors := make(NodeSet, 0)

	if n.I > 0 {
		neighbors = append(neighbors, g[n.I-1][n.J])
	}

	if n.I < len(g)-1 {
		neighbors = append(neighbors, g[n.I+1][n.J])
	}

	if n.J > 0 {
		neighbors = append(neighbors, g[n.I][n.J-1])
	}

	if n.J < len(g[0])-1 {
		neighbors = append(neighbors, g[n.I][n.J+1])
	}

	return neighbors
}
