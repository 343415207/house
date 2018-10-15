package model

import (
	"time"
)

type COrder struct {
	ID          int64     `gorm:"column:id;primary_key" json:"id"`
	CName       string    `gorm:"column:c_name" json:"c_name"`
	CPhone      string    `gorm:"column:c_phone" json:"c_phone"`
	CSex        int       `gorm:"column:c_sex" json:"c_sex"`
	PID         int64     `gorm:"column:p_id" json:"p_id"`
	IsFirst     int       `gorm:"column:is_first" json:"is_first"`
	EscorteType int       `gorm:"column:escorte_type" json:"escorte_type"`
	OStatus     int       `gorm:"column:o_status" json:"o_status"`
	IsDeleted   int       `gorm:"column:is_deleted" json:"is_deleted"`
	AccessTime  time.Time `gorm:"column:access_time" json:"access_time"`
	Desc        string    `gorm:"column:desc" json:"desc"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (c *COrder) TableName() string {
	return "c_order"
}
