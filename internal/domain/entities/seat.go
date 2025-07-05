package entities

type Seat struct {
	ID     uint   `gorm:"primaryKey"`
	RoomID uint   `gorm:"not null;index"`
	Row    string `gorm:"size:5;not null"` // Ex: "A", "B"
	Number int    `gorm:"not null"`        // Ex: 1, 2, 3

	Room Room `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE;"`
}
