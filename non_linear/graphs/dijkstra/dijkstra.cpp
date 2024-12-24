//{ Driver Code Starts
#include<bits/stdc++.h>
using namespace std;

// } Driver Code Ends

#include<limits.h>

class Solution
{
public:
  //Function to find the shortest distance of all the vertices
  //from the source vertex S.

  int selectMin(vector<int> &distance, vector<bool>&visited) {

    int node;
    int minval = INT_MAX;

    for (int i = 0; i < distance.size(); i++) {
      if (!visited[i] && minval > distance[i]) {
        minval = distance[i];
        node = i;
      }
    }
    return node;
  }
  vector <int> dijkstra(int V, vector<vector<int>> adj[], int S)
  {
    vector<bool> visited(V, false);
    vector<int> distance(V, INT_MAX);

    distance[S] = 0;

    for (int i = 0; i < V - 1; i++) {

      int u = selectMin(distance, visited);
      visited[u] = true;

      for (auto i : adj[u]) {
        int neighbour = i[0];
        int weight = i[1];

        if (!visited[neighbour] && distance[u] + weight < distance[neighbour]) {
          distance[neighbour] = distance[u] + weight;
        }
      }
    }
    return distance;
  }
};


//{ Driver Code Starts.


int main()
{
  int t;
  cin >> t;
  while (t--) {
    int V, E;
    cin >> V >> E;
    vector<vector<int>> adj[V];
    int i = 0;
    while (i++ < E) {
      int u, v, w;
      cin >> u >> v >> w;
      vector<int> t1, t2;
      t1.push_back(v);
      t1.push_back(w);
      adj[u].push_back(t1);
      t2.push_back(u);
      t2.push_back(w);
      adj[v].push_back(t2);
    }
    int S;
    cin >> S;

    Solution obj;
    vector<int> res = obj.dijkstra(V, adj, S);

    for (int i = 0; i < V; i++)
      cout << res[i] << " ";
    cout << endl;
  }

  return 0;
}

// } Driver Code Ends
