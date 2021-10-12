package entity

type Product struct {
	ID          int    `gorm:"primary_key;not null"`
	ProductName string `gorm:"type:varchar(200);not null"`
	Memo        string `gorm:"type:varchar(400);not null"`
	Status      string `gorm:"type:char(2);not null"`
}
