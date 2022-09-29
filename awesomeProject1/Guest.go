package main

import "fmt"

type Guest struct {
	id   string
	city string
}

func (c *Guest) update(Ticket string) {
	fmt.Printf("Sending email to Guest %s for availability to buy Ticket %s\n", c.id, Ticket)

}

func (c *Guest) getID() string {
	return c.id
}
