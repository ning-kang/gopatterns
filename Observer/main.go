package main

import "fmt"

type Observable interface {
	Subscribe(observer Observer)
	Unsubscribe(observer Observer)
	NotifyContent() string
	NotifyAll()
}

type Observer interface {
	Name() string
	Notify(observable Observable)
}

type Movie struct {
	name string
}

// Observable
type MovieLibrary struct {
	subscriptions []Observer
	movieList     []Movie
	newMovie      Movie
}

func (m *MovieLibrary) Subscribe(observer Observer) {
	m.subscriptions = append(m.subscriptions, observer)
}

func (m *MovieLibrary) Unsubscribe(observer Observer) {
	subLength := len(m.subscriptions)
	for i, sub := range m.subscriptions {
		if sub.Name() == observer.Name() {
			m.subscriptions[subLength-1], m.subscriptions[i] = m.subscriptions[i], m.subscriptions[subLength-1]
			m.subscriptions = m.subscriptions[:subLength-1]
		}
	}
}

func (m *MovieLibrary) NotifyContent() string {
	return m.newMovie.name
}

func (m *MovieLibrary) NotifyAll() {
	for _, sub := range m.subscriptions {
		sub.Notify(m)
	}
}

func (m *MovieLibrary) AddNewMovie(movie Movie) {
	m.movieList = append(m.movieList, movie)
	m.newMovie = movie
	fmt.Printf("Added new movie: %s\n", movie.name)
	m.NotifyAll()
}

// Observer
type Person struct {
	name string
}

// Name implements Observer.
func (p Person) Name() string {
	return p.name
}

// Notify implements Observer.
func (p Person) Notify(observable Observable) {
	fmt.Printf("Hi %s, your movie library just got new movie %s\n", p.Name(), observable.NotifyContent())
}

func main() {
	ml := &MovieLibrary{}
	user1 := &Person{"Tony"}
	user2 := &Person{"Boris"}
	ml.Subscribe(user1)
	ml.Subscribe(user2)

	ml.AddNewMovie(Movie{"The Godfather"})
}
