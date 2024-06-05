package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
	Паттерн "состояние" используется когда объект меняет свое поведение в зависимости от состояния. Это делает его
похожим на конечный автомат.
	Плюсы:
		- Избавляет от множества больших условных операторов, связанных с различными состояниями объекта.
		- Позволяет переключаться между состояниями объекта во время выполнения.
		- При добавлении нового состояния может потребоваться изменение только предыдущего состояния.
	Минусы:
		- Может привести к созданию большого количества классов.
		- Усложняет отладку программы из-за зависимости от текущего состояния.
	Примером использования паттерна "состояние" может быть видео плеер. У видео будет два состояния "воспроизведение" и "пауза".
	Когда видео находится в состоянии "воспроизведение", нажатие на кнопку "пробел" переведет его в состояние "пауза", и наоборот.
*/

// State определяет интерфейс для состояния.
type State interface {
	Handle(context *Context)
}

// PlayVideo представляет собой состояние воспроизведения видео.
type PlayVideo struct {
}

func (PlayVideo) Handle(context *Context) {
	fmt.Println("Setting pause")
	context.SetState(PauseVideo{})
}

// PauseVideo представляет собой состояние паузы видео.
type PauseVideo struct {
}

func (PauseVideo) Handle(context *Context) {
	fmt.Println("Starting video")
	context.SetState(PlayVideo{})
}

// Context представляет собой контекст, использующий состояние.
type Context struct {
	state State
}

// SetState устанавливает состояние для контекста.
func (c *Context) SetState(state State) {
	c.state = state
}

// Request выполняет запрос, который делегируется текущему состоянию.
func (c *Context) Request() {
	c.state.Handle(c)
}

func main() {
	context := Context{
		state: PlayVideo{},
	}

	context.Request() // Имитация нажатия пробела
	context.Request()
	context.Request()
	context.Request()
	context.Request()
}
