package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
	Паттерн "цепочка вызовов" применяется в случаях, когда запрос можно обработать по принципу "сам или передать дальше".
Если объект не может обработать запрос, он передает его следующему объекту в цепочке. В конечном итоге запрос будет
обработан или не обработан вообще.
	Плюсы:
		- уменьшение зависимости между клиентом и обработчиками
		- возможность добавления новых обработчиков для обработки исключительных ситуаций
	Минусы:
		- не гарантируется, что запрос будет обработан
	Примером использования паттерна "цепочка вызовов" может служить обработка запросов веб-сервера. Например, в сервисе
реализована система уровней доступа, где каждый уровень обрабатывает запрос, если не может обработать, передает его дальше.
*/

type Handler interface {
	SetNext(Handler)
	Handle(string)
}

type ConcreteHandler1 struct {
	successor Handler
}

func (h *ConcreteHandler1) SetNext(handler Handler) {
	h.successor = handler
}

func (h *ConcreteHandler1) Handle(request string) {
	if request == "request1" {
		fmt.Println("ConcreteHandler1")
		return
	} else if h.successor != nil {
		h.successor.Handle(request)
	}
	return
}

type ConcreteHandler2 struct {
	successor Handler
}

func (h *ConcreteHandler2) SetNext(handler Handler) {
	h.successor = handler
}

func (h *ConcreteHandler2) Handle(request string) {
	if request == "request2" {
		fmt.Println("ConcreteHandler2")
		return
	} else if h.successor != nil {
		h.successor.Handle(request)
	}
	return
}

func main() {
	h1 := &ConcreteHandler1{}
	h2 := &ConcreteHandler2{}

	h1.SetNext(h2)

	h1.Handle("request1")
	h1.Handle("request2")
	h1.Handle("request3") // Выведено не будет
}
