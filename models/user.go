package models

type User struct {
	BaseModel
	FullName string `gorm:"varchar(300);not null" json:"nama_lengkap"`
	Email    string `gorm:"varchar(300);not null" json:"email"`
	Password string `gorm:"varchar(300);not null" json:"password"`
	Role     string `gorm:"varchar(30);not null" json:"role"`
}