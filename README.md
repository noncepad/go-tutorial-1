# go-tutorial-1
Tutorial for go


# Setup

[Install Golang](https://go.dev/doc/install).

```bash
cd $MYWORKSPACE
git clone https://github.com/noncepad/go-tutorial-1
cd go-tutorial-1
```

## Starting From Scratch

```bash
go mod init github.com/noncepad/go-tutorial-1
```

This creates a `go.mod` file that will contain dependencies.


# Points

## Summary

Golang is a typed language that compiles to a binary.  It was designed by Google to run backend Rest API services in the cloud.  Golang simplifies concurrency programming by using channels.  Python programmers have the easiest time transitioning to Golang.

Unlike Rust and Typescript, there is no async.  Go-routines, a kind of light thread, are used instead.

## Basic Syntax


```go
x:=2
y:=x
var z int
z = x + y
if z < 0{
    log.Printf("z=%d",z)
} else {
    panic("z is too big")
}
```

## Structs

Pass-by-value:

```go
type Fruit struct{
    Name string `json:"name"`
    Price float64 `json:"price"`
}

apple:=Fruit{
    Name: "apple", Price: 34.244,
}

```

Pass-by-reference:

```go
apple:=new(Fruit)
apple.Name="apple"
apple.Price=34.244
if apple == nil{
    panic("apple should have been set")
}
```

## Functions

Pass-by-value:

```go
func EatFruit(f Fruit) bool{
    if f.Name == "apple"{
        return true
    } else {
        // throw up everything else
        return false
    }
}
```

Pass-by-reference

```go
type Fruit struct{
    Name string `json:"name"`
    Price float64 `json:"price"`
    IsEaten bool `json:"is_eaten"`
}
// both styles are ok, but the second one is more convenient in keeping functions clear.
// func EatFruitKeepState(f *Fruit) bool 
func (f *Fruit)EatFruitKeepState() bool{
    if f==nil{
        panic("no fruit")
    }
    if f.Name=="apple"{
        f.IsEaten = true
        return true
    } else {
        return false
    }
}
```

# Programming Pattern

Separate pass-by-reference from everything else.  

Put the pass-by-reference (state) in a single go-routine.

```go
// this is safe with f=Fruit{}
go EatFruit(f)

// this is not safe with f=&Fruit{} or f:=new(Fruit)
go f.EatFruitKeepState()
```

```go
func loopInternal(ctx context.Context, internalC <-chan func(*Fruit), initialState *Fruit){
    log.Print("entering loop")
    in:=intialState
out:
    for{
        select{
            case <-ctx.Done():
                break out
            case req:=<-internalC:
                req(in)
        }
    }
    log.Print("leaving loop")
}
fruit:=&Fruit{
    Name:"apple",
    Price: 34.4324,
    IsEaten: false,
}
ctxOutside,cancelOutside:=context.WithCancel(context.Background()
internalC:=make(chan func(*Fruit),1)
// safe to pass by reference because we will not touch fruit again in this goroutine
go loopInternal(ctxOutside,internalC,fruit)

// how do we touch the fruit then?
ansC:=make(chan bool,1)
internalC<-func(f *Fruit){
    ansC<-f.EatFruitKeepState()
}
didIEatTheFruit:=<-ansC
log.Printf("did I eat the fruit? %b",didIEatTheFruit)
```

`ctx` is the common way to refer to a *context*.  In Golang, developers are expected to spawn many goroutines.  Manually using channels to make sure all of the goroutines have not had an error or properly exit when the program is exiting is difficult.  Context simplifies this task.  

Context objects are hereditary.

```go
ctxChild,cancelChild:=context.WithCancel(ctxParent)
```

To close a context:

```go
cancelChild()
```

Canceling means that:

```go
<-ctxChild.Done()
log.Print("context has been canceled")
```

See the loopInternal function above.  Once the context has been canceled, the goroutine will exit.  It is safe to pass context objects around as they pass-by-value.