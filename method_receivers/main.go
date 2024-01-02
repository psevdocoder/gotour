package main

import "fmt"

type CustomStruct struct {
	On    bool
	Ammo  int
	Power int
}

func (c *CustomStruct) Shoot() bool {
	if !c.On || c.Ammo <= 0 {
		return false
	}
	c.Ammo--
	return true
}

func (c *CustomStruct) RideBike() bool {
	if !c.On || c.Ammo <= 0 {
		return false
	}
	c.Power--
	return true
}

func main() {
	myStruct := *new(CustomStruct)
	testStruct := &myStruct
	myStruct.On, myStruct.Ammo, myStruct.Power = true, 10, 10
	fmt.Println(testStruct.Ammo)
}
