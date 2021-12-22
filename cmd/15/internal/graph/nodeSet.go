package graph

import (
	"math"
)

type NodeSet []Node

func (v NodeSet) GetNodeWithMinRisk(risks [][]int) Node {
	minI, minRisk := 0, math.MaxInt

	for i, ver := range v {
		if risks[ver.I][ver.J] < minRisk {
			minI = i
			minRisk = risks[ver.I][ver.J]
		}
	}

	return v[minI]
}

func (v NodeSet) Empty() bool {
	return len(v) == 0
}

func (v *NodeSet) Remove(n Node) {
	for i, ver := range *v {
		if ver == n {
			*v = append((*v)[:i], (*v)[i+1:]...)
		}
	}
}

func (v NodeSet) Contains(n Node) bool {
	for _, ver := range v {
		if ver == n {
			return true
		}
	}

	return false
}

func (v NodeSet) Intersect(w NodeSet) NodeSet {
	intersection := make(NodeSet, 0)

	for _, ver := range w {
		if v.Contains(ver) {
			intersection = append(intersection, ver)
		}
	}

	return intersection
}
