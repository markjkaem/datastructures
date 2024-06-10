package datastructures

import "fmt"

// Graph represents the graph
type Graph struct {
	nodes map[string]map[string]int
}

// NewGraph creates a new graph
func NewGraph() *Graph {
	return &Graph{nodes: make(map[string]map[string]int)}
}

// AddNode adds a new node to the graph
func (g *Graph) AddNode(name string) {
	if _, exists := g.nodes[name]; !exists {
		g.nodes[name] = make(map[string]int)
	}
}

// AddEdge adds a new edge (connection) between two nodes
func (g *Graph) AddEdge(from, to string, weight int) {
	g.AddNode(from)
	g.AddNode(to)
	g.nodes[from][to] = weight
}

// PrintGraph prints the graph
func (g *Graph) PrintGraph() {
	for node, edges := range g.nodes {
		fmt.Printf("%s -> ", node)
		for edge, weight := range edges {
			fmt.Printf("(%s: %d) ", edge, weight)
		}
		fmt.Println()
	}
}

func MainGraph() {
	graph := NewGraph()

	// Add nodes (cities) to the graph
	graph.AddNode("Amsterdam")
	graph.AddNode("Bahrein")
	graph.AddNode("Cadmandu")

	// Add edges (roads) between cities with distances
	graph.AddEdge("Amsterdam", "Bahrein", 20)
	graph.AddEdge("Amsterdam", "Cadmandu", 20)
	graph.AddEdge("Bahrein", "Cadmandu", 15)
	graph.AddEdge("Bahrein", "Amsterdam", 20)
	graph.AddEdge("Cadmandu", "Amsterdam", 20)
	graph.AddEdge("Cadmandu", "Bahrein", 15)

	// Print the graph
	graph.PrintGraph()
}
