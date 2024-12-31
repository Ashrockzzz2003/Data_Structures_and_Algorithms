import java.util.*;
public class GraphNode {
    /**
     * Performs Breadth-First Search (BFS) on a graph represented as an adjacency list.
     *
     * @param adjacencyList  The graph represented as an adjacency list where keys are node values
     *                       and values are lists of connected nodes (neighbors).
     * @param startingNode   The node from which BFS traversal starts.
     * @return A list of node values in BFS traversal order.
     */
    public static List<Integer> performBFS(Map<Integer, List<Integer>> adjacencyList, int startingNode) {
        List<Integer> bfsOrder = new ArrayList<>(); // To store nodes in BFS order
        Set<Integer> visitedNodes = new HashSet<>(); // To track nodes that have been visited
        Queue<Integer> nodesToVisit = new LinkedList<>(); // Queue to process nodes in BFS order

        // Initialize BFS with the starting node
        nodesToVisit.add(startingNode);
        visitedNodes.add(startingNode);

        // Process nodes until there are no more nodes to visit
        while (!nodesToVisit.isEmpty()) {
            int currentNode = nodesToVisit.poll(); // Get the next node to process
            bfsOrder.add(currentNode); // Add the current node to the BFS result

            // Explore all neighbors of the current node
            for (int neighborNode : adjacencyList.getOrDefault(currentNode, new ArrayList<>())) {
                // If the neighbor has not been visited, mark it as visited and add it to the queue
                if (!visitedNodes.contains(neighborNode)) {
                    visitedNodes.add(neighborNode);
                    nodesToVisit.add(neighborNode);
                }
            }
        }

        return bfsOrder;
    }

    public static void main(String[] args) {
        // Define the graph using an adjacency list representation
        // Each key represents a node, and the value is a list of its neighbors
        Map<Integer, List<Integer>> graphAdjacencyList = new HashMap<>();
        graphAdjacencyList.put(0, Arrays.asList(1, 2)); // Node 0 is connected to nodes 1 and 2
        graphAdjacencyList.put(1, Arrays.asList(3, 4)); // Node 1 is connected to nodes 3 and 4
        graphAdjacencyList.put(2, Arrays.asList(3));    // Node 2 is connected to node 3
        graphAdjacencyList.put(3, Arrays.asList(4, 5)); // Node 3 is connected to nodes 4 and 5
        graphAdjacencyList.put(4, new ArrayList<>());   // Node 4 has no neighbors
        graphAdjacencyList.put(5, new ArrayList<>());   // Node 5 has no neighbors

        // Perform BFS starting from node 0
        List<Integer> bfsTraversalResult = performBFS(graphAdjacencyList, 0);

        // Print the BFS Traversal
        System.out.println("BFS Traversal of the graph:");
        for (int nodeValue : bfsTraversalResult) {
            System.out.print(nodeValue + " ");
        }
    }
}
