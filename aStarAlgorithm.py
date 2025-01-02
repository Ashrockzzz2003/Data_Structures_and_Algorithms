# Defining the graph as a dictionary where each node has a list of connected nodes with weights
graphNodes = {
    'A': [('B', 6), ('F', 3)],
    'B': [('C', 3), ('D', 2)],
    'C': [('D', 1), ('E', 5)],
    'D': [('C', 1), ('E', 8)],
    'E': [('I', 5), ('J', 5)],
    'F': [('G', 1), ('H', 7)],
    'G': [('I', 3), ('K', 2)],
    'H': [('I', 2), ('L', 4)],
    'I': [('E', 5), ('J', 3), ('M', 2)],
    'J': [('N', 6)],
    'K': [('M', 4), ('O', 2)],
    'L': [('M', 3), ('P', 5)],
    'M': [('N', 1), ('Q', 4)],
    'N': [('R', 6)],
    'O': [('Q', 3), ('S', 4)],
    'P': [('R', 2), ('T', 6)],
    'Q': [('U', 5)],
    'R': [('V', 7)],
    'S': [('U', 4), ('W', 3)],
    'T': [('V', 3), ('X', 5)],
    'U': [('Y', 2)],
    'V': [('Z', 4)],
    'W': [('Y', 3)],
    'X': [('Z', 6)],
    'Y': [('Z', 1)],
    'Z': []
}

# Function to fetch the neighbors of a given node
def get_neighbors(v):
    if v in graphNodes:
        return graphNodes[v]
    else:
        return None

# Heuristic function: This function has a value/cost assigned to each node. Usually this is desgined to calculate the cost of the route designed
def heuristic(n):
    heuristic_distance = {
        'A': 25, 'B': 24, 'C': 23, 'D': 22, 'E': 21, 'F': 20,
        'G': 19, 'H': 18, 'I': 17, 'J': 16, 'K': 15, 'L': 14,
        'M': 13, 'N': 12, 'O': 11, 'P': 10, 'Q': 9, 'R': 8,
        'S': 7, 'T': 6, 'U': 5, 'V': 4, 'W': 3, 'X': 2, 'Y': 1, 'Z': 0
    }
    return heuristic_distance[n]

# Implementation of the A* Algorithm
def aStarAlgorithm(start_node, stop_node):
    # Initialize the open set with the start node
    open_set = set(start_node)
    # Initialize the closed set as empty
    closed_set = set()
    # Dictionary to store the cost to reach each node
    g = {}
    # Dictionary to store the parent of each node for remaking the path 
    parents = {}
    #The staring node always has the staring cost to be 0
    g[start_node] = 0
    # The parent of the start node is itself
    parents[start_node] = start_node

    # Loop until the open set is empty or the destination is found
    while len(open_set) > 0:
        n = None
        # Select the node with the lowest f(n) value
        for v in open_set:
            if n == None or g[v] + heuristic(v) < g[n] + heuristic(n):
                n = v

        # If the destination node is reached or there are no neighbors
        if n == stop_node or graphNodes[n] == None:
            pass
        else:
            # Explore the neighbors of the current node
            for (m, weight) in get_neighbors(n):
                # If the neighbor is not in open or closed set, add it to the open set
                if m not in open_set and m not in closed_set:
                    open_set.add(m)
                    parents[m] = n
                    g[m] = g[n] + weight
                # If the new path is shorter, update the cost and parent
                else:
                    if g[m] > g[n] + weight:
                        g[m] = g[n] + weight
                        parents[m] = n
                        if m in closed_set:
                            closed_set.remove(m)
                            open_set.add(m)

        #If there is not path, the required mesasage is printed
        if n == None:
            print('Path does not exist!')
            return None

        # If the destination is reached, remake and return the path
        if n == stop_node:
            path = []
            while parents[n] != n:
                path.append(n)
                n = parents[n]
            path.append(start_node)
            path.reverse()
            print('Path found: {}'.format(path))
            return path

        # Move current node from open set to closed set
        open_set.remove(n)
        closed_set.add(n)

    # If the loop exits without finding a path, it prints an error message
    print('Path does not exist!')
    return None

# Display available nodes
print("\nNodes Available:")
print()
print(graphNodes)
print()

# Take user input for start and destination nodes
point_a = input("Enter the starting point: ")
point_b = input("Enter the destination point: ")

# Call the main algorithm function to find the shortest path of the given starting point and ending point
aStarAlgorithm(point_a, point_b)
