package main

import "fmt"

/*
	Реализовать паттерн «команда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Паттерн "команда" используется для инкапсуляции запроса в виде объекта, что позволяет параметризовать клиентов с
различными запросами, организовать очередь или протоколировать запросы.
	Плюсы:
		- возможность создания очереди запросов
		- возможность протоколирования запросов
	Минусы:
		- увеличение количества структур за счет необходимости описания реализации каждой команды
	Примером использования паттерна "команда" может служить реализация пультов управления для умного дома, где каждая
кнопка пульта является командой, которая управляет устройством (свет, кондиционер, телевизор и т.д.)
*/

// Command

type Command interface {
	Execute()
}

// Receiver

type Light struct {
	On bool
}

func (l *Light) Switch() {
	if l.On {
		l.On = false
		fmt.Println("Light switched off")
	} else {
		l.On = true
		fmt.Println("Light switched on")
	}
}

type FlipUpCommand struct {
	Light *Light
}

func (f *FlipUpCommand) Execute() {
	f.Light.Switch()
}

type FlipDownCommand struct {
	Light *Light
}

func (f *FlipDownCommand) Execute() {
	f.Light.Switch()
}

// Invoker

type Switch struct {
}

func (s *Switch) Act(cmd Command) {
	cmd.Execute()
}

func main() {
	light := &Light{}
	switchUp := &FlipUpCommand{Light: light}
	switchDown := &FlipDownCommand{Light: light}
	s := &Switch{}

	s.Act(switchUp)
	s.Act(switchDown)
}
