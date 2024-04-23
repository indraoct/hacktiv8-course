package entities

import (
	"time"
)

type Orders struct {
	OrderId      uint      `gorm:"primaryKey" json:"orderId,omitempty"`
	CustomerName string    `gorm:"not null; unique; type:varchar(70)" json:"customerName,omitempty"`
	OrderedAt    time.Time `gorm:"type:timestamp" json:"orderAt,omitempty"`
	CreatedAt    time.Time `gorm:"type:timestamp" json:"createdAt,omitempty"`
	UpdatedAt    time.Time `gorm:"type:timestamp" json:"updatedAt,omitempty"`
	Items        []Items   `gorm:"foreignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"items,omitempty"`
}
