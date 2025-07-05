package entities

import "time"

type Reservation struct {
	ID        uint   `gorm:"primaryKey"`
	SessionID uint   `gorm:"not null;index"`
	SeatID    uint   `gorm:"not null;index"`
	UserID    string `gorm:"size:255;not null"`
	CreatedAt time.Time

	Session Session `gorm:"foreignKey:SessionID;constraint:OnDelete:CASCADE;"`
	Seat    Seat    `gorm:"foreignKey:SeatID;constraint:OnDelete:CASCADE;"`
}
