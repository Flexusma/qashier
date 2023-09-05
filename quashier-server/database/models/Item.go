package models

import (
	"flexusma.de/quashier-server/util"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name  string
	Code  string
	Price float64
	Image string
}

func NewItem(item Item) Item {
	util.DBI.Create(&item)
	return item
}

// Get all items that exist
func AllItems(items []Item) *gorm.DB {
	return util.DBI.Find(&items)
}
