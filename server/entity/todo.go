package entity

type Todo struct {
	ID     int    `gorm:"primary_key;not null"`
	Name   string `gorm:"type:varchar(50);not null"`
	Memo   string `gorm:"type:text;not null"`
	Status string `gorm:"type:tinyint(1);not null"`
}
