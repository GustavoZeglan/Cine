package entities

type Room struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255;not null"`
	Capacity int    `gorm:"not null"`
}
