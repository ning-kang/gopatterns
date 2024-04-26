package main

import "fmt"

type class string

const (
	mage    class = "mage"
	worrior class = "worrior"
	hunter  class = "hunter"
)

type Player struct {
	name  string
	class class
	game  *Game
}

func NewPlayer(name string, class class) *Player {
	return &Player{name: name, class: class}
}

func (p *Player) String() string {
	return p.name + " the " + string(p.class)
}

func (p *Player) Receive(source Player, message string) {
	fmt.Printf("%s -- [%s]: %s\n", p, &source, message)
}

func (p *Player) Broadcast(message string) {
	for _, dest := range p.game.players {
		if dest.name != p.name {
			dest.Receive(*p, message)
		}
	}
}

func (p *Player) Wisper(dest Player, message string) {
	dest.Receive(*p, message)
}

type Game struct {
	players []*Player
}

func (g *Game) Join(p *Player) {
	p.game = g
	g.players = append(g.players, p)
	fmt.Println("[Game]:", p, "has joined the game")
}

func main() {
	g := Game{}
	p1 := NewPlayer("John", mage)
	p2 := NewPlayer("James", worrior)
	p3 := NewPlayer("Jazz", hunter)
	g.Join(p1)
	g.Join(p2)
	p2.Broadcast("Hello there")
	g.Join(p3)
	p3.Broadcast("Hello Everyone")
	p2.Wisper(*p1, "Hello John")
}
