package entities

import "time"

// User represents the User table
type User struct {
	ID          uint          `gorm:"primaryKey" json:"id,omitempty"`
	Username    string        `json:"username,omitempty"`
	Title       string        `json:"title,omitempty"`
	Email       string        `gorm:"not null, unique" json:"email,omitempty"`
	Password    string        `json:"-"`
	Age         int           `json:"age,omitempty"`
	CreatedAt   time.Time     `json:"-"`
	UpdatedAt   time.Time     `json:"-"`
	Token       string        `json:"token,omitempty" gorm:"-"`
	Comment     []Comment     `json:"comment,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Photo       []Photo       `json:"photo,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SocialMedia []SocialMedia `json:"socialMedia,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Comment represents the Comment table
type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id,omitempty"`
	Message   string    `json:"message,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	UserID    uint
	PhotoID   uint
}

// Photo represents the Photo table
type Photo struct {
	ID        uint      `gorm:"primaryKey" json:"id,omitempty"`
	Caption   string    `json:"caption,omitempty"`
	PhotoURL  string    `json:"photo_url,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	UserID    uint
}

// SocialMedia represents the Social Media table
type SocialMedia struct {
	ID             uint      `gorm:"primaryKey" json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	SocialMediaURL string    `json:"social_media_url,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	UserID         uint
}
