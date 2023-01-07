package models

type AddCartJson struct {
	Id int
}

type ModifyCartJson struct {
	Op     string
	Id     int
	Amount int
}
