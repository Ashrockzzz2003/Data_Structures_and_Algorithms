package main

import "fmt"

type Graph struct {
	vertices int
	adjList  map[int][]int
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adjList:  make(map[int][]int),
	}
}

func (g *Graph) AddEdge(v, w int) {
	g.adjList[v] = append(g.adjList[v], w)
}

func (g *Graph) TarjanSCC() [][]int {
	index := 0
	indices := make([]int, g.vertices)
	lowLink := make([]int, g.vertices)
	onStack := make([]bool, g.vertices)
	stack := []int{}
	result := [][]int{}

	for i := 0; i < g.vertices; i++ {
		indices[i] = -1
	}

	var strongConnect func(v int)
	strongConnect = func(v int) {
		indices[v] = index
		lowLink[v] = index
		index++
		stack = append(stack, v)
		onStack[v] = true

		for _, w := range g.adjList[v] {
			if indices[w] == -1 {
				strongConnect(w)
				lowLink[v] = min(lowLink[v], lowLink[w])
			} else if onStack[w] {
				lowLink[v] = min(lowLink[v], indices[w])
			}
		}

		if lowLink[v] == indices[v] {
			scc := []int{}
			for {
				w := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				onStack[w] = false
				scc = append(scc, w)
				if w == v {
					break
				}
			}
			result = append(result, scc)
		}
	}

	for i := 0; i < g.vertices; i++ {
		if indices[i] == -1 {
			strongConnect(i)
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	g := NewGraph(8)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(5, 3)
	g.AddEdge(6, 4)
	g.AddEdge(6, 7)
	g.AddEdge(7, 6)

	fmt.Println("Strongly Connected Components:")
	sccs := g.TarjanSCC()
	for i, scc := range sccs {
		fmt.Printf("SCC %d: %v\n", i+1, scc)
	}
}
