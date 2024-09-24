package main

import (
	"errors"
	"fmt"
	"strconv"
)

//Model

type Trans struct {
	Name   string
	Weight int
}
type Wendaphos struct {
	Name   string
	Weight int
}
type M100FT struct {
	Name   string
	Weight int
}
type ISP struct {
	Name   string
	Weight int
}
type Karagenan struct {
	Name   string
	Weight int
}
type M100F struct {
	Name   string
	Weight int
}

func NewTG(name string, amount int) *Trans {
	return &Trans{
		Name:   name,
		Weight: amount,
	}
}
func NewWP(name string, amount int) *Wendaphos {
	return &Wendaphos{
		Name:   name,
		Weight: amount,
	}
}
func NewM100FT(name string, amount int) *M100FT {
	return &M100FT{
		Name:   name,
		Weight: amount,
	}
}
func NewISP(name string, amount int) *ISP {
	return &ISP{
		Name:   name,
		Weight: amount,
	}
}
func NewKaragenan(name string, amount int) *Karagenan {
	return &Karagenan{
		Name:   name,
		Weight: amount,
	}
}
func NewM100F(name string, amount int) *M100F {
	return &M100F{
		Name:   name,
		Weight: amount,
	}
}

// Interface
type Weight interface {
	readWeight() int
}

func checkWeight(item Weight) {
	fmt.Printf("Weight of %v is %d grams", item.readWeight())
}

//TODO: Implement read name

// Read weight
func (tg Trans) readWeight() int {
	return tg.Weight
}
func (wp Wendaphos) readWeight() int {
	return wp.Weight
}
func (ft M100FT) readWeight() int {
	return ft.Weight
}
func (isp ISP) readWeight() int {
	return isp.Weight
}
func (kar Karagenan) readWeight() int {
	return kar.Weight
}
func (mr M100F) readWeight() int {
	return mr.Weight
}

func main() {
	//Data input
	dataTG, err := GetInputInt("TG weight: ")
	if err != nil {
		fmt.Println(err)
	}
	dataWP, err := GetInputInt("WP weight: ")
	if err != nil {
		fmt.Println(err)
	}
	dataFT, err := GetInputInt("M100FT weight: ")
	if err != nil {
		fmt.Println(err)
	}
	dataISP, err := GetInputInt("ISP weight: ")
	if err != nil {
		fmt.Println(err)
	}
	dataKar, err := GetInputInt("Karagenan weight: ")
	if err != nil {
		fmt.Println(err)
	}
	dataMR, err := GetInputInt("M100F weight: ")
	if err != nil {
		fmt.Println(err)
	}
	//Create Initial Instance
	tg := NewTG("Trans", 0)
	wp := NewWP("Wendaphos", 0)
	ft := NewM100FT("M100FT", 0)
	isp := NewISP("ISP", 0)
	kar := NewKaragenan("Karagenan", 0)
	mr := NewM100F("MR", 0)

	//Read value
	checkWeight(tg)
	checkWeight(wp)
	checkWeight(ft)
	checkWeight(isp)
	checkWeight(kar)
	checkWeight(mr)
}

func GetInputInt(prompt string) (int, error) {
	var x string
	var value int
	fmt.Println(prompt)
	_, err := fmt.Scanln(&x)
	if err != nil {
		return 0, err
	} else if x == "" {
		return 0, errors.New("value must not be empty")
	}
	value, err = strconv.Atoi(x)
	return value, nil
}
