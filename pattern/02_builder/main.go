package pattern

import (
	"fmt"
)

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	Паттерн "строитель" необходим для создания экземпляров одной сущности с разными конфигурациями без необходимости
изменять код.
	Плюсы:
		- Разделение процесса создания экземпляра класса на несколько этапов, упрощая понимание и процесс отладки
		- Использование одной сущности без необходимости создания дополнительных
	Минусы:
		- Необходимость создания Builder для каждой конфигурации

    На практике паттерн можно использовать при разработке приложения для автосалона. Машина состоит из трех основных
компонентов: двигатель, цвет и трансмиссия. Таким образом, можно использовать одну сущность, а конфигураций может быть несколько
десятков.

*/

// Car представляет собой конечный продукт, который строит строитель.
type Car struct {
	Brand        string
	Model        string
	Year         int
	BodyStyle    string
	Transmission string
	Color        string
}

// CarBuilder представляет интерфейс строителя для создания автомобиля.
type CarBuilder interface {
	SetBrand(brand string) CarBuilder
	SetModel(model string) CarBuilder
	SetYear(year int) CarBuilder
	SetBodyStyle(bodyStyle string) CarBuilder
	SetTransmission(transmission string) CarBuilder
	SetColor(color string) CarBuilder
	Build() Car
}

// carBuilder реализует интерфейс CarBuilder.
type carBuilder struct {
	brand        string
	model        string
	year         int
	bodyStyle    string
	transmission string
	color        string
}

// NewCarBuilder создаёт новый экземпляр carBuilder.
func NewCarBuilder() CarBuilder {
	return &carBuilder{}
}

func (b *carBuilder) SetBrand(brand string) CarBuilder {
	b.brand = brand
	return b
}

func (b *carBuilder) SetModel(model string) CarBuilder {
	b.model = model
	return b
}

func (b *carBuilder) SetYear(year int) CarBuilder {
	b.year = year
	return b
}

func (b *carBuilder) SetBodyStyle(bodyStyle string) CarBuilder {
	b.bodyStyle = bodyStyle
	return b
}

func (b *carBuilder) SetTransmission(transmission string) CarBuilder {
	b.transmission = transmission
	return b
}

func (b *carBuilder) SetColor(color string) CarBuilder {
	b.color = color
	return b
}

func (b *carBuilder) Build() Car {
	return Car{
		Brand:        b.brand,
		Model:        b.model,
		Year:         b.year,
		BodyStyle:    b.bodyStyle,
		Transmission: b.transmission,
		Color:        b.color,
	}
}

func main() {
	builder := NewCarBuilder()
	car := builder.SetBrand("Toyota").
		SetModel("Camry").
		SetYear(2021).
		SetBodyStyle("Sedan").
		SetTransmission("Automatic").
		SetColor("Blue").
		Build()

	fmt.Printf("Car: %+v\n", car)
}
