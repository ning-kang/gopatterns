package main

import "fmt"

type Student struct {
	name string
}

type Collection interface {
	createIterator() Iterator
}

// Class is a Collection of Student
type Class struct {
	students []*Student
}

func (c *Class) createIterator() Iterator {
	return &StudentIterator{students: c.students}
}

type Iterator interface {
	hasNext() bool
	getNext() *Student
}

type StudentIterator struct {
	students []*Student
	index    int
}

func (i *StudentIterator) hasNext() bool {
	return i.index < len(i.students)
}

func (i *StudentIterator) getNext() *Student {
	if i.hasNext() {
		s := i.students[i.index]
		i.index++
		return s
	}
	return nil
}

func main() {
	c := &Class{
		students: []*Student{{"James"}, {"David"}, {"Allan"}},
	}
	for i := c.createIterator(); i.hasNext(); {
		s := i.getNext()
		fmt.Printf("Student %s\n", s.name)
	}
}
