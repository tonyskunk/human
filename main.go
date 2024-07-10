package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Speak() {
	fmt.Println("Привет меня зовут", h.Name, "мне", h.Age, "лет")
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
