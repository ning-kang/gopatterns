package main

import "fmt"

type Shape interface {
	Render()        // Original Method
	Accept(visitor) // Visitor Method
}

type Square struct {
}

type Triangle struct {
}

type Circle struct {
}

func (s *Square) Render() {
	fmt.Println("Rendering Square...")
}

func (s *Square) Accept(v visitor) {
	v.visitSquare(*s)
}

func (t *Triangle) Render() {
	fmt.Println("Rendering Triangle...")
}

func (t *Triangle) Accept(v visitor) {
	v.visitTriangle(*t)
}

func (c *Circle) Render() {
	fmt.Println("Rendering Circle...")
}

func (c *Circle) Accept(v visitor) {
	v.visitCircle(*c)
}

// New Request: Get Area of each Shape

type visitor interface {
	visitSquare(Square)
	visitTriangle(Triangle)
	visitCircle(Circle)
}

// New Visitor
type AreaVisitor struct {
}

func (a *AreaVisitor) visitCircle(c Circle) {
	fmt.Println("Calculating Circle area...")
}

func (a *AreaVisitor) visitSquare(s Square) {
	fmt.Println("Calculating Square area...")
}

func (a *AreaVisitor) visitTriangle(t Triangle) {
	fmt.Println("Calculating Triangle area...")
}

func main() {
	s := &Square{}
	t := &Triangle{}
	c := &Circle{}

	// Call native methods
	s.Render()
	t.Render()
	c.Render()

	// Call visitor methods
	av := &AreaVisitor{}
	s.Accept(av)
	t.Accept(av)
	c.Accept(av)
}
