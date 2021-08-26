package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostData struct{
	Id primitive.ObjectID `bson:"_id, omitempty"`
	ImageUrl string
	Title string
	Description string
	Content string
	CreatorName string
	CreatedAt string
}
