package model

type Collection struct {
	Model
	Title string `json:"title" gorm:"size:50;not null;unique;comment:'收藏标题'"`
	CollectionItems []CollectionItem `json:"collection_items"`
}