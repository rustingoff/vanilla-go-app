package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductPGDB struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Price       float32
	Stock       uint
}

type ProductMGDB struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Price       float32            `json:"price,omitempty" bson:"price,omitempty"`
	Stock       uint               `json:"stock,omitempty" bson:"stock,omitempty"`
}
