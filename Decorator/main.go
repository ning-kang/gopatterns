package main

import "fmt"

type Human interface {
	Intro() string
}

type NamedHuman struct {
	name string
}

func (n *NamedHuman) Intro() string {
	return fmt.Sprintf("Hi, my name is %s", n.name)
}

// Decorator 1
type Worker struct {
	NamedHuman NamedHuman
	job        string
}

func (w *Worker) Intro() string {
	return fmt.Sprintf("%s, I work as %s", w.NamedHuman.Intro(), w.job)
}

// Decorator 2
type AgedHuman struct {
	Worker Worker
	age    int
}

func (a *AgedHuman) Intro() string {
	return fmt.Sprintf("%s, I am %v year old", a.Worker.Intro(), a.age)
}

func main() {
	person := &AgedHuman{Worker{NamedHuman{"Lisa"}, "Engineer"}, 30}
	fmt.Println(person.Intro())
}
