package main

type UFInterface interface {
	New()
	GetGraph() Graph
	Union(p int, q int) Graph
	Connected(p int, q int) Graph
}

type UF struct {
	graph Graph
}

type Graph struct {
}

func main() {

}
