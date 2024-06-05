package main

import (
	"fmt"
	"time"
)

/*
	Паттерн "посетитель" позволяет добавлять новые операции к объектам без изменения их структур. Вместо этого новая
операция добавляется в структуру посетителя.
	Плюсы:
		- позволяет добавлять новые операции к объектам без изменения их структур
		- позволяет собрать родственные операции в одной структуре
		- позволяет сделать операции над объектами разных структур
	Минусы:
		- усложняет структуру программы
		- усложняет добавление новых элементов, т.к. требуется изменение всех посетителей
	Примером использования паттерна "посетитель" может быть работа с документами. Документы могут быть разных типов, но
все они могут быть посещены менеджером, который может совершать разные действия над документами. Также, могут быть и
разные посетители, которых можно добавлять в программу без изменения структуры документов.
*/

// Element

type Document interface {
	Accept(Visitor)
}

type Application struct {
	Topic string
	Date  time.Time
}

func (a *Application) Accept(v Visitor) {
	v.VisitApplication(a)
}

type Contract struct {
	Party1   string
	Party2   string
	Duration time.Duration
}

func (c *Contract) Accept(v Visitor) {
	v.VisitContract(c)
}

// Visitor

type Visitor interface {
	VisitApplication(*Application)
	VisitContract(*Contract)
}

type Manager struct {
	Name string
}

func (m *Manager) VisitApplication(a *Application) {
	fmt.Printf("Manager %s is visiting application with topic: %s\n", m.Name, a.Topic)
}

func (m *Manager) VisitContract(c *Contract) {
	fmt.Printf("Manager %s is visiting contract between: %s and %s for %d years\n",
		m.Name,
		c.Party1,
		c.Party2,
		int(c.Duration.Hours()/24/365))
}

func main() {
	app := &Application{
		Topic: "Letter of resignation",
		Date:  time.Now(),
	}

	contract := &Contract{
		Party1:   "WB Tech",
		Party2:   "OZON Tech",
		Duration: time.Hour * 24 * 365 * 2,
	}

	manager := &Manager{Name: "Konovalov A."}

	app.Accept(manager)
	contract.Accept(manager)
}
