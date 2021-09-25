# String Tricks in Go

byte can be used directly for arithmetics. for example:

```
s := "123"

fmt.Println(s[1]-'0') // Will print 1

fmt.Println(s[2]-'0') // Will print 2

fmt.Println(s[2]-s[0]) // Will print 2
```