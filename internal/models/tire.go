package models

type Tire struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Brand   string `json:"brand"`
	Size    string `json:"size"`
	Price   int    `json:"price"`
	InStock bool   `json:"in_stock"`
}
