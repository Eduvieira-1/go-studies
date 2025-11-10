package main

import "fmt"

type Address struct {
	StreetAddress string
	Number        int
	City          string
	State         string
}

type Person interface {
	Disable()
}

type Company struct {
	Name string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func (company Company) Disable() {

}

func (c Client) Disable() {
	c.Active = false
	fmt.Printf("O client %s foi desativado", c.Name)
}

func Deactivation(person Person) {
	person.Disable()
}

func main() {
	eduardo := Client{
		Name:   "Eduardo",
		Age:    25,
		Active: true,
	}

	eduardo.Active = false
	eduardo.City = "Uberaba"
	eduardo.Disable()

	myCompany := Company{}

	Deactivation(myCompany)

	fmt.Printf("Name: %s, Age: %d, Active: %t", eduardo.Name, eduardo.Age, eduardo.Active)
}
