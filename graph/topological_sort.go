package graph

import "container/list"

// TopologicalSortDFS implements the Topological Sorting algorithm using the DFS approach.
func TopologicalSortDFS(n int, edges [][]int) []int {
	ans := []int{}
	g := genDirectedAdj(n, edges)
	stack := list.New()
	visited := make([]int, n)

	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			dfsTopologicalSort(g, i, stack, visited)
		}
	}
	for stack.Len() > 0 {
		m := stack.Back()
		stack.Remove(m)
		ans = append(ans, m.Value.(int))
	}

	return ans
}

func dfsTopologicalSort(g [][]int, n int, s *list.List, visited []int) {
	visited[n] = 1
	for _, c := range g[n] {
		if visited[c] == 0 {
			dfsTopologicalSort(g, c, s, visited)
		}
	}
	s.PushBack(n)
}

// TopologicalSortKahn implements the Topological Sorting algorithm using Kahn's algorithm.
func TopologicalSortKahn(n int, edges [][]int) []int {
	graph := genDirectedAdj(n, edges)
	queue := []int{}
	degrees := make([]int, n)
	res := []int{}

	for _, edge := range edges {
		degrees[edge[1]]++ 
	}

	for node, degree := range degrees {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		peek := queue[0]
		queue = queue[1:]
		res = append(res, peek)
		for _, node := range graph[peek] {
			degrees[node]--
			if degrees[node] == 0 {
				queue = append(queue, node)
			}
		}
	}

	return res
}


func genDirectedAdj(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], edge[1])
	}
	return g
}
