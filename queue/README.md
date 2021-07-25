# queue

queue implement a simple FIFO queue backed by a LinkedList

## Operations

```
type Operations interface {
	Add(interface{})
	Remove() interface{}
	Peek() interface{}
	IsEmpty() bool
}
```

## Usage

initialise queue

```
    q := queue.New()
```

add element to queue (accepting interface type)

```
    queue.Add("anything")
```

Peek the first element in queue

```
    val := queue.Peek()
```

Check if queue is empty

```
    isEmpty := queue.IsEmpty()
```

remove first element from queue

```
    queue.Remove()
```