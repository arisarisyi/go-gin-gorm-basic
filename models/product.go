package models

type Product struct {
	BaseModel
	NamaProduct string `gorm:"type:varchar(300)" json:"nama_product"`
	Description string `gorm:"type:text" json:"description"`
}