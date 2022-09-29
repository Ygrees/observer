package main

import "fmt"

type Ticket struct {
	observerList []Observer
	artist       string
	city         string
	tickets      int
	inStock      bool
}

func newTicket(artist string, city string, tickets int) *Ticket {
	return &Ticket{
		artist:  artist,
		city:    city,
		tickets: tickets,
	}
}
func (i *Ticket) updateAvailability() {
	fmt.Printf("Concert %s is now in %s \n", i.artist, i.city)
	i.inStock = true

	i.notifyAll()
}
func (i *Ticket) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Ticket) deregister(o Observer) {
	i.observerList = removeFromslice(i.observerList, o)
}

func (i *Ticket) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.artist)
	}
}

func removeFromslice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
