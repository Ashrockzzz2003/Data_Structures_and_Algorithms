# Priority Queue Implementation
class PriorityQueue:
    def __init__(self):
        self.queue = []
        self.size = 0

    def parent(self, i):
        return (i - 1) // 2

    def minHeapify(self, i):
        left = 2 * i + 1
        right = 2 * i + 2
        smallest = i

        if left < self.size and self.queue[left][0] < self.queue[smallest][0]:
            smallest = left
        if right < self.size and self.queue[right][0] < self.queue[smallest][0]:
            smallest = right

        if smallest != i:
            self.queue[i], self.queue[smallest] = self.queue[smallest], self.queue[i]
            self.minHeapify(smallest)

    def insert(self, key, value):
        self.queue.append((key, value))
        self.size += 1
        i = self.size - 1

        while i > 0 and self.queue[self.parent(i)][0] > self.queue[i][0]:
            self.queue[i], self.queue[self.parent(i)] = self.queue[self.parent(i)], self.queue[i]
            i = self.parent(i)

    def extractMin(self):
        if self.size < 1:
            return None
        root = self.queue[0]
        self.queue[0] = self.queue[self.size - 1]
        self.size -= 1
        self.queue.pop()
        self.minHeapify(0)
        return root

    def buildMinHeap(self, elements):
        self.queue = elements
        self.size = len(elements)
        for i in range(self.size // 2 - 1, -1, -1):
            self.minHeapify(i)


def dijkstra(graph, start):
    distances = {vertex: float('inf') for vertex in graph}
    distances[start] = 0

    pq = PriorityQueue()
    pq.buildMinHeap([(distances[vertex], vertex) for vertex in graph])

    while pq.size > 0:
        _, u = pq.extractMin()
        for v, weight in graph[u]:
            if distances[v] > distances[u] + weight:
                distances[v] = distances[u] + weight
                pq.insert(distances[v], v)

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
