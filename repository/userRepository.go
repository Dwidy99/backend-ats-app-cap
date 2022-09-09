package repository

import (
	"log"
	"mini-project/entity"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	IsDuplicateEmail(email string) (tx *gorm.DB)
	VerifyCredential(email string, password string) interface{}
	InsertUser(user entity.User, applicant entity.Applicant) entity.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) InsertUser(user entity.User, applicant entity.Applicant) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	applicant.UserID = int(user.ID)
	db.connection.Save(&applicant)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}