package entity

import (
	"gorm.io/gorm"
	"time"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Model struct {
	ID        string    `sql:"type:uuid; default:uuid_generate_v4();size:100; not null"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
func (u *Model) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	if u.ID == "" {
		err = errors.New("can't save invalid data")
	}
	return
}
type User struct {
	Model
	FirstName      string     `json:"first_name" gorm:"not null" binding:"required" form:"first_name"`
	LastName       string     `json:"last_name" gorm:"not null" binding:"required" form:"last_name"`
	Email          string     `json:"email" gorm:"not null" binding:"required" form:"email"`
	Password       string     `json:"password" gorm:"-" binding:"required" form:"password"`
	HashedPassword string     `json:"-,omitempty" gorm:"not null"`
	Phone          string     `json:"phone,omitempty"`
	Bio            string     `json:"bio,omitempty"`
	Followers      []Follower `gorm:"many2many:user_followers" json:"followers,omitempty"`
	Verified       bool       `json:"verified" gorm:"default:false"`
	Active         bool       `json:"active" gorm:"default:false"`
	Blogs          []Blog     `gorm:"-" json:"blogs,omitempty"`
	ImageUrl       string     `json:"image_url,omitempty"`
	Profession     string     `json:"profession,omitempty"`
}

type Follower struct {
	Model
	FollowerID string `json:"follower_id" gorm:"foreignkey:User(id)"`
	FollowedID string `json:"followed_id" gorm:"foreignkey:User(id)"`
}
