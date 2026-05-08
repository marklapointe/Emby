package repository

import (
	"fmt"
	"time"

	"github.com/emby/emby-go/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID                          string    `json:"Id"`
	Name                        string    `json:"Name"`
	Username                    string    `json:"Username,omitempty"`
	EmailAddress                string    `json:"Email,omitempty"`
	PrimaryImageTag             string    `json:"PrimaryImageTag,omitempty"`
	HasConfiguredPassword       bool      `json:"HasConfiguredPassword"`
	HasConfiguredEasyPassword   bool      `json:"HasConfiguredEasyPassword"`
	EnableAutoLogin             bool      `json:"EnableAutoLogin"`
	LastLoginDate               time.Time `json:"LastLoginDate,omitempty"`
	CreatedDate                 time.Time `json:"CreatedDate"`
	ConnectUserName             string    `json:"ConnectUserName,omitempty"`
}

type UserRepository struct {
	*BaseRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{NewBaseRepository(db)}
}

func (r *UserRepository) CreateUser(user *model.GORMUser) error {
	if user.Id == "" {
		user.Id = fmt.Sprintf("user-%d", time.Now().UnixNano())
	}
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByID(id string) (*model.GORMUser, error) {
	var user model.GORMUser
	err := r.db.First(&user, "Id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByName(name string) (*model.GORMUser, error) {
	var user model.GORMUser
	err := r.db.Where("Name = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUsers() ([]*model.GORMUser, error) {
	var users []*model.GORMUser
	err := r.db.Select("Id, Name, PrimaryImageTag").Order("Name").Find(&users).Error
	if err != nil {
		return nil, err
	}
	if users == nil {
		users = []*model.GORMUser{}
	}
	return users, nil
}

func (r *UserRepository) UpdateUser(user *model.GORMUser) error {
	return r.db.Model(user).Updates(map[string]interface{}{
		"Name": user.Name,
		"PrimaryImageTag": user.PrimaryImageTag,
	}).Error
}

func (r *UserRepository) DeleteUser(id string) error {
	return r.db.Delete(&model.GORMUser{}, "Id = ?", id).Error
}

func (r *UserRepository) SetPassword(userID, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}
	return r.db.Model(&model.GORMUser{}).Where("Id = ?", userID).Updates(map[string]interface{}{
		"LoginPassword": string(hash),
		"InvalidLoginAttemptCount": 0,
	}).Error
}

func (r *UserRepository) ValidatePassword(userID, password string) (bool, error) {
	var user model.GORMUser
	err := r.db.Select("LoginPassword").Where("Id = ?", userID).First(&user).Error
	if err != nil {
		return false, err
	}

	if user.LoginPassword == "" {
		return false, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.LoginPassword), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *UserRepository) Authenticate(username, password string) (*model.GORMUser, error) {
	var user model.GORMUser
	err := r.db.Where("Name = ? OR Username = ? OR EmailAddress = ?", username, username, username).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	if user.LoginPassword == "" {
		if password == "" {
			return &user, nil
		}
		return nil, fmt.Errorf("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.LoginPassword), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}
	return &user, nil
}

func (r *UserRepository) UpdateLastLogin(userID string) error {
	return r.db.Model(&model.GORMUser{}).Where("Id = ?", userID).Update("LastLoginDate", time.Now()).Error
}

func (r *UserRepository) UserCount() (int64, error) {
	var count int64
	err := r.db.Model(&model.GORMUser{}).Count(&count).Error
	return count, err
}