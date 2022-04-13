package entities

type Rent struct {
	Return_date string `gorm:"type:date;not null"`
}
