package main

import "fmt"

type Handler interface {
	SetNext(handler Handler) Handler
	HandleRequest(request string)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) Handler {
	b.next = handler
	return handler
}

func (b *BaseHandler) HandleNext(request string) {
	if b.next != nil {
		b.next.HandleRequest(request)
	}
}

type BasicSupport struct {
	BaseHandler
}

func (b *BasicSupport) HandleRequest(request string) {
	if request == "pregunta básica" {
		fmt.Println("BasicSupport: Resolviendo la solicitud básica.")
	} else {
		fmt.Println("BasicSupport: No puedo manejar esto, pasando al siguiente nivel.")
		b.HandleNext(request)
	}
}

type TechnicalSupport struct {
	BaseHandler
}

func (t *TechnicalSupport) HandleRequest(request string) {
	if request == "problema técnico" {
		fmt.Println("TechnicalSupport: Resolviendo el problema técnico.")
	} else {
		fmt.Println("TechnicalSupport: No puedo manejar esto, pasando al siguiente nivel.")
		t.HandleNext(request)
	}
}

// Supervisor maneja solicitudes complejas o críticas.
type Supervisor struct {
	BaseHandler
}

func (s *Supervisor) HandleRequest(request string) {
	fmt.Println("Supervisor: Evaluando y resolviendo la solicitud crítica.")
}

func main() {
	basic := &BasicSupport{}
	technical := &TechnicalSupport{}
	supervisor := &Supervisor{}

	basic.SetNext(technical).SetNext(supervisor)

	fmt.Println("Solicitud 1: pregunta básica")
	basic.HandleRequest("pregunta básica")

	fmt.Println("\nSolicitud 2: problema técnico")
	basic.HandleRequest("problema técnico")

	fmt.Println("\nSolicitud 3: problema complejo")
	basic.HandleRequest("problema complejo")
}
