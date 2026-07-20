package model

type URL struct {
	ID       int      `gorm:"primaryKey;autoIncrement"`
	UserID   int      `gorm:"not null"`
	URL      string   `gorm:"size:250;not null"`
	Period   int      `gorm:"not null"`
	Statuses []Status `gorm:"foreignKey:URLID"`
}
