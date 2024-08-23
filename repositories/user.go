package repositories

import (
	"time"

	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	User struct {
		db    *gorm.DB
		redis *redis.Client
	}
	UserInterface interface {
		FindOne(id uuid.UUID) (*model.User, error)
		FindEmail(email string) (*model.User, error)
		Create(user *model.UserCreate) (*model.User, error)
		Update(id uuid.UUID, user *model.UserUpdate) (*model.User, error)
		Delete(id uuid.UUID) error
		VerifyOTP(otp string) (*model.User, error)
		ResendEmailOTP(email string) (*model.User, error)
		Login(email, password string) (*model.User, error)
		Logout(token string, expired time.Duration) error
	}
)

// NewUserRepository will create an object that represent the User.Repository interface
func NewUserRepository(db *gorm.DB, redis *redis.Client) UserInterface {
	return &User{db, redis}
}

// FindOne will return a user by id
func (u *User) FindOne(id uuid.UUID) (*model.User, error) {
	var user model.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, utils.NewError(utils.ErrNotFound, "user tidak ditemukan")
	}
	return &user, nil
}

// FindEmail will return a user by email
func (u *User) FindEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, utils.NewError(utils.ErrNotFound, "user tidak ditemukan")
	}
	return &user, nil
}

// Create will create a new user
func (u *User) Create(user *model.UserCreate) (*model.User, error) {
	if user.Password != user.ConfirmPassword {
		return nil, utils.NewError(utils.ErrValidation, "password and confirm password must be the same")
	}

	id := uuid.New()

	// ambil otp dari id sebanyak character 5
	otp := id.String()[0:5]

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, utils.NewError(utils.ErrUnknown, "gagal membuat password")
	}

	newUser := model.User{
		ID:         id,
		Name:       user.Name,
		Email:      user.Email,
		Password:   string(hashPassword),
		Province:   user.Province,
		City:       user.City,
		OTP:        otp,
		Role:       "user",
		IsVerified: false,
	}

	var searchUser model.User
	if err := u.db.Where("email = ?", user.Email).First(&searchUser).Error; err != nil {
		if err := u.db.Create(&newUser).Error; err != nil {
			return nil, utils.NewError(utils.ErrUnknown, "gagal membuat user")
		}
	} else {
		// update user yang sudah ada
		searchUser.Name = user.Name
		searchUser.Province = user.Province
		searchUser.City = user.City
		searchUser.OTP = otp
		searchUser.Role = "user"
		searchUser.IsVerified = false
		if err := u.db.Save(&searchUser).Error; err != nil {
			return nil, utils.NewError(utils.ErrUnknown, "gagal membuat user")
		}
		newUser = searchUser
	}

	return &newUser, nil
}

// Update will update a user by id dengan check field yang tidak dirubah maka tidak diupdate
func (u *User) Update(id uuid.UUID, user *model.UserUpdate) (*model.User, error) {
	var oldUser model.User
	if err := u.db.First(&oldUser, id).Error; err != nil {
		return nil, utils.NewError(utils.ErrNotFound, "user tidak ditemukan")
	}
	if user.Name != "" {
		oldUser.Name = user.Name
	}
	if user.Province != "" {
		oldUser.Province = user.Province
	}
	if user.City != "" {
		oldUser.City = user.City
	}
	if err := u.db.Save(&oldUser).Error; err != nil {
		return nil, utils.NewError(utils.ErrUnknown, "gagal update user")
	}
	return &oldUser, nil
}

// Delete will delete a user by id
func (u *User) Delete(id uuid.UUID) error {
	if err := u.db.Delete(&model.User{}, id).Error; err != nil {
		return utils.NewError(utils.ErrNotFound, "user tidak ditemukan")
	}
	return nil
}

// VerifyOTP will verify otp
func (u *User) VerifyOTP(otp string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("otp = ?", otp).First(&user).Error; err != nil {
		return nil, utils.NewError(utils.ErrNotFound, "otp tidak ditemukan")
	}
	user.IsVerified = true
	if err := u.db.Save(&user).Error; err != nil {
		return nil, utils.NewError(utils.ErrUnknown, "gagal verifikasi otp")
	}
	return &user, nil
}

func (u *User) ResendEmailOTP(email string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, utils.NewError(utils.ErrNotFound, "pengguna tidak ditemukan")
	}

	// ambil otp dari id sebanyak character 5
	otp := user.ID.String()[0:5]

	user.OTP = otp
	if err := u.db.Save(&user).Error; err != nil {
		return nil, utils.NewError(utils.ErrUnknown, "gagal mengirim ulang otp")
	}
	return &user, nil
}

// Login will login a user
func (u *User) Login(email, password string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, utils.NewError(utils.ErrNotFound, "pengguna tidak ditemukan")
	}

	if !user.IsVerified {
		return nil, utils.NewError(utils.ErrValidation, "email belum diverifikasi")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, utils.NewError(utils.ErrValidation, "password salah")
	}

	return &user, nil
}

func (u *User) Logout(token string, expired time.Duration) error {
	err := u.redis.Set(u.db.Statement.Context, token, true, expired).Err()
	if err != nil {
		return utils.NewError(utils.ErrUnknown, "gagal logout")
	}
	return nil
}
