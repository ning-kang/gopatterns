package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address
}

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)
	return &result
}

var mainOfficeEmployeePrototype = Employee{
	"",
	Address{
		0, "123 East RD", "London",
	},
}

var auxOfficeEmployeePrototype = Employee{
	"",
	Address{
		0, "66 West RD", "London",
	},
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

// prototype factory
func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOfficeEmployeePrototype, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOfficeEmployeePrototype, name, suite)
}

func main() {
	john := NewMainOfficeEmployee("John", 100)
	jane := NewAuxOfficeEmployee("Jane", 200)
	fmt.Println(john, jane)
}
