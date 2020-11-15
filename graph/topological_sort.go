package graph

import "container/list"

func TopologicalSort(n int, edges [][]int) []int {
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

func genDirectedAdj(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], edge[1])
	}
	return g
}
