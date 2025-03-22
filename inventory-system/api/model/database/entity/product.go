package entity

type Product struct {
	ProductID          uint `gorm:"primaryKey"`
	ProductName        string
	ProductDescription string
	ProductMinStock    int32
}
