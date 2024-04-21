package main

import "fmt"

type Position struct {
	X, Y int
}

type Hero struct {
	faceTo   string
	position Position
}

func (h *Hero) Turn(direction string) {
	h.faceTo = direction
}

func (h *Hero) Move(distance int) {
	switch h.faceTo {
	case "East":
		h.position.X += distance
	case "West":
		h.position.X -= distance
	case "North":
		h.position.Y += distance
	case "South":
		h.position.Y -= distance
	}
}

// command pattern
type Command interface {
	Call()
	Undo()
}

type Direction int

const (
	East Direction = iota
	West
	North
	South
)

type HeroTurnCommand struct {
	hero              *Hero
	previousDirection string
	direction         Direction
}

type HeroMoveCommand struct {
	hero     *Hero
	distance int
}

func NewHeroTurnCommand(hero *Hero, previousDirection string, direction Direction) *HeroTurnCommand {
	return &HeroTurnCommand{hero: hero, previousDirection: previousDirection, direction: direction}
}

func NewHeroMoveCommand(hero *Hero, distance int) *HeroMoveCommand {
	return &HeroMoveCommand{hero: hero, distance: distance}
}

func (h *HeroTurnCommand) Call() {
	h.previousDirection = h.hero.faceTo
	switch h.direction {
	case East:
		h.hero.faceTo = "East"
	case West:
		h.hero.faceTo = "West"
	case North:
		h.hero.faceTo = "North"
	case South:
		h.hero.faceTo = "South"
	}
}

func (h *HeroTurnCommand) Undo() {
	h.hero.faceTo = h.previousDirection
}

func (h *HeroMoveCommand) Call() {
	switch h.hero.faceTo {
	case "East":
		h.hero.position.X += h.distance
	case "West":
		h.hero.position.X -= h.distance
	case "North":
		h.hero.position.Y += h.distance
	case "South":
		h.hero.position.Y -= h.distance
	}
}

func (h *HeroMoveCommand) Undo() {
	switch h.hero.faceTo {
	case "East":
		h.hero.position.X -= h.distance
	case "West":
		h.hero.position.X += h.distance
	case "North":
		h.hero.position.Y -= h.distance
	case "South":
		h.hero.position.Y += h.distance
	}
}

func main() {
	hero := Hero{}
	tcmd1 := NewHeroTurnCommand(&hero, hero.faceTo, West)
	tcmd2 := NewHeroTurnCommand(&hero, hero.faceTo, West)
	mcmd := NewHeroMoveCommand(&hero, 100)

	tcmd1.Call()
	fmt.Printf("Hero is at %v facing to %s\n", hero.position, hero.faceTo)
	tcmd2.Call()
	fmt.Printf("Hero is at %v facing to %s\n", hero.position, hero.faceTo)
	tcmd2.Undo()
	fmt.Printf("Hero is at %v facing to %s\n", hero.position, hero.faceTo)

	mcmd.Call()
	fmt.Printf("Hero is at %v facing to %s\n", hero.position, hero.faceTo)
	mcmd.Call()
	fmt.Printf("Hero is at %v facing to %s\n", hero.position, hero.faceTo)
}
