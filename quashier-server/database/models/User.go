package models

import (
	"flexusma.de/quashier-server/util"
	"gorm.io/gorm"
)

// Represents a user that is saved in quashier database
type User struct {
	gorm.Model
	Name    string
	balance float64
}

func NewUser(name string) User {
	var u = User{balance: 0, Name: name}
	util.DBI.Create(&u)
	return u
}

func (u User) GetBalance() float64 {
	return u.balance
}

func (u User) SubBalance(val float64) float64 {
	util.DBI.Model(&u).Update("balance", u.balance-val)
	return u.balance
}

func (u User) AddBalance(val float64) float64 {
	util.DBI.Model(&u).Update("balance", u.balance+val)
	return u.balance
}
