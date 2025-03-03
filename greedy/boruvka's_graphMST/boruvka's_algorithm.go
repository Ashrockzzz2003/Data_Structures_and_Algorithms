package main

/*
Boruvka's Algorithm:
	1) Initialize all vertices as individual components.
	2) Initialize MST = [] empty.
	3) While there are more than one component do:
		For each component:
			i. Find closest weight edge that connect the component to another.
			ii. Add this edge to the MST if not already added.
	4) Return MST.
*/

/*
s - Source
d - Destination
w - Weight
v - Vertex
e - Edge
g - Graph
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Stores the information about one edge
type Edge struct {
	s, d, w int // Source, Destination, Weight
}

// Stores the vertices and edge list of the graph
type Graph struct {
	v int
	e []Edge
}

// Creating a graph
func create_graph(v int) *Graph {
	return &Graph{v: v, e: []Edge{}}
}

// Adding an edge to the graph
func (g *Graph) add_edge(s, d, w int) {
	g.e = append(g.e, Edge{s: s, d: d, w: w})
}

// Find the root of a given vertex (Used to identify the component)
func (g *Graph) find_root(component []int, v int) int {
	root := v
	for component[root] != root {
		root = component[root]
	}
	for v != root {
		parent := component[v]
		component[v] = root
		v = parent
	}
	return root
}

func (g *Graph) Boruvka_algo() []Edge {

	// Initialize all vertex as its own component
	component := make([]int, g.v)
	for i := 0; i < g.v; i++ {
		component[i] = i
	}

	// Initialize MST = [] empty
	mst := []Edge{}
	components := g.v

	// While there are more than one component do:
	for components > 1 {
		// For each component:
		shortest := make([]*Edge, g.v)

		// Find the shortest edge that connects the component to another
		for _, edge := range g.e {
			s_root := g.find_root(component, edge.s)
			d_root := g.find_root(component, edge.d)

			// If it is 2 different components
			if s_root != d_root {
				if shortest[s_root] == nil || shortest[s_root].w > edge.w {
					e := edge
					shortest[s_root] = &e
				}

				if shortest[d_root] == nil || shortest[d_root].w > edge.w {
					e := edge
					shortest[d_root] = &e
				}
			}
		}

		// Add this edge to the MST if not already added
		for i := 0; i < g.v; i++ {
			if shortest[i] != nil {
				s_root := g.find_root(component, shortest[i].s)
				d_root := g.find_root(component, shortest[i].d)
				edge := shortest[i]

				// Merging components and adding edge to MST
				if s_root != d_root {
					mst = append(mst, *edge)
					root1 := g.find_root(component, s_root)
					root2 := g.find_root(component, d_root)
					component[root1] = root2
					components--
				}
			}
		}
	}
	// Return MST
	return mst
}

func main() {
	read := bufio.NewReader(os.Stdin)

	// Input the number of vertices and edges
	fmt.Print("Enter number of vertices and edges: ")
	input, _ := read.ReadString('\n')
	input = strings.TrimSpace(input)
	v_e := strings.Split(input, " ")
	v, _ := strconv.Atoi(v_e[0])
	e, _ := strconv.Atoi(v_e[1])

	// Create the graph
	graph := create_graph(v)

	// Getting input for edges
	fmt.Println("Enter edge info:")
	for i := 0; i < e; i++ {
		edge, _ := read.ReadString('\n')
		edge = strings.TrimSpace(edge)
		edge_lis := strings.Split(edge, " ")
		s, _ := strconv.Atoi(edge_lis[0])
		d, _ := strconv.Atoi(edge_lis[1])
		w, _ := strconv.Atoi(edge_lis[2])

		graph.add_edge(s, d, w) // Inserting into graph
	}

	// Finding the MST
	mst := graph.Boruvka_algo()

	// Printing the MST
	fmt.Println("Edges in the Minimum Spanning Tree:")
	total_weight := 0
	for _, edge := range mst {
		fmt.Printf("%d,%d = %d\n", edge.s, edge.d, edge.w)
		total_weight += edge.w
	}
	fmt.Printf("Total weight: %d\n", total_weight)
}

/* Test case:
Enter number of vertices and edges: 9 14
Enter edge info:
0 1 4
1 2 8
2 3 7
3 4 9
4 5 10
5 6 2
6 7 1
7 0 8
1 7 11
3 5 14
7 8 7
8 6 2
2 8 2
2 5 4
Output:
Edges in the Minimum Spanning Tree:
0,1 = 4
2,8 = 2
2,3 = 7
3,4 = 9
5,6 = 2
6,7 = 1
8,6 = 2
1,2 = 8
Total weight: 35 */