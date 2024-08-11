package factory

import "fmt"

// The Factory Method defines a method, which should be used for
//  creating objects instead of using a direct constructor call (new operator).
//   Subclasses can override this method to change the class of objects that will be created.

// Factory method is a creational design pattern
//
//	which solves the problem of creating product objects without specifying their concrete classes.
type IGun interface {
	setName(string)
	setPower(int)
	getPower() int
	getName() string
}

// Concrete prouct
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name

}
func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

func (g *Gun) getName() string {
	return g.name

}

type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun{
			name:  "Musket",
			power: 123,
		},
	}
}

type AK47 struct {
	Gun
}

func newGun() IGun {
	return &AK47{
		Gun{
			name:  "AK$&",
			power: 47,
		},
	}
}

// Factory method

func getType(types string) IGun {
	if types == "AK47" {
		return newGun()
	} else {
		return newMusket()
	}

}

func InitDemo() {

}

func DemoFactory() {
	ak46 := getType("AK47")
	fmt.Println(" ", ak46.getName(), ak46.getPower())
	musket := getType("Musket")
	fmt.Println(" ", musket.getName(), musket.getPower())

}
