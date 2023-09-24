package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
}

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CategoryID  primitive.ObjectID `bson:"category_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Brand       string             `bson:"brand"`
	Description string             `bson:"description,omitempty"`
	Image       []string           `bson:"image,omitempty"`
	Price       int64              `bson:"price,omitempty"`
	IsBlocked   bool               `bson:"is_blocked,omitempty"`
}

// Not a db model, it's  used for passing page info
type PageInfo struct {
	Skip  int64
	Limit int64
}
