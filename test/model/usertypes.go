package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FNAME    string             `bson:"fname,omitempty" json:"fname,omitempty"`
	LNAME    string             `bson:"lname,omitempty" json:"lname,omitempty"`
	AGE      int32              `bson:"age,omitempty" json:"age,omitempty"`
	BDATE    time.Time          `bson:"Bdate,omitempty" json:"Bdate,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
