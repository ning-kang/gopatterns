package main

import "fmt"

type Design struct {
	style string
}

func (d *Design) Design() {
	fmt.Printf("Design the house in %s style\n", d.style)
}

type Material struct {
	material string
}

func (m *Material) Purchase() {
	fmt.Printf("Purchase %s material\n", m.material)
}

type BuiltBy struct {
	builder string
}

func (b *BuiltBy) Build() {
	fmt.Printf("The house is built by %s\n", b.builder)
}

type House struct {
	Design   Design
	Material []Material
	BuildBy  BuiltBy
}

func NewHouse(style string, materials []string, builder string) *House {
	d := Design{style: style}
	m := make([]Material, 2)
	for _, material := range materials {
		m = append(m, Material{material: material})
	}
	b := BuiltBy{builder: builder}
	return &House{d, m, b}
}

func main() {
	h := NewHouse("modern", []string{"Wood", "Brick"}, "Awesome Builder")
	h.Design.Design()
	for _, m := range h.Material {
		m.Purchase()
	}
	h.BuildBy.Build()
}
