package main

func main() {

	Artist := newTicket("Instasamka", "Astana", 100)

	observerFirst := &Guest{id: "dora@gmail.com"}
	observerSecond := &Guest{id: "moneylen@gmail.com"}

	Artist.register(observerFirst)
	Artist.register(observerSecond)

	Artist.updateAvailability()
}
