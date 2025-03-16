package spec

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	CommentModel struct {
		ID          primitive.ObjectID `json:"_id" bson:"_id"`
		Text        string             `json:"text" bson:"text"`
		CreatedBy   string             `json:"created_by" bson:"created_by"`
		CreatedDate time.Time          `json:"created_date" bson:"created_date"`
	}
)
