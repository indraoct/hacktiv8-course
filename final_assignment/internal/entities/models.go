package entities

import "time"

// User represents the User table
type User struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	Username    string        `json:"username"`
	Title       string        `json:"title"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	Age         int           `json:"age"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Comment     []Comment     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Photo       []Photo       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SocialMedia []SocialMedia `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Comment represents the Comment table
type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Comment   string    `json:"comment"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint
	PhotoID   uint
}

// Photo represents the Photo table
type Photo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint
}

// SocialMedia represents the Social Media table
type SocialMedia struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	UserID         uint
}
