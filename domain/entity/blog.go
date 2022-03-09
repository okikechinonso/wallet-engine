package entity

type Blog struct {
	ID       string    `bson:"blog_id"`
	UserID   string    `json:"user_id" bson:"user_id" `
	Author   string    `json:"author" bson:"author"`
	Title    string    `json:"title" gorm:"not null"  bson:"title" binding:"required" form:"title"`
	Body     string    `json:"body" gorm:"not null" bson:"body" binding:"required" form:"body"`
	Images   []Image   `gorm:"many2many:blog_images" bson:"images" json:"images,omitempty"`
	Comments []Comment `json:"comments,omitempty" bson:"comments,omitempty"`
	Likes    []Like    `json:"like,omitempty"`
}

type Comment struct {
	ID       string
	BlogID   string `json:"blog_id" bson:"blog_id" binding:"required" gorm:"foreignkey:Blog(id)"`
	UserID   string `json:"user_id" bson:"user_id" binding:"required" gorm:"foreignkey:User(id)"`
	UserName string `json:"user_name" bson:"user_name" binding:"required" gorm:"not null"`
	Body     string `json:"body" bson:"body" gorm:"not null" binding:"required" form:"body"`
}

type Image struct {
	BlogID   string `bson:"blog_id" binding:"required"`
	ImageUrl string `gorm:"not null" binding:"reqiured"`
}
type Like struct {
	UserID string`bson:"blog_id" binding:"required"`
	Status bool `bson:"status" binding:"required"`
}
