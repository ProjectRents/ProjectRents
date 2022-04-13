package entities

type Book struct {
	Title     string `gorm:"type:varchar(50);not null"`
	Isbn      string `gorm:"type:varchar(50);not null"`
	Author    string `gorm:"type:varchar(50);not null"`
	
}
