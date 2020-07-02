package main

import (
	"fmt"
)


type Sgu struct{
	Led string
	Sirene int
	Voltage int
}

type SguAction interface {
	HornAction(variant string)
	LedAction(variant ...string)
	SireneAction(variant string)
}

func (s *Sgu) HornAction(variant string) {
	fmt.Println("Горн делает:", variant)
}

func (s *Sgu) LedAction(variant ...string) {
	fmt.Println("Балка мигает:", variant)
}

func (s *Sgu) SireneAction(variant string) {
	fmt.Println("Сирена делает:", variant)
}

func NewPoliceCar() SguAction {
	return &Sgu{
		Led: "red&blue", //color led bar
		Sirene: 120, //watt
		Voltage: 12, //volt
	}
}

func main()  {
	myPoliceCar, ok := NewPoliceCar().(*Sgu)
	if ok != true {
		fmt.Println("Sorry, i have problems")
	}
	fmt.Println(myPoliceCar, ok)

	myPoliceCar.HornAction("Кря-Кря!!!")
	myPoliceCar.LedAction("Мигает вперед и назад!!!", myPoliceCar.Led)
	myPoliceCar.SireneAction("Виуууу,Виуууу!!!")

}