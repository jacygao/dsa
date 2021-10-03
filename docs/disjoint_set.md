# Disjoint Set

## Other names

- Union-find
- Merge-find

## Use Case

- Detect cycle in an **undirected** graph
- Find consecutive elements in an unsorted list

## Techniques

Using a 2 steps approach to solve these problems using Disjoint set:

- Find: a function that finds the absolute root of a node

- Union: merge if 2 nodes are disjointed

## Graph

Detect cycle in an **undirected** graph is the most popular problem solve by Disjoint Set.

## Array

Find consecutive elements in an unsorted list

Input: [100,4,200,1,3,2]
Output: max length of consecutive elements (4 in this case as in [1,2,3,4])

Solution:
Consider every element is a node in a tree, every element will point to a root node or itself.

With this technique in mind, we can create a hashmap to look up every element:

```
    func solve(input []int) int {
        set := make(map[int]struct{}, len(input))
        for val := range input {
            set[val] = struct{}{}
        }
        
        max := 0
        for _, val := range set {
            if _, ok := set[val-1]; ok {
                continue
            }
            
            cur := 0
            for _, ok := set[val]; ok {
                cur++
                val+=1
            }
            if cur > max {
                max = cur
            }
        }
        return max
    }
```