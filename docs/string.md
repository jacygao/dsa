# string

## Techniques

- A slice with 26 length can be leveraged as a more efficient map

- When using byte letters as index, substract `'a'` to get their position in the alphabet. For example `'c'-'a'= 2`

- When solving string problems using iteration, also consider iterating string in reverse order. For example [Backspace String Compare](https://leetcode.com/problems/backspace-string-compare/)