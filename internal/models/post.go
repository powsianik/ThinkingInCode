package models

import "time"

type PostData struct{
	ImageUrl string
	Title string
	Content string
	CreatorName string
	CreatedAt time.Time
}
