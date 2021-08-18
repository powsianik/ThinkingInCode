package models

import "time"

type PostData struct{
	Id string `bson:"_id,omitempty"`
	ImageUrl string
	Title string
	Description string
	Content string
	CreatorName string
	CreatedAt time.Time
}
