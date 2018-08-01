package main

import (
	"fmt"
	"github.com/Workiva/go-datastructures/graph"
)

//UFInterface exposes a interface that implements Union-Find methods
type UFInterface interface {
	New()
	GetGraph() *graph.SimpleGraph
	Union(p int, q int) *graph.SimpleGraph
	Connected(p int, q int) *graph.SimpleGraph
}

//UF is a struct that has a graph and implements all methods of a UFInterface
type UF struct {
	graph *graph.SimpleGraph
}

//New is a method that initializes a graph
func (uf *UF) New() {
	uf.graph = graph.NewSimpleGraph()
}

//GetGraph is a method that returns a graph
func (uf *UF) GetGraph() *graph.SimpleGraph {
	return uf.graph
}

//Union is a method that adds a edge between two points
func (uf *UF) Union(p int, q int) *graph.SimpleGraph {
	uf.graph.AddEdge(p, q)
	return uf.graph
}

//Connected is a method that verifies if a component is connected
func (uf *UF) Connected(p int, q int) *graph.SimpleGraph {
	// TODO Implement graph connectivity
	return uf.graph
}

func main() {
	fmt.Println("Dynamic Connectivity")
}
