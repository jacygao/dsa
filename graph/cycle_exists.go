package graph

var (
	white int = 0
	grey  int = 1
	black int = 2
)

// For a disconnected graph, we get the DFS forest as output.
// To detect cycle, we can check for cycle in individual trees by checking back edges with colors approach.
func cycleExists(n, e int, edges [][]int) bool {
	// using white, grey, black approach to detect cycles in a graph.
	// process input, create adj list.
	g := make([][]int, n)
	for _, v := range edges {
		g[v[0]] = append(g[v[0]], v[1])
	}

	// pick starting node from the white list, conduct dfs
	// when visiting a node, mark as grey
	// after visiting all children, mark as black
	// detect back edges when a node is pointing to a grey node, cycle exists
	colors := make([]int, n)
	for index, color := range colors {
		if color == white {
			return check(g, colors, index)
		}
	}
	return false
}

func check(g [][]int, colors []int, node int) bool {
	if colors[node] != white {
		return false
	}
	colors[node] = grey
	if len(g[node]) > 0 {
		for _, child := range g[node] {
			if colors[child] == grey {
				return true
			}
			if check(g, colors, child) {
				return true
			}
		}
	}
	colors[node] = black
	return false
}
