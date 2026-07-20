package model

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"size:250;unique;not null"`
	Password string `gorm:"size:250;not null"`
	Urls     []URL  `gorm:"foreignKey:UserID"`
}
