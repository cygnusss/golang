package main

import (
	"fmt"
)

type Graph struct {
	Val    int
	Edges  []*Graph
	Weight int
}

func (g *Graph) AddEdge(v, w int) Graph {
	node := Graph{v, []*Graph{}, w}
	g.Edges = append(g.Edges, &node)
	return node
}

func (g *Graph) ShowEdges() {
	for _, node := range g.Edges {
		fmt.Println(*node)
	}
}

func main() {
	graph := Graph{10, []*Graph{}, 6}
	edge1 := graph.AddEdge(6, 2)

	edge1.AddEdge(1, 1)
	edge1.AddEdge(2, 1)

	graph.ShowEdges()
}
