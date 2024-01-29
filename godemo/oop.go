package godemo

import (
	"errors"
	"fmt"
)

// Structs (Objects)
type Person struct {
	Name    string
	Age     int
	Address string
	Married bool
	Sex     string
}

// Move method allows a Person to change their address.
func (p *Person) Move(newAddress string) {
	p.Address = newAddress
	fmt.Printf("%s has moved to %s\n", p.Name, newAddress)
}

// GetMarried method updates a Person's marital status and, if applicable, their name prefix.
func (p *Person) GetMarried() error {
	if p.Age < 18 {
		return errors.New("person is too young to get married")
	}
	p.Married = true
	if p.Sex == "female" {
		p.Name = "Mrs. " + p.Name
	}
	fmt.Printf("%s has gotten married\n", p.Name)
	return nil
}

func ObjectOrientedProgramming() {

	// Demo here
}
