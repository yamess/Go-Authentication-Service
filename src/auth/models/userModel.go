package models

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/yamess/auth/database"
	"github.com/yamess/auth/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

type User struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey:unique"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Email     string    `json:"email" validate:"email" gorm:"unique"`
	IsActive  *bool     `json:"is_active" validate:"required"`
	IsAdmin   *bool     `json:"is_admin" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	Base
}
type UserUpdate struct {
	FirstName string
	LastName  string
	IsActive  *bool
	IsAdmin   *bool
}

type UpdatePassword struct {
	OldPassword     string `json:"old_password" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password"`
}

// Validate User data
func (u *User) Validate() error {
	validation := validator.New()
	return validation.Struct(u)
}
func (reset *UpdatePassword) Validate() error {
	validation := validator.New()
	return validation.Struct(reset)
}

// CreateRecord CRUD
func (u *User) CreateRecord() *gorm.DB {
	hashedPwd, err := utils.HashPassword(u.Password)
	if err != nil {
		log.Fatalf(err.Error())
	}
	u.Password = hashedPwd
	u.Id = uuid.New()

	res := database.MyDB.Conn.
		Model(&u).
		Clauses(clause.Returning{}).
		Create(&u)
	return res
}

func (u *User) GetRecord(searchField string, searchValue any) *gorm.DB {
	query := fmt.Sprintf("%s = ?", searchField)
	res := database.MyDB.Conn.Model(&u).Find(&u, query, searchValue)
	return res
}

func (u *User) UpdateRecord(id uuid.UUID, userId uuid.UUID) *gorm.DB {
	u.CreatedBy = userId
	u.Id = id
	res := database.MyDB.Conn.Model(&u).Omit("CreatedBy", "CreatedAt", "Email").Updates(&u)
	return res
}

func (u *User) UpdatePassword(recordId uuid.UUID, userId uuid.UUID) *gorm.DB {
	u.UpdatedBy = userId
	u.Id = recordId
	//result := database.MyDB.Conn.Model(&u).Select( "Password", "UpdatedBy", "UpdatedAt").Updates(&u)
	result := database.MyDB.Conn.
		Model(&u).
		Clauses(clause.Returning{}).
		Select("Password", "UpdatedBy", "UpdatedAt").
		Updates(u)
	return result
}

func (u *User) DeleteRecord() *gorm.DB {
	res := database.MyDB.Conn.Delete(&u, "where email = ?", u.Email)
	return res
}

// BeforeCreate Hooks
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.CreatedAt = time.Now().UTC()
	return nil
}
