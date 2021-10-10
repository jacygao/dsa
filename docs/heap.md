# Heap

Heap is a semi-ordered tree-based data structure.

## Use Cases

- Kth Largest Element in an Array

## Types

- Max Heap: Each parent key is greater than its children (largest element on top)
- Min Heap: Each parent key is less than ts children (smallest element on top)

## Operations

- Insert
- Removal

## Storage

Using an `array` with starting index as `1`

## Arithmetic

- Get left child index: `cur * 2`
- Get right child index: `cur * 2 + 1`
- Get parent index: `floor(cur / 2)`

## Priority Queue
