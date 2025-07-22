package entities

import "time"

type Session struct {
	ID      uint      `gorm:"primaryKey"`
	MovieID uint      `gorm:"not null;index"`
	RoomID  uint      `gorm:"not null;index"`
	StartAt time.Time `gorm:"not null"`

	Movie Movie `gorm:"foreignKey:MovieID;constraint:OnDelete:CASCADE;"`
	Room  Room  `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE;"`
}
