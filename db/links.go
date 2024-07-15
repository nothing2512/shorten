package db

import "time"

type Link struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Original  string    `json:"original"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime,column:updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

func (*Link) TableName() string {
	return "links"
}
