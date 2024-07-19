package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Speak() {
	fmt.Println("Hello my name is", h.Name, "i'm", h.Age, "years old")
}

type Action struct {
	Human
}

func main() {
	action := Action{
		Human: Human{
			Name: "Anton",
			Age:  27,
		},
	}
	action.Speak()
}
