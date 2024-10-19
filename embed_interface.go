package main

import "fmt"

type Fooer interface {
	Foo() string
}

type Wrapper struct {
	Fooer
	Counter int
}

// sink takes a value implementing the Fooer interface.
func sink(f Fooer) {
	fmt.Println("Fooer", &f)
	wrapper, ok := f.(Wrapper)
	if ok {
		fmt.Println("Counter", wrapper.Counter)
	} else {
		fmt.Println("Not a Wrapper")
	}
	fmt.Println("sink:", f.Foo())
}

// TheRealFoo is a type that implements the Fooer interface.
type TheRealFoo struct {
}

func (trf TheRealFoo) Foo() string {
	return "TheRealFoo Foo"
}

func main() {
	wrapper := Wrapper{
		Fooer:   TheRealFoo{},
		Counter: 1,
	}
	fmt.Println("wrapper", &wrapper)
	sink(wrapper)

	messages := make(chan Fooer)

	go func() { messages <- wrapper }()

	f := <-messages
	wrapper1, ok := f.(Wrapper)
	if ok {
		fmt.Println("Read from channel: Counter", wrapper1.Counter)
	} else {
		fmt.Println("Read from channel: Not a Wrapper")
	}
}
