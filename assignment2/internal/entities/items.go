package entities

type Items struct {
	ItemId      uint   `gorm:"primaryKey" json:"item_id"`
	ItemCode    string `gorm:"not null; unique; type:varchar(35)" json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	ItemRefer   uint
}
