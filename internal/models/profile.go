package models

type Profile struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"not null" json:"first_name" binding:"required"`
	LastName  string `gorm:"not null" json:"last_name" binding:"required"`
	Birthdate string `gorm:"not null" json:"birthdate" binding:"required"`
	Email     string `gorm:"unique;not null" json:"email" binding:"required,email"`
}

func (Profile) TableName() string {
	return "public.profiles"
}
