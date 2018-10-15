package model

import "time"

type User struct {
	ID        int64     `gorm:"column:id;primary_key" json:"id" pk:"true" ai:"true"`
	Name      string    `gorm:"column:name" json:"name"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Pwd       string    `gorm:"column:pwd" json:"pwd"`
	Role      int       `gorm:"column:role" json:"role"`
	IsDeleted int       `gorm:"column:is_deleted" json:"is_deleted"`
	CreatedAt time.Time `gorm:"-" json:"created_at" `
	UpdatedAt time.Time `gorm:"-" json:"updated_at" ignore:"true"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "user"
}
