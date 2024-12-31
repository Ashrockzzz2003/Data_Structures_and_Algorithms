import java.util.*;

public class GraphNode {
    /**
     * Recursive function to perform Depth-First Search (DFS).
     *
     * @param currentNode       The current node being visited.
     * @param visitedNodes      Array to track visited nodes.
     * @param adjacencyList     The adjacency list representing the graph.
     * @param dfsTraversalOrder The list to store the order of DFS traversal.
     */
    public static void performDFS(int currentNode, boolean[] visitedNodes, ArrayList<ArrayList<Integer>> adjacencyList, ArrayList<Integer> dfsTraversalOrder) {
        visitedNodes[currentNode] = true; // Mark the current node as visited
        dfsTraversalOrder.add(currentNode); // Add the current node to the traversal order

        // Recursively visit all unvisited neighbors of the current node
        for (Integer neighbor : adjacencyList.get(currentNode)) {
            if (!visitedNodes[neighbor]) {
                performDFS(neighbor, visitedNodes, adjacencyList, dfsTraversalOrder);
            }
        }
    }

    /**
     * Performs DFS traversal of a graph from the starting node (node 0).
     *
     * @param numberOfVertices  The number of vertices in the graph.
     * @param adjacencyList     The adjacency list representing the graph.
     * @return A list of nodes in the order they are visited during DFS.
     */
    public ArrayList<Integer> dfsTraversal(int numberOfVertices, ArrayList<ArrayList<Integer>> adjacencyList) {
        boolean[] visitedNodes = new boolean[numberOfVertices]; // To track visited nodes
        ArrayList<Integer> dfsTraversalOrder = new ArrayList<>(); // To store the traversal order

        // Start DFS from node 0
        performDFS(0, visitedNodes, adjacencyList, dfsTraversalOrder);
        return dfsTraversalOrder;
    }

    public static void main(String[] args) {
        // Define the number of vertices in the graph
        int numberOfVertices = 9;

        // Initialize the adjacency list to represent the graph
        ArrayList<ArrayList<Integer>> adjacencyList = new ArrayList<>();
        for (int i = 0; i < numberOfVertices; i++) {
            adjacencyList.add(new ArrayList<>());
        }

        // Add edges to the graph (directed edges)
        adjacencyList.get(0).add(1);
        adjacencyList.get(0).add(2);
        adjacencyList.get(1).add(3);
        adjacencyList.get(1).add(4);
        adjacencyList.get(2).add(5);
        adjacencyList.get(3).add(6);
        adjacencyList.get(3).add(7);
        adjacencyList.get(4).add(8);

        // Create a GraphNode object to perform DFS
        GraphNode graph = new GraphNode();
        ArrayList<Integer> dfsResult = graph.dfsTraversal(numberOfVertices, adjacencyList);

        // Print the result of DFS traversal
        System.out.println("DFS Traversal of the graph: " + dfsResult);
    }
}
