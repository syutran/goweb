package main

import (
	"fmt"
)

type Human struct {
	name string
	age int 
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

func (h *Human) SayHi(){
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h *Human) Sing(lyrics string){
	fmt.Println("La la la, la la la", lyrics)
}

func (h *Human) Guzzle(beerStein string){
	fmt.Println("Guzzle guzzle guzzle... ", beerStein)
}

func (e *Employee) SayHi(){
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}

func (e *Employee) SpeedSalary(amount float32){
	e.money -= amount
}

type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface{
	SayHi()
	Sing(lyrics string)
	BorrowMoney(amount float32)
}

type ElderLyGent interface{
	SayHi()
	Sing(lyrics string)
	SpendSalary(amount float32)
}