package entities

import "time"

type Mahasiswa struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement"` // gorm tag jus for try, because field ID default is primary key
	Nim       string    `json:"nim" gorm:"not null;unique;size=10"`
	Name      string    `json:"name" gorm:"not null:size=100"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
