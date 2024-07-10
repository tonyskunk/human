package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Speak() {
	fmt.Println("Прив, меня зовут", h.Name, "\nМне", h.Age)
}

type Action struct {
	Human
}

func main() {
	action := Action{
		Human: Human{
			Name: "Антон",
			Age:  27,
		},
	}
	action.Speak()
}
