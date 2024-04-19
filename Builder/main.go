package main

import "fmt"

type Person struct {
	// address
	StreetAddress, Postcode, City string

	// job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		&Person{},
	}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder { // fluent
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder { // fluent
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder { // fluent
	b.person.Postcode = postcode
	return b
}

func (b *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	b.person.CompanyName = companyName
	return b
}

func (b *PersonJobBuilder) As(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}

func (b *PersonJobBuilder) Earning(income int) *PersonJobBuilder {
	b.person.AnnualIncome = income
	return b
}

func (b *PersonBuilder) Build() *Person { // yield the person after building
	return b.person
}

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives(). // switch to PersonAddressBuilder
		At("123 London Road").
		In("London").
		WithPostcode("SW123").
		Works(). // switch to PersonJobBuilder
		At("Fabrikam").
		As("Programmer").
		Earning(1423650)
	person := pb.Build()
	fmt.Println(person)
}
