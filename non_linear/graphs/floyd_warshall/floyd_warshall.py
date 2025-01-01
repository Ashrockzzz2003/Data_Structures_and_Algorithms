def adjacencyListToMatrix(adjList, vertices):
    matrix = [[float('inf')] * vertices for _ in range(vertices)]
    for i in range(vertices):
        matrix[i][i] = 0  # Distance to self is 0
    for u, edges in adjList.items():
        for v, weight in edges:
            matrix[u][v] = weight
    return matrix

def floydWarshall(vertices, graph):
    dist = [row[:] for row in graph]

    for k in range(vertices):
        for i in range(vertices):
            for j in range(vertices):
                if dist[i][k] + dist[k][j] < dist[i][j]:
                    dist[i][j] = dist[i][k] + dist[k][j]
    
    # Check for negative weight cycles
    for i in range(vertices):
        if dist[i][i] < 0:
            raise ValueError("Negative weight cycle detected.")

    return dist

# Main

adjList = {
    0: [(1, 3), (3, 7)],
    1: [(0, 8), (2, 2)],
    2: [(0, 5), (3, 1)],
    3: [(0, 2)]
}

vertices = len(adjList)

try:
    graph = adjacencyListToMatrix(adjList,vertices)

    result = floydWarshall(vertices, graph)

    print("Shortest distances between every pair of vertices:")
    for num, row in enumerate(result):
        print(f"Vertex {num}: {row}")

except ValueError as e:
    print(e)