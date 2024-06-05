package main

import "fmt"

/*
Подсистемы: DVDPlayer, SoundSystem и Projector представляют собой различные компоненты домашнего кинотеатра.
Фасад: HomeTheaterFacade предоставляет упрощенный интерфейс для управления всеми компонентами домашнего кинотеатра.
Методы фасада: WatchMovie и EndMovie обеспечивают простой и интуитивно понятный способ взаимодействия с системой.
*/

/*
	Паттерн "фасад" необходим для упрощения взаимодействия с системой. Например, клиент не знает о существовании отдельных
компонентов компьютера. Ему важно включить компьютер, поработать на нём и выключить.
	Примером реального использования паттерна Фасад может быть авторизация пользователя. Клиент вводит логин и пароль,
нажимает кнопку "Войти". За простым интерфейсом и одним действием от него скрываются многие процессы работы с
введенными данными - поиск аккаунта в базе дынных, хэширование пароля и проверка его с указанным в базе, создаение
куки или JWT, авторизация, перенаправление на следующую страницу.
	К плюсам можно отнести:
		- Упрощение взаимодействия с системой путем сокрытия сложной логики
		- Облегчение поддержки кода с возможностью изменения реализации системы без изменения интерфейса клиента
	К минусам можно отнести:
		- Может привести к нарушению принципа единой ответственности, если фасад выполняет слишком много функций
		- Могут возникнуть трудности с отладкой кода, т.к. основная логика спрятана за фасадом
*/

// Подсистема 1: DVD проигрыватель
type DVDPlayer struct{}

func (d *DVDPlayer) On() {
	fmt.Println("DVD Player is On")
}

func (d *DVDPlayer) Play(movie string) {
	fmt.Printf("Playing movie: %s\n", movie)
}

func (d *DVDPlayer) Off() {
	fmt.Println("DVD Player is Off")
}

// Подсистема 2: Звук
type SoundSystem struct{}

func (s *SoundSystem) On() {
	fmt.Println("Sound System is On")
}

func (s *SoundSystem) SetVolume(volume int) {
	fmt.Printf("Setting volume to %d\n", volume)
}

func (s *SoundSystem) Off() {
	fmt.Println("Sound System is Off")
}

// Подсистема 3: Проектор
type Projector struct{}

func (p *Projector) On() {
	fmt.Println("Projector is On")
}

func (p *Projector) Off() {
	fmt.Println("Projector is Off")
}

// Фасад
type HomeTheaterFacade struct {
	dvdPlayer   *DVDPlayer
	soundSystem *SoundSystem
	projector   *Projector
}

func NewHomeTheaterFacade(dvd *DVDPlayer, sound *SoundSystem, proj *Projector) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		dvdPlayer:   dvd,
		soundSystem: sound,
		projector:   proj,
	}
}

func (ht *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")
	ht.projector.On()
	ht.soundSystem.On()
	ht.soundSystem.SetVolume(5)
	ht.dvdPlayer.On()
	ht.dvdPlayer.Play(movie)
}

func (ht *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting movie theater down...")
	ht.dvdPlayer.Off()
	ht.soundSystem.Off()
	ht.projector.Off()
}

func main() {
	dvd := &DVDPlayer{}
	sound := &SoundSystem{}
	projector := &Projector{}

	homeTheater := NewHomeTheaterFacade(dvd, sound, projector)

	homeTheater.WatchMovie("Inception")
	homeTheater.EndMovie()
}
