package entities

type Items struct {
	ItemId      uint   `gorm:"primaryKey" json:"item_id"`
	ItemCode    string `gorm:"not null; unique; type:varchar(35)" json:"item_code"`
	Description string `gorm:"type:varchar(100)" json:"description"`
	Quantity    int    `gorm:"type:varchar(100)" json:"quantity"`
	OrderId     uint
}
