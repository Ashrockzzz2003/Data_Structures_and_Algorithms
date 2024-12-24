def bellmanFord(vertices, edges, source):
    distances = [float('inf')] * vertices
    distances[source] = 0

    for _ in range(vertices - 1):
        for u, v, weight in edges:
            if distances[u] != float('inf') and distances[v] > distances[u] + weight:
                distances[v] = distances[u] + weight

    for u, v, weight in edges:
        if distances[u] != float('inf') and distances[v] > distances[u] + weight:
            print("Negative weight cycle detected.")
            return [-1]

    return distances

def adjacenyListToEdges(adj_list):
    edges = []
    for u in adj_list:
        for v, weight in adj_list[u]:
            edges.append((u, v, weight))
    return edges


# Main

adjList = {
    0: [(1, -1), (2, 4)],
    1: [(2, 3), (3, 2), (4, 2)],
    2: [],
    3: [(2, 5), (1, 1)],
    4: [(3, -3)]
}

vertices = len(adjList)

# Convert adjacency list to edge list
edges = adjacenyListToEdges(adjList)

source = 0

result = bellmanFord(vertices, edges, source)

if result != [-1]:
    print("Shortest distances from source: " ,source)
    print(result)
