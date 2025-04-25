package models

type Tire struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Brand  string `json:"brand"`
	Size   string `json:"size"`
	Price  int    `json:"price"`
	UserID uint   `json:"user_id"`
}
