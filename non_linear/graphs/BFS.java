

import java.util.ArrayList;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.Queue;

public class BFS {
    public ArrayList<Integer> bfsOfGraph(int vertex,ArrayList<ArrayList<Integer>> adjacent){
        ArrayList<Integer> bfs = new ArrayList<>();//the list stores the elements
        boolean visit[] = new boolean[vertex];//visited array is created
        Queue<Integer> q = new LinkedList<>();

        q.add(0);
        visit[0]=true;

        while(!q.isEmpty()){
            Integer node = q.poll();
            bfs.add(node);

            for(Integer it: adjacent.get(node)){
                if(visit[it]==false){
                    visit[it]=true;
                    q.add(it);
                }
            }
        }
        return bfs;
    }

    public static void main(String[] args) {
        // Number of vertices
        int vertex = 6;

        // Creating an adjacency list representation of the graph
        ArrayList<ArrayList<Integer>> adjacent = new ArrayList<>();

        // Initialize adjacency list for each vertex
        for (int i = 0; i < vertex; i++) {
            adjacent.add(new ArrayList<>());
        }

        // Adding edges to the graph
        adjacent.get(0).add(1);
        adjacent.get(0).add(2);
        adjacent.get(1).add(3);
        adjacent.get(1).add(4);
        adjacent.get(2).add(3);
        adjacent.get(3).add(4);
        adjacent.get(3).add(5);

        // Displaying the graph
//        System.out.println("Adjacency List of the Graph:");
//        for (int i = 0; i < vertex; i++) {
//            System.out.println(i + ": " + adjacent.get(i));
//        }

        // Create an object of BFS and call bfsOfGraph
        BFS bfsSolver = new BFS();
        ArrayList<Integer> bfsResult = bfsSolver.bfsOfGraph(vertex, adjacent);

        // Printing the BFS Traversal
        System.out.println("\nBFS Traversal of the graph:");
        for (Integer node : bfsResult) {
            System.out.print(node + " ");
        }
    }
}
