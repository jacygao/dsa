# Prefix Sum

## Use case

- Contiguous subarray problems
- Calculate sum or product of subarray
- Find all subarrays with sum or product equal to a given value

## Sum

Input:  1 -> 2 -> 3 -> 4

Prefix Sum: 1 -> 3 -> 6 -> 10

To calculate sum between indexes 2 and 3, use the value of index 3 to subtract the value of index 1 in the Perfix Sum list:

```
Sum between index 2 and 3 from the input list = Prefix Sum [3] - Prefix Sum[1] => 10 - 3 = 7
```

## Multiply

Input:  1 -> 2 -> 3 -> 4

Prefix Product: 1 -> 2 -> 6 -> 24

```
Product between index 2 and 3 from the input list = Prefix Product [3] / Prefix Product[1] => 24 / 2 = 12
```

## Subtract

Input:  1 -> 2 -> 3 -> 4

Prefix Sum: 1 -> 3 -> 6 -> 10

To Look for a subarray sum, take the value of the right index, subtract the value of the left index - 1.

For example, to get the sum between index 1 and 3:

```
prefixSum[3] - prefixSum[0] => 10 - 1 = 9 

```

## Techniques and tricks

When using Presfix Sum as cumulative sum table, it is often helpful to create an array with n+1 size. The additional slot is used to store 0 as the first element to reduce further if statement when index is 0.

for example:

```
Input:  1 -> 2 -> 3 -> 4

Prefix Sum: 0 -> 1 -> 3 -> 6 -> 10
```

To calculate all sums of sub arrays:

```
func allSums(input []int) []int {
    sums := []int{}
    for i:=0; i<len(input); i++ {
        for j:=0; j<len(input); j++ {
            sums = append(sums, prefixSum[j+1]-prefixSum[i])
        }
    }
    return sums
}
```
