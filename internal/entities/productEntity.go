package entities

type ProductPGDB struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Price       float32
	Stock       uint
}
