package datasource

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contact struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name"`
	Number string             `bson:"number"`
	Group  primitive.ObjectID `bson:"group_id"`
}

type Group struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

type ContactGroup struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name"`
	Number string             `bson:"number"`
	Group  Group              `bson:"group_id"`
}
