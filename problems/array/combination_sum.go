package array

func CombinationSumSol1(candidates []int, target int) [][]int {
	res := &Result{}
	res.calculate(0, candidates, target, 0, nil)
	return res.res
}

type Result struct {
	res [][]int
}

func (r *Result) calculate(i int, candidates []int, target int, sum int, cur []int) {
	if sum > target {
		return
	}
	if sum == target {
		r.res = append(r.res, append([]int{}, cur...))
		return
	}
	for ; i < len(candidates); i++ {
		r.calculate(i, candidates, target, sum+candidates[i], append(cur, candidates[i]))
	}
}
