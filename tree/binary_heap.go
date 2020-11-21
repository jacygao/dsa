package tree

type HeapNode struct {
	val      int
	priority int
}

type MinHeap struct {
	size     int
	capacity int
	tree     []*HeapNode
}

func NewMinHeap(cap int) *MinHeap {
	return &MinHeap{
		capacity: cap,
		tree:     []*HeapNode{},
	}
}

func (mh *MinHeap) Insert(priority int, val int) {
	// additional node beyond capacity is ignored
	if mh.size == mh.capacity {
		return
	}

	newIndex := mh.size
	mh.size++
	mh.tree = append(mh.tree, &HeapNode{
		val:      val,
		priority: priority,
	})
	if newIndex == 0 {
		return
	}
	parIndex := mh.parentIndex(newIndex)
	for mh.tree[newIndex].priority < mh.tree[parIndex].priority {
		mh.swap(newIndex, parIndex)
		if parIndex == 0 {
			break
		}
		newIndex = parIndex
		parIndex = mh.parentIndex(newIndex)
	}
}

func (mh *MinHeap) ExtractMin() *HeapNode {
	parIndex := 0
	for mh.rightChild(parIndex) < mh.size &&
		mh.tree[parIndex].priority < mh.tree[mh.leftChild(parIndex)].priority &&
		mh.tree[parIndex].priority < mh.tree[mh.rightChild(parIndex)].priority {
		if mh.tree[mh.leftChild(parIndex)].priority <= mh.tree[mh.rightChild(parIndex)].priority {
			mh.swap(parIndex, mh.leftChild(parIndex))
			parIndex = mh.leftChild(parIndex)
			continue
		}
		if mh.tree[mh.leftChild(parIndex)].priority > mh.tree[mh.rightChild(parIndex)].priority {
			mh.swap(parIndex, mh.rightChild(parIndex))
			parIndex = mh.rightChild(parIndex)
			continue
		}
	}
	res := mh.tree[parIndex]
	mh.tree = append(mh.tree[:parIndex], mh.tree[parIndex+1:]...)
	mh.size--
	return res
}

func (mh *MinHeap) swap(i1, i2 int) {
	t := mh.tree[i1]
	mh.tree[i1] = mh.tree[i2]
	mh.tree[i2] = t
}

func (mh *MinHeap) parentIndex(index int) int {
	return (index - 1) / 2
}

func (mh *MinHeap) leftChild(index int) int {
	return index*2 + 1
}

func (mh *MinHeap) rightChild(index int) int {
	return index*2 + 2
}
