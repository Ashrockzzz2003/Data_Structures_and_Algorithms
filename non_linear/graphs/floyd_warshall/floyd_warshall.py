def adjacencyListToMatrix(adjList, vertices):
    matrix = [[float('inf')] * vertices for _ in range(vertices)]
    for i in range(vertices):
        matrix[i][i] = 0  # Distance to self is 0
    for u in adjList:
        for v, weight in adjList[u]:
            matrix[u][v] = weight
    return matrix


def floydWarshall(vertices, graph):
    dist = [[graph[i][j] for j in range(vertices)] for i in range(vertices)]

    for k in range(vertices):
        for i in range(vertices):
            for j in range(vertices):
                if dist[i][k] + dist[k][j] < dist[i][j]:
                    dist[i][j] = dist[i][k] + dist[k][j]

    return dist


# Main

adjList = {
    0: [(1, 3), (3, 7)],
    1: [(0, 8), (2, 2)],
    2: [(0, 5), (3, 1)],
    3: [(0, 2)]
}

vertices = 4

graph = adjacencyListToMatrix(adjList, vertices)

result = floydWarshall(vertices, graph)

print("Shortest distances between every pair of vertices:")
for num,row in enumerate(result):
    print(num ,":",row)
