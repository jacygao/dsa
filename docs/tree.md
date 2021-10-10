# Tree

## Depth First Search (DFS)

### Use Cases

- Return kth smallest node in BST

### Techniques

- Inorder DFS of a BST results in an ascending order

## Breath First Search (BFS)

### Use Cases

- Return all left side or right side nodes of a binary tree
- Return some form of nodes per level (Zigzag level order traversal)

## Remove a tree node from BST

There are 3 situations when removing a tree node from BST:
- Remove a leaf node
- Remove a node with one child
- Remove a node with two children

### Remove a leaf node

- update reference in its parent node to `nil`.

### Remove a node with one child

- update reference in its parent node to its child node.

### Remove a node with two children

#### Option 1

update its parent node to the smallest node in its right sub tree or the largest node in its left sub tree.

#### Option 2

Update its parent node to the right child node, then update the child of the smallest node in the right sub tree to the left child node of the parent node.

or

Update its parent node to the left child node, then update the child of the largest node in the left sub tree to the right child node of the parent node.

#### Techniques

- Recursion

- Due to the natual of BST, all nodes in the left sub tree are less than the root node and all nodes in the right sub tree are larger than the root node.

Therefore

- If target node value is less than the root node value, traverse to the left sub tree

- If target node value is larger than the root node value, traverse to the right sub tree

- The smallest node is the most left leaf node of the right sub tree

- The largest node is the most right leaf node of the left sub tree

#### Time Complexity

`O(h)` where h is the height of the tree

## Counting nodes

Counting how many nodes in a binary tree can be simplified to counting both left and right sub tree plus the root node itself.

For example:

```
    func countNodes(root *TreeNode) int {
        if root == nil {
            return 0
        }
        return countNodes(root.Left) + countNodes(root.Right) + 1
    }
```

### Time Complexity 

`O(n)`

### Optimisation

- Total nodes of a perfect tree is `2^h-1` where `h` is the height of the tree
- Based on the above technique, total node of a complete tree is total nodes of h-1 plus total nodes of the leaf nodes on `h`
- Using this technique can optimise time complexity to under `O(n)`