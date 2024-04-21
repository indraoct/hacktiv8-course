package entities

import "time"

type Orders struct {
	OrderId      uint      `gorm:"primaryKey" json:"order_id"`
	CustomerName string    `gorm:"not null; unique; type:varchar(70)" json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Items        []Items   `gorm:"foreignKey:ItemRefer" json:"items"`
}
