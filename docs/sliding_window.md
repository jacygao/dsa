# Sliding Window

In some problems, Sliding Window can convert a nested loop to a single for loop to reduce the time complexity from `O(n^2)` to `O(n)`.

## Use Cases

- sub-array / sub-set problems

## Techniques

- For solving string problems, maintain a live character counts by using arrays with 26 length ([26]int{})
- 2 pointers are similar to sliding window for solving sub-array problems. In the case of 2 pointers the window is resizable with its left and right indexes