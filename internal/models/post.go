package models

import "time"

type PostData struct{
	ImageUrl string
	Title string
	Description string
	Content string
	CreatorName string
	CreatedAt time.Time
}
