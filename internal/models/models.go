package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CategoryID  primitive.ObjectID `bson:"category_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Image       []string           `bson:"image"`
	Price       int32              `bson:"price"`
	IsBlocked   bool               `bson:"is_blocked"`
}
