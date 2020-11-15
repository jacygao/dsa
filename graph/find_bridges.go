package graph

var (
	time int = 0
	ans  [][]int
)

func findBridges(n int, edges [][]int) [][]int {
	// disc tracks the time when a node is visited
	disc := make([]int, n)
	// low tracks the lower value of a comparason of two nodes
	low := make([]int, n)
	// visitied tracks whether a node has been visited by dfs
	visited := make([]int, n)
	// parent keeps the parent node of a node
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
	}
	// answer slice
	ans = [][]int{}
	// converts edges to adjacency list
	adj := genUndirectedAdj(n, edges)

	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			dfsFindBridges(adj, i, disc, low, visited, parent)
		}
	}

	return ans
}

// genGraph converts a list of edges to an adjacency list for
// processing graph related problems. n represents the total amount
// of given nodes. Assume edges do not contain duplicates.
func genUndirectedAdj(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], edge[1])
		g[edge[1]] = append(g[edge[1]], edge[0])
	}
	return g
}

func dfsFindBridges(g [][]int, n int, disc, low, visited, parent []int) {
	time += 1
	disc[n] = time
	low[n] = time
	visited[n] = 1
	for _, v := range g[n] {
		if visited[v] == 0 {
			parent[v] = n
			dfsFindBridges(g, v, disc, low, visited, parent)

			// compare low value and update parent low to lower value
			low[n] = min(low[n], low[v])
			// compare low value to disc value
			if low[v] > disc[n] {
				// bridge found
				ans = append(ans, []int{n, v})
			}
		} else if v != parent[n] {
			low[n] = min(low[n], disc[v])
		}
	}

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
