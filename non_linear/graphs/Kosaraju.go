// Kosaraju's algorithm is used to find strongly connected components. It is applicable only for directed graphs
package main

import "fmt"

type Graph struct {
	adj  [][]int // Adjacency list
	adjT [][]int // Transposed graph adjacency list
	vstd []bool  // Visited array
}

// DFS for the first pass to populate the stack
func (g *Graph) dfs(node int, stack *[]int) {
	g.vstd[node] = true
	for index := 0; index < len(g.adj[node]); index++ {
		neighbor := g.adj[node][index]
		if !g.vstd[neighbor] {
			g.dfs(neighbor, stack)
		}
	}
	*stack = append(*stack, node)
}

// DFS for the second pass to collect SCC components
func (g *Graph) dfsTranspose(node int, component *[]int) {
	g.vstd[node] = true
	*component = append(*component, node)
	for index := 0; index < len(g.adjT[node]); index++ {
		neighbor := g.adjT[node][index]
		if !g.vstd[neighbor] {
			g.dfsTranspose(neighbor, component)
		}
	}
}

// Kosaraju's algorithm to count and print SCCs
func (g *Graph) kosaraju(V int) (int, [][]int) {
	// Step 1: First DFS to populate the stack
	stack := []int{}
	g.vstd = make([]bool, V)
	for i := 0; i < V; i++ {
		if !g.vstd[i] {
			g.dfs(i, &stack)
		}
	}

	// Step 2: Create the transpose graph
	g.adjT = make([][]int, V)
	for i := 0; i < V; i++ {
		for index := 0; index < len(g.adj[i]); index++ {
			neighbor := g.adj[i][index]
			g.adjT[neighbor] = append(g.adjT[neighbor], i)
		}
	}

	// Step 3: Perform DFS on the transposed graph in stack order
	for i := range g.vstd {
		g.vstd[i] = false
	}

	scc := 0
	var components [][]int

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !g.vstd[node] {
			scc++
			component := []int{}
			g.dfsTranspose(node, &component)
			components = append(components, component)
		}
	}

	return scc, components
}

func main() {
	var V, E int

	// User input for vertices and edges
	fmt.Print("Enter the number of vertices: ")
	fmt.Scan(&V)
	fmt.Print("Enter the number of edges: ")
	fmt.Scan(&E)

	// Initialize the graph
	graph := Graph{
		adj: make([][]int, V),
	}

	// User input for edges
	fmt.Println("Enter edges (format: from to):")
	for i := 0; i < E; i++ {
		var u, v int
		fmt.Printf("Edge %d: ", i+1)
		fmt.Scan(&u, &v)
		graph.adj[u] = append(graph.adj[u], v)
	}

	// Find SCCs
	count, components := graph.kosaraju(V)
	fmt.Printf("\nThe number of strongly connected components is: %d\n", count)

	fmt.Println("\nThe strongly connected components are:")
	for i, component := range components {
		fmt.Printf("Component %d: %v\n", i+1, component)
	}
}
