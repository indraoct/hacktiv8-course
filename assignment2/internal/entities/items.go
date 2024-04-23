package entities

import "time"

type Items struct {
	ItemId      uint      `gorm:"primaryKey" json:"itemId,omitempty"`
	ItemCode    string    `gorm:"not null; unique; type:varchar(35)" json:"itemCode,omitempty"`
	Description string    `gorm:"type:varchar(100)" json:"description,omitempty"`
	Quantity    int       `gorm:"type:varchar(100)" json:"quantity,omitempty"`
	CreatedAt   time.Time `gorm:"type:varchar(100)" json:"createdAt,omitempty"`
	UpdatedAt   time.Time `gorm:"type:varchar(100)" json:"updatedAt,omitempty"`
	OrderId     uint
}
