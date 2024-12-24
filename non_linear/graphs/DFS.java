

import java.util.ArrayList;

public class DFS {
    public static void dfs(int node, boolean vis[], ArrayList<ArrayList<Integer>> adj, ArrayList<Integer> ls) {
        vis[node] = true;
        ls.add(node);

        for (Integer it : adj.get(node)) {
            if (vis[it] == false) {
                dfs(it, vis, adj, ls);
            }
        }
    }

    public ArrayList<Integer> dfsOfGraph(int V, ArrayList<ArrayList<Integer>> adj) {
        boolean vis[] = new boolean[V + 1];
        vis[0] = true;
        ArrayList<Integer> ls = new ArrayList<>();
        dfs(0, vis, adj, ls);
        return ls;
    }
    public static void main(String[] args) {
        // Create a graph with 5 nodes (0 to 4)
        int V = 9;
        ArrayList<ArrayList<Integer>> adj = new ArrayList<>();

        // Initialize adjacency list
        for (int i = 0; i < V; i++) {
            adj.add(new ArrayList<>());
        }

        // Add edges
        adj.get(0).add(1);
        adj.get(0).add(2);
        adj.get(1).add(3);
        adj.get(1).add(4);
        adj.get(2).add(5);
        adj.get(3).add(6);
        adj.get(3).add(7);
        adj.get(4).add(8);

        // Perform DFS
        DFS graph = new DFS();
        ArrayList<Integer> result = graph.dfsOfGraph(V, adj);

        // Print the result
        System.out.println("DFS Traversal of the graph: " + result);
    }
}

