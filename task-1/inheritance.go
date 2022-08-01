package main

import "fmt"

//Go prefers to use composition instead of inheritance. This allows you to embed one structure within another.
//Go does not support type inheritance.

// Action - We insert the Human into the Action structure, thus the Action structure inherits the Human methods
type Action struct {
	Human
}

// Human - Define the structure Human with some set of fields
type Human struct {
	age    int
	height int
}

// Talk - Define a method for the Human structure
func (h *Human) Talk() {
	fmt.Printf("my age %d and my height %d \n", h.age, h.height)
}

func main() {

	//Initialize the Action structure with the Human structure embedded in it
	action := &Action{Human{
		age:    15,
		height: 176,
	}}

	//Call the method of the Human structure for the Action structure via embedding
	action.Talk()
}
