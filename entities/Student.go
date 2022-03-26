package entities

import "time"

type Student struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement"` // gorm tag jus for try, because field ID default is primary key
	Nim       string    `json:"nim" gorm:"not null;unique;size:10" validate:"required"`
	Name      string    `json:"name" gorm:"not null;size:100" validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" validate:"required"`
}
