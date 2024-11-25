package auth

import (
	"database/sql"
	"github.com/RikiIhsan/lib/database"
	"gorm.io/gorm"
	"time"
)

type (
	User struct {
		ID         uint64    `gorm:"autoIncrement;primaryKey"`
		Username   string    `gorm:"unique;size:10"`
		Name       string    `gorm:"size:255;unique;not null"`
		Password   string    `gorm:"size:255"`
		Email      string    `gorm:"size:255;unique;not null"`
		CreatedAt  time.Time `gorm:"autoCreateTime"`
		UpdatedAt  time.Time `gorm:"autoUpdateTime"`
		Active     bool      `gorm:"default:false"`
		ActivateAt sql.NullTime
	}
	Key struct {
		Token      string    `gorm:"size:120;not null;primaryKey"`
		Created_at time.Time `gorm:"autoCreateTime"`
		Expired_at time.Time `gorm:"not null"`
		Details    string    `gorm:"not null"`
	}
)

func (User) TableName() string { return "asrs_users" }
func (Key) TableName() string  { return "asrs_key" }
func (u *User) Create() (tx *gorm.DB) {
	return database.Session["sdp"].DB.Create(&u)
}
func (u *User) FindByUsername() (tx *gorm.DB) {
	return database.Session["sdp"].DB.Where("username = ?", u.Username).First(&u)
}
func (k *Key) Create() (tx *gorm.DB) {
	return database.Session["sdp"].DB.Create(&k)
}
func (k *Key) Find() (tx *gorm.DB) {
	return database.Session["sdp"].DB.First(&k)
}
