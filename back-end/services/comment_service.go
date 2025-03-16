package services

import (
	"back-end/spec"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type CommentService struct {
	commentModel *mongo.Collection
}

func NewCommentService() *CommentService {
	client := NewMongoDB("mongodb://root:test12345@0.0.0.0:27017")
	return &CommentService{
		commentModel: client,
	}
}

func (c CommentService) Insert(model spec.CommentModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := c.commentModel.InsertOne(ctx, model)
	return err
}

func (c CommentService) Find() ([]spec.CommentModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := c.commentModel.Find(ctx, bson.M{},
		&options.FindOptions{Sort: bson.M{"created_date": -1}})
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	results := []spec.CommentModel{}
	for cur.Next(ctx) {
		var model spec.CommentModel
		if err = cur.Decode(&model); err != nil {
			return nil, err
		}
		results = append(results, model)
	}

	return results, nil
}

func NewMongoDB(connection string) *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))
	if err != nil {
		fmt.Sprintf("can't initialize MongoDB: %v", err)
	}

	collection := client.Database("test").Collection("comments")

	return collection
}
