package domain

type ProductPrice struct {
	ProductID int `json:"ProductId" bson:"ProductId" gorm:"foreign_key"`
	Currency  string
	Value     float64
}
