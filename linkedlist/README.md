# linkedlist

linkedlist implements a basic doubly linked list in Go.

## Usage

### Create a new linked list

```go
ll := linkedlist.New()
```

### Append value to the end of a linked list

```go
// Append accepts any interface.
ll.Append("your value")
```

### Prepend value to the start of a linked list

```go
// Prepend accepts any interface.
ll.Prepend("your value")
```

### Delete a node from a linked list

```go
// Retrieve the first node from linked list
node := ll.First

ll.Delete(node)
```

### Shift a node to the end of a linked list

```go
// Retrieve the first node from linked list
node := ll.First

ll.ShiftToEnd(node)
```