package sort

type Merger struct {
	in []int
}

func MergeSort(input []int) []int {
	l := len(input)
	temp := make([]int, l)
	merger := &Merger{input}
	merger.sort(temp, 0, l-1)

	return merger.in
}

func (m *Merger) sort(temp []int, low, high int) {
	if low < high {
		mid := int((low + high) / 2)
		m.sort(temp, low, mid)
		m.sort(temp, mid+1, high)
		m.merge(temp, low, mid, high)
	}
}

func (m *Merger) merge(temp []int, low, mid, high int) {
	for i := low; i <= high; i++ {
		temp[i] = m.in[i]
	}
	cur := low
	hl := low
	hr := mid + 1
	for hl <= mid && hr <= high {
		if temp[hl] > temp[hr] {
			m.in[cur] = temp[hr]
			hr++
		} else {
			m.in[cur] = temp[hl]
			hl++
		}
		cur++
	}

	for i := 0; i <= mid-hl; i++ {
		m.in[cur+i] = temp[hl+i]
	}
}
