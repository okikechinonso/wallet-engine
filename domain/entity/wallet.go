package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
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
	FirstName      string `json:"first_name" gorm:"not null" binding:"required" form:"first_name"`
	LastName       string `json:"last_name" gorm:"not null" binding:"required" form:"last_name"`
	Email          string `json:"email" gorm:"not null" binding:"required" form:"email"`
	Password       string `json:"password" gorm:"-" binding:"required" form:"password"`
	HashedPassword string `json:"-,omitempty" gorm:"not null"`
	Phone          string `json:"phone,omitempty"`
	UserWallet     Wallet `json:"user_wallet"`
}

type Wallet struct {
	Model
	WalletAddress string `json:"wallet_address"`
	UserID        string `json:"user_id"`
	Balance       int64  `json:"balance"`
	Active        bool   `json:"active" gorm:"default:false"`
}
