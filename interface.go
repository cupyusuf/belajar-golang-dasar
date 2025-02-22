package main

type hasName interface {
	getName() string
}

func sayHello(hasName hasName) {
	println("Hello", hasName.getName())
}

type person struct {
	name string
}

func (p person) getName() string {
	return p.name
}

type animal struct {
	name string
}

func (a animal) getName() string {
	return a.name
}

func main() {
	p := person{name: "Yusuf"}
	sayHello(p)

	a := animal{name: "Kucing"}
	sayHello(a)
}
