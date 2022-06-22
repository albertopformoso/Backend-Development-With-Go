# Sharing Memory with Communication

## CSP Model

+ input/output primitives to communicate processes in concurrent code.
+ `!` sends information to a proces
+ `?` reads the output of a proces

**Protected Command `->`**
+ Gard(Conditional) -> Declaration

Example:
```
p1?c -> p2!c
```
The output of `p1` is sent to a varaibale `c` and the input of `p2` is received from the same variable.

## Channels

The communication by channel is the principal method of syncronization between go routines.

To create a chanel:
```go
ch := make(chan string)
```

Send "Hello!" to the chanel `ch`
```go
ch <- "Hello!"
```

Recieve from the chanel `ch` and assign the value to the variable message.
```go
message := <-ch
```

With channels you don't have to lock and unlock the access to the shared memory because you are using this channel to communicate the data and the channels are in charge of locking and unlocking the access.

Example:
```go
func main() {
    message := make(chan string)

    go func() {
        message <- "Hello!"
    }()

    fmt.Println(<-message)
}
```

Representative diagram:

<img src="../media/fork-join-model-ch.png" width="400px"/>
