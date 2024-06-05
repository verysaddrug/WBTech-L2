package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
	Паттер "стратегия" используется для определения семейства алгоритмов, инкапсуляции каждого из них и обеспечения
их взаимозаменяемости. Он позволяет изменять алгоритмы независимо от клиентов, которые ими пользуются.
	Плюсы:
		- Позволяет добавлять и изменять алгоритмы, не изменяя код контекста.
		- Все алгоритмы наследуются от одного интерфейса, что упрощает их использование и делает их универсальными.
	Минусы:
		- Усложняет код программы.
		- Необходимо создавать отдельную структуру для каждого алгоритма.
	Примером использования паттерна "стратегия" может служить реализация приложения навигации, где разные алгоритмы
используются для построения маршрута в зависимости от выбранного транспорта (пешком, на автомобиле, на автобусе).
*/

// Strategy определяет интерфейс стратегии.
type Strategy interface {
	Algorithm(num1 int, num2 int) int
}

// OperationAdd представляет собой конкретную стратегию сложения.
type OperationAdd struct {
}

func (OperationAdd) Algorithm(num1 int, num2 int) int {
	return num1 + num2
}

// OperationSubstract представляет собой конкретную стратегию вычитания.
type OperationSubstract struct {
}

func (OperationSubstract) Algorithm(num1 int, num2 int) int {
	return num1 - num2
}

// OperationMultiply представляет собой конкретную стратегию умножения.
type OperationMultiply struct {
}

func (OperationMultiply) Algorithm(num1 int, num2 int) int {
	return num1 * num2
}

// Context представляет собой контекст, использующий стратегию.
type Context struct {
	strategy Strategy
}

// SetStrategy устанавливает стратегию для контекста.
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// ExecuteOperation выполняет операцию с использованием установленной стратегии.
func (c *Context) ExecuteOperation(num1 int, num2 int) int {
	return c.strategy.Algorithm(num1, num2)
}

func main() {
	context := &Context{}

	context.SetStrategy(OperationAdd{})
	fmt.Println("1 + 3 = ", context.ExecuteOperation(1, 3))

	context.SetStrategy(OperationSubstract{})
	fmt.Println("1 - 3 = ", context.ExecuteOperation(1, 3))

	context.SetStrategy(OperationMultiply{})
	fmt.Println("3 * 3 = ", context.ExecuteOperation(3, 3))
}
