package main

import "fmt"

type transport interface {
	navigateToDestination()
}

type client struct {
}

// client method startNavigation can only take parameter that implements transport interface
func (c *client) startNavigation(trans transport) {
	fmt.Println("client starting the navigation process")
	trans.navigateToDestination()
}

type boat struct{}

func (b *boat) navigateToDestination() {
	fmt.Println("boat is navigating to island")
}

type car struct{}

// car does not implement transport interface, so it cannot be passed into client.startNavigation as a parameter
func (c *car) driveToDestination() {
	fmt.Println("car is going to the destination")
}

// adapter struct used by car to go to island
type carAdapter struct {
	carTransport *car
}

// carAdapter allows car to implement transport interface by adding a navigateToDestination method
// internally, carAdapter contains a car, which can driveToDestination
func (c *carAdapter) navigateToDestination() {
	fmt.Println("Adapter modify car to allow navigation.")
	c.carTransport.driveToDestination()
}

func main() {
	client := client{}
	boat := boat{}
	client.startNavigation(&boat)
	// An assignment to a variable of interface type is valid if the value being assigned
	// implements the interface it is assigned to. It implements it if its method set is
	// a superset of the interface. The method set of pointer types includes methods with
	// both pointer and non-pointer receiver. The method set of non-pointer types only
	// includes methods with non-pointer receiver.

	car := car{}
	carAdapter := carAdapter{
		carTransport: &car,
	}
	// by calling carAdapter instead of car, client can now navigate to island
	client.startNavigation(&carAdapter)
}
