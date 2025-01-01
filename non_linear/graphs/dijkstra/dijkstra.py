# Priority Queue Implementation
import heapq

def dijkstra(graph, start):
    distances = {vertex: float('inf') for vertex in graph}
    distances[start] = 0

    pq = []
    heapq.heappush(pq,(0,start))

    while pq:
        curr_dist, u = heapq.heappop(pq)

        for v, weight in graph[u]:
            if distances[v] > curr_dist + weight:
                distances[v] = curr_dist + weight
                heapq.heappush(pq, (distances[v], v))

    # Check for disconnected vertices
    for vertex in distances:
        if distances[vertex] == float('inf'):
            distances[vertex] = "Inf"

    return distances

graph = {
    "LAX": [("SFO", 337), ("DFW", 1235), ("MIA", 2342)],
    "SFO": [("LAX", 337), ("DFW", 1464), ("ORD", 1846), ("BOS", 2704)],
    "DFW": [("LAX", 1235), ("SFO", 1464), ("ORD", 802), ("MIA", 1121), ("JFK", 1391)],
    "MIA": [("LAX", 2342), ("DFW", 1121), ("JFK", 1090), ("BWI", 946), ("BOS", 1258)],
    "ORD": [("SFO", 1846), ("DFW", 802), ("JFK", 740), ("BWI", 621), ("BOS", 867), ("PVD", 849)],
    "BWI": [("MIA", 946), ("ORD", 621), ("JFK", 184)],
    "JFK": [("DFW", 1391), ("MIA", 1090), ("ORD", 740), ("BWI", 184), ("BOS", 187), ("PVD", 144)],
    "BOS": [("SFO", 2704), ("MIA", 1258), ("ORD", 867), ("JFK", 187)],
    "PVD": [("ORD", 849), ("JFK", 144)]
}

print("Dijkstra's shortest paths:")
print(dijkstra(graph, "BWI"))
