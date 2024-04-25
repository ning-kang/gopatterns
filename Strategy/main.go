package main

import "fmt"

// Core strategy interface
type NavigationStrategy interface {
	Navigate(dest string)
}

type Transport int

const (
	Drive Transport = iota
	Walk
)

type DriveNavigationStrategy struct{}

func (d *DriveNavigationStrategy) Navigate(dest string) {
	fmt.Printf("Navigate to %s by driving\n", dest)
}

type WalkNavigationStrategy struct{}

func (d *WalkNavigationStrategy) Navigate(dest string) {
	fmt.Printf("Navigate to %s by walking\n", dest)
}

type Navigator struct {
	navigationStrategy NavigationStrategy
}

func NewNavigator(navigationStrategy NavigationStrategy) *Navigator {
	return &Navigator{navigationStrategy: navigationStrategy}
}

func (n *Navigator) SetTransport(t Transport) {
	switch t {
	case Drive:
		n.navigationStrategy = &DriveNavigationStrategy{}
	case Walk:
		n.navigationStrategy = &WalkNavigationStrategy{}
	}
}

func (n *Navigator) GoTo(dest string) {
	n.navigationStrategy.Navigate(dest)
}

func main() {
	n := NewNavigator(&DriveNavigationStrategy{}) // Initiate with a default strategy
	n.GoTo("Paris")

	n.SetTransport(Walk)
	n.GoTo("London")
}
