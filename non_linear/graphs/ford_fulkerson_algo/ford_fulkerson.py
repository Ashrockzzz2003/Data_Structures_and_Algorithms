from collections import defaultdict

class Graph:
    def __init__(self, vertices):
        self.vertices = vertices
        # now by defaultdict we can automatically initialize the empty lists for the new keys that we addd
        self.graph = defaultdict(lambda: defaultdict(int))

    def add_edge(self, u, v, capacity):
        # used to add an edge to the graph by a capacity , u:source vertex ,v:destination vertex,capacity:capcity of the edge
        self.graph[u][v] = capacity

    def dfs(self, source, sink, parent, visited):
        # this returns true if there is a path from source to sink in a residual graph and it even fills parent[] to store its path

        # the currene vertex marked as visisted
        visited[source] = True

        # checking all the adhacent vertices of the dequeued vertex u
        for neigbour, capacity in self.graph[source].items():
            if not visited[neigbour] and capacity > 0:
                parent[neigbour] = source
                # if reached sink return true
                if neigbour == sink:
                    return True
                # lse recursion for neighbor
                if self.dfs(neigbour, sink, parent, visited):
                    return True

        return False

    def ford_fulkerson(self, source, sink):
        # returns max flow from ssource to sink

        # initialize variables needed
        parent = [-1] * self.vertices
        max_flow = 0

        # augment the flow till there is a path from source to sink
        while self.dfs(source, sink, parent, [False] * self.vertices):
            # to find min residual capacity of the edges
            path_flow = float("inf")
            s = sink
            while s != source:
                path_flow = min(path_flow, self.graph[parent[s]][s])
                s = parent[s]

            # Add path flow to overall flow
            max_flow += path_flow

            # for updating residual capacities of the edges and reverse edges
            v = sink
            while v != source:
                u = parent[v]
                self.graph[u][v] -= path_flow  # Forward edge
                self.graph[v][u] += path_flow  # Backward edge
                v = parent[v]

        return max_flow


def example_usage():
    # graph creating
    g = Graph(6)

    # adding the edges
    g.add_edge(0, 1, 16)
    g.add_edge(0, 2, 13)
    g.add_edge(1, 2, 10)
    g.add_edge(1, 3, 12)
    g.add_edge(2, 1, 4)
    g.add_edge(2, 4, 14)
    g.add_edge(3, 2, 9)
    g.add_edge(3, 5, 20)
    g.add_edge(4, 3, 7)
    g.add_edge(4, 5, 4)

    source = 0  # source vertex
    sink = 5  # sink vertex

    # printing
    max_flow = g.ford_fulkerson(source, sink)
    print("Maximum flow from vertex", source, "to vertex", sink, "is", max_flow)


if __name__ == "__main__":
    example_usage()
