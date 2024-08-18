package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// User represents a user model.
// swagger:model User
type User struct {
	ID uuid.UUID `json:"-" gorm:"type:uuid;primary_key"`
	// Email bersifat unik dan tidak boleh kosong
	Email string `json:"-" gorm:"type:varchar(255);unique;not null"`
	// Name tidak boleh kosong
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	// Role merupakan enum yang berisi "admin" dan "user" dan not null
	Role string `json:"role" gorm:"type:varchar(255);not null"`
	// Password disimpan dalam bentuk hash
	Password string `json:"-" gorm:"type:text;not null"`
	// OTP digunakan untuk konfirmasi email
	OTP string `json:"-" gorm:"type:varchar(5)"`
	// IsVerified menandakan apakah email sudah diverifikasi
	IsVerified bool `json:"-" gorm:"type:boolean;default:false"`
	// Provinsi tempat tinggal user
	Province string `json:"province" gorm:"type:varchar(255)"`
	// Kota tempat tinggal user
	City string `json:"city" gorm:"type:varchar(255)"`
	// CreatedAt menandakan waktu user dibuat
	CreatedAt time.Time `json:"-" gorm:"type:timestamp without time zone;default:now()"`
	// UpdatedAt menandakan waktu user terakhir diupdate
	UpdatedAt time.Time `json:"-" gorm:"type:timestamp without time zone;default:now()"`
	// DeletedAt menandakan waktu user dihapus
	DeletedAt gorm.DeletedAt `json:"-" gorm:"type:timestamp without time zone"`
}

func (u User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate() (err error) {
	// Generate id dengan versi 6
	u.ID = uuid.New()
	return
}

type UserCreate struct {
	Email           string `json:"email" binding:"required,email"`
	Name            string `json:"name" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	Province        string `json:"province" binding:"required"`
	City            string `json:"city" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserUpdate struct {
	Name     string `json:"name"`
	Province string `json:"province"`
	City     string `json:"city"`
}
