package models

import (
	"flexusma.de/quashier-server/util"
	"gorm.io/gorm"
)

type PaymentMethod int8

const (
	Cash PaymentMethod = iota
	Balance
)

type Purchase struct {
	gorm.Model
	ItemName  string
	ItemPrice float64
	ItemType  PaymentMethod
	UserID    uint
}

func NewPurchase(name string) Purchase {
	var u = Purchase{}
	util.DBI.Create(&u)
	return u
}
