package main

import "fmt"

// Handler define la interfaz para procesar una solicitud o pasarla al siguiente manejador.
type Handler interface {
	SetNext(handler Handler) Handler
	HandleRequest(request string)
}

// BaseHandler proporciona la estructura común para los manejadores.
type BaseHandler struct {
	next Handler
}

// SetNext configura el siguiente manejador en la cadena.
func (b *BaseHandler) SetNext(handler Handler) Handler {
	b.next = handler
	return handler
}

// HandleNext pasa la solicitud al siguiente manejador si existe.
func (b *BaseHandler) HandleNext(request string) {
	if b.next != nil {
		b.next.HandleRequest(request)
	}
}

// BasicSupport maneja solicitudes básicas.
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

// TechnicalSupport maneja solicitudes técnicas.
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

// main configura la cadena y envía solicitudes.
func main() {
	// Crear los manejadores
	basic := &BasicSupport{}
	technical := &TechnicalSupport{}
	supervisor := &Supervisor{}

	// Configurar la cadena: Basic -> Technical -> Supervisor
	basic.SetNext(technical).SetNext(supervisor)

	// Enviar solicitudes a la cadena
	fmt.Println("Solicitud 1: pregunta básica")
	basic.HandleRequest("pregunta básica")

	fmt.Println("\nSolicitud 2: problema técnico")
	basic.HandleRequest("problema técnico")

	fmt.Println("\nSolicitud 3: problema complejo")
	basic.HandleRequest("problema complejo")
}
