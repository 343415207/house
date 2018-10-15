package model

import "time"

type Project struct {
	ID        int64     `gorm:"column:id;primary_key" json:"id"`
	PName     string    `gorm:"column:p_name" json:"p_name"`
	Desc      string    `gorm:"column:desc" json:"desc"`
	URL       string    `gorm:"column:url" json:"url"`
	PType     int       `gorm:"column:p_type" json:"p_type"`
	IsDeleted int       `gorm:"column:is_deleted" json:"is_deleted"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (p *Project) TableName() string {
	return "project"
}
