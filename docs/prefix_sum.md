# Prefix Sum

# Sum

Input:  1 -> 2 -> 3 -> 4

Prefix Sum: 1 -> 3 -> 6 -> 10

To calculate sum between indexes 2 and 3, use the value of index 3 to subtract the value of index 1 in the Perfix Sum list:

```
Sum between index 2 and 3 from the input list = Prefix Sum [3] - Prefix Sum[1] => 10 - 3 = 7
```

# Multiply

Input:  1 -> 2 -> 3 -> 4

Prefix Product: 1 -> 2 -> 6 -> 24

```
Product between index 2 and 3 from the input list = Prefix Product [3] / Prefix Product[1] => 24 / 2 = 12
```

## Use case

- Calculate sum or product between 2 indexes of an array

