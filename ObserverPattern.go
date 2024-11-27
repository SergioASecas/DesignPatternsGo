package main

import "fmt"

type Observer interface {
	Update(message string)
}

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyAll()
}

type NewsChannel struct {
	observers []Observer
	message   string
}

func (n *NewsChannel) Register(observer Observer) {
	n.observers = append(n.observers, observer)
}

func (n *NewsChannel) Unregister(observer Observer) {
	for i, obs := range n.observers {
		if obs == observer {
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
			break
		}
	}
}

func (n *NewsChannel) NotifyAll() {
	for _, observer := range n.observers {
		observer.Update(n.message)
	}
}

func (n *NewsChannel) AddMessage(message string) {
	n.message = message
	n.NotifyAll()
}

type User struct {
	name string
}

func (u *User) Update(message string) {
	fmt.Printf("Usuario %s recibió el mensaje: %s\n", u.name, message)
}

func main() {

	newsChannel := &NewsChannel{}

	user1 := &User{name: "Alice"}
	user2 := &User{name: "Bob"}
	user3 := &User{name: "Charlie"}

	newsChannel.Register(user1)
	newsChannel.Register(user2)
	newsChannel.Register(user3)

	newsChannel.AddMessage("¡Se lanzó una nueva funcionalidad!")

	newsChannel.Unregister(user1)

	newsChannel.AddMessage("¡Nueva actualización disponible!")

}
