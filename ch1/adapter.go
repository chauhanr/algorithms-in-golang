package ch1

import "fmt"

//IProcess interface
type IProcess interface {
	process()
}

//Adapter struct
type Adapter struct {
	adaptee Adaptee
}

func (adapter Adapter) process() {
	fmt.Printf("Adapter Process\n")
	adapter.adaptee.convert()
}

//Adaptee Struct
type Adaptee struct {
	adapterType int
}

func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}
